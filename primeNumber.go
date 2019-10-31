package main

func PrimeNumber(i int) bool {
	for j := 2; j < i; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}

//func main() {
//	//求200~800之间的素数
//	for i := 200; i <= 800; i++ {
//		if v := PrimeNumber(i); v {
//			fmt.Println(i)
//		}
//	}
//}
