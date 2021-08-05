package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
			fmt.Println(err)
			os.Exit(1)
	}
	var (
		a string
		b string
		c string
	)
	counter := 0

	for {
		_, err := fmt.Fscanf(file, "%s %s %s\n", &a,&b,&c) // give a patter to scan
		if err != nil {
				if err == io.EOF {
						break // stop reading the file
				}
				fmt.Println(err)
				os.Exit(1)
		}
		min,_ := strconv.Atoi(strings.Split(a,"-")[0])
		max,_ := strconv.Atoi(strings.Split(a,"-")[1])
		letter := strings.Split(b,":")[0]
		min--
		max--
		matchA := min<len(c) && string(c[min]) == letter
		matchB := max<len(c) && string(c[max]) == letter
		if (matchA || matchB) && !(matchA && matchB) {
			fmt.Println(min,max,letter,c)
			counter++
			}
	}
	fmt.Println(counter)
}
