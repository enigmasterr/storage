package main

import (
	"fmt"
	"log"

	"github.com/enigmasterr/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	restoredFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(st.Files)

	fmt.Println("It is restored", restoredFile)

}
