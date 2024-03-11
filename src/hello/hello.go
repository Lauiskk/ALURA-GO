package main

import (
	"fmt"
	"reflect"
)

func main(){
	//You can choose to type the variable in GO
	name := "Luis"
	version := 1.1

	fmt.Println("Olá sr.", name)
	fmt.Println("Este programa está na versão", version)

	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit")

	var choice int
	fmt.Scan( &choice )
	fmt.Println("O comando escolhido foi", choice)
	fmt.Println("O tipo da variavel da escolha é", reflect.TypeOf(choice) )
}
