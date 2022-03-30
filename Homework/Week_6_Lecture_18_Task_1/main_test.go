package main

import (
	"testing"
	"time"
)

const primeNums = 100

func Benchmark100PrimesWith0MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primesAndSleep(primeNums, 0*time.Millisecond)
	}
}

func Benchmark100PrimesWith5MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primesAndSleep(primeNums, 5*time.Millisecond)
	}
}

func Benchmark100PrimesWith10MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primesAndSleep(primeNums, 10*time.Millisecond)
	}
}

func Benchmark100GoPrimesWith0MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primesAndSleep(primeNums, 0*time.Millisecond)
	}
}

func Benchmark100GoPrimesWith5MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primesAndSleep(primeNums, 5*time.Millisecond)
	}
}

func Benchmark100GoPrimesWith10MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primesAndSleep(primeNums, 10*time.Millisecond)
	}
}
