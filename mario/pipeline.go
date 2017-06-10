package mario

// Pipeline holds our information for our tasks
type Pipeline struct {
	Name    string `yaml:"name"`
	Summary string `yaml:"summary"`
	Dir     string `yaml:"dir"`
}
