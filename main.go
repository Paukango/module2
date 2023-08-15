package main

import "fmt"

func main() {
	var b int = 1
	var a *int
	a = &b
	fmt.Println(*a)
	*a = 2
	fmt.Print(b, "\n")
}
package main

import "fmt"

func main() {

	var b int = 1
	var a *int    
	a = &b
	fmt.Println(*a) 
	*a = 2
	fmt.Print(b, "\n")

}