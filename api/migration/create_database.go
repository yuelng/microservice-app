package migration

import (
	"fmt"
	"net/url"
	"os/exec"
)

func CreateDatabase() {
	p := fmt.Println
	// databse_url, err := utils.ConfigGet("database", "url")
	database_url := "postgresql://go_web_dev:go_web_dev@localhost/go_web_dev?sslmode=disable"

	u, _ := url.Parse(database_url)
	owner := u.User.Username()
	pass, _ := u.User.Password()
	path := u.Path

	// create user
	create_user := fmt.Sprintf("psql postgres -c \"create user  %s with password '%s';\"", owner, pass)
	// create database
	create_database := fmt.Sprintf("psql postgres -c \"create database %s with owner %s;\"", path[1:], owner)
	// superuser
	super_user := fmt.Sprintf("psql postgres -c \"ALTER USER %s WITH SUPERUSER;\"", owner)

	create_user_cmd := exec.Command("/bin/sh", "-c", create_user)
	create_database_cmd := exec.Command("/bin/sh", "-c", create_database)
	super_user_cmd := exec.Command("/bin/sh", "-c", super_user)

	if err := create_user_cmd.Run(); err != nil {
		p("database user " + owner + " already existed")
	}
	if err := create_database_cmd.Run(); err != nil {
		p("database " + path[1:] + " already existed")
	}
	if err := super_user_cmd.Run(); err != nil {
		p("already superuser")
	}
}
