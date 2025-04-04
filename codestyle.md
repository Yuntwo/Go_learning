# Naming
- 总的来说首字母大小写决定了变量、类型、函数等的可见性，大写表示exported包外可见，小写表示unexported包内可见。
- 这种大小写决定可见性编译规则简化了如Java中的public、private等关键字。
## Package Names
- Use lowercase for package names.
## File Names
- Use lowercase for file names.
## Parameter Names
- Use first letter of the type name as the parameter name.
## Others
- Use `camelCase` for function names, variable names, and public properties.
## Initialisms/Acronyms
- Use lowercase or uppercase for all based on variable visibility.
- Keep the initialism's inherent case if it starts with the required case.
- Examples: `iOS` (private), `DDoS` (public), `ddos` (private).
- See [more examples](https://google.github.io/styleguide/go/decisions#initialisms)
# Loop
## 只有For循环，没有While循环，For循环可以实现While循环和限循环
- `for init; condition; post { }` 实现For循环
- `for condition { }` 实现While循环
- `for { }` 实现无限循环
# 括号
- 循环条件、if条件、case条件、select条件等等都不需要圆括号，这与Java等语言不同
# 类型
- 一般是名称在前，类型在后，如`var a int`

# References
[https://google.github.io/styleguide/go/index](https://google.github.io/styleguide/go/index)

[https://google.github.io/styleguide/go/decisions#naming](https://google.github.io/styleguide/go/decisions#naming)

# Testing
- 测试文件名以`_test.go`结尾，强制性的，否则不会被识别为测试文件
- 测试函数名不是强制性的，常见规范是`TestXxx`(单元测试)，`BenchmarkXxx`(性能测试)，`ExampleXxx`(使用示例)
- 单元测试通常指针对代码中最小的可测试单元(如函数或方法)进行的测试，目的是验证其功能是否符合预期。虽然理想情况下单元测试应独立于其他函数，但在实际中，可能会调用其他函数或依赖外部资源