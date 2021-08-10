package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/medinfocentral/openfda"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("please enter a directory path")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("mic")

	dirPath := os.Args[1]

	fileInfos, err := ioutil.ReadDir(dirPath + "/drug-label")
	if err != nil {
		log.Fatal(err)
	}

	entityFileInfos, err := ioutil.ReadDir(dirPath + "/nsde")
	if err != nil {
		log.Fatal(err)
	}

	nsde := make(map[string]openfda.NSElement)
	for _, info := range entityFileInfos {
		f, _ := os.Open(path.Join(dirPath, "nsde", info.Name()))
		data := struct {
			Results []openfda.NSElement `json:"results"`
		}{}
		if err := json.NewDecoder(f).Decode(&data); err != nil {
			log.Fatal(err)
		}
		for _, e := range data.Results {
			if e.ProductType == "HUMAN PRESCRIPTION DRUG" {
				nsde[e.PackageNDC] = e
			}
		}
	}

	// of, err := os.Create("drugs.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer of.Close()
	// ofe := json.NewEncoder(of)

	for _, info := range fileInfos {
		fmt.Println(info.Name())

		filePath := path.Join(dirPath, "drug-label", info.Name())

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
			var (
				start, end *time.Time
			)
			if d.Drug.IsHumanPrescriptionDrug() {
				for _, id := range d.Drug.PackageNDC {
					if en, ok := nsde[id]; ok {
						nEnd, err := time.Parse("20060102", en.MarketingEndDate)
						if err == nil && (end == nil) || (end != nil && nEnd.After(*end)) {
							end = &nEnd
						}
						nStart, err := time.Parse("20060102", en.MarketingStartDate)
						if err == nil && (start == nil) || nStart.Before(*start) {
							start = &nStart
						}
						d.Drug.MarketingCategory = en.MarketingCategory
					}
				}
				d.Drug.MarketingEndDate = end
				d.Drug.MarketingStartDate = start
			}

			if len(d.Drug.SPLSetID) == 1 {
				d.Drug.ID = d.Drug.SPLSetID[0]
				if _, err := db.Collection("drugs").InsertOne(context.Background(), d.Drug); err != nil {
					fmt.Printf("%q %q", err.Error(), d.Drug.ID)
				}
			} else {
				if d.Drug.IsHumanPrescriptionDrug() {
					fmt.Println(d.Drug)
				}
			}
		}

		_ = f.Close()
	}
}
