package main

import (
	"fmt"
	"reflect"
)

func main(){
	//You can choose to type the variable in GO
	name := "Luis"
	age := 20
	version := 1.1
	fmt.Println("Olá sr.", name);
	fmt.Println("Este programa está na versão", version)

	fmt.Println("O tipo da variavel idade é", reflect.TypeOf(age))
	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(name))
	fmt.Println("O tipo da variavel versao é", reflect.TypeOf(version))
}
