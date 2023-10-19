package config

type KafkaConfig struct {
	URL   string
	Topic string
}

func LoadKafkaConfig() *KafkaConfig {
	return &KafkaConfig{
		URL:   GetEnv("KAFKA_URL", "localhost:9093"),
		Topic: GetEnv("KAFKA_TOPIC", "member"),
	}
}
