package main

import (
	"fmt"
	"math"
	"strconv"
)

func catAndMouse(c1 int32, c2 int32, m int32) {

	c1s := math.Abs(float64(m - c1))
	c2s := math.Abs(float64(m - c2))

	if c1s < c2s {
		fmt.Println("Cat A")
		return
	}
	if c1s > c2s {
		fmt.Println("Cat B")
		return
	}

	fmt.Println("Mouse C")
	return
}

func getMoneySpent(keyboards []int32, drives []int32, b int32) {
	max := int32(-1)

	for _, k := range keyboards {
		for _, d := range drives {
			if max < k+d && k+d <= b {
				max = k + d
			}
		}
	}
	fmt.Println("- ELECTRONICS SHOP -")
	fmt.Println(max)
	fmt.Println()
}

func countingValleys(n int32, s string) {

	l := int32(0)
	v := int32(0)

	for _, i := range s {
		if rune(i) == 'U' {
			l++
		} else {
			l--
		}

		if l == 0 && rune(i) == 'U' {
			v++
		}
	}
	fmt.Println("- COUNTING VALLEYS -")
	fmt.Println(v)
	fmt.Println()
}

func pageCount(n int32, p int32) {
	b := n - p

	if n%2 == 0 {
		b++
	}

	a := int32(0)

	if n/2 >= p {
		a = p - ((p + 1) / 2)
	} else {
		a = b - ((b + 1) / 2)
	}
	fmt.Println("- DRAWING BOOK -")
	fmt.Println(a)
	fmt.Println()
}

func sockMerchant(n int32, ar []int32) {

	var contains = func(n int32, arr []int32) bool {
		found := false
		if len(arr) > 0 {
			for _, v := range arr {
				if v == n {
					found = true
				}
			}
		}
		return found
	}

	pair := int32(0)
	skips := []int32{}
	for i := int32(0); i < n; i++ {

		if contains(i, skips) {
			continue
		}

		count := int32(0)

		for j := int32(0); j < n; j++ {
			if ar[i] == ar[j] {
				count++
				skips = append(skips, j)
			}
		}

		if count >= 2 {
			pair += (count / 2)
		}
	}

	fmt.Println("- SOCK MERCHANT -")
	fmt.Println(pair)
	fmt.Println()

}

func bonAppetit(bill []int32, k int32, b int32) {

	total := int32(0)

	for i, it := range bill {
		if int32(i) == k {
			continue
		}
		total += it
	}

	split := (total / 2)

	fmt.Println("- BON APPETIT -")

	if split != b {
		fmt.Println(b - split)
	} else {
		fmt.Println("Bon Appetit")
	}

	fmt.Println()
}

func dayOfProgrammer(year int32) {

	fmt.Println("- DAY OF THE PROGRAMMER -")

	if year >= 1919 {
		y := strconv.Itoa(int(year))
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			fmt.Println("12.09." + y)
			fmt.Println()
			return
		}
		fmt.Println("13.09." + y)
		fmt.Println()
		return
	}

	if year >= 1700 && year <= 1917 {
		y := strconv.Itoa(int(year))
		if year%4 == 0 {
			fmt.Println("12.09." + y)
			fmt.Println()
			return
		}
		fmt.Println("13.09." + y)
		fmt.Println()
		return
	}

	if year == 1918 {
		y := strconv.Itoa(int(year))
		fmt.Println("26.09." + y)
		fmt.Println()
		return
	}
}

func migratoryBirds(arr []int32) {

	occ := make(map[string]int32)

	max := int32(0)
	temp := int32(0)

	for _, b := range arr {
		bs := strconv.Itoa(int(b))
		if occ[bs] == 0 {
			occ[bs] = 1
		} else {
			occ[bs] = occ[bs] + 1
		}

		if occ[bs] > max {
			max = occ[bs]
			temp = b
		}
	}

	min := int32(1000000)

	for k, v := range occ {
		n, _ := strconv.Atoi(k)
		if occ[strconv.Itoa(int(temp))] == v && n < int(min) {
			min = int32(n)
		}
	}

	fmt.Println(min)

}
