package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/globalsign/mgo/bson"
	"github.com/medinfocentral/openfda"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("please enter a directory path")
	}

	dirPath := os.Args[1]

	fileInfos, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	of, err := os.Create("drugs.json")
	if err != nil {
		log.Fatal(err)
	}
	defer of.Close()
	ofe := json.NewEncoder(of)

	var result []openfda.Drug

	for _, info := range fileInfos {
		fmt.Println(info.Name())

		filePath := path.Join(dirPath, info.Name())

		f, err := os.Open(filePath)
		if err != nil {
			log.Print(err)
			continue
		}

		data := struct {
			Results []struct {
				Drug openfda.Drug `json:"openfda"`
			} `json:"results"`
		}{}

		if err := json.NewDecoder(f).Decode(&data); err != nil {
			log.Print(err)
			continue
		}

		for _, d := range data.Results {
			result = append(result, d.Drug)
		}

		f.Close()
	}

	for _, r := range result {
		r.ID = bson.NewObjectId()
		if err := ofe.Encode(r); err != nil {
			log.Fatal(err)
			continue
		}
	}
}
