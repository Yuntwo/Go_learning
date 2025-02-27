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
- 循环条件、case条件、select条件等等都不需要括号
# 类型
- 一般是名称在前，类型在后，如`var a int`

# References
[https://google.github.io/styleguide/go/index](https://google.github.io/styleguide/go/index)

[https://google.github.io/styleguide/go/decisions#naming](https://google.github.io/styleguide/go/decisions#naming)

