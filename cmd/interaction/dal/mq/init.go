package mq

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	config "github.com/lizaiganshenmo/mixStew/cmd/interaction/configs"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	CommentQueueName = "comment_queue"
)

var (
	CommentMQ *RabbitMQ
)

func Init() {
	CommentMQ = newRabbitMQ("rabbitmq_mixStew", CommentQueueName, "")
	go ConsumeCommentMsg() // 开个协程消费mq msg
}

type RabbitMQ struct {
	Conn         *amqp.Connection
	Channel      *amqp.Channel
	MQUrl        string
	QueueName    string
	ExchangeName string
}

func newRabbitMQ(srvName, queueName, exchangeName string) *RabbitMQ {
	mqUrl, err := config.GetRabbitMQUrl(srvName)
	if err != nil {
		klog.Error(err)
		panic(err)
	}
	if mqUrl == "" {
		panic(fmt.Sprintf("invalid mq url. srvName: %s", srvName))
	}

	conn, err := amqp.Dial(mqUrl)
	if err != nil {
		klog.Error(err)
		panic(err)
	}

	channel, err := conn.Channel()
	if err != nil {
		klog.Error(err)
		panic(err)
	}
	err = channel.Qos(1, 0, false)
	if err != nil {
		klog.Error(err)
		panic(err)
	}

	return &RabbitMQ{
		Conn:         conn,
		Channel:      channel,
		MQUrl:        mqUrl,
		QueueName:    queueName,
		ExchangeName: exchangeName,
	}

}

func (q *RabbitMQ) Publish(ctx context.Context, message string) error {
	_, err := q.Channel.QueueDeclare(
		q.QueueName,
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
		return err
	}

	err = q.Channel.PublishWithContext(ctx,
		q.ExchangeName,
		q.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		klog.CtxWarnf(ctx, "mq publish err:%v", err)
		return err
	}
	return nil
}

func (q *RabbitMQ) Consume() {

}
