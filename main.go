package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/abdelilahakebli/creditcard-checker-go/algorithms"
	"github.com/fatih/color"
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
		fail := color.New(color.FgBlack).Add(color.BgRed).Add(color.Bold)
		fail.Print(" FAIL ")
		fmt.Println(" Invalid Creadit Card Number.")
		return
	}

	pass := color.New(color.FgBlack).Add(color.BgGreen).Add(color.Bold)
	pass.Print(" PASS ")
	fmt.Println(" Valid Creadit Card Number.")

}
