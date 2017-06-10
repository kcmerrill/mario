package mario

import (
	"errors"
	"fmt"
	"strings"
)

// CLI contains eveyrthing we need for our arguments
type CLI struct {
	raw  []string
	task string
	name string
	file string
	args []string
}

// NewCLI returns a CLI object
func NewCLI(params []string) *CLI {
	c := &CLI{}
	c.Parse(params)
	return c
}

// Parse will take os.args are parse it out appropriately
func (cli *CLI) Parse(args []string) error {
	cli.raw = args

	if len(cli.raw) == 0 {
		// weird ... really, you should've passed in os.Args
		return errors.New("Hmmm ... Not sure what to do here")
	}

	// used to determine the filename
	cli.name = cli.raw[0]

	// set some defaults
	cli.file = ":local"
	cli.task = ":list"

	if len(cli.raw) == 1 {
		// no need to go on.
		return nil
	}

	if len(cli.raw) >= 2 {
		// ok, so a few things here. Remote task. Local task. What?
		// mario pipeline::stage
		// mario taskname
		// mario kcmerrill/alfred:taskname
		// mario http://example.com/whatever/mario.yml:taskname
		cli.file, cli.task = cli.ParseFileAndTask(cli.raw[1], cli.name)

		if len(cli.raw) >= 3 {
			// meaning we have arguments
			cli.args = cli.raw[2:]
		} else {
			// boo! no args!
			cli.args = make([]string, 0)
		}
	}

	// no errors ... hopefully
	return nil
}

// ParseFileAndTask takes in a string, and parses it to figure out if it's remote or local task
func (cli *CLI) ParseFileAndTask(param, name string) (string, string) {
	// does it start with http?
	if strings.HasPrefix(param, "http") {
		// we have to get the http: colon out of the way :(
		bits := strings.SplitN(param, ":", 3)
		fmt.Println(bits)
		url := strings.Join(bits[0:2], ":")
		if len(bits) >= 3 {
			// alright, so we have tasks and args ...
			return url, bits[2]
		}

		return url, ":list"
	}

	// lets check if this is a github file
	if strings.Contains(param, "/") {
		bits := strings.Split(param, ":")
		url := fmt.Sprintf("https://raw.githubusercontent.com/%s/master/%s.yml", bits[0], name)
		if len(bits) >= 2 {
			return url, bits[1]
		}
		return url, ":list"
	}

	// alright, so it's not a url, it's not a github repo, it must be just a regular local task
	return ":local", param
}
