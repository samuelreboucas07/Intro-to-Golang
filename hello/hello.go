package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome string = "samuel"
	idade := 50 //Declarando variável :=
	fmt.Println("Olá senhor", nome)
	fmt.Println("Sua idade é", idade)

	fmt.Println("O tipo da variável idade é", reflect.TypeOf(idade))
}
