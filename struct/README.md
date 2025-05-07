## Golang 如何實踐OOP概念

- Go沒有class, 使用struct來執行類似的概念
- 沒有繼承、constructor, destruct, this等概念
- 還是有 封裝、多型等概念，但實現的方法不同
    - 繼承透過匿名字段實現
- 透過interface來讓耦合性低

- struct可以比擬為class, 他只是定義object裡面要有哪些資訊
- struct variable比擬為instance, 是實作struct的實際變數

## Struct variable在記憶體的分配
假設有一個struct:
```go
type Cat struct {
	Name  string //大寫開頭 代表public
	Age   int
	Color string
}
```
會在記憶體分配三個空間，存放這三個變數
當cat1實作struct時
```go
var cat1 Cat
```
cat1實際上是直接包含了整個struct的值，所以他是值類型（因為他不是存放一個記憶體位址）

- struct的所有變數，在記憶體裡面是連續分布的

## Public/Private struct
透過struct名字來區分。
大寫開頭代表可以在別的package使用。
小寫開頭代表這是一個private struct，無法在其他package使用。