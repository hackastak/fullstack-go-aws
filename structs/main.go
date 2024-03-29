package main

import (
  "fmt"
)

type Person struct {
  Name string
  Age int
}

func NewPerson(name string, age int) Person {
  return Person{
    Name: name,
    Age: age,
  }
}

func main() {
  myPerson := Person{
    Name: "hackastak",
    Age: 31,
  }

  myPerson = NewPerson("Hackastak", 31)

  fmt.Printf("This is my person %+v", myPerson)
}


