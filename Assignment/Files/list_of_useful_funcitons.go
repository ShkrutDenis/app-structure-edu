package main

import (
	"fmt"
	"math/rand"
)

func GenerateUniqueID() string {
	return "unique_id"
}

func RemoveAllDigits(str string) string {
	return fmt.Sprintf("%s_without_digits", str)
}

func CheckLuck(list []string) string {
	return list[rand.Intn(len(list)+1)]
}
