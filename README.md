# lara-go
因为是从laravel转到gin的 :) .  所以对gin-admin进行了一部分优化，增加了一些自定义配置项，只支持前后端分离。仅做学习使用，谨慎用于生产环境。

- 自定义ShouldBindJSON的错误内容，自定义位置为internal/app/validator/rule.go
  - 注意要使用validator代替binding
    - 自己新增校验的tag和其对应的方法,注意需要strings.ToLower(方法名) = e.Tag()。例如：
        
      ```
      func (*Errors) Required(field, param string) string {
          return fmt.Sprintf("%s字段不能为空", field)
      }
      ```
    - 具体支持的规则需要参考[validator.v10](https://github.com/go-playground/validator)的文档
- 把路由配置进行了抽离，路由配置地址为internal/app/router
  - 目录下文件名代表了路由的前缀，例如：router_api.go代表/api的路由。会自动调用get{Api}Routes方法
  - 相关中间件的位置是internal/app/middleware
- 新增了计划任务相关
  - 自定义Job类型(实现Run方法即可)。位置是internal/app/cron
  - 创建完Job需要在internal/app/cron/cron.go中注册，具体请看示例。
- 其他功能基本都是延用[gin-admin](https://github.com/LyricTian/gin-admin)，具体实现请查看文档