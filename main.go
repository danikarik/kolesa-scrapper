package main

import (
	"flag"
	"log"
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

	switch *PARSEMODE {
	case "old":
		GetCars(OLDCARSURL, "out/old_cars.csv")
	case "new":
		GetCars(NEWCARSURL, "out/new_cars.csv")
	case "all":
		GetCars(OLDCARSURL, "out/old_cars.csv")
		GetCars(NEWCARSURL, "out/new_cars.csv")
	default:
		log.Println("specify PARSEMODE")
	}
}
