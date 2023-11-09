package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/abdelilahakebli/creadircard-checker-go/algorithms"
)

func main() {

	arg := os.Args[1]
	cc, err := strconv.Atoi(arg)

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}

	check := algorithms.CalculateLuhn(cc / 10)

	if check != cc%10 {
		fmt.Println("Invalid Creadit Card Number")
		return
	}

	fmt.Println("Valid Creadit Card Number")

}
