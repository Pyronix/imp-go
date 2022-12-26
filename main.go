package main

import (
	"bufio"
	"fmt"
	"imp/helper"
	"imp/parser"
	"imp/types"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	//text := "1+2"

	p := parser.NewParser(text)
	exp, err := p.ParseExpression()

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Pretty: " + exp.Pretty())
		fmt.Println("Inferred Type: " + helper.StructToJson(exp.Infer(types.TypeState{})))
		fmt.Println("Evaluated Value: " + helper.StructToJson(exp.Eval(types.ValueState{})))
	}
}
