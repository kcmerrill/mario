package mario

import (
	"fmt"
	"testing"
)

func TestCliParseFileAndTaskURL(t *testing.T) {
	c := &CLI{}
	file, task := c.ParseFileAndTask("http://github.com/kcmerrill/alfred/alfred.yml:bingowashisnameo", "mario")

	if file != "http://github.com/kcmerrill/alfred/alfred.yml" {
		t.Logf("The github url was not parsed properly")
		t.Logf(file)
		t.FailNow()
	}

	if task != "bingowashisnameo" {
		t.Logf("The task was not parsed properly")
		t.FailNow()
	}

	_, task = c.ParseFileAndTask("http://github.com/kcmerrill/alfred/alfred.yml", "mario")
	if task != ":list" {
		t.Logf("The task was not parsed properly")
		t.FailNow()
	}
}

func TestCliParseFileAndTaskGithub(t *testing.T) {
	c := &CLI{}
	file, task := c.ParseFileAndTask("kcmerrill/alfred:taskname", "mario")

	if file != "https://raw.githubusercontent.com/kcmerrill/alfred/master/mario.yml" {
		t.Logf("The github project url was not parsed properly")
		t.Logf(file)
		t.FailNow()
	}

	if task != "taskname" {
		t.Logf("The task was not parsed properly")
		t.FailNow()
	}

	// lets try again, but with no task name
	_, task = c.ParseFileAndTask("kcmerrill/alfred", "mario")
	if task != ":list" {
		t.Logf("The task was not parsed properly")
		t.FailNow()
	}
}

func TestCliParseFileAndTaskLocal(t *testing.T) {
	c := &CLI{}
	file, task := c.ParseFileAndTask("sometask", "mario")

	if file != ":local" {
		t.Logf("This is a local task, file = :local")
		t.Logf(file)
		t.FailNow()
	}

	if task != "sometask" {
		t.Logf("The task was not parsed properly")
		t.FailNow()
	}
}

func TestCLIParse(t *testing.T) {
	c := NewCLI([]string{})
	c.Parse([]string{"mario", "http://github.com/kcmerrill/alfred/alfred.yml:my.cool.task", "arg1", "arg2"})

	if c.name != "mario" {
		t.Logf("The name of the application should be mario")
		t.FailNow()
	}

	if c.file != "http://github.com/kcmerrill/alfred/alfred.yml" {
		t.Logf("The name of the file should be github.com(ish)")
		t.FailNow()
	}

	if c.task != "my.cool.task" {
		t.Logf("The name of the task parsed should be 'my.cool.task")
		t.FailNow()
	}

	if len(c.args) != 2 {
		t.Logf("The args should be arg1 and arg2")
		t.FailNow()
	}

	c = NewCLI([]string{})
	c.Parse([]string{"mario"})
	if c.file != ":local" {
		t.Logf("Simply calling the application should yield a file of :local")
		t.FailNow()
	}

	if c.name != "mario" {
		t.Logf("The name of the application should be 'mario'")
		t.FailNow()
	}

	if c.task != ":list" {
		t.Logf("The default task should be :list")
		t.FailNow()
	}

	if len(c.args) != 0 {
		t.Logf("No arguments should be present")
		fmt.Println(c.args)
		t.FailNow()
	}
}
