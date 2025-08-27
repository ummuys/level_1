package main

import (
	"math/big"
	"sync"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		testNum  int
		a, b     string
		expected string
	}{
		{1, "0", "0", "0"},
		{2, "5", "7", "12"},
		{3, "1000000000000000", "1", "1000000000000001"},
		{4, "340282366920938463463374607431768211455", "1", "340282366920938463463374607431768211456"},
	}

	wg := sync.WaitGroup{}
	for _, tc := range tests {
		wg.Add(1)
		tc := tc
		func() {
			defer wg.Done()
			a, _ := new(big.Int).SetString(tc.a, 10)
			b, _ := new(big.Int).SetString(tc.b, 10)
			want, _ := new(big.Int).SetString(tc.expected, 10)

			got := addBI(a, b)
			if got.Cmp(want) != 0 {
				t.Errorf("testNum=%d: Add(%s,%s)=%s, want %s",
					tc.testNum, tc.a, tc.b, got.String(), tc.expected)
			}
		}()
	}
	wg.Wait()
}

func TestSub(t *testing.T) {
	tests := []struct {
		testNum  int
		a, b     string
		expected string
	}{
		{1, "7", "5", "2"},
		{2, "0", "0", "0"},
		{3, "1000000000000001", "1", "1000000000000000"},
		{4, "1", "340282366920938463463374607431768211456", "-340282366920938463463374607431768211455"},
	}

	wg := sync.WaitGroup{}
	for _, tc := range tests {
		wg.Add(1)
		tc := tc
		func() {
			defer wg.Done()
			a, _ := new(big.Int).SetString(tc.a, 10)
			b, _ := new(big.Int).SetString(tc.b, 10)
			want, _ := new(big.Int).SetString(tc.expected, 10)

			got := subBI(a, b)
			if got.Cmp(want) != 0 {
				t.Errorf("testNum=%d: Sub(%s,%s)=%s, want %s",
					tc.testNum, tc.a, tc.b, got.String(), tc.expected)
			}
		}()
	}
	wg.Wait()
}

func TestMul(t *testing.T) {
	tests := []struct {
		testNum  int
		a, b     string
		expected string
	}{
		{1, "0", "123456789", "0"},
		{2, "12", "11", "132"},
		{3, "123456789", "987654321", "121932631112635269"},
		{4, "18446744073709551616", "18446744073709551616", "340282366920938463463374607431768211456"},
	}

	wg := sync.WaitGroup{}
	for _, tc := range tests {
		wg.Add(1)
		tc := tc
		func() {
			defer wg.Done()
			a, _ := new(big.Int).SetString(tc.a, 10)
			b, _ := new(big.Int).SetString(tc.b, 10)
			want, _ := new(big.Int).SetString(tc.expected, 10)

			got := mulBI(a, b)
			if got.Cmp(want) != 0 {
				t.Errorf("testNum=%d: Mul(%s,%s)=%s, want %s",
					tc.testNum, tc.a, tc.b, got.String(), tc.expected)
			}
		}()
	}
	wg.Wait()
}

func TestDiv(t *testing.T) {
	tests := []struct {
		testNum  int
		a, b     string
		expected string
	}{
		{1, "10", "3", "3"},
		{2, "100", "5", "20"},
		{3, "0", "7", "0"},
		{4, "340282366920938463463374607431768211456", "2", "170141183460469231731687303715884105728"},
	}

	wg := sync.WaitGroup{}
	for _, tc := range tests {
		wg.Add(1)
		tc := tc
		func() {
			defer wg.Done()
			a, _ := new(big.Int).SetString(tc.a, 10)
			b, _ := new(big.Int).SetString(tc.b, 10)
			want, _ := new(big.Int).SetString(tc.expected, 10)

			got, _ := divBI(a, b)
			if got.Cmp(want) != 0 {
				t.Errorf("testNum=%d: Div(%s,%s)=%s, want %s",
					tc.testNum, tc.a, tc.b, got.String(), tc.expected)
			}
		}()
	}
	wg.Wait()
}
