package messageBroker

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
	"github.com/sirupsen/logrus"
)

func WriteMessage(topic string, message *kafka.Message, partition int) (int, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", KafkaEndpoint, topic, partition)
	if err != nil {
		logrus.Fatal("failed to dial leader:", err)
		return 0, err
	}

	topicConfigs := []kafka.TopicConfig{{Topic: topic, NumPartitions: 1, ReplicationFactor: 1}}

	err = conn.CreateTopics(topicConfigs...)
	if err != nil {
		logrus.Fatal("failed to create topic:", err)
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	result, err := conn.WriteCompressedMessages(compress.Lz4.Codec(), *message)

	if err := conn.Close(); err != nil {
		logrus.Fatal("failed to close writer:", err)
	}
	return result, err
}
