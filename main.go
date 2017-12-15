package main

import (
	"flag"
	"log"
	"sync"
	"time"
)

const (
	// OLDCARSURL ссылка на раздел "Авто с пробегом"
	OLDCARSURL = "https://kolesa.kz/cars/"
	// NEWCARSURL ссылка на раздел "Легковые авто"
	NEWCARSURL = "https://kolesa.kz/new-cars/search/"
	// SHOWURL ссылка конкретного объявления
	SHOWURL = "https://kolesa.kz/a/show/"
	// PHONEURL ссылка на телефон
	PHONEURL = "https://kolesa.kz/a/ajaxPhones/"
)

var (
	// PARSEMODE режим парсинга
	PARSEMODE = flag.String("mode", "", "Parse MODE")
)

func main() {
	flag.Parse()
	start := time.Now()
	switch *PARSEMODE {
	case "old":
		var wg sync.WaitGroup
		wg.Add(1)
		go GetCars(OLDCARSURL, "out/old_cars.csv", &wg)
		wg.Wait()
	case "new":
		var wg sync.WaitGroup
		wg.Add(1)
		go GetCars(NEWCARSURL, "out/new_cars.csv", &wg)
		wg.Wait()
	case "all":
		var wg sync.WaitGroup
		wg.Add(2)
		go GetCars(OLDCARSURL, "out/old_cars.csv", &wg)
		go GetCars(NEWCARSURL, "out/new_cars.csv", &wg)
		wg.Wait()
	default:
		log.Println("specify PARSEMODE")
	}
	elapsed := time.Since(start)
	log.Printf("parsing took %s", elapsed)
}
