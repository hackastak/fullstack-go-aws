package main

import (
  "fmt"
)


func main() {
  animalArray := [2]string{}


  animalArray[0] = "dog"
  animalArray[1] = "cat"


  animalSlice := []string{}
  animalSlice.append(animalSlice, "moose")

  fmt.Println(animals)

  for i := 0; i < len(animals); i++ {
    fmt.Printf("this is my animal %s\n", animals[i])
  }

  for index, value := range animals {
    fmt.Printf("this is my index %d and this is my animal %s\n", index, value)
  }
  
  for i := 0; i < 10; i++ {
    fmt.Println(i)
  }

  for value := range 10 {
    fmt.Println(value)
  }
}
