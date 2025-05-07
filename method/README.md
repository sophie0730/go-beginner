## Method

Golang也有method的概念，只要是自定義的type都可以有method，不僅限於method
```go
type Person struct {
	Name string
}

func (person Person) test() {
	person.Name = "jack" //實際上是調用一個副本，所以跟調用這個方法的struct本人沒有關係
	fmt.Println("test()", person.Name)
}

func main() {
	var p Person
	p.Name = "sophie"
	p.test() //調用方法
	fmt.Println("main() p.Name= ", p.Name)
}
```

- test就是Person struct的method
- test是和Person類型綁定的
- test方法只能透過Person類型的變數調用，而不能直接調用，也不能使用其他變數來調用
- person表示哪個Person變數調用，這個person就是他的副本，跟函數的參數傳遞很像

## 參數傳遞
在method中，參數傳遞是pass by value, 當呼叫method時，他會為了傳入的參數再分派獨立的記憶體空間
同樣地，method所屬的struct也會另外copy一份，也會有一份自己獨立的記憶體空間
```go
func (person Person) calculate2(n int) { //n, person會再獨立被分配一塊記憶體空間
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}

	fmt.Println(person.Name, "calculate result is ", result)
}
```

## 在method中修改struct的值
- 因為struct是pass by value, 所以一班情況下，method針對struct的改動不影響原始struct
- 如果要修改，傳入的struct必須是指標的形式，讓method使用實際struct的記憶體位址
- see exercise/