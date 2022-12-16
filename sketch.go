package main

import "fmt"
import . "imp/types"

func run(e Expression) {
	s := make(map[string]Value)
	t := make(map[string]Type)
	fmt.Printf("\n ******* ")
	fmt.Printf("\n %s", e.Pretty())
	fmt.Printf("\n %s", ShowVal(e.Eval(s)))
	fmt.Printf("\n %s", ShowType(e.Infer(t)))
}

func ex1() {
	ast := Plus(Mult(Number(1), Number(2)), Number(0))

	run(ast)
}

func ex2() {
	ast := And(Bool(false), Number(0))
	run(ast)
}

func ex3() {
	ast := Or(Bool(false), Number(0))
	run(ast)
}

func main() {

	fmt.Printf("\n")

	ex1()
	ex2()
	ex3()
}
