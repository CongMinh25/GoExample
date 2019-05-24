package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"
)

type Person struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func worker(id int, people <-chan Person, results chan<- int) {
	for p := range people {
		fmt.Println("Workder ", id, "Done people", p)
		time.Sleep(time.Second * 5)
		results <- id
	}
}

func task(people chan<- Person) {
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
}

func getError(errorId string) string {
	errors := map[string]string{
		"1": "Error 1",
		"2": "Error 2",
	}
	outputChannel := make(chan string)
	go func() {
		time.Sleep(time.Second)
		if r, ok := errors[errorId]; ok {
			outputChannel <- nil}
		} else {
			outputChannel <- "Error not found"}
		}
	}()

	return <-outputChannel
}

func main() {
	people := make(chan Person)
	results := make(chan int)
	fmt.Println(results)
	go task(people)
	for i := 0; i < 10; i++ {
		go worker(i, people, results)
	}
	//close(people)

	for a := 0; a < 10; a++ {
		<-results
	}
	// var waitgroup sync.WaitGroup
	// waitgroup.Add(1)
	// go task(c, &waitgroup)
	// go work(c)
	// go work(c)
	// go work(c)
	// go work(c)
	// waitgroup.Wait()
	// fmt.Println("Finished Execution")
	//peopleJson, _ := json.Marshal(people)
	//fmt.Println(string(peopleJson))
}

func work(c chan Person) {
	for i := range c {
		//fmt.Print(i)
		s := i.Age
		writeFile(s)
	}
}

func writeFile(content string) {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString("aaaa")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
