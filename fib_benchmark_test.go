package main

import (
	"testing"
	"math/big"
)

// 再帰的
func FibonacciRecursive(n int) *big.Int {
	if n < 2 {
		return big.NewInt(int64(n))
	}
	x := new(big.Int).Set(FibonacciRecursive(n - 1))
	y := new(big.Int).Set(FibonacciRecursive(n - 2))
	return x.Add(x, y)
}

// 反復的
func FibonacciIterative(n int) *big.Int {
	x, y := big.NewInt(0), big.NewInt(1)
	for i := 0; i < n; i++ {
		x, y = y, new(big.Int).Add(x, y)
	}
	return x
}

// メモ化
func FibonacciMemoized(n int, memo map[int]*big.Int) *big.Int {
	if n < 2 {
		return big.NewInt(int64(n))
	}
	if _, ok := memo[n]; !ok {
		memo[n] = new(big.Int).Add(FibonacciMemoized(n-2, memo), FibonacciMemoized(n-1, memo))
	}
	return memo[n]
}

var test_count = 30

func BenchmarkFibonacciRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciRecursive(test_count)
	}
}

func BenchmarkFibonacciIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciIterative(test_count)
	}
}

func BenchmarkFibonacciMemoized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciMemoized(test_count, make(map[int]*big.Int))
	}
}
