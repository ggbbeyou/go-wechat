# Golang Wechat SDK

用 `golang` 写的一个用于微信小程序支付以及小程序后端接口调用的`sdk`，现在只实现了部分常用接口，后面有时间再加吧

## 接口预览

### 支付部分接口

- 统一下单接口（`SendUnifiedOrder`）

- 订单查询（`QueryOrder`）

- 关闭订单（`CloseOrder`）

### 小程序功能部分接口

- 获取`AccessToken`的接口（`QueryAccessToken`）

- `code`换`session`接口（`QuerySessionByCode`）

- 发送模板消息接口（`SendTemplateMessage`）

- 无限获取小程序码接口（`GetWXACodeUnlimit`）

### 用于微信接口调用的工具方法

- 生成签名的方法（`SignWechat`）

- 生成小程序调用微信支付的预支付数据方法（`GenPrepayData`）

- 解析微信回调接口数据的方法（`ParseNotifyReq`）

- 生成微信回调接口响应数据的方法（`GenNotifyResp`）

- 校验签名的方法（`VerifyNotifySign`）

## 安装

```sh
> go get github.com/lujin123/go-wechat
```

## 使用

### 初始化

```go
func init() {
    mchid := ""
    apikey := ""
    appid := ""
    appsecret := ""
    notifyUrl := ""
    auth := ""
    wechat = NewWechat(appid, appsecret, mchid, notifyUrl, apikey, auth)
}
```

先把微信需要的一些数据准备好，然后生成一个对象即可，如果这个参数调用的接口不需要的话，是可以不传的

`auth` 这个参数是在调用小程序功能接口的时候需要添加在`header`中

### 统一下单

```go
order := &UnifiedOrder{
    //参数自己填上
}
resp, err := wechat.SendUnifiedOrder(order)
```

其中`UnifiedOrder`是一个对象，这个库里面自定义的，创建好了参数丢进去即可，返回的`resp`也是一个自定义的对象实例

### 订单查询

```go
resp, err := wechat.QueryOrder("orderNo", "transactionId")
...
```

### 关闭订单

```go
resp, err := wechat.CloseOrder("orderNo")
...
```

### 小程序获取 AccessToken

_小程序的接口有些需要 token 的，需要先调用`AccessToken`接口获取之后赋值为`wechat`对象的`AccessToken`属性，或者先准备好`AccessToken`直接赋值也行_

_这个接口每次调用都会刷新`token`，导致原来老的`token`失效，所以需要注意，不要随便调用，具体的看[这里](https://developers.weixin.qq.com/miniprogram/dev/api-backend/getAccessToken.html)_

```go
resp, err := wechat.QueryAccessToken()
...
```

### 小程序 code 换 session

```go
resp, err := wechat.QuerySessionByCode("jsCode")
...
```

### 小程序发送模板消息

```go
// 消息内容的填充
data := map[string]interface{}{
    "keyword1": map[string]interface{}{
        "value": "title",
    },
    "keyword2": map[string]interface{}{
        "value": "name",
    },
}
req := SendTemplateReq{
    ToUser:          "发送对象openid",
    TemplateId:      "消息模板ID，需要提前配置好",
    Page:            "着陆页",
    FormId:          "提前手机的用户FormId",
    EmphasisKeyword: "加粗关键字",
    Data:            data,
}
resp, err := wechat.SendTemplateMessage(req)
...
```

### 获取无限小程序码

```go
req := WXACodeUnlimitReq{
    Scene:     "",
    Page:      "",
    Width:     430,
    AutoColor: false,
    LineColor: map[string]int64{},
    IsHyaline: false,
}
resp, err := wechat.GetWXACodeUnlimit(req)
...
```

### 工具方法就不列了，直接源码

## 最后

这是项目中需要用到，所以归总了下，方便其他的项目调用，现在直接用这个做个服务给外面调用，尤其是关于`token`的过期问题，服务需要保证`token`有效，其他的服务调用即可，否则每次都要关心这个接口调用是否因为`token`失效原因失败了，再获取`token`再调用，而且还要做好锁的问题，否则多个线程调用会导致老的`token`又失效的情况，灰常麻烦
