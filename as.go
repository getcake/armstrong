package main

import ("fmt"
        _"strconv"
        _"reflect"
        "math")

// get number of digits in number
func digitCount(num int) int {
  var count int = 0
  for num != 0 {
    num /= 10
    count += 1
  }
  return count
}

// returns math.pow as int
func intPowers(x, y int) int {
  return int(math.Pow(float64(x), float64(y)))
}

// raise each digit to the power of digitCount(number)
func armstrong(number int) int {
  var count int = digitCount(number)
  remainder := 0
  sumResult := 0
  for number != 0 {
    remainder = number % 10
    sumResult += intPowers(remainder, count)
    number = number / 10
  }
  return sumResult
}

// calculate whether the number given is an armstrong number
func main() {
  var number int
  fmt.Scanln(&number)
  count := digitCount(number)
  sum := armstrong(number)
  fmt.Printf("The sum of all digits in the number '%d' to the power of '%d' is %d\n", number, count, sum)

  if sum == number {
    fmt.Printf("The number '%d' is an armstrong number", number)
  } else {
    fmt.Printf("The number '%d' is not an armstrong number", number)
  }
}
