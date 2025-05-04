package main

// 使用pointer
func testPtr(num *int) {
	*num = 20
}

// 返回多個值
func getSumAndSub(n1 int, n2 int) (int, int) {
	return n1 - n2, n1 + n2
}

func main() {

}
