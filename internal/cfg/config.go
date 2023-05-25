package cfg

type Config struct {
	Current_service string

	Restaurant_grpc_service_port int32
	Restaurant_http_service_port int32
	Restaurant_host              string

	Customer_grpc_service_port int32
	Customer_http_service_port int32
	Customer_host              string

	Statistics_grpc_service_port int32
	Statistics_http_service_port int32
	Statistics_host              string

	Rabbitmq_host       string
	Rabbitmq_port       int32
	Rabbitmq_queue_name string

	Postgres_username string
	Postgres_password string
	Postgres_host     string
	Postgres_port     int32
	Postgres_ssl      string
	Postgres_dbname   string
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
