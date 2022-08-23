package main

import "fmt"

func isValid(s string) bool {
	pMap := map[string]string{"(": ")", "{": "}", "[": "]"}
	stackTemp := make([]string, 0)
	for index := range s {
		strItem := string(s[index])
		if _, ok := pMap[strItem]; ok {
			stackTemp = append(stackTemp, strItem)
		} else {

			if len(stackTemp) > 0 {
				stackDoubleItem :=pMap[stackTemp[len(stackTemp)-1]]
				if stackDoubleItem == strItem {
					stackTemp = stackTemp[:len(stackTemp)-1]
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}
	return true
}
func main() {
	str := "([})"

	res := isValid(str)
	fmt.Println(res)
}
