package main

import (
	"fmt"
	"reflect")

// Variable can be declared in var block
var (
	name string
	age int
	location string
)

// Similar data type can be put together
var (
	name1, location1 string
	age1 int
)

// Variable can be declared witout var block also
var name2 string
var age2 string
var location2 string

// Variable declaration can include initialization, each per variable
var (
	name3 string = "pankaj"
	age3 int = 18
	location3 = "bangalore"
)

// Variable can infer type if it is initialized
var (
	name4 = "pankaj"
	age4 = 18
	location4 = "bangalore"	
)

// Variables can be initialize in same line
var (
	name5, age5, location5 = "pankaj", 18, "bangalore"
)

// Inside a function, := assignment can be used instead of var
func SomeFunc() {
	name6, location6 := "pankaj", "bangalore"
	age6 := 18
	fmt.Printf("%s %d %s \n",name6, age6, location6)
}

// Varaible can also be of function type
var action = func(aString string) {
	//doing something
	fmt.Println(aString)
}

func main() {
	fmt.Println(reflect.TypeOf(age4)) //int
	fmt.Println(reflect.TypeOf(age3)) //int

	SomeFunc()
	action("hello world")
}