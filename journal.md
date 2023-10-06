# 日志
## 10.1、10.2
### 用户注册登录和修改信息的api   
- 使用viper读取配置文件
- 使用gorm连接数据库
- 使用gin框架
- 使用zap日志库

### 遇到的奇怪bug
在utils函数里写了一个验证密码正确性的函数，调用时会出现空指针使用错误，但是单独看这个函数好像并没有错误，debug了好久暂时没找到原因。   
~~~go
func PasswordVerify(hashedPwd string, plainPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	return err == nil
}
~~~

### TODO
头像功能、校验密码bug解决、邮件功能

### 总结
1. 用go有点生疏了，写了蛮多指针引用、空指针的错误，debug了好久，邮件的部分还没做，想先推推进度    
2. 第一次用viper,zap这些库，调试了一段时间，使用viper解析配置文件一直有问题，还是应该先看样例再开始写  


## 10.3
### Debug
1. zap日志指针初始化大小写写错了，导致出现空指针    
2. viper配置文件，不是config.go相对的路径，而是go程序运行位置相对的路径，之前写的是```viper.SetConfigName("config.json")```，应该是```viper.SetConfigName("config/config.json")```

### 记录
1. 昨天的登录和读取配置bug修好了，写了更新用户信息api,头像功能     
2. 继续了解OAuth2.0, 了解了流程，写了一个与OAuth2提供商进行交互的demo，对流程更熟悉了   
3. 找资料，找到了一个搭建OAuth2.0的仓库  


## 10.4
### 记录
1. 开始做OAuth2，定义数据类型、路由、服务器、数据库等
2. 想的是先把授权码模式基本逻辑写出来，再结合到用户管理api里 
3. TODO: 
- 申请OAuth2的用户id在哪个阶段传送？放在哪个结构体里比较好
- 继续完善路由

## 10.5
### 记录
1. 逻辑基本完成，但是还有很多细节没处理，比如scope,state,token的过期时间等
2. 测试已经完成的逻辑

### TODO
- scope,state,token的过期时间等
- 细节补充：redirect_url校验，error和error_description等的处理符合文档
- 是否可以利用response_type=token或者code来实现代码复用？
- 现在做的oauth2是一个相对独立的，便于测试，后续需要把它结合到之前写的用户管理api里