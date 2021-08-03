package sumOfTwoNumbers

func getSum(a, b int) int {
	var result []struct{}
	for i:=0;i<a;i++{
		result = append(result, struct{}{})
	}
	for i:=0;i<b;i++{
		result = append(result, struct{}{})
	}
	return len(result)
}
