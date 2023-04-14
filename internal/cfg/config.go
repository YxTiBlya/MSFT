package cfg

type Config struct {
	General_host string
	General_port int32

	Restaurant_service_port int32
	Customer_service_port   int32
	Statistics_service_port int32

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
