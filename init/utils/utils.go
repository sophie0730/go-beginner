package utils

var Age int
var Name string

// 為了在main.go引入package時能直接使用初始化過後的全局變數，需要使用init
func init() {
	Age = 10
	Name = "sophie"
}
