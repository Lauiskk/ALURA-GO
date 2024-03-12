package main

import (
	"fmt"
)

func main(){
	fmt.Println("----------------------")
	fmt.Println("Mortal Clash")
	fmt.Println("----------------------")
	fmt.Println("Choose your character:")
	fmt.Println("----------------------")
	fmt.Println("1 - Sub-One ")
	fmt.Println("2 - Scorpen ")
	fmt.Println("3 - Katana ")
	fmt.Println("4 - Raidon ")
	fmt.Println("5 - Liu Keng ")
	fmt.Println("6 - Johnny Sage ")
	fmt.Println("----------------------")

	var choice int

	fmt.Print("Make your choice: ")
	fmt.Scan(&choice)

	//if choice == 1 {
	//	fmt.Println("Sub-One, freezing")
	//}else if choice == 2{
	//	fmt.Println("Scorpen, too hot")
	//}else if choice == 3{
	//	fmt.Println("Katana, sharp as f")
	//}else if choice == 4{
	//	fmt.Println("Raidon, electabuzz")
	//}else if choice == 5{
	//	fmt.Printf("Liu keng, charizard")
	//}else if choice == 6{
	//	fmt.Println("Johnny Sage, punch some")
	//}else{
	//	fmt.Println("This character does not exist!! yet")
	//}

	switch choice{
	case 1:
		fmt.Println("Sub-One, freezing")
	case 2:
		fmt.Println("Scorpen, too hot")
	case 3:
		fmt.Println("Katana, sharp as f")
	case 4:
		fmt.Println("Raidon, electabuzz")
	case 5:
		fmt.Printf("Liu keng, charizard")
	case 6:
		fmt.Println("Johnny Sage, punch some")
	default:
		fmt.Println("This character does not exist!! yet")
	}



}
