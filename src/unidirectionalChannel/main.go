// unidirectionalChannel project main.go
package main

import "bufio"
import "log"
import "strings"

var sentence string = ""
var inputReader *bufio.Reader
var err error

type Person struct {
	Name    string
	Age     uint8
	Address Addr
}

type Addr struct {
	city     string
	district string
}

type PersonHandler interface {
	Batch(origs <-chan Person) <-chan PersonHandler
	Handle(orig Person)
}

type PersonHandlerImpl struct{}

func (handler PersonHandlerImpl) Batch(origs <-chan Person) <-chan Person {
	dests := make(chan Person, 100)
	go func() {
		for {
			p, ok := <-origs
			if !ok {
				close(dests)
				break
			}

			handler.Handle(p)
			dests <- p
		}
	}()
	return dests
}

func (handle PersonHandlerImpl) Handle(orig Person) {

}

func main() {
	porigs := make(chan Person, 100)
	//pdests := make(chan Person, 100)

	p1 := Person{"Harry", 32, Addr{"Beijing", "Haidian"}}
	//p2 := Person{"Sam", 30, Addr{"Guangzhou", "Yuexiu"}}
	//p3 := Person{"Jack", 35, Addr{"Guangzhou", "Haizhu"}}

	var phImpl PersonHandlerImpl

	porigs <- p1

	pdests := phImpl.Batch(porigs)

	for {
		log.Println("Please input sentence.")

		sentence, err = inputReader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		if strings.Contains(sentence, "p1") {
			porigs <- p1
		}

		if strings.Contains(sentence, "close") {
			close(porigs)

			who := <-pdests
			log.Println(who)
			break
		}
	}
}
