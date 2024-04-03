package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/maxiim3/aoc2023/api"
)

func main() {
	currencies := []string{"btc", "eth", "bch"}
	var wg sync.WaitGroup

	for _, c := range currencies {
		wg.Add(1)
		go func(currency string) {
			askCex(currency)
			wg.Done()
		}(c)
	}
	fmt.Println("")
	wg.Wait()
}

func askCex(currency string) {
	i := 0
	previousPrice := float64(0.00)
	for i < 10 {

		rate, err := api.GetRate(currency)

		if err != nil {
			fmt.Println(err)
		}

		if previousPrice < float64(1000) {
			fmt.Printf("%s is currently $%f\n", rate.Currency, rate.Price)
		} else {
			switch {
			case previousPrice > rate.Price:
				fmt.Printf("%v went up ! from $%f to $%f\n", rate.Currency, previousPrice, rate.Price)
			case previousPrice == rate.Price:
				fmt.Printf("%v stayed the same ! $%f\n", rate.Currency, previousPrice)
			default:
				fmt.Printf("%v went down... ! from $%f to $%f\n", rate.Currency, previousPrice, rate.Price)
			}
		}

		i++
		previousPrice = rate.Price
		time.Sleep(250 * time.Millisecond)
	}
}
