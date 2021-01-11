package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/kshvakov/clickhouse"
	"io/ioutil"
	"os"
	"os/user"
	"sshit/pkg/asciinema"
	"sshit/pkg/database"
	"sshit/pkg/watcher"
	"strconv"
	"time"
)

var (
	debug   bool
	db      *sqlx.DB
	a       asciinema.Asciinema
	session database.Session
)

func main() {
	// Set DEBUG
	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))
	if err != nil {
		debug = false
	}
	if debug {
		fmt.Println("DEBUG MODE ON")
	}

	// Set DB
	strConnection := os.Getenv("DB")
	if strConnection == "" {
		strConnection = "tcp://127.0.0.1:9000?database=mydb"
	}
	db = sqlx.MustConnect("clickhouse", strConnection)
	defer db.Close()

	// Set username
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	session.User = u.Username
	if debug {
		fmt.Println("username is ", session.User)
	}

	showLogo()

	// Where you want to connect via SSH?
	a.AskRemote()
	a.AskPort()
	// We had to connect via root
	a.Login = "root" // a.AskLogin()
	if debug {
		fmt.Println(a.Login)
		fmt.Println(a.Remote)
		fmt.Println(a.Port)
	}

	// Start Session
	session.Created = time.Now()
	session.UUID, _ = database.GenUUID(db)

	// Create full path if not exists
	dir := "./var/log/sshit/" + session.User + "/" + session.Remote
	a.Filename = dir + fmt.Sprintf(
		"%d-%02d-%02dT%02d:%02d:%02d-00:00.cast",
		session.Created.Year(), session.Created.Month(), session.Created.Day(),
		session.Created.Hour(), session.Created.Minute(), session.Created.Second(),
	)
	if debug {
		fmt.Println(a.Filename)
		fmt.Println(session.UUID)
	}
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		panic(err)
	}
	// Registration session in DB
	err = session.Insert(db)
	if err != nil {
		panic(err)
	}

	// OnChange asciinema cast
	go watcher.OnChange(a.Filename, func(s string) {
		log := &database.Log{
			Doc:         s,
			Created:     time.Now(),
			SessionUUID: session.UUID,
		}
		err = log.Insert(db)
		if err != nil {
			panic(err)
		}
	})

	a.Run(debug)

	fmt.Println("See you soon!")
	os.Exit(0)
}

func showLogo() {
	logo, err := ioutil.ReadFile("assets/logo.txt")
	if err == nil {
		fmt.Println(string(logo))
	}
}
