package main

import "fmt" 

func main() {
	fmt.Println("Hello, World!")
	var n int
	fmt.Print("Angka :")
	fmt.Scan(&n)

	for i:=0;i<n;i++{
		fmt.Print(i++)
	}
}
