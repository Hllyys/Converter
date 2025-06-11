package scraper

type Field struct {
	Selector  string   `yaml:"selector"`
	Attr      string   `yaml:"attr"`
	Transform []string `yaml:"transform"`
}

type Config struct {
	Selector string            `yaml:"selector"`
	Fields   map[string]Field  `yaml:"fields"`
}

type Item map[string]string
