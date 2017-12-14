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
		// id := 38982249
		// phone, err := GetPhoneNumber(id)
		// if err != nil {
		// 	log.Fatalln(err)
		// }
		// log.Println(phone)
		GetOLDCars()
	case "new":
		log.Println("not implemented")
	case "all":
		log.Println("also not implemented")
	default:
		log.Println("specify PARSEMODE")
	}
}
