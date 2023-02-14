package kafka

import (
	"flag"
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"time"
)

var (
	//brokers       = flag.String("brokers", "150.158.237.39:50004", "The Kafka brokers to connect to, as a comma separated list")
	userName      = flag.String("username", "ckafka-jamo8r7b#doutok", "The SASL username")
	passwd        = flag.String("passwd", "doutokno1", "The SASL password")
	algorithm     = flag.String("algorithm", "", "The SASL SCRAM SHA algorithm sha256 or sha512 as mechanism")
	topic         = flag.String("topic", "test", "The Kafka topic to use")
	certFile      = flag.String("certificate", "", "The optional certificate file for client authentication")
	keyFile       = flag.String("key", "", "The optional key file for client authentication")
	caFile        = flag.String("ca", "", "The optional certificate authority file for TLS client authentication")
	tlsSkipVerify = flag.Bool("tls-skip-verify", false, "Whether to skip TLS server cert verification")
	useTLS        = flag.Bool("tls", false, "Use TLS to communicate with the cluster")
	mode          = flag.String("mode", "produce", "Mode to run in: \"produce\" to produce, \"consume\" to consume")
	logMsg        = flag.Bool("logmsg", false, "True to log consumed messages to console")

	logger = log.New(os.Stdout, "[Producer] ", log.LstdFlags)
)

type msgConsumerGroup struct{}

func (m msgConsumerGroup) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (m msgConsumerGroup) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (m msgConsumerGroup) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message topic:%q partition:%d offset:%d  value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))

		// 其他数据落库操作

		// 标记，sarama会自动进行提交，默认间隔1秒
		sess.MarkMessage(msg, "")
	}
	return nil
}

var ConsumerGroup msgConsumerGroup

func InitSynProducer(brokers []string) sarama.SyncProducer {
	conf := sarama.NewConfig()
	conf.Consumer.Offsets.AutoCommit.Enable = true
	conf.Consumer.Offsets.AutoCommit.Interval = time.Second * 1
	conf.Producer.Retry.Max = 1
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.Metadata.Full = true
	conf.Version = sarama.V0_10_2_0
	conf.ClientID = "sasl_scram_client"
	conf.Metadata.Full = true
	conf.Net.SASL.Enable = true
	conf.Net.SASL.User = *userName
	conf.Net.SASL.Password = *passwd
	conf.Net.SASL.Handshake = true

	// 使用同步producer，异步模式下有更高的性能，但是处理更复杂，这里建议先从简单的入手
	producer, err := sarama.NewSyncProducer(brokers, conf)
	if err != nil {
		panic(err.Error())
	}
	return producer
}

func InitConsumerGroup(brokers []string, groupId string) sarama.ConsumerGroup {
	conf := sarama.NewConfig()
	conf.Consumer.Offsets.AutoCommit.Enable = true
	conf.Consumer.Offsets.AutoCommit.Interval = time.Second * 1
	conf.Producer.Retry.Max = 1
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.Metadata.Full = true
	conf.Version = sarama.V0_10_2_0
	conf.ClientID = "sasl_scram_client"
	conf.Metadata.Full = true
	conf.Net.SASL.Enable = true
	conf.Net.SASL.User = *userName
	conf.Net.SASL.Password = *passwd
	conf.Net.SASL.Handshake = true

	cGroup, err := sarama.NewConsumerGroup(brokers, groupId, conf)
	if err != nil {
		panic(err)
	}
	return cGroup
}
