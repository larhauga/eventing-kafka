package env

// Package Constants
const (
	// Eventing-Kafka Configuration
	ServiceAccountEnvVarKey = "SERVICE_ACCOUNT"
	MetricsPortEnvVarKey    = "METRICS_PORT"
	MetricsDomainEnvVarKey  = "METRICS_DOMAIN"
	HealthPortEnvVarKey     = "HEALTH_PORT"

	// Kafka Authorization
	KafkaBrokerEnvVarKey   = "KAFKA_BROKERS"
	KafkaUsernameEnvVarKey = "KAFKA_USERNAME"
	KafkaPasswordEnvVarKey = "KAFKA_PASSWORD"

	// Kafka Configuration
	KafkaProviderEnvVarKey = "KAFKA_PROVIDER"
	KafkaTopicEnvVarKey    = "KAFKA_TOPIC"

	// Knative Logging Configuration
	KnativeLoggingConfigMapNameEnvVarKey = "CONFIG_LOGGING_NAME" // Note - Matches value of configMapNameEnv constant in Knative.dev/eventing/pkg/logging !

	// Dispatcher Configuration
	ChannelKeyEnvVarKey           = "CHANNEL_KEY"
	ServiceNameEnvVarKey          = "SERVICE_NAME"
	ExponentialBackoffEnvVarKey   = "EXPONENTIAL_BACKOFF"
	InitialRetryIntervalEnvVarKey = "INITIAL_RETRY_INTERVAL"
	MaxRetryTimeEnvVarKey         = "MAX_RETRY_TIME"
)