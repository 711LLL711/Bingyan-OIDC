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
