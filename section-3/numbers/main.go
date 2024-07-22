package main

func main() {
	initArray := []int{1, 2, 3, 4, 5}

	for _, v := range initArray {
		if v%2 == 0 {
			println(v, "is even")
		} else {
			println(v, "is odd")
		}
	}
}
