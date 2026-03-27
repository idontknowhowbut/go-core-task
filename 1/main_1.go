package main

import (
	"crypto/sha256"
	"fmt"
	"reflect"
)

type vars struct {
	NumDecimal     int       // Десятичная система
	NumOctal       int       // Восьмеричная система
	NumHexadecimal int       // Шестнадцатиричная система
	Pi             float64   // Тип float64
	Name           string    // Тип string
	IsActive       bool      // Тип bool
	ComplexNum     complex64 // Тип complex64
}

func main() {
	v := reflect.ValueOf(vars{
		NumDecimal:     42,
		NumOctal:       052,
		NumHexadecimal: 0x2A,
		Pi:             3.14,
		Name:           "Golang",
		IsActive:       true,
		ComplexNum:     1 + 2i,
	})

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%T\n", v.Field(i).Interface())
	}

	varsString := ""
	for i := 0; i < v.NumField(); i++ {
		varsString += fmt.Sprint(v.Field(i).Interface())
	}

	fmt.Println(varsString)

	runeSlice := []rune(varsString)

	saltedSlice := insertSalt(runeSlice, "go-2024")

	fmt.Println(string(saltedSlice))

	hash := sha256.Sum256([]byte(string(saltedSlice)))
	fmt.Printf("%x\n", hash)

}

func insertSalt(s []rune, salt string) []rune {
	mid := len(s) / 2
	res := make([]rune, 0, len(s)+len([]rune(salt)))
	res = append(res, s[:mid]...)
	res = append(res, []rune(salt)...)
	res = append(res, s[mid:]...)
	return res
}
