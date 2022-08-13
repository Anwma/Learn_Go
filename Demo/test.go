package Demo

/**
 * @param score: Each student's grades
 * @param ask: A series of inquiries
 * @return: Find the percentage of each question asked
 */

//func EnglishSoftware(score []int, ask []int) []int {
//	// write your code here
//	studentScore := make([]int, 200)
//	numScore := make([]int, 155)
//	n, scores := 0, 0
//	fmt.Scan(n)
//	for i := 1; i <= n; i++ {
//		fmt.Scan(n)
//		studentScore[i] = scores
//		numScore[scores]++
//	}
//	for i := 1; i <= 150; i++ {
//		numScore[i] += numScore[i-1]
//	}
//	q, id := 0, 0
//	fmt.Scan(q)
//	for q > 0 {
//		fmt.Scan(id)
//		scores = studentScore[id]
//		ans := float64(100.0 * (numScore[scores] - 1) / n)
//		fmt.Printf("%.6f", ans)
//		q--
//	}
//	return ask[]
//}
