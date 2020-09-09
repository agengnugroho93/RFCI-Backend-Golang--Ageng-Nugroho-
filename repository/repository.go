package repository

import (
	"fmt"
	"log"
	"os"
	"time"
)

const (
	datelayout = "2006-01-02T15:04:05Z07:00"
)

// Swap repo declaration
func Swap(numbers []int) {
	var swap, no int
	swap = 0
	for i := len(numbers); i > 0; i-- {
		for j := 1; j < i; j++ {
			if numbers[j-1] > numbers[j] {
				no++
				threshold := numbers[j]
				fmt.Print(no, ". ", "[", threshold, ",", numbers[j-1], "]", "->")
				numbers[j] = numbers[j-1]
				numbers[j-1] = threshold
				fmt.Println(numbers)
				swap++
			}
		}
	}
	print("Jumlah swap : ", swap)
}

// Logger function to create log system
func Logger(authorstr, titlestr, messagestr, status, scheme, baseurl string) error {
	currentTime := time.Now()
	timeNow := currentTime.Format(datelayout)

	a, err := os.OpenFile("author.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer a.Close()

	logger := log.New(a, "", log.LstdFlags)
	logger.Println("[", timeNow, "]", status, scheme, baseurl, authorstr)

	t, err := os.OpenFile("title.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer t.Close()

	logger = log.New(t, "", log.LstdFlags)
	logger.Println("[", timeNow, "]", status, scheme, baseurl, titlestr)

	m, err := os.OpenFile("message.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer m.Close()

	logger = log.New(m, "", log.LstdFlags)
	logger.Println("[", timeNow, "]", status, scheme, baseurl, messagestr)

	return err
}

// Foo fix the broken thing
func Foo(input int, numbers []int) (int, [][]int) {
	var result [][]int
	var nextIndex int

	for idx, number := range numbers {
		nextIndex = idx + 1

		for j := nextIndex; j < len(numbers); j++ {
			if nextIndex >= len(numbers) {
				break
			}
			nextNumber := numbers[j]
			if number+nextNumber == input {
				var match = []int{number, nextNumber}
				result = append(result, match)
			}

		}

	}
	return input, result
}
