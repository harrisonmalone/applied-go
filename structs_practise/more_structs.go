package main

type Person struct {
	Name string;
	Age int;
	Address
}

type Address struct {
	Name string;
	Suburb string;
}