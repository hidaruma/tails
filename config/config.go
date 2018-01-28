package config

type Config struct {

	MediaDirectories []string `json:"mediadirectories"`
	DB *DBConfig `json:"db"`
}

type DBConfig struct {
	File string `json:"file"`
}