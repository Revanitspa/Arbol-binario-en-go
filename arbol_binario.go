/*
Beta version 1.0
RevanJZ
*/

package main

import (
    "fmt"
    
)
// se cre la estructura
type arbol struct {
	pinpollo int
	izq,der *arbol
}
//Crendo el nodo
func Crearpinpollo(x int) *arbol {
mao := new (arbol)
	mao.pinpollo = x
	mao.izq = nil
	mao.der = nil
	return mao
}
// Insertando los datos
func Insertarpinpollo(maokai *arbol, x int) *arbol  {

	if maokai == nil {
		
		return Crearpinpollo(x)
		

		
	}else if x == maokai.pinpollo {
		fmt.Println("EL valor ya existe")
		return maokai

	}else if x < maokai.pinpollo {
		maokai.izq=Insertarpinpollo(maokai.izq, x)
		return maokai
		
	}else{
		maokai.der=Insertarpinpollo(maokai.der, x)
		return maokai
	}
	
}
func PreOrden(maokai *arbol) {
	
	if maokai != nil {
		fmt.Println(maokai.pinpollo)
		PreOrden(maokai.izq)
		PreOrden(maokai.der)
		
	}

	
}
func EnOrden(maokai *arbol) {
	
	if maokai != nil {
		
		EnOrden(maokai.izq)
		fmt.Println(maokai.pinpollo)
		EnOrden(maokai.der)
		}

}
func PostOrden(maokai *arbol) {
	
	if maokai != nil {
		
		PostOrden(maokai.izq)
		PostOrden(maokai.der)
		fmt.Println(maokai.pinpollo)
		}
}
func VerArbol(maokai *arbol, x int) {
	if maokai != nil {
		VerArbol(maokai.der, x+1)
		for i := 0; i < x; i++ {
	
			 fmt.Println(maokai.pinpollo)

			}
		VerArbol(maokai.izq, x+1)
	}else{
		return;
	}
	
}
func main() {
var mao *arbol
var n int
var x int
fmt.Println(" Cuantos nodos a crear ")
fmt.Scanf("%v \n",&n)
for i := 0; i < n; i++ {
	fmt.Println("Agregar numero al arbol")
	fmt.Scanf("%v \n",&x)
	mao= Insertarpinpollo(mao,x)
	
}

fmt.Println("\n Ver el arbol \n\n")
VerArbol(mao, 0)
fmt.Println("Preorden")
PreOrden(mao)
fmt.Println("Postorden")
PostOrden(mao)
fmt.Println("Enorden")
EnOrden(mao)
 

}
