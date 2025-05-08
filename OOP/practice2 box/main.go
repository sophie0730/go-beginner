package main

import "fmt"

//創建一個Box sturct, 聲明立方體的長寬高，長寬高要從終端獲得
//聲明一個方法取得立方體的體積
//創建一個Box struct variable, 印出給訂尺寸的立方體體積

type Box struct {
	len    float64
	width  float64
	height float64
}

func (box *Box) getVolumn() float64 {
	return box.len * box.width * box.height
}

func main() {
	box := Box{
		len:    5.0,
		width:  6.0,
		height: 1.1,
	}

	res := box.getVolumn()
	fmt.Println(res)
}
