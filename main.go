package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli/v2"
	"imp/parser"
	"imp/types"
	"io"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		UseShortOptionHandling: true,
		Action: func(cCtx *cli.Context) error {
			repl()
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "runs the given program",
				Action: func(cCtx *cli.Context) error {
					input := cCtx.Args().Get(0)

					file, err := os.Open(input)
					if err != nil {
						return err
					}

					program := parser.NewParserFromReader(file).ParseProgram()

					typeScope := &types.TypeState{}
					types.PushTypeScope(typeScope)

					valueScope := &types.ValueState{}
					types.PushValueScope(valueScope)

					if !program.Check(typeScope) {
						panic(fmt.Errorf("statement is ill typed"))
					}

					program.Eval(valueScope)

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func repl() {
	reader := bufio.NewReader(os.Stdin)

	typeScope := &types.TypeState{}
	types.PushTypeScope(typeScope)

	valueScope := &types.ValueState{}
	types.PushValueScope(valueScope)

	for {
		fmt.Print("> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}

			panic(err)
		}

		func() {
			defer func() {
				err := recover()
				if err != nil {
					fmt.Println(err)
				}
			}()

			program := parser.NewParser(text).ParseProgram()

			if !program.Check(typeScope) {
				panic(fmt.Errorf("statement is ill typed"))
			}

			v := program.Eval(valueScope)
			fmt.Printf(" => %s\n", types.ShowVal(v))
		}()
	}
}
