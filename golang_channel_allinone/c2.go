package main

import "fmt"

type Person struct {
  Name    string
  Age     uint8
  Address Addr
}

type Addr struct {
  city        string
  district    string
}

func testTranslateStruct() {
  //personChan := make(chan Person)     // 如果只写不读，或者只读不写，就会死锁
  personChan := make(chan Person, 1)     // 死锁

  person := Person{Name: "xiaoming", Age: 10, Address: Addr{city: "shenzhen", district: "longgang"}}
  personChan <-person

  person.Address = Addr{"guangzhou", "huadu"}
  fmt.Printf("person: %+v\n", person)

  newPerson := <-personChan
  fmt.Printf("new person: %+v", newPerson)
}

func main() {
  testTranslateStruct()
}