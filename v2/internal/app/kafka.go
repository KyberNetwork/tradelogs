package app

import (
	"github.com/KyberNetwork/tradelogs/v2/pkg/kafka"
	"github.com/urfave/cli"
)

var (
	kafkaAddresses = cli.StringSliceFlag{ // nolint: gochecknoglobals
		Name:   "kafka-addresses",
		Usage:  "Kafka hosts to connect",
		EnvVar: "KAFKA_ADDRESSES",
		Value:  &cli.StringSlice{"127.0.0.1:9092"},
	}
	kafkaAuthentication = cli.BoolFlag{ // nolint: gochecknoglobals
		Name:   "kafka-authentication",
		Usage:  "Kafka authentication enable",
		EnvVar: "KAFKA_AUTHENTICATION",
		//Destination: false, // nolint: gomnd
	}
	kafkaUsername = cli.StringFlag{ // nolint: gochecknoglobals
		Name:   "kafka-username",
		Usage:  "Kafka username to authenticate with",
		EnvVar: "KAFKA_USERNAME",
		Value:  "",
	}
	kafkaPassword = cli.StringFlag{ // nolint: gochecknoglobals
		Name:   "kafka-password",
		Usage:  "Kafka password to authenticate with",
		EnvVar: "KAFKA_PASSWORD",
		Value:  "",
	}
	KafkaBroadcastTopic = cli.StringFlag{
		Name:   "kafka-broadcast-topic",
		Usage:  "Kafka broadcast topic",
		EnvVar: "KAFKA_BROADCAST_TOPIC",
		//Required: true,
		Value: "trade-logs",
	}
)

func KafkaFlag() []cli.Flag {
	return []cli.Flag{
		kafkaAddresses, kafkaAuthentication, kafkaUsername, kafkaPassword, KafkaBroadcastTopic,
	}
}

func KafkaConfigFromFlags(c *cli.Context) *kafka.Config {
	authenticationEnabled := c.Bool(kafkaAuthentication.Name)
	var username, password string
	if authenticationEnabled {
		username = c.String(kafkaUsername.Name)
		password = c.String(kafkaPassword.Name)
	}
	return &kafka.Config{
		Addresses:         c.StringSlice(kafkaAddresses.Name),
		UseAuthentication: authenticationEnabled,
		Username:          username,
		Password:          password,
	}
}
