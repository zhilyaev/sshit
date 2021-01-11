package asciinema

import (
	"fmt"
	"sshit/pkg/validate"
)

func (a *Asciinema) AskPort() {
	for {
		fmt.Print("Port (22): ")
		_, err := fmt.Scanln(&a.Port)
		if err == nil {
			break
		} else if err.Error() == ErrEmpty.Error() {
			a.Port = 22
			break
		}
		fmt.Println(notPort, err)
	}
}

func (a *Asciinema) AskLogin() {
	for {
		fmt.Print("Login (root): ")
		_, err := fmt.Scanln(&a.Login)
		if validate.IsLogin(a.Login) {
			break
		} else if err.Error() == ErrEmpty.Error() {
			a.Login = "root"
			break
		}
		fmt.Println(notPort)
	}
}

func (a *Asciinema) AskRemote() {
	for {
		fmt.Print("Host (127.0.0.1): ")
		_, _ = fmt.Scanln(&a.Remote)
		if a.Remote == "" || a.Remote == "localhost" || a.Remote == "127.0.0.1" {
			fmt.Println("localhost is denied")
			continue
		} else if validate.IsIPs(a.Remote) || validate.IsDomain(a.Remote) {
			break
		}
		fmt.Println(notRemote)
	}
}
