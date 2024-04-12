# 抖音开放平台RPC服务

## 抖音授权流程

1. 第三方发起抖音授权登录请求，抖音用户允许授权第三方应用后（扫码确认或账号密码授权登录），确认通过后，会重定向到第三方网站（回调接口）。并且附带授权临时票据（code）

> 说明：抖音开放平台现已开放WEB、Android、iOS和JS四种渠道的授权方式能力，您可以根据需要自行选择符合您需求的授权方法四种方式的区别详见“应用场景”，四种接入方式，详见[各渠道接入说明](https://developer.open-douyin.com/docs/resource/zh-CN/dop/ability/user-authorization/solution)

2. 第三方通过code参数，以及ClientKey和ClientSecret等参数，通过API换取access_token

> 详细操作方法：[授权令牌接入说明](https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/account-permission/get-access-token)

3. 通过access_token进行接口调用，获取用户基本信息及其他操作等。可以根据您的需求，选择OpenAPi的接口接入。

> [抖音开放平台OpenApi接口文档地址](https://developer.open-douyin.com/docs/resource/zh-CN/dop/develop/openapi/list)

## 已实现的功能

- **JSSDK授权**：
> 1. 获取client_key和签名接口

