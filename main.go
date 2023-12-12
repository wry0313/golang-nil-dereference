package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type OuterStruct struct {
	InnerStruct   InnerStruct
	CauseErr bool
}

type InnerStruct struct {
	Field string
}

func (inner InnerStruct) String() string {
	return "Hello"
}

func getInnerStruct(p OuterStruct) *InnerStruct {
	if !p.CauseErr {
		return &p.InnerStruct
	} else {
		return nil
	}
}

func main() {
	var p OuterStruct
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("do you want to cause nil pointer error? (true/false): ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	text = text[:len(text)-1]
	input, err := strconv.ParseBool(text)
	if err != nil {
		log.Fatal(err)
	}
	p.CauseErr = input
	InnerStruct := getInnerStruct(p)
	fmt.Println(InnerStruct.Field)
}
