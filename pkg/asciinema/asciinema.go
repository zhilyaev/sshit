package asciinema

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

type Asciinema struct {
	Filename string
	Login    string
	Remote   string
	Port     uint16
}

const (
	notPort   = "Port is incorrect!"
	notLogin  = "Login is incorrect!"
	notRemote = "Remote doesnt look like a ip or a domain!"
)

var ErrEmpty = errors.New("unexpected newline")

// Contains input for "asciinema rec -c" arg
func (a *Asciinema) cmd() string {
	return fmt.Sprintf("-c \"ssh %s@%s -p %d\"", a.Login, a.Remote, a.Port)
}

func (a *Asciinema) Run(debug bool) {
	if debug {
		run("asciinema", "rec", a.cmd(), a.Filename)
	} else {
		run("asciinema", "rec", "-q", "-y", a.cmd(), a.Filename)
	}
}

func run(program string, args ...string) {
	c := exec.Command(program, args...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		panic(err)
	}
}
