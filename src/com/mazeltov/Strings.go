package main

import (
	"fmt"
	"unicode/utf8"
)

func notObvious() {
	helloWorldStrin := "Привет мир, а не ютф, кеку %№!23&&"

	byteLen := len(helloWorldStrin)
	fmt.Println(fmt.Sprintf("Lenth in bytes %d", byteLen))

	symbLen := utf8.RuneCountInString(helloWorldStrin)

	fmt.Println(fmt.Sprintf("Symble len %d", symbLen))
}
