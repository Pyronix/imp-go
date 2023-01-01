package main

import (
	"bufio"
	"fmt"
	"imp/parser"
	"imp/types"
	"os"
)

func main() {
	repl()
}

func repl() {
	reader := bufio.NewReader(os.Stdin)

	typeScope := types.TypeState{}
	valueScope := types.ValueState{}

	for {
		fmt.Print("> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		fmt.Print(" => ")

		program, err := parse(text)
		if err != nil {
			fmt.Println(err)
			continue
		}

		_, err = typeCheck(program, typeScope)
		if err != nil {
			fmt.Println(err)
			continue
		}

		v := program.Eval(valueScope)
		fmt.Println(types.ShowVal(v))
	}
}

func parse(input string) (types.Expression, error) {
	return parser.NewParser(input).ParseExpression()
}

func typeCheck(program types.Expression, scope types.TypeState) (bool, error) {
	t := program.Infer(scope)

	if t == types.TypeIllTyped {
		return false, fmt.Errorf("statement is ill typed")
	}

	return true, nil
}

func evaluate(program types.Expression, scope types.ValueState) types.Value {
	return program.Eval(scope)
}
