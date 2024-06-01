package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var maxLines = 1000000000 // 1 billion: 1000000000

func main() {
	calculateAverages()
}

func calculateAverages() {
	measurements := make(map[string]Measurement)
	startTime := time.Now()

	log.Println("Reading file")
	readFile(&measurements)
	readTime := time.Since(startTime)
	log.Printf("Read %d measurements in %s\n", len(measurements), readTime)

	log.Println("Calculating average")
	calculateAverage(measurements)
	calculateTime := time.Since(startTime) - readTime

	log.Printf("Calculated average in %s\n", calculateTime)
}

type Measurement struct {
	min   int
	max   int
	sum   int
	count int
}

type FinalMeasurement struct {
	min  float32
	max  float32
	mean float32
}

type Avergage struct {
	min  float64
	mean float64
	max  float64
}

func readFile(m *map[string]Measurement) {
	file, err := os.Open("../measurements.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0

	for scanner.Scan() {
		if lines >= maxLines {
			break
		}

		text := scanner.Text()
		rawMeasurement := strings.Split(text, ";")
		strInt := strings.Replace(rawMeasurement[1], ".", "", -1)

		i, err := strconv.Atoi(strInt)
		if err != nil {
			panic(err)
		}

		location := rawMeasurement[0]

		measurement, ok := (*m)[location]
		if !ok {
			measurement = Measurement{min: i, max: i, sum: i, count: 1}
		} else {
			measurement.min = min(measurement.min, i)
			measurement.max = max(measurement.max, i)
			measurement.sum += i
			measurement.count++
		}

		(*m)[location] = measurement

		lines++
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func calculateAverage(measurements map[string]Measurement) {
	final := make(map[string]FinalMeasurement)

	keys := make([]string, 0)
	for k := range measurements {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	result := "{"
	for _, location := range keys {
		average := measurements[location]
		final[location] = FinalMeasurement{
			min:  float32(average.min) / 10,
			max:  float32(average.max) / 10,
			mean: float32(average.sum) / float32(average.count*10),
		}
		result += fmt.Sprintf("%s=%.1f/%.1f/%.1f, ", location, final[location].min, final[location].mean, final[location].max)
	}

	result = strings.TrimSuffix(result, ", ") + "}"

	log.Printf("Result: %s\n", result)
}
