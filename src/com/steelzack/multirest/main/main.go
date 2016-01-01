package main

import (
	"com/steelzack/multirest/searcher"
	"path/filepath"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)


type Value1 struct {
	Value1 string
}

type Value2 struct {
	Value2 string
}

type Value1s []Value1

type Value2s []Value2

func main() {
	t, err := GetConfiguration()
	if err != nil {
		panic(err)
	}

	readerea := new(searcher.ReaderEA)
	err = readerea.WalkThrough(t)

	// TODO Make calls to the webservice
}


func GetConfiguration() (searcher.Config, error) {
	configuration := searcher.Config{}
	filename, err := filepath.Abs("./properties.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
	}
	err = yaml.Unmarshal([]byte(string(yamlFile)), &configuration)
	if err != nil {
		log.Println(err)
	}
	return configuration, err
}

