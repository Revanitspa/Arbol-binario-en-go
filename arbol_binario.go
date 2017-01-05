/*
Alpha version 2.1
RevanJZ
*/

package main

import "fmt"
type arbol struct {
	pinpollo int
	izq,der *arbol
}
func Crearpinpollo(x int) *arbol {
mao := new (arbol)
	mao.pinpollo = x
	mao.izq = nil
	mao.der = nil
	return mao
}
func Insertarpinpollo(maokai *arbol, x int) *arbol  {

	if maokai == nil {
		
		return Crearpinpollo(x)
		

		
	}else if x < maokai.pinpollo {
		maokai.izq=Insertarpinpollo(maokai.izq, x)
		return maokai
		
	}else{
		maokai.der=Insertarpinpollo(maokai.der, x)
		return maokai
	}
	
}
func main() {
var mao *arbol
 mao =Insertarpinpollo(mao,12)
 fmt.Println(mao.pinpollo)
 fmt.Println(mao.izq)
 fmt.Println(mao.der)
 mao =Insertarpinpollo(mao,1)
 fmt.Println(mao.pinpollo)
 fmt.Println(mao.izq)
 fmt.Println(mao.der)

 

}