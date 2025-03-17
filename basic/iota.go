package basic

const (
	i = 1 << iota
	j = 3 << iota
	// 省略赋值表达式，沿用上一行的表达式形式
	k
	l
)

//func main() {
//    fmt.Println("i=",i)
//    fmt.Println("j=",j)
//    fmt.Println("k=",k)
//    fmt.Println("l=",l)
//}
