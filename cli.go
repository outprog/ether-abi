package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/outprog/ether-abi/utils"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ether-abi"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "file", EnvVar: "FILE", Usage: "load your json file(abi)"},
		cli.StringFlag{Name: "type", EnvVar: "TYPE", Value: "encode", Usage: "encode or decode, default encode"},
		cli.StringFlag{Name: "method", EnvVar: "METHOD", Usage: "method name you want to encode or decode"},
		cli.StringFlag{Name: "args", EnvVar: "ARGS", Usage: "args of method"},
	}
	app.Action = run
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

func run(c *cli.Context) (err error) {
	json, err := ioutil.ReadFile(c.String("file"))
	if err != nil {
		return err
	}
	wabi := utils.NewWABI(string(json))

	aArgs := strings.Split(c.String("args"), ",")

	output := ""
	if c.String("type") == "decode" {
		// output, err = wabi.Decode()
	} else {
		output, err = wabi.Encode(c.String("method"), aArgs...)
	}
	if err != nil {
		return err
	}
	fmt.Println(output)

	return nil
}
