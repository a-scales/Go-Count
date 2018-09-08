package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Filename: ")
	input, _ := reader.ReadString('\n')
	fmt.Print("test.txt" == strings.TrimSpace(input))
	data, err := ioutil.ReadFile(strings.TrimSpace(input))
	check(err)
	words := strings.Fields(string(data))

	wordMap := make(map[string]int)
	for _, w := range words {
		if val, ok := wordMap[w]; ok {
			wordMap[w] = val + 1
		} else {
			wordMap[w] = 1
		}
	}
	for k, v := range wordMap {
		fmt.Println(k, v)
	}
}
