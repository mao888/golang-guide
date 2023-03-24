[OAuth](http://en.wikipedia.org/wiki/OAuth)是一个关于授权（authorization）的开放网络标准，在全世界得到广泛应用，目前的版本是2.0版。Marketing-API的授权流程基于OAuth2.0进行。

### 应用场景

相较于传统的账号密码登陆进行权限控制，OAuth2.0更能满足第三方操作的场景，避免账号密码暴露而造成风险，也能通过简单的授权操作使得第三方应用能够读取或操作账号下的相关信息。常见的第三方登录就是基于OAuth2.0进行的。

### 涉及的角色

1. **授权服务端**
2. Marketing-API中开放平台的角色
3. 在开放平台中会处理用户的授权记录。
4. **资源服务端**
5. Marketing-API中API接口的角色
6. API提供的所有接口就是用户的资源，应用获取到访问的凭证后即可调用API接口，访问用户的资源。
7. **用户**
8. Marketing-API中涉及的各类广告账号，包括代理商、管家、广告主，这些角色的登陆者都是用户
9. 用户在授权服务端（开放平台，或其他头条系广告账号平台登陆后跳转至开放平台）登录，并选择相应的账号进行授权，开放平台会记录相应的授权关系，生成相应的授权码。
10. **接入的第三方应用**
11. Marketing-API中的应用（开发者申请的APP_ID）
12. 用户通过给该第三方应用进行授权，授权后第三方应用可进行资源访问凭证的获取，而后进行资源的访问。

### 授权模式——授权码模式

OAuth2.0有许多模式：简易模式、授权码模式、密码模式，其中授权码模式是最常用的一种模式，也是Marketing-API采用的模式，也是功能最完整，流程最严密的模式。

#### 相关术语

| 字段                        | 含义                                                         | Marketing-API图示                                            |
| --------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| app_id                      | 应用id： 应用申请审核通过后生成，API应用的唯一标识，可通过“[APPID管理](https://ad.oceanengine.com/openapi/appid/list.html)”界面进行获取。 | ![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1679642199366-1a5f53cd-fc5f-4a69-98a0-a3cb9ec77649.png) |
| app_secret/secret           | 应用密钥： 应用申请审核通过后生成，再进行授权数据相关操作时需要给定该密钥，可通过“APPID管理-应用详情”进行获取。 | ![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1679642199306-a4a6a648-3c1a-4c3c-8eec-1c2125be2831.png) |
| 回调地址                    | 用户授权完成后，会将auth_code等信息通过该接口进行返回。 回调地址可在“APPID管理-应用详情”界面进行修改。 | ![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1679642199299-8af4a43e-8feb-4df9-bc2a-0112d56ffc21.png) |
| auth_code                   | 授权码： 用户授权完成后会通过拼接入回调地址返回，该授权码可作为换取Access-Token的凭证，通过“[获取Access Token](https://ad.oceanengine.com/openapi/doc/index.html?id=1696710505596940)”接口使用。 | ![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1679642199573-9a8c58c5-3c43-4a84-8d3a-919aa37e2c59.png) |
| access_token/Access-Token   | 资源访问令牌： 访问资源的凭证，会验证令牌的有效性，一个令牌的有效期为1天，过期将不可继续访问。 可通过“[获取Access Token](https://ad.oceanengine.com/openapi/doc/index.html?id=1696710505596940)”接口与“[刷新Refresh Token](https://ad.oceanengine.com/openapi/doc/index.html?id=1696710506097679)”接口获取 |                                                              |
| refresh_token/Refresh-Token | 刷新access token令牌： 因access_token存在有效期，为防止access_token过期而造成不可用，提供了Refresh-Token用于token的刷新，只要按规律刷新，可保证资源的长期可访问。refresh_token有效期为30天，一旦使用即失效。 可通过“[获取Access Token](https://ad.oceanengine.com/openapi/doc/index.html?id=1696710505596940)”接口与“[刷新Refresh Token](https://ad.oceanengine.com/openapi/doc/index.html?id=1696710506097679)”接口获取 |                                                              |
| permissions                 | 权限 授权成功后并不是可以操作所有的资源，而是在授权时设置的权限，授权成功后的access_token仅能访问权限范围内的资源信息。 | ![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1679642199505-f6b5cdd9-14e4-4cc5-9505-a9c43582ec29.png) |

#### 流程

1. 用户给指定的第三方应用app_id授权
2. 授权服务端生成auth_code，并回调给第三方应用
3. 第三方应用通过auth_code、app_id、app_secret请求授权服务端，授权服务器在验证完auth_code是否失效以及接入的用户信息是否有效（通过传递的app_id和app_secret信息和服务端已经保存的用户信息进行匹配）之后，授权服务端生成Access Token和Refresh Token并返回给客户端。
4. 第三方应用通过得到的Access Token请求资源服务应用，获取需要的且在申请的Access Token权限范围内的资源信息。

#### 基本流程图

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1679642199963-9de7a41c-6d4e-49c8-b436-a4df41498ffd.png)

#### 序列图

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1679642200005-e613d26d-e8e9-4db3-a4f5-d9782ce0ab33.png)

#### token刷新

为了资源的安全，减小Access-Token泄漏而导致的数据安全问题的风险，所以Access-Token并不是长期有效的，设置了过期时间，所以需要通过刷新获得新的access_token。只要保持正常的刷新，在用户不取消授权的情况下，将能一直访问用户的资源。

- 如下示例代码所示，应用可以设置定时任务进行token的刷新，从而保证资源的长期访问
- 定时任务，由于access_token仅一天有效期，故而至少一天刷新一次，才可保证token有效访问。若担心临界时间点，可设置一天刷新两次。
- 获取到的access_token与refresh_token均需要进行保存，刷新时取最近的access_token进行刷新操作，使用后即失效。
- refresh_token有效期30天，过期后需重新授权（使用后会生成新的refresh_token，新的token会重置有效期，依然有30天有效期），若正常刷新不会出现refresh_token过期的情况。

```go
# -*- coding: utf-8 -*-
import requests
from apscheduler.schedulers.blocking import BlockingScheduler

APP_ID=0
APP_SECRET="xxx"

# 刷新access-token
def job():
    # 从存储中获取最新的refresh_token
    refresh_token = get_refresh_token_from_db()

    open_api_url_prefix = "https://ad.oceanengine.com/open_api/"
    uri = "oauth2/refresh_token/"
    url = open_api_url_prefix + uri
    data = {
        "app_id": APP_ID,
        "secret": APP_SECRET,
        "grant_type": refresh_token,
        "refresh_token": "xxx",
    }
    rsp = requests.post(url, json=data)
    rsp_data = rsp.json()

    # 将刷新后的token进行存储，包括access_token & refresh_token
    save_access_token(rsp_data)


if __name__ == "__main__":
   scheduler = BlockingScheduler()
    # 设置定时任务，每天晚上8点进行刷新
    scheduler.add_job(job, 'cron', day_of_week='0-6', hour=20, minute=00)
    scheduler.start()
```