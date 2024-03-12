package main

import (
	"fmt"
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

	//if choice == 1 {
	//	fmt.Println("Monitoring")
	//}else if choice == 2{
	//	fmt.Println("Showing Logs")
	//}else if choice == 0{
	//	fmt.Println("Ending programm")
	//}else{
	//	fmt.Println("Null comand")
	//}

	switch choice{
	case 1:
		fmt.Println("Monitoring")
	case 2:
		fmt.Println("Showing Logs")
	case 0:
		fmt.Println("Ending programm")
	default:
		fmt.Println("Null comand")
	}
}
