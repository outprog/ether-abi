package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/outprog/ether-abi/utils"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ether-abi"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "file, f", Usage: "load abi(json) from file"},
		cli.StringFlag{Name: "type, t", Value: "encode", Usage: "encode or decode"},
		cli.StringFlag{Name: "method, m", Usage: "method name for encode or decode"},
		cli.StringFlag{Name: "args, a", Usage: "args of method, split of comma(,)"},
	}
	app.Action = run
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) (err error) {
	cfile := c.String("file")
	ctype := c.String("type")
	cmethod := c.String("method")
	cargs := c.String("args")
	if cfile == "" || ctype == "" || cmethod == "" {
		fmt.Print("wrong arguments\nusage: ether-abi -h")
		return
	}

	json, err := ioutil.ReadFile(cfile)
	if err != nil {
		return err
	}
	wabi := utils.NewWABI(string(json))

	aArgs := strings.Split(cargs, ",")

	output := ""
	if ctype == "decode" {
		// output, err = wabi.Decode()
	} else {
		output, err = wabi.Encode(cmethod, aArgs...)
		output = "0x" + output
	}
	if err != nil {
		return err
	}
	fmt.Println(output)

	return nil
}
