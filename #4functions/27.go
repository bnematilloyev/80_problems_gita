package main

func lcm(a int, b int) int {
	return a * b / gcd(a, b)
}

//func main() {
//	var a, b int
//	fmt.Scan(&a, &b)
//	a, b = max(a, b), min(a, b)
//	fmt.Println(lcm(a, b))
//}
