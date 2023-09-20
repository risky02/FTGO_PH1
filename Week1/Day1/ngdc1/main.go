package main

import "fmt"

func main() {
	// NG Challenge 1 : Variabel 1
	x, y := 10, 5
	z := x + y
	fmt.Println(z)

	// NG Challenge 1 : Variabel 2
	var myNum int32 = 50
	fmt.Println(myNum)
	var myNum2 float32 = 51
	fmt.Println(myNum2)
	var myNumStr string = "52"
	fmt.Println(myNumStr)

	// NG Challenge 1 : CLI
	fmt.Println("masukkan nama anda: " )
	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello", name)

	// NG Challenge 1 : Array & Slice 1
	people := []string {"Walt", "Jesse", "Skyler", "Saul"}
	fmt.Println(people)
	fmt.Println(len(people))
	people = append(people, "Hank", "Marie")
	fmt.Println(people)
	fmt.Println(len(people))

	// NG Challenge 1 : Array & Slice 2
	arr1 := map[string]string{
		"name":   "Hank",
		"gender": "M",
	}
	arr2 := map[string]string{
		"name":   "Heisenberg",
		"gender": "M",
	}
	arr3 := map[string]string{
		"name":   "Skyler",
		"gender": "F",
	}

	identity := [3]map[string]string{}
	identity[0] = arr1
	identity[1] = arr2
	identity[2] = arr3
	fmt.Println(identity)

	for i := range identity {
		if identity[i]["gender"] == "M" {
			identity[i]["name"] = "Mr " + identity[i]["name"]
		} else if identity[i]["gender"] == "F" {
			identity[i]["name"] = "Mrs " + identity[i]["name"]
		}
	}
	fmt.Println(identity)
}