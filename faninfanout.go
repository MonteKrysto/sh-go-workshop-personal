package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type City struct {
	Name       string
	Population int
}

func main() {
	f, err := os.Open("cities.csv")
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
	upperRowsChannel1 := upperCities(filterOutMinPopulations1(rows))
	upperRowsChannel2 := upperCities(filterOutMinPopulations1(rows))
	upperRowsChannel3 := upperCities(filterOutMinPopulations1(rows))
	for c := range fanIn(upperRowsChannel1, upperRowsChannel2, upperRowsChannel3) {
		log.Println(c)
	}

}

func genRows(r io.Reader) chan City {
	out := make(chan City)

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
			population, err := strconv.Atoi(row[9])
			if err != nil {
				continue
			}
			out <- City{
				Name:       row[1],
				Population: population,
			}
		}
		close(out)
	}()
	return out
}

func upper(city City) City {
	return City{Name: strings.ToUpper(city.Name), Population: city.Population}
}

// upperCities takes a read-only channel of City and returns a read-only channel of City
func upperCities(cities <-chan City) <-chan City {
	out := make(chan City)
	// This is a concurrent pipeline so the work of ingesting and transforming the data is done in parallel
	// rather than done sequentially on every row that is read in
	go func() {
		for city := range cities {
			out <- City{Name: strings.ToUpper(city.Name), Population: city.Population}
		}
		close(out)
	}()
	return out
}

func filterOutMinPopulations1(cities <-chan City) <-chan City {
	out := make(chan City)
	// This is a concurrent pipeline so the work of ingesting and transforming the data is done in parallel
	// rather than done sequentially on every row that is read in
	go func() {
		for city := range cities {
			if city.Population > 40000 {
				out <- City{Name: strings.ToUpper(city.Name), Population: city.Population}
			}
		}
		close(out)
	}()
	return out
}

func filterOutMinPopulations(min int) func(cities <-chan City) <-chan City {
	return func(cities <-chan City) <-chan City {
		out := make(chan City)
		// This is a concurrent pipeline so the work of ingesting and transforming the data is done in parallel
		// rather than done sequentially on every row that is read in
		go func() {
			for city := range cities {
				if city.Population > 40000 {
					out <- City{Name: strings.ToUpper(city.Name), Population: city.Population}
				}
			}
			close(out)
		}()
		return out
	}
}

// fanIn takes a variable number of read-only channels of City and returns a read-only channel of City
// it reconciles all the channels into a single channel to be read from in the main function (the main channel)
func fanIn(channels ...<-chan City) <-chan City {
	out := make(chan City)

	wg := &sync.WaitGroup{}
	wg.Add(len(channels))

	for _, c := range channels {
		fmt.Println(runtime.NumGoroutine())
		go func(city <-chan City) {
			for city := range city {
				out <- city
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
