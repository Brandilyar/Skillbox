package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	task_1()
}

func task_1() {
	var b bytes.Buffer
	var slicestring []string
	flag.Parse()
	args := flag.Args()
	if len(args) == 3 {
		for _, one_file := range args[:2] {
			file, err := os.Open(one_file)
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()
			byf, err := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println(err)
			}
			slicestring = append(slicestring, string(byf))
		}
		resultdata := strings.Join(slicestring, " ")

		b.WriteString(resultdata)
		if err := ioutil.WriteFile(args[2], b.Bytes(), 0644); err != nil {
			fmt.Println(err)
		}
	} else if len(args) == 2 {
		for _, one_file := range args {
			file, err := os.Open(one_file)
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()
			byf, err := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println(err)
			}
			slicestring = append(slicestring, string(byf))
		}
		resultdata := strings.Join(slicestring, " ")
		fmt.Println(resultdata)
	} else if len(args) == 1 {

		file, err := os.Open(args[0])
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		byf, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(byf))

	}
}
