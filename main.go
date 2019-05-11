package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	. "github.com/sanksons/riverntorch/riverntorch"
	"gopkg.in/yaml.v2"
)

func main() {

	m, err := readInput()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	people := RiverCrossers(m)
	if people.Len() == 0 {
		fmt.Printf("Need atleast 1 Person to cross")
		return
	}

	fmt.Println(
		//Get an object of RiverCrosser and use it to cross river.
		GetRiverCrosser(people.Sort()).Cross(),
	)
}

func readInput() ([]*Person, error) {

	var dataFilePath string
	flag.StringVar(&dataFilePath, "file", "data.yml", "data file path")

	flag.Parse()

	fmt.Printf("Reading Input from: %s\n", dataFilePath)
	dataBytes, err := ioutil.ReadFile(dataFilePath)
	if err != nil {
		return nil, err
	}
	m := make([]*Person, 0)
	err = yaml.Unmarshal(dataBytes, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
