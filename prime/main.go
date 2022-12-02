package main

// Created by @nicobove94 at 2021-01-01 19:00:00
// Last modified by @nicobove94 at 2021-01-01 19:00:00

import (
  "strconv"
	"fmt"
	"math"
  "os"
)

func generatePrimes(n int) {
  for i := 0; i < n; i++ {
    if isPrime(i) {
      fmt.Printf("%d ", i)
    }
  }
}

// Function to check if a number is prime
func isPrime(n int) bool {

	// Special case for n = 2
	if n == 2 {
		return true
	}

	// Check if n is even or n is less than 2
	if n % 2 == 0 || n < 2 {
		return false
	}

	// Check if n is divisible by any odd number between 3 and the square root of n
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	// If none of the above conditions are met, n is prime
	return true
}

func checkIfPrime (n int) {
  if isPrime(n) {
    fmt.Printf("%d is prime", n)
  } else {
    fmt.Printf("%d is not prime", n)
  }
}

func main() {
  allArgs := os.Args[1:]
  if len(allArgs) == 0 {
    println("PRIME NUMBERS:\nCommands:\n\t-c:\t\t\tCheck if the number is prime.\n\t-g:\t\t\tGenerate a list of prime numbers up to the inputted number.")
  }

  //fmt.Printf(allArgs[0] + " " + allArgs[1])
  for _, arg := range allArgs {
    if arg == "-c" {
      i, err := strconv.Atoi(allArgs[1])
      if err != nil {
        fmt.Println(err)
      }
      checkIfPrime(i)
    } else if arg == "-g" {
      i, err := strconv.Atoi(allArgs[1])
      if err != nil {
        fmt.Println(err)
      }
      generatePrimes(i)
    }
  }
}
// README.md

