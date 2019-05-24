package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func worker(id int, people <-chan Person, results chan<- int) {
	for j := range people {
		//fmt.Println("worker", id, "started  job", j)
		//time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- id
	}
}

func task(people chan<- Person) {

}

func main() {
	people := make(chan Person, 10)
	results := make(chan int, 10)

	for w := 0; w < 10; w++ {
		go worker(w, people, results)
	}

	cvsFile, _ := os.Open("people.csv")
	reader := csv.NewReader(bufio.NewReader(cvsFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		p := Person{
			Name: line[0],
			Age:  line[1],
		}
		people <- p
	}
	close(people)

	for a := 0; a < 10; a++ {
		<-results
	}
}
