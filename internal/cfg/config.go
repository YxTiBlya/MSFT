package cfg

type Config struct {
	General_host    string
	Logging_in_file bool

	Restaurant_grpc_service_port int32
	Restaurant_http_service_port int32

	Customer_grpc_service_port int32
	Customer_http_service_port int32

	Statistics_grpc_service_port int32
	Statistics_http_service_port int32

	Rabbitmq_host       string
	Rabbitmq_port       int32
	Rabbitmq_queue_name string

	Postgres_username string
	Postgres_password string
	Postgres_host     string
	Postgres_port     int32
	Postgres_ssl      string
	Postgres_dbname   string
	// Postgres_restaurant_dbname string
	// Postgres_consumer_dbname   string
	// Postgres_statistics_dbname string
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}

func UpdateConfig(newConfig *Config) {
	config = newConfig
}
