# 引用类型(Reference Types)(Q:概念不是很有感觉)
- 是指那些在赋值或传递时，传递的是对底层数据结构的引用，而不是复制整个数据的类型
- 包括
  - Slice(切片)，修改切片的元素会影响底层数组和其他引用该数组的切片
  - Map(映射)，修改映射的元素会影响其他引用该映射的引用
  - Channel(通道)
  - Function(函数)
  - Pointer(指针)，修改指针指向的值会影响所有引用该内存地址的变量，`这也是为什么参数传递值类型时都要用其指针的原因`
  - Interface(接口)
# 变量声明
- 无初始化的变量声明(会被默认初始化为零值，但是结构不能省略)，必须`var 变量名 类型`
- 有初始化的变量声明
  - 可以省略类型，会根据初始化值的类型自动推断类型
  - 可以省略var用`:=`代替，但是只能用于函数内部，不能用于包级变量

