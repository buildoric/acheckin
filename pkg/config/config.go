package config

type CreateConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Config struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Create      CreateConfig `json:"create"`
}

func NewConfig() {
	// yamlFile, err := os.ReadFile("conf.yaml")
}
