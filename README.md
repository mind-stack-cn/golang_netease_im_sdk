# Golang 网易云信SDK
 
[网易云信开发手册](http://dev.netease.im/docs)

## 实现功能
* 创建云信ID
* 云信ID更新
* 更新并获取新token
* 封禁云信ID
* 解禁云信ID
* 更新用户名片
* 获取用户名片
* TODO...

## 使用说明
````
    ...
    util.Init("#YOUR APP KEY", "#YOUR APP SEC KEY")// 初始化
    res, err := util.CreateAccid("zhangsan", "", "", "", "")
    ...
````

