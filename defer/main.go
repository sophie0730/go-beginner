package main

import "fmt"

/*
defer:
函數中經常需要創建資源，defer能確保函數執行完畢後，能即時釋放資源
*/

func sum(n1 int, n2 int) int {
	//當執行到defer, 暫時不執行，會把defer後面的語句放入另一個獨立的stack
	//當函數執行完畢後，會再從defer stack裡面，按照first in last out的方式把資料拿出stack
	defer fmt.Println("ok n1=", n1)
	defer fmt.Println("ok n2=", n2)

	//因為稍早已經把n1, n2 copy進defer stack, 所以後續的值變化只會在原本的stack寫入，defer的stack的值不會變
	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok res=", res)

	return res
}
func main() {
	res := sum(10, 20)
	fmt.Println("main: ", res)
}

/*
順序
ok res= 30
ok n2= 20
ok n1= 10
main:  30

要注意defer 是以函式為scope，函式執行完，就會馬上去看defer stack
*/
