package main

import "fmt"
import "os"

func main() {
	data := os.args[1:]
	for i := 0; i < StrLen2(data); i++ {
		fmt.Println(i)
	}
}
