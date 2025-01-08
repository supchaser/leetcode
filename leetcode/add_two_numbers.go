package leetcode

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbersHard(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	mas1 := make([]int, 0, 100)
	mas2 := make([]int, 0, 100)

	headResult := result
	headL1 := l1
	headL2 := l2

	// закинуть все цифры из списка в массив
	// fmt.Printf("Печать первого списка\n")
	for headL1 != nil {
		// fmt.Print(headL1.Val)
		mas1 = append(mas1, headL1.Val)
		headL1 = headL1.Next
	}

	// fmt.Printf("\n")

	// fmt.Printf("Печать второго списка\n")
	for headL2 != nil {
		// fmt.Print(headL2.Val)
		mas2 = append(mas2, headL2.Val)
		headL2 = headL2.Next
	}

	// fmt.Printf("\n")

	// отладочная печать массивов
	// for i, val := range mas1 {
	// 	if i == 0 {
	// 		fmt.Printf("Первый массив:\n")
	// 	}
	// 	fmt.Print(val)
	// }

	// fmt.Printf("\n")

	// for i, val := range mas2 {
	// 	if i == 0 {
	// 		fmt.Printf("Второй массив:\n")
	// 	}
	// 	fmt.Print(val)
	// }

	// fmt.Printf("\n")

	str1 := ""
	for i := len(mas1) - 1; i >= 0; i-- {
		str1 += strconv.Itoa(mas1[i])
	}

	// fmt.Printf("Строка 1: %s", str1)

	str2 := ""
	for i := len(mas2) - 1; i >= 0; i-- {
		str2 += strconv.Itoa(mas2[i])
	}
	// fmt.Printf("\n")

	// fmt.Printf("Строка 2: %s", str2)

	// fmt.Printf("\n")

	number1 := new(big.Int)
	number1, ok := number1.SetString(str1, 10)
	if !ok {
		fmt.Printf("ошибка приведения типов: %v \n", ok)
		return nil
	}

	number2 := new(big.Int)
	number2, ok = number2.SetString(str2, 10)
	if !ok {
		fmt.Printf("ошибка приведения типов: %v \n", ok)
		return nil
	}
	// fmt.Printf("Число 1: %v, Число 2: %v\n", number1, number2)

	resultNumber := new(big.Int)
	resultNumber.Add(number1, number2)
	// fmt.Printf("Результат сложения: %v\n", resultNumber)
	resultNumberString := resultNumber.String()
	// fmt.Printf("Результат сложения в виде строки: %s\n", resultNumberString)

	masString := strings.Split(resultNumberString, "")

	// // отладочная печать
	// for _, val := range masString {
	// 	fmt.Println(val)
	// }

	for i := len(masString) - 1; i >= 0; i-- {
		intVal, err := strconv.Atoi(masString[i])
		if err != nil {
			return nil
		}
		// fmt.Println(result.Val, intVal)
		result.Val = intVal
		if i != 0 {
			result.Next = &ListNode{}
			result = result.Next
		}
	}

	return headResult
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	resultHead := result
	ostatok := 0

	for l1 != nil || l2 != nil || ostatok != 0 {
		x := 0
		if l1 != nil {
			x = l1.Val
		}

		y := 0
		if l2 != nil {
			y = l2.Val
		}

		sum := ostatok + x + y
		ostatok = sum / 10

		resultHead.Val = sum
		resultHead.Next = &ListNode{}
		resultHead = resultHead.Next
		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	return result
}
