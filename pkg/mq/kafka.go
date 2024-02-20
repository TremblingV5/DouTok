package mq

import (
	"github.com/Shopify/sarama"
	"sync"
)

type KafkaConfig struct {
	topic  string
	addr   []string
	config *sarama.Config
}

type KafkaProducer struct {
	config *KafkaConfig
	client sarama.SyncProducer
}

type KafkaConsumer struct {
	config *KafkaConfig
	client sarama.Consumer
}

type ProducerResp struct {
	Partition int32
	Offset    int64
}

func NewKafkaConfig(topic string, config *sarama.Config) *KafkaConfig {
	return &KafkaConfig{
		topic:  topic,
		config: config,
	}
}

func NewKafkaProducer(config *KafkaConfig) (*KafkaProducer, func(), error) {
	client, err := sarama.NewSyncProducer(config.addr, config.config)
	shutdown := func() {
		client.Close()
	}
	if err != nil {
		return nil, shutdown, err
	}

	return &KafkaProducer{
		config: config,
		client: client,
	}, shutdown, nil
}

func NewKafkaConsumer(config *KafkaConfig) (*KafkaConsumer, func(), error) {
	consumer, err := sarama.NewConsumer(config.addr, nil)
	shutdown := func() {
		consumer.Close()
	}
	if err != nil {
		return nil, shutdown, err
	}
	return &KafkaConsumer{
		config: config,
		client: consumer,
	}, shutdown, nil
}

func (producer *KafkaProducer) Produce(data []byte) (*ProducerResp, error) {
	partition, offset, err := producer.client.SendMessage(&sarama.ProducerMessage{
		Topic: producer.config.topic,
		Value: sarama.ByteEncoder(data),
	})
	if err != nil {
		return nil, err
	}

	return &ProducerResp{
		Partition: partition,
		Offset:    offset,
	}, nil
}

func (consumer *KafkaConsumer) Consume(op func(b []byte)) error {
	var wg sync.WaitGroup

	partitionList, err := consumer.client.Partitions(consumer.config.topic)
	if err != nil {
		return err
	}

	for partition := range partitionList {
		pc, err := consumer.client.ConsumePartition(consumer.config.topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			return err
		}
		wg.Add(1)

		go func(sarama.PartitionConsumer) {
			defer pc.AsyncClose()
			for message := range pc.Messages() {
				op(message.Value)
			}
		}(pc)
	}

	wg.Wait()
	return nil
}
