package main

/*
	例：寻找最长不含有重复字符的子串

	对于每一个字母x
	1.lastOccurred[x]不存在，或者 < start->无需操作
	2.lastOccurred[x]>= start ->更新start
	3.更新lastOccurred[x]，更新maxLength
*/

//func lengthOfNonRepeatingSubStr(s string) int {
//	lastOccurred := make(map[byte]int) // 一种是 uint8 类型，或者叫 byte 型，代表了 ASCII 码的一个字符。
//	start := 0
//	maxLength := 0
//	for i, ch := range []byte(s) {
//		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
//			start = lastI + 1
//		}
//		if i-start+1 > maxLength {
//			maxLength = i - start + 1
//		}
//		lastOccurred[ch] = i
//	}
//	return maxLength
//}

// 支持中文版
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

//func main() {
//	fmt.Println(
//		lengthOfNonRepeatingSubStr("abcabcbb"),
//		lengthOfNonRepeatingSubStr("bbbbb"),
//		lengthOfNonRepeatingSubStr("pwwkew"),
//		lengthOfNonRepeatingSubStr(""),
//		lengthOfNonRepeatingSubStr("b"),
//		lengthOfNonRepeatingSubStr("abcdef"),
//		lengthOfNonRepeatingSubStr("测试"),
//		lengthOfNonRepeatingSubStr("12321"),
//		lengthOfNonRepeatingSubStr("12321"),
//	)
//}