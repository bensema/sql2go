package main

import (
	"bufio"
	"fmt"
	"github.com/bensema/sql2go/core"
	"github.com/bensema/sql2go/database"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

var (
	app *cli.App
)

func init() {
	app = cli.NewApp()
	HostFlag := &cli.StringFlag{
		Name:  "host",
		Usage: "Database address",
		Value: "127.0.0.1",
	}
	PortFlag := &cli.IntFlag{
		Name:  "port",
		Usage: "Database port",
		Value: 3306,
	}
	UsernameFlag := &cli.StringFlag{
		Name:  "username",
		Usage: "Database username",
		Value: "root",
	}
	PasswordFlag := &cli.StringFlag{
		Name:  "password",
		Usage: "Database password",
		Value: "123456",
	}
	DatabaseFlag := &cli.StringFlag{
		Name:     "database",
		Usage:    "Use database",
		Required: true,
	}
	app.Flags = []cli.Flag{
		HostFlag,
		PortFlag,
		UsernameFlag,
		PasswordFlag,
		DatabaseFlag,
	}
	app.Action = action
}

func action(c *cli.Context) error {
	dbConf := database.Config{
		Active:      20,
		Idle:        10,
		IdleTimeout: 3600,
		Username:    c.String("username"),
		Password:    c.String("password"),
		Host:        c.String("host"),
		Port:        c.Int("port"),
		Database:    c.String("database"),
	}
	if c.NumFlags() > 0 {
		if err := cmd(&dbConf); err != nil {
			return cli.NewExitError(err, 1)
		}
	} else {
		_ = cli.ShowAppHelp(c)
	}
	return nil
}

func cmd(dbConf *database.Config) error {
	s2g := &core.S2G{
		Db: database.New(dbConf),
	}
	cmd := core.Cmd{}
	handlers := cmd.Handlers()
	cmd.Help()
	br := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> # ")
		line, _, _ := br.ReadLine()
		if len(line) == 0 {
			continue
		}
		tokens := strings.Split(string(line), " ")
		if handler, ok := handlers[strings.ToLower(tokens[0])]; ok {
			ret := handler(s2g, tokens)
			if ret != 0 {
				break
			}
		} else {
			fmt.Println("Unknown command>>", tokens[0])
		}
	}
	return nil
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
