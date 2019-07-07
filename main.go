package main

import (
	"AutoReservationSys/crawler"
	"log"
	"os"
	"time"
)

var logger *log.Logger

func main() {

	//logger
	f, err := os.OpenFile("/var/log/sakazuki_alert.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer func() {
		fe := f.Close()
		logger.Println(fe)
	}()
	logger = log.New(f, "[SAKAZUKI_ALERT]", log.LstdFlags|log.LUTC)

	for {
		// start crawling
		start := time.Now()

		// スケジュールの確認を開始
		sc := crawler.ScheduleCrawler{}

		err = sc.StartCrawling()

		if err != nil {
			logger.Println(err)
		}

		elapsed := time.Since(start)
		log.Printf("Binomial took %s", elapsed)

		sleepTillNextMin()
	}

}

func sleepTillNextMin() {
	currentTime := time.Now()
	sleepTime := 180 - currentTime.Second()

	logger.Printf("sleep : %v\n", sleepTime)

	time.Sleep(time.Duration(sleepTime) * time.Second)
}
