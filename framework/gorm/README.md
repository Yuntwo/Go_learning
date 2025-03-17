# Gorm ORM 框架
## 优势
- 简洁易用：通过链式调用简化了常见数据库操作。
- 功能丰富：支持复杂关系和事务。
- 高扩展性：支持插件开发和自定义扩展。
## 高级特性
### 预加载：通过 `Preload` 方法一次性加载关联数据。
### Association
Association 是 GORM 对关联关系操作的便捷抽象，不是强制的类型检查定义，让开发者可以通过 GORM 提供的 API 在高层次直接操作关联，而不需要手动编写复杂的 SQL；
否则需要手动管理所有关联操作。例如，Append、Replace、Delete 等操作，实际上在后台会生成相应的 SQL，直接作用于相关的表。
## 特点
- 通过在方法中传入变量指针指明将查询结果存在哪个变量中、删除哪种数据对应的记录等。
## 初始化配置
```text
struct &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info) // 打印所有 SQL 和调试信息，默认logger.Silent
}
```
### 日志级别
- `logger.Silent`：默认级别，只记录严重错误和必要信息（例如 record not found）。
- `logger.Error`：记录错误信息，不输出成功的 SQL。
- `logger.Warn`：记录警告信息和错误信息。
- `logger.Info`：记录所有 SQL 语句和调试信息（适合开发阶段）。

设置合适的日志级别有助于排查问题(开发环境)，减少日志的冗余量(尤其是生产环境)
## 反向工程
- 通过 `gormgen` 工具可以根据数据库表生成对应的 Gorm 模型/类型安全 API。
