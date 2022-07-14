package modules

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func WriteJson(jsonFile string, data any) {
	jsonEncode, _ := json.MarshalIndent(data, "", " ")
	err := ioutil.WriteFile(jsonFile, jsonEncode, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadJson(jsonFile string, i interface{}) {
	file, _ := os.Open(jsonFile)
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(i)
	if err != nil {
		log.Fatal(err)
	}
}
