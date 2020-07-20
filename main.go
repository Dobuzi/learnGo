package main

import (
	"fmt"

	"github.com/Dobuzi/learnGo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	word := "hello"
	word2 := "hallo"
	definition := "Greeting"
	definition2 := "Insa"
	definition3 := "German Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	err22 := dictionary.Add(word2, definition3)
	if err22 != nil {
		fmt.Println(err22)
	}
	hello, _ := dictionary.Search(word)
	fmt.Println("found", word, ", definition:", hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
	err3 := dictionary.Update(word, definition2)
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(dictionary)
	}

	err4 := dictionary.Delete(word)
	if err4 != nil {
		fmt.Println(err4)
	} else {
		fmt.Println(dictionary)
	}

}
