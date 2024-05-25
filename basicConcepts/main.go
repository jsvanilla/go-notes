// cada programa de go necesita al menos un paquete y una funcion main
package main

import (
	"fmt"
	"strconv"
	"time"
)

func main(){
	var x int
	x = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)
	myValue, err := strconv.ParseInt("7", 0, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}
	m := make(map[string]int)
	m["Key"] = 6
	fmt.Println(m["Key"])
	// los array tienen un length fijo, pero los slices no
	mySlice := []int{1,2,3}
	for index, value := range mySlice {
		fmt.Println(index)
		fmt.Println(value)
	}
	// append agrega un elemento al final del slice
	mySlice = append(mySlice, 16)
	fmt.Print(mySlice[len(mySlice)-1])

	
	// el canal es el que ejecuta la go routine. la palabra go solo crea la rutina
	// ejemplo de canal 
	c := make(chan int)

	// usar la palabra resevada go para llamar a una go routine
	// ejecutando la funcion en una rutina diferente a la de main
	go doSomething(c)
	<-c  // con esto, la rutina se bloquea hasta que este canal devuelva este mensaje

	// APUNTADORES
	// & ACCEDE A LA DIRECCION EN MEMORIA DE LA VARIABLE
	// * RETORNA EL VALOR ALOJADO EN ESA DIRECCION DE MEMORIA DEL APUNTADOR
	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}


// go routine
func doSomething(c chan int){
	time.Sleep(3*time.Second)
	fmt.Println(("Done"))
	c <- 1
}

