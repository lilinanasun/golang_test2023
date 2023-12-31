package main

import "fmt"

const female string = "female"

func main() {
	var age = 30
	var sex string = "ladyboy"

	education := "bachelor"

	fmt.Println(29 - 3854)
	fmt.Println(true && false)
	fmt.Println(age > 65 || false)
	fmt.Println(!true)
	fmt.Println(3.1415 + float64(age))
	fmt.Println(3.1415 + 30)
	fmt.Println(age > 65 && sex == "male")
	fmt.Println(education)

	for i := 0; i < 10; i = i + 2 {
		if sex == "female" {
			fmt.Println(i)
		} else if sex == "male" {
			fmt.Println(i * 2)
		} else if sex == "bisexual" {
			fmt.Println(i * 3)
		} else {
			fmt.Println(i * 4)
		}

		switch sex {
		case "female":
			fmt.Println(i)
		case "male":
			fmt.Println(i * 2)
		case "bisexual":
			fmt.Println(i * 3)
		default:
			fmt.Println(i * 4)
		}
	}

}
