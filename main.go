package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"

	"github.com/claudiu/gocron"
)

func main() {
	s := gocron.NewScheduler()
	s.Every(5).Seconds().Do(UpdateWeatherData)
	<-s.Start()
}

func WaterDesc(number int) string {
	switch {
	case number <= 5:
		return "Aman"
	case 6 <= number && number <= 8:
		return "Siaga"
	case number > 8:
		return "Bahaya"
	}
	return "Not Defined"
}

func WindDesc(number int) string {
	switch {
	case number <= 6:
		return "Aman"
	case 7 <= number && number <= 15:
		return "Siaga"
	case number > 15:
		return "Bahaya"
	}
	return "Not Defined"
}

func UpdateWeatherData() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	water := r1.Intn(20)
	waterDesc := WaterDesc(water)
	waterStr := strconv.Itoa(water)
	wind := r1.Intn(20)
	windDesc := WindDesc(wind)
	windStr := strconv.Itoa(wind)

	file := fmt.Sprintf(`
	<!DOCTYPE html>
		<html>
		<head>
			<link rel="stylesheet" href="css.bootstrap.css">
			<script src="jquery.js"></script>
		</head>
		<body>
			<h1 center>Prakiraan Cuaca</h1>
			Water : <span id="water">%s</span>
			<br>
			Description : <span id="water-description">%s</span>
			<br>
			Wind : <span id="wind">%s</span>
			<br>
			Description : <span id="wind-description">%s</span>
		</body>
		<script>
			
		</script>
	</html>
	`, waterStr, waterDesc, windStr, windDesc)

	_ = ioutil.WriteFile("weather.html", []byte(file), 0644)
}
