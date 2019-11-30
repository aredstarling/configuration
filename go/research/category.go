package research

// Category structure
type Category struct {
	Name string   `yaml:"name"`
	Type string   `yaml:"type"`
	URLs []string `yaml:"urls"`
}
