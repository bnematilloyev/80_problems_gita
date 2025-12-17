package main

func gcd(a int, b int) int {
	for b > 0 {
		c := a % b
		a = b
		b = c
	}
	return a
}

//func main() {
//	var a, b int
//	fmt.Scan(&a, &b)
//	a, b = max(a, b), min(a, b)
//	fmt.Println(gcd(a, b))
//}
