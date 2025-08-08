package main

//Standard Library
import "fmt"

func main() {
	var name string
	var age int

	fmt.Println("Enter name : ")
	fmt.Scan(&name)

	fmt.Println("Enter age : ")
	fmt.Scan(&age)

	fmt.Printf("Hello, %s! you are %d years old.", name, age)

}
