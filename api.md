# 用户管理api文档

## 统一响应格式
|字段|类型|说明|
|:----:|:----:|:----:|
|success|bool|请求是否成功|
|hint|string|提示信息|
|data|object|响应数据|

## 用户注册
~~~
POST /registration
~~~
### 请求参数
| 参数名 | 类型 | 说明 |
| :----: | :----: | :----: |
| username | string | 用户名 |
| password | string | 密码 |
| email | string | 邮箱 |
### 响应
1. 注册成功-->提示发送邮件
2. 注册失败-->hint传输错误信息

## 用户登录
~~~
POST /login
~~~
### 请求参数
| 参数名 | 类型 | 说明 |
| :----: | :----: | :----: |
|email | string | 邮箱 |
|password | string | 密码 |

### 响应
登录成功-->返回用户信息id,username,avatar
## 用户更新
~~~
POST /update
~~~
### 请求参数
| 参数名 | 类型 | 说明 |
| :----: | :----: | :----: |
|username | string | 用户名 |
|bio | string | 简介 |
|avatar | file | 头像 |

用户id:从jwt中获取