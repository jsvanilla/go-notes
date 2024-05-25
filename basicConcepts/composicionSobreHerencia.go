package main

type Person struct {
	name string
	age int
}


type Developer struct {
	id int
}


type FullTimeEmployee struct {
	Developer
	Person
}