/*
Alpha version 1.2
*/

package main

import "fmt"
type arbol struct {
	pinpollo int
	izq,der *arbol
}
func crearpinpollo(x int) *arbol {
mao := new (arbol)
	mao.pinpollo = x
	mao.izq = nil
	mao.der = nil
	return mao
}
func insertarpinpollo(maokai *arbol, x int) {
	
}
func main() {
mao := new (arbol)
mao= crearpinpollo(12)
 fmt.Println(mao.pinpollo)
 fmt.Println(mao.izq)	
}