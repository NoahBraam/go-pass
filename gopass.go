package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var passLen, numPasses int

const passChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*-_=+"

func init() {
	rand.Seed(time.Now().UnixNano())
	flag.IntVar(&passLen, "l", 12, "length of the password")
	flag.IntVar(&numPasses, "n", 1, "number of passwords")
	flag.Parse()
}

func main() {
	fmt.Println("Password Length: ", passLen, " Num Passwords: ", numPasses)
	for i := 0; i < numPasses; i++ {
		pass := randPass(passLen)
		fmt.Println(pass)
	}
}

func randPass(l int) string {
	b := make([]byte, l)
	for i := range b {
		b[i] = passChars[rand.Intn(len(passChars))]
	}
	return string(b)
}
