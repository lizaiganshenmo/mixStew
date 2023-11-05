package logrus

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/lizaiganshenmo/mixStew/library/constants"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

// hook for ElasticSearch
type ElasticHook struct {
	client    *elastic.Client    // es的客户端
	host      string             // es 的 host
	index     IndexNameFunc      // 获取索引的名字
	levels    []logrus.Level     // 日志的级别 info，error
	ctx       context.Context    // 上下文
	ctxCancel context.CancelFunc // 上下文cancel的函数，
	fireFunc  fireFunc           // 需要实现这个
}

// 发送到es的信息结构
type message struct {
	Host      string
	Timestamp string `json:"@timestamp"`
	Message   string
	Data      logrus.Fields
	Level     string
}

// IndexNameFunc get index name
type IndexNameFunc func() string

type fireFunc func(entry *logrus.Entry, hook *ElasticHook) error

// NewElasticHook 新建一个es hook对象
func NewElasticHook(client *elastic.Client, host string, level logrus.Level, index string) (*ElasticHook, error) {
	return NewElasticHookWithFunc(client, host, level, func() string { return index })
}

func NewElasticHookWithFunc(client *elastic.Client, host string, level logrus.Level, indexFunc IndexNameFunc) (*ElasticHook, error) {
	return newHookFuncAndFireFunc(client, host, level, indexFunc, syncFireFunc)
}

// 新建一个hook
func newHookFuncAndFireFunc(client *elastic.Client, host string, level logrus.Level, indexFunc IndexNameFunc, fireFunc fireFunc) (*ElasticHook, error) {
	var levels []logrus.Level
	for _, l := range []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	} {
		if l <= level {
			levels = append(levels, l)
		}
	}

	ctx, cancel := context.WithCancel(context.TODO())

	return &ElasticHook{
		client:    client,
		host:      host,
		index:     indexFunc,
		levels:    levels,
		ctx:       ctx,
		ctxCancel: cancel,
		fireFunc:  fireFunc,
	}, nil
}

// createMessage 创建信息
func createMessage(entry *logrus.Entry, hook *ElasticHook) *message {
	level := entry.Level.String()

	if e, ok := entry.Data[logrus.ErrorKey]; ok && e != nil {
		if err, ok := e.(error); ok {
			entry.Data[logrus.ErrorKey] = err.Error()
		}
	}

	// add request_id to log data
	logid, ok := metainfo.GetPersistentValue(entry.Context, constants.RequestIdKey)
	if ok {
		entry.Data[constants.RequestIdKey] = logid
	}

	return &message{
		hook.host,
		entry.Time.UTC().Format(time.RFC3339Nano),
		entry.Message,
		entry.Data,
		strings.ToUpper(level),
	}
}

// syncFireFunc 异步发送
func syncFireFunc(entry *logrus.Entry, hook *ElasticHook) error {
	bulkReq := hook.client.Bulk()

	data := createMessage(entry, hook)
	doc := elastic.NewBulkIndexRequest().
		Index(hook.index()).Doc(&data)

	bulkReq.Add(doc)

	_, err := bulkReq.Do(hook.ctx)
	if err != nil {
		fmt.Printf("Error send log info to es .err: %s\n", err)
	}

	return err
}

// Fire 实现 logrus hook 必须要的接口函数
func (hook *ElasticHook) Fire(entry *logrus.Entry) error {
	return hook.fireFunc(entry, hook)
}

// Levels 实现 logrus hook 必须要的接口函数
func (hook *ElasticHook) Levels() []logrus.Level {
	return hook.levels
}
