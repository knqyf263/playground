package main

import "fmt"

type Company interface {
	Info(name string) (zip, address, phone string, employee int)
}

type IsraeliCompany struct{}

func (c IsraeliCompany) Info(name string) (string, string, string, int) {
	// select company information from database and return them
	return "Zip Code", "Address", "Phone", 0
}

type Foo struct {
	Company Company
}

func (f Foo) DoSomething() {
	fmt.Println(f.Company.Info("Aqua"))
}

func main() {

}
