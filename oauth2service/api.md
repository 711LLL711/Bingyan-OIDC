# oauth2 api文档  
## 客户端注册
### 请求
```
POST /application
```
### 参数
| 参数名 | 类型 |
| --- | --- |
| name | string |
| redirect_uri | string |
| domain | string |

## 请求授权码
### 请求
```
GET /authorize
```
### 参数
| 参数名 | 类型 |
| --- | --- |
|response_type=code|string|
| client_id | string |
| redirect_uri | string |
| scope | string |

## 用户登录
### 请求
```
POST /auth/login
```
### 参数
| 参数名 | 类型 |
| --- | --- |
| userid | string |
| password | string |


## 申请token
### 请求
```
POST /token
```
### 参数
| 参数名 | 类型 |
| --- | --- |
| code | string |
| client_id | string |
| client_secret | string |

## 使用token访问用户api
### 请求
```
GET /api/user
```
### 参数
| 参数名 | 类型 |
| --- | --- |
| access_token | string |

