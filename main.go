package main

import "fmt" 

func main() {
	fmt.Println("Hello, World!")
	var n int
	fmt.Print("Angka :",n)
	fmt.Scan(&n)

	for i:=0;i<n;i++{
		fmt.Print("hasil",i++)
	}
}
