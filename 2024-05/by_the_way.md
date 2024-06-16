# 不要把屎一样的代码留给编译器猜

```go
func main() {
	left := 2
	right := 10

	fmt.Println(left, right, (right-left)>>1, (right-left)>>1+left)
	// 2 10 4 6
}
```

```python
left = 2
right = 10

print(left, right, (right - left) >> 1, (right - left) >> 1 + left)
# 2 10 4 1
```

Python 的算数运算符优先级大于位移运算

当不确定优先级或者知道可能有歧义的时候，加上括号。不要把屎一样的代码给编译器猜。

# 设计对调用者友好型函数

- 不要让调用者传用不上的参数
- 使用空字段防止结构体无键字面量 https://twitter.com/func25/status/1770792643892519079
- 传参方式阳间一点，坏的例子参考同级目录的 `main.go` 文件

