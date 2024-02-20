package mq

import (
	"context"

	"github.com/bytedance/sonic"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/lizaiganshenmo/mixStew/cmd/interaction/dal/db"
	"github.com/rabbitmq/amqp091-go"
)

const (
	CreateOperation = iota
	DeleteOperation
)

type CommentMiddle struct {
	Comment   *db.CommentInfo
	Operation int // 操作 : 新建、删除
}

func PublishComment(ctx context.Context, message string) error {
	return CommentMQ.Publish(ctx, message)
}

// 消费 mq评论消息
func ConsumeCommentMsg() {
	_, err := CommentMQ.Channel.QueueDeclare(
		CommentMQ.QueueName,
		// 是否持久化
		true,
		// 是否为自动删除
		false,
		// 是否具有排他性
		false,
		// 是否阻塞
		false,
		// 额外属性
		nil,
	)
	if err != nil {
		klog.Errorf("consume err. err:%v", err)
		return
	}

	msgCh, err := CommentMQ.Channel.Consume(
		CommentMQ.QueueName,
		"",
		false, // autoAck
		false,
		false,
		true,
		nil,
	)

	if err != nil {
		klog.Errorf("consume err. err:%v", err)
		return
	}

	for msg := range msgCh {
		t := msg
		go dealCommentInfo(&t)
	}

}

func dealCommentInfo(msg *amqp091.Delivery) {
	var commentMid CommentMiddle
	err := sonic.Unmarshal(msg.Body, &commentMid)
	if err != nil {
		klog.Warnf("err happens.err:%v", err)
		return
	}

	switch commentMid.Operation {
	case CreateOperation:
		// 添加数据到db
		err = db.CreateComment(context.TODO(), commentMid.Comment)
		if err != nil {
			klog.Warnf("db.CreateComment fail. err:%+v", err)
			return
		}
	case DeleteOperation:
		// 删除数据
		err = db.DeleteComment(context.TODO(), commentMid.Comment.CommentId)
		if err != nil {
			klog.Warnf("db.DeleteComment fail. err:%+v", err)
			return
		}
	default:
		klog.Warnf("invalid commentMid.Operation :commentMid %#v", commentMid)
	}

	msg.Ack(false)

}
