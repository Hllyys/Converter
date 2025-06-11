package scraper

type FieldConfig struct {
	Selector  string      `yaml:"selector"`
	Transform interface{} `yaml:"transform"` // string veya []string olabilir
}

type Config struct {
	Selector string                 `yaml:"selector"`
	Fields   map[string]FieldConfig `yaml:"fields"`
}

type Item map[string]string
