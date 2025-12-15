package main

/* First attempt
func divisorsSum(N int) int {
	sum := 0
	for i := 1; i < N; i++ {
		if N%i == 0 {
			sum += i
		}
	}
	return sum
}

func isAmicable(X int) (bool, int) {
	d1 := divisorsSum(X)
	d2 := divisorsSum(d1)
	return d2 == X, d1
}

func main() {
	var N int
	fmt.Scanf("%d", &N)
	for i := 100; i <= N; i++ {
		ok, friend := isAmicable(i)
		if ok {
			fmt.Println(i, friend)
		}
	}
}
*/

/*
My above code solution rate by GPT
Umumiy baho:
	To‘g‘rilik: 8/10
	Go uslubi: 8/10
	Algoritmik samaradorlik: 5/1
*/

// 2nd attempt
/*
func divisorsSum(N int) int {
	if N <= 1 {
		return 0
	}
	sum := 1
	for i := 2; i*i <= N; i++ {
		if N%i == 0 {
			sum += i
			if i != N/i {
				sum += N / i
			}
		}
	}
	return sum
}

func isAmicable(X int) (bool, int) {
	d1 := divisorsSum(X)
	d2 := divisorsSum(d1)
	if d1 <= X {
		return false, 0
	}
	return d2 == X, d1
}

func main() {
	var N int
	fmt.Scanf("%d", &N)
	for i := 100; i <= N; i++ {
		ok, friend := isAmicable(i)
		if ok {
			fmt.Println(i, friend)
		}
	}
}
*/

// Geminining end advanced yechimi
func advancedFindAmicable(N int) []struct{ A, B int } {
	if N < 220 { // Eng kichik dost son 220
		return nil
	}

	// 1. Bo'luvchilar Yig'indisini Hisoblash (Sieving)
	// S[i] = i ning o'zidan boshqa bo'luvchilari yig'indisi.
	aliquotSums := make([]int, N+1)
	for i := range aliquotSums {
		aliquotSums[i] = 1 // 1 har doim bo'luvchi
	}
	aliquotSums[0] = 0 // 0 uchun to'g'ri

	for i := 2; i <= N/2; i++ { // i ni bo'luvchi sifatida olamiz
		for j := 2 * i; j <= N; j += i { // i ning karralilari bo'ylab yuramiz
			aliquotSums[j] += i
		}
	}

	// 2. Dost Sonlarni Tekshirish
	var amicablePairs []struct{ A, B int }
	for a := 2; a <= N; a++ {
		b := aliquotSums[a]

		// a != b (dost son ta'rifi)
		if b > a && b <= N {
			// b ning bo'luvchilar yig'indisi a ga tengmi?
			if aliquotSums[b] == a {
				amicablePairs = append(amicablePairs, struct{ A, B int }{A: a, B: b})
			}
		}
	}
	return amicablePairs
}
