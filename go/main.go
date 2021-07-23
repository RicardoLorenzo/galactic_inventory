package main

import (
	"flag"

	"github.com/RicardoLorenzo/server-test/go/config"
	"github.com/RicardoLorenzo/server-test/go/inventory"
)

func main() {
	c := &config.Config{}

	flag.StringVar(&c.DBServer, "db_server", "localhost", "database server")
	flag.StringVar(&c.DBPort, "db_port", "3306", "database port")
	flag.StringVar(&c.DBUser, "db_user", "root", "database user")
	flag.StringVar(&c.DBPassword, "db_password", "secret", "database password")
	flag.StringVar(&c.DBName, "db_name", "starfleet", "database name")

	server := inventory.NewServer()
	handler, err := inventory.NewHandler(c)
	if err != nil {
		panic(err)
	}

	server.AddHandler(inventory.HandlerGet, "/list", handler.List)
	server.AddHandler(inventory.HandlerPost, "/create", handler.Create)
	server.AddHandler(inventory.HandlerPost, "/update", handler.Update)
	server.AddHandler(inventory.HandlerDelete, "/delete", handler.Delete)

	// This will block wile the server is running
	server.Run()
}
