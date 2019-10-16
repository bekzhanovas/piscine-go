package main

 

import "fmt"

 

func main() {

    final:=""

    for i:=97; i<123; i++ {

   

                character := rune(i)

    final+=string(character)

    }

    fmt.Println(final)

			
			}
				
