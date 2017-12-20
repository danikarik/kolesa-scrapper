package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

// GetPhoneNumber получает номера
func GetPhoneNumber(id int) (string, error) {

	phoneURL := fmt.Sprintf("%s?id=%d", PHONEURL, id)

	client := &http.Client{}
	req, err := http.NewRequest("GET", phoneURL, nil)
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	if err != nil {
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	ajaxModel := &AjaxModel{}
	err = json.Unmarshal(body, ajaxModel)
	if err != nil {
		return "", err
	}

	return strings.Replace(ajaxModel.Data.Model.Phone, ", ", "\n", -1), nil
}

// GetCars парсит страницу с авто
func GetCars(CARSURL, FILENAME string, wg *sync.WaitGroup) error {

	defer wg.Done()

	file, err := os.Create(FILENAME)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"Brand", "Name", "Year", "City", "Volume", "Phone"})

	var cnt int

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML("div.pager li a", func(e *colly.HTMLElement) {
		cnt, _ = strconv.Atoi(e.Text)
	})

	c.Visit(CARSURL)

	for i := 1; i <= cnt; i++ {
		url := fmt.Sprintf("%s?page=%d", CARSURL, i)
		log.Println(url)
		c = colly.NewCollector()
		c.OnHTML("div[id^='advert']", func(e *colly.HTMLElement) {
			divID := e.Attr("id")
			if !strings.Contains(divID, "note-editor") {
				dataID, _ := strconv.Atoi(e.Attr("data-id"))
				url = fmt.Sprintf("%s%d", SHOWURL, dataID)
				log.Println(url)
				car := Car{}
				c = colly.NewCollector()
				headidx := 0
				infoidx := 0
				c.OnHTML("div.product header h1 span", func(e *colly.HTMLElement) {
					if headidx == 0 {
						car.Brand = e.Text
					}
					if headidx == 1 {
						car.Name = e.Text
					}
					if headidx == 2 {
						car.Year = strings.Trim(e.Text, " ")
					}
					headidx++
				})
				c.OnHTML("div.description-body dl dd", func(e *colly.HTMLElement) {
					if infoidx == 0 {
						if e.Text == "На заказ" {
							infoidx--
						} else {
							car.City = e.Text
						}
					}
					if infoidx == 2 {
						vol := strings.Trim(strings.Replace(e.Text, "\n", "", -1), " ")
						car.Volume = vol
					}
					infoidx++
				})
				c.Visit(url)
				// GET PHONE NUMBER
				phone, err := GetPhoneNumber(dataID)
				if err != nil {
					log.Println("no phone number")
				}
				// WRITE ROW
				writer.Write([]string{
					car.Brand,
					car.Name,
					car.Year,
					car.City,
					car.Volume,
					phone,
				})
			}
		})
		c.Visit(url)
	}

	return nil
}
