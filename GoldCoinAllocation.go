package main

import (
	"strings"
)

var all int = 100

func GoldCoinAllocation(name string) int {
	var Gold int
	for _, v := range strings.Split(name, "") {
		if v == "a" || v == "A" {
			all = all - 1
			Gold = Gold + 1
			continue
		}
		if v == "e" || v == "E" {
			all = all - 1
			Gold = Gold + 1
			continue
		}
		if v == "i" || v == "I" {
			all = all - 2
			Gold = Gold + 2
			continue
		}
		if v == "o" || v == "O" {
			all = all - 3
			Gold = Gold + 3
			continue
		}
		if v == "u" || v == "U" {
			all = all - 5
			Gold = Gold + 5
			continue
		}
	}
	return Gold
}

//func main() {
//	fmt.Printf("用户: Matthew 持有%d金币\n", GoldCoinAllocation("Matthew"))
//	fmt.Printf("用户: Augustus 持有%d金币\n", GoldCoinAllocation("Augustus"))
//	fmt.Printf("用户: Emilie 持有%d金币\n", GoldCoinAllocation("Emilie"))
//	fmt.Printf("用户: Giana 持有%d金币\n", GoldCoinAllocation("Giana"))
//	fmt.Printf("用户: Elizabeth 持有%d金币\n", GoldCoinAllocation("Elizabeth"))
//	fmt.Printf("用户: Sarah 持有%d金币\n", GoldCoinAllocation("Sarah"))
//	fmt.Printf("用户: Heidi 持有%d金币\n", GoldCoinAllocation("Heidi"))
//	fmt.Printf("用户: Peter 持有%d金币\n", GoldCoinAllocation("Peter"))
//	fmt.Printf("用户: Adriano 持有%d金币\n", GoldCoinAllocation("Adriano"))
//	fmt.Printf("用户: Aaron 持有%d金币\n", GoldCoinAllocation("Aaron"))
//	fmt.Printf("总共剩余金币: %d\n", all)
//}
