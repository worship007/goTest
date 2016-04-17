// unidirectionalChannel project main.go
package main

import "bufio"

import "log"
import "time"
import "os"
import "strings"
import "strconv"

//import "runtime"

//import "strings"

var sentence string = ""
var inputReader *bufio.Reader
var err error

type PersonHandler interface {
	Batch(origs <-chan int) <-chan PersonHandler
	Handle(orig int)
}

type PersonHandlerImpl struct{}

func (handler PersonHandlerImpl) Batch(origs <-chan int) <-chan int {
	dests := make(chan int, 100)
	go func() {
		for {
			p, ok := <-origs
			if !ok {
				close(dests)
				log.Println("the channel is closed.")
				break
			}

			handler.Handle(&p)
			log.Println("send:", p)
			dests <- p
		}
	}()
	return dests
}

func (handle PersonHandlerImpl) Handle(orig *int) {
	*orig = *orig + 100
}

func main() {
	inputReader = bufio.NewReader(os.Stdin)

	var ptemp int
	var ok bool
	porigs := make(chan int, 100)
	//pdests := make(chan inta, 100)

	p1 := 1
	p2 := 2
	p3 := 3
	p4 := 4
	p5 := 5

	var phImpl PersonHandlerImpl

	porigs <- p1
	porigs <- p2
	porigs <- p3

	pdests := phImpl.Batch(porigs)

	porigs <- p4
	porigs <- p5

	time.Sleep(2000000000)

	for {
		log.Println("Please input sentence.")

		sentence, err = inputReader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		if strings.Contains(sentence, "complete") {
			close(porigs)

			for {
				ptemp, ok = <-pdests

				if !ok {
					break
				} else {
					log.Println(ptemp)
				}
			}

			break
		} else if strings.Contains(sentence, "insert") {
			log.Println("Please input the integer.")

			sentence, err = inputReader.ReadString('\n')

			if err != nil {
				log.Fatal(err)
			}

			ptemp, err = strconv.Atoi(sentence[:len(sentence)-2])

			if err != nil {
				log.Fatal(err)
			}

			porigs <- ptemp

			time.Sleep(1000000000)
		} else if strings.Contains(sentence, "quit") {
			break
		} else {
			log.Println("Please input 'insert' or 'complete' or 'quit'.")
		}
	}

	/*log.Println("Please input sentence.")

	sentence, err = inputReader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(sentence, "close") {
		ptemp, ok = <-pdests
		if !ok {
			log.Println(ptemp)
		}
	}*/
}
