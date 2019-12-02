package main

import ( 
	"fmt"
	"time"
	"math"
	"strconv"
)

func fanIn(input1, input2, input3 <-chan string) <-chan string {
	c := make(chan string)
	timeout := time.After(10 * time.Second)
	go func() {
		for {
			select {
			case s := <-input1: c<- s
			case s := <-input2: c<- s
			case s := <-input3: c<- s
			case <-timeout:
				fmt.Println("Fan in classifications timeout.")
				return 
			}
		}
	}()
	return c
}

func isPrime(n int) <-chan string {
	c := make(chan string)
	go func() {
		var is_prime = true
		if n <= 2 || n % 2 == 0 {
			is_prime = false
		} else {
			max := int( math.Sqrt( float64(n) ) )
			for i := 3; i <= max ; i += 2 {
				if n % i == 0 {
					is_prime = false
					break
				}
			}
		}
		c <- "Is prime #? " + strconv.FormatBool(is_prime)
	}()
	return c
}

func isPerfectSquare(n int) bool {
	s := int( math.Sqrt(float64(n)) )
	return s*s == n
}
func isFibonacci(n int) <-chan string {
	c := make(chan string)
	go func() {
		is_fibonacci := isPerfectSquare(5*n*n + 4) || isPerfectSquare(5*n*n - 4)
		c <- "Is Fibonacci #? " + strconv.FormatBool(is_fibonacci)	       
	}()
	return c
}

func collatzDistance(n int) <-chan string {
	c := make(chan string)
	go func(){
		step := 1
		for n != 1 {
			if n % 2 == 0 {
				n /= 2
			} else {
				n *= 3
				n++
			}
			step++
		}
		c <- "Collatz Distance: " + strconv.Itoa(step)
	}()
	return c
}

func numberSuffix(n int) string {
	v := "th"
	if n == 1 {
		v = "st"
	} else if n == 2 {
		v = "nd"
	} else if n == 3 {
		v = "rd"
	} 
	return v
}
func calcPolygon(s int, n int, c chan<- string) {
	go func() {
		m := 1
		for {
			// candidate is the mth s-sided polygonal number
			// candidate := (m*m*(s-2)-m*(s-4))/2
			candidate := (s-2)*(m*(m+1)/2)-(s-3)*m
			if n == candidate {
				c<- "Is the " + strconv.Itoa(m) + numberSuffix(m) + " " + strconv.Itoa(s) + "-sided regular polygon."
			}
			if n <= candidate {
				return
			}
			m++
		}
	}()
}
func whichPolygons(n int) {
	c := make(chan string)
	timeout := time.After(20 * time.Second)
	go func() {
		s := 3
		for s <= n {
			calcPolygon(s, n, c)
			s++
		}
	}()
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("Polygon classification timeout.")
			return
		}
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	c := fanIn(isPrime(n), isFibonacci(n), collatzDistance(n))
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(<-c)
		}
	}()
	whichPolygons(n)
}