package main

import "fmt"

// 或许这道题不适合我，因为我一旦数组开大了编译器就给我报错，那我就只说一下思路了。首先先用一个二维数组将数据存储起来
// 然后我想在用一个一维数组先全部初始化为0，再将蛋的品质附到数组里，出现一次该数组对应元素就加一，再来一个变量记录有多少
// 非0的数，这样就得到了蛋的种类。但是要考虑区间(t-86400,t]，也就是说如果超出一天就又要拿一个新数组重新记录。
func main() {
	var n int
	fmt.Scanf("%d", &n)
	var arr [30][30]int
	for i := 0; i < 30; i++ {
		for j := 0; j < 30; j++ {
			arr[i][j] = 0
		}
	}
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i][0])
		fmt.Scanf("%d", &arr[i][1])
		for j := 2; j < arr[i][1]+2; j++ {
			fmt.Scanf("%d", &arr[i][j])
		}
	}
}
