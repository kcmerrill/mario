package mario

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func (m *Mario) load(file string) ([]byte, error) {
	if strings.HasPrefix(file, "http") {
		//alright, it's a url Lets go fetch it
		resp, rerr := http.Get(file)
		if rerr != nil {
			return []byte{}, fmt.Errorf("Unable to fetch %s", file)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return []byte{}, fmt.Errorf("Unable to read instructions of %s", file)
		}
		// return the goods!
		return body, nil
	}

	// Grab it locally
	dir, err := os.Getwd()
	instructions := []byte{}
	if err == nil {
		// Just keep going ...
		for {
			// Did we find a bunch of alfred files?
			patterns := []string{
				dir + fmt.Sprintf("/%s.yml", m.Name),
				dir + fmt.Sprintf("/.%s/*%s.yml", m.Name, m.Name),
				dir + fmt.Sprintf("/%s/*%s.yml", m.Name, m.Name)}
			for _, pattern := range patterns {
				if configFiles, filesErr := filepath.Glob(pattern); filesErr == nil && len(configFiles) > 0 {
					for _, configFile := range configFiles {
						if instruct, readErr := ioutil.ReadFile(configFile); readErr == nil {
							// Sweet. We found an config file. Lets save it off and return
							instructions = append(instructions, []byte("\n\n")...)
							instructions = append(instructions, instruct...)
							// Be sure that we ar relative to where we found the config file
							m.RootDir = dir
						}
					}
					return instructions, nil
				}
			}
			dir = filepath.Dir(dir)
			if dir == "/" || dir == "C:\\" {
				// We've gone too far ...
				break
			}
		}
	}

	// We didn't find anything. /cry
	return []byte{}, nil
}
