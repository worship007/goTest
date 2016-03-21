package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

type Result struct {
	Person []Person `xml:"person"`
}
type Person struct {
	Name      string    `xml:"name,attr"`
	Age       int       `xml:"age,attr"`
	Career    string    `xml:"career"`
	Interests Interests `xml:"interests"`
}
type Interests struct {
	Interest []string `xml:"interest"`
}

func main() {
	content, err := ioutil.ReadFile("studygolang.xml")
	if err != nil {
		log.Fatal(err)
	}

	var result Result
	err = xml.Unmarshal(content, &result)

	if err != nil {
		log.Fatal(err)
	}

	for i, person := range result.Person {
		log.Println("No.", i+1, "person is:", person.Name)
		log.Println(person.Age)
		log.Println(person.Career)
		log.Println(person.Interests)
	}
}
