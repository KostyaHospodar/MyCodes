package main

import (
	"fmt"
	"time"
)

func main() {
	firstDate := time.Now()
	secondDate := time.Date(1789, 5, 5, 0, 0, 0, 0, time.UTC)
	difference := firstDate.Sub(secondDate)

	fmt.Printf("Years: %d\n", int64(difference.Hours()/24/365))
	fmt.Printf("Months: %d\n", int64(difference.Hours()/24/30))
	fmt.Printf("Days: %d\n", int64(difference.Hours()/24))
}
