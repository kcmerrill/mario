package mario

// Mario will handle all of our setup to get our tasks running
type Mario struct {
	cli          *CLI
	Name         string
	RootDir      string
	Instructions []byte
}

// New will create a new Mario object
func New(args []string) *Mario {
	m := &Mario{
		cli: NewCLI(args),
	}
	// set some defaults up
	m.Name = m.cli.name

	// load our instructions
	// TODO -> error handling?
	m.Instructions, _ = m.load(m.cli.file)
	return m
}
