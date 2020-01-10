# 推送网关对接文档

版本：`v2`

以下推送网关简称 `网关`，接入设备简称 `终端`

支持接入的终端：**WebSocket**、**IOS**、**Android**

* 网关域名：
  * 测试域名：openservice.dev.jsjinfo.cn
  * 生产域名：openservice.jsjinfo.cn

* 服务名：
  * WebSocket连接服务：**notification-comet**
  * 网关操作服务：**notification-gateway**

## 鉴权

在请求的Header中增加`Authorization: md5(md5(body+salt))`

## 1. WebSocket

WebSocket接入端点：wss://域名/notification-comet/v1/ws/{targetType}/{jsjUniqueId}

| 变量        | 说明                                                         |
| :---------- | ------------------------------------------------------------ |
| jsjUniqueId | 金色世纪会员编号或员工唯一编号                               |
| targetType  | 目标类型；<br/>金色世纪APP=100；<br/>空铁管家APP=200；<br/>员工端APP=300；<br/>空铁管家微信小程序=400；<br/>空铁1号600 |

### 1.1 初始化

`wss://域名/服务名/v1/ws/{targetType}/{jsjUniqueId}`

### 1.2 心跳

我们约定每2秒通过ws协议向推送网关发送一个`+`字符，推送网关到后，返回一个`-`字符

### 1.3 断开

当终端置于`后台`或者`退出`后，终端应当主动断开WebSocket连接



## 2. 终端注册、终端注销

注册接口地址：[POST] https://域名/notification-gateway/v1/terminal/register

注销接口地址：[POST] https://域名/notification-gateway/v1/terminal/unRegister

**请求参数：**

| 变量         | 类型   | 备注                                                         |
| ------------ | ------ | ------------------------------------------------------------ |
| JsjUniqueId  | int    | 推送的目标，用户编号、员工编号                               |
| PushToken    | string | 目标的唯一识别码                                             |
| PlatformType | int    | 平台类型：<br />10：IOS<br />20：Android<br />30：MiniProgram<br />40：WebSocket<br />50：DingDing |
| TargetType   | int    | 应用类型：<br />100：金色世纪<br />200：空铁管家<br />300：员工端<br />400：微信小程序<br />500：钉钉<br />600：空铁1号 |



## 3. 推送消息

接口地址：[POST] https://域名/notification-gateway/v1/notification/send

**请求参数：**

| 变量        | 类型     | 备注                                                         |
| ----------- | -------- | ------------------------------------------------------------ |
| JsjUniqueId | int      | 推送的目标，用户编号、员工编号                               |
| TargetType  | int      | 推送的应用<br />100：金色世纪<br />200：空铁管家<br />300：员工端<br />400：微信小程序<br />500：钉钉<br />600：空铁1号 |
| Title       | string   | 标题                                                         |
| Message     | string   | 摘要                                                         |
| GroupName   | string   | 自定义消息分组名称，请一类消息保持使用同一个名称             |
| Route       | string   | 跳转路由：<br />1、微信小程序 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。<br />2、APP中，点击弹出的系统通知栏，跳转页面的路由 |
| Data        | []string | 推送附加的自定义节点数据；<br />字典结构，请勿出现重复的key<br />不支持嵌套JSON，仅支持简单的键值数据结构 |



## 4. 获取推送历史记录

接口地址：[POST] https://域名/notification-gateway/v1/msg

### 4.1 请求

#### 4.1.1 **请求参数**

| 变量        | 类型   | 备注                                      |
| ----------- | ------ | ----------------------------------------- |
| JsjUniqueId | int    | 目标的唯一用户识别码，和PushToken二选一   |
| PushToken   | string | 目标的唯一设备识别码，和JsjUniqueId二选一 |
| Offset      | int    | 指定开始返回记录前要跳过的记录数 (最小1)  |
| Limit       | int    | 检索出的最大记录数 (最大20)               |

#### 4.1.2 **请求数据样例**

```json
{
    "JsjUniqueId": 0,
    "PushToken": "3d77bb9867b1d955110f026729ed7f647bb51d5beb12c59aeb648f8489f6741a",
    "Offset": 1,
    "Limit": 10
}
```



### 4.2 响应

#### 4.2.1 响应参数

| 变量        | 类型   | 备注             |
| ----------- | ------ | ---------------- |
| Id          | int    | 编号             |
| PushToken   | string | 目标的唯一识别码 |
| Title   | string | 标题 |
| Message   | string | 简要描述 |
| Router | string | 路由键         |
| DataContent | string | 消息内容         |
| GroupName   | string | 分组名           |
| CreateTime  | time   | 发送时间         |

#### 4.2.2 **响应数据样例**

```json
{
    "baseHead": {
        "code": "0000",
        "message": "success"
    },
    "Msg": [
        {
            "Id": 2,
            "JsjUniqueId": 2059797,
            "PushToken": "2059797",
            "Title": "分享消息",
            "Message": "扫描成功",
            "DataContent": "{\"jsjUniqueId\":\"2059797\",\"linkUrl\":\"\",\"message\":\"扫描成功\",\"opType\":\"1\",\"orderNumber\":\"0\",\"title\":\"分享消息\"}",
            "PlatformTypeId": 0,
            "PlatformTypeName": "",
            "TargetTypeId": 400,
            "TargetTypeName": "空铁管家微信小程序",
            "GroupName": "测试组",
            "CreateTime": "2020-01-10 15:26:40"
        }
    ]
}
```

