package deck

type Deck struct {
	Name     string   `yaml:"name"`
	Contents []string `yaml:"contents"`
}
