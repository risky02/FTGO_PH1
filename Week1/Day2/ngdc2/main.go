package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main()  {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("List:")
	fmt.Println("1. Conditional 1")
	fmt.Println("2. Conditional 2")
	fmt.Print("Input nomor: ")
	scanner.Scan()
	fmt.Println()
	i := scanner.Text()

	switch i {
	case "1": conditional1()
	case "2": conditional2()
	default:
		fmt.Println("Anda memasukkan angka diluar list")
		return
	}
}

func conditional1()  {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Masukkan nama anda: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("Selamat datang di angka random,", name)
	randomNumber := rand.Intn (100) + 1

	switch {
	case randomNumber > 80:
		fmt.Println("Anda sangat beruntung ", name)
	case randomNumber > 60:
		fmt.Println("Anda beruntung ", name)
	case randomNumber > 40:
		fmt.Println("Anda kurang beruntung ", name)
	default:
		fmt.Println("Anda sedang sial ", name)
	}
	os.Exit(0)
}


func conditional2()  {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Masukkan nama anda: ")
	scanner.Scan()
	name := scanner.Text()
	name = strings.TrimSpace(name)

	fmt.Printf("Masukkan umur anda: ")
	scanner.Scan()
	umurStr := scanner.Text()
	umurStr = strings.TrimSpace(umurStr)

	umur, err := strconv.Atoi(umurStr) 

	if err != nil {
		fmt.Println("Masukkan umur dengan angka")
		return
	}
	
	if umur < 0 || umur > 100 {
		fmt.Println("Umur anda lebih dari rentang 0 - 100")
	} else if umur > 18 {
		fmt.Println("Silahkan Masuk, ", name)
	} else {
		fmt.Println("Dilarang Masuk")
	}	
}