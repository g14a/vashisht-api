package config

// AppConfig represents the whole of config.yml
type AppConfig struct {
	MongoConfig *MongoConfig `yaml:"mongo"`
}

// MongoConfig reads the credentials of mongodb
type MongoConfig struct {
	Host        string                 `yaml:"host"`
	Port        int                    `yaml:"port"`
	Database    string                 `yaml:"database"`
	Collections *MongoCollectionConfig `yaml:"collections"`
}

// MongoCollectionConfig reads the collections part of the config.yml file
type MongoCollectionConfig struct {
	UserCollectionName         string `yaml:"users"`
	EventCollectionName        string `yaml:"events"`
	RegistrationCollectionName string `yaml:"registrations"`
}
