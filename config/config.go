package config

// AppConfig represents the whole of config.yml
type AppConfig struct {
	MongoConfig *MongoConfig `yaml:"mongo"`
}

// MongoConfig reads the credentials of mongodb
type MongoConfig struct {
	Hosts       string                 `yaml:"url"`
	Collections *MongoCollectionConfig `yaml:"collections"`
}

// MongoCollectionConfig reads the collections part of the config.yml file
type MongoCollectionConfig struct {
	UserCollectionName         string `yaml:"users"`
	EventCollectionName        string `yaml:"events"`
	RegistrationCollectionName string `yaml:"registrations"`
}
