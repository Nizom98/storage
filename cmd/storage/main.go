package main

import (
	"fmt"
	"github.com/Nizom98/storage/v2/internal/storage"
)

func main() {
	fmt.Println("dsfsf")
	st := storage.NewStorage()
	fmt.Println(st)
}

 //set GOPROXY=https://proxy.golang.org,direct