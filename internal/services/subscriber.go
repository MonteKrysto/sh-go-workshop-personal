package services

import (
	"encoding/csv"
	"github.com/SpringCare/sh-go-workshop/internal/interfaces"
	"github.com/SpringCare/sh-go-workshop/internal/models"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

type SubscriberService struct {
	repository interfaces.FileUpload
}

func NewSubscriberService(repository interfaces.FileUpload) *SubscriberService {
	return &SubscriberService{
		repository: repository,
	}
}

func (ss SubscriberService) Upload(fileName string, subscriber models.Subscriber) error {
	ss.process(fileName)
	//return es.repository.Create(file, ctx)
	return nil
}

func (ss SubscriberService) process(fileName string) {
	f, err := os.Open(fileName) //Open("cities.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rows := genRows(f)
	// single go routine pipeline
	//upperRowsChannel := upperCities(rows)
	//// consume the rows from the channel
	//for row := range upperRowsChannel {
	//	//r := upper(row)
	//	log.Println(row)
	//}
	// fan-out pattern - split the work into multiple go routines which will run concurrently
	// and compete to consume from the channel to process the data

	// different workers will consume from the channel
	//       __ worker1
	// rows /__ worker2
	//      \__ worker3
	//       __ workern...
	//filterPopulation := filterOutMinPopulations(40000)
	//var list []models.Subscriber
	upperRowsChannel1 := upperCities(rows)
	upperRowsChannel2 := upperCities(rows)
	upperRowsChannel3 := upperCities(rows)

	//wg := &sync.WaitGroup{}
	//wg.Add(1000)
	//sem := make(chan bool, 10)
	for c := range fanIn(upperRowsChannel1, upperRowsChannel2, upperRowsChannel3) {

		//sem <- true
		//go func() {
		//	defer wg.Done()
		//	defer func() { <-sem }()
		ss.repository.Upload(fileName, c)
		//}()
		//wg.Wait()
	}

}

func genRows(r io.Reader) chan models.Subscriber {
	out := make(chan models.Subscriber)

	go func() {
		reader := csv.NewReader(r)
		_, err := reader.Read()
		if err != nil {
			log.Fatal(err)
		}
		for {
			row, err := reader.Read()
			if err != nil {
				log.Fatal(err)
			}
			if err == io.EOF {
				break
			}

			out <- models.Subscriber{
				FirstName: row[0],
				LastName:  row[1],
				Email:     row[2],
				DOB:       row[3],
				Phone:     row[4],
				SSN:       row[5],
				City:      row[6],
				Country:   row[7],
				Interest1: row[8],
				Interest2: row[9],
				Interest3: row[10],
			}
		}
		close(out)
	}()
	return out
}

//func upper(city models.Subscriber) models.Subscriber {
//	return models.Subscriber{Name: strings.ToUpper(city.Name), Population: city.Population}
//}

// upperCities takes a read-only channel of models.Subscriber and returns a read-only channel of models.Subscriber
func upperCities(subscribers <-chan models.Subscriber) <-chan models.Subscriber {
	out := make(chan models.Subscriber)
	// This is a concurrent pipeline so the work of ingesting and transforming the data is done in parallel
	// rather than done sequentially on every row that is read in
	go func() {
		for subscriber := range subscribers {
			out <- models.Subscriber{
				FirstName: strings.ToUpper(subscriber.FirstName),
				LastName:  strings.ToUpper(subscriber.LastName),
				Email:     subscriber.Email,
				DOB:       subscriber.DOB,
				Phone:     subscriber.Phone,
				SSN:       subscriber.SSN,
				City:      subscriber.City,
				Country:   subscriber.Country,
				Interest1: subscriber.Interest1,
				Interest2: subscriber.Interest2,
				Interest3: subscriber.Interest3,
			}
		}
		close(out)
	}()
	return out
}

//func filterOutMinPopulations1(cities <-chan models.Subscriber) <-chan models.Subscriber {
//	out := make(chan models.Subscriber)
//	// This is a concurrent pipeline so the work of ingesting and transforming the data is done in parallel
//	// rather than done sequentially on every row that is read in
//	go func() {
//		for city := range cities {
//			if city.Population > 40000 {
//				out <- models.Subscriber{Name: strings.ToUpper(city.Name), Population: city.Population}
//			}
//		}
//		close(out)
//	}()
//	return out
//}

// fanIn takes a variable number of read-only channels of models.Subscriber and returns a read-only channel of models.Subscriber
// it reconciles all the channels into a single channel to be read from in the main function (the main channel)
func fanIn(channels ...<-chan models.Subscriber) <-chan models.Subscriber {
	out := make(chan models.Subscriber)

	wg := &sync.WaitGroup{}
	wg.Add(len(channels))

	for _, c := range channels {
		//fmt.Println(runtime.NumGoroutine())
		go func(subscriber <-chan models.Subscriber) {
			for subscriber := range subscriber {
				out <- subscriber
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
