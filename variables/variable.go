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

func main() {
	fmt.Println(reflect.TypeOf(age4)) //int
	fmt.Println(reflect.TypeOf(age3)) //int
}