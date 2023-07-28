## WARNING: REMOTE HOST IDENTIFICATION HAS CHANGED!

@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

@    WARNING: REMOTE HOST IDENTIFICATION HAS CHANGED!     @

@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1690253807510-269163b9-bb2e-43a1-aa78-3d046ae5876e.png)

这个问题出现的原因是远程主机的密钥信息在您的本地计算机上发生了变化，可能是由于服务器更新或者配置更改导致的。由于密钥发生了变化，SSH 连接认证失败，为了确保安全，SSH 客户端会拒绝连接。

要解决这个问题，您可以采取以下步骤：

1. 首先，请确保您的网络环境是可信的，确保没有遭到中间人攻击。
2. 确认服务器管理员对服务器进行了更改并已更新密钥。如果您是服务器管理员，请确保正确地更新了服务器密钥。
3. 打开终端或命令提示符，进入到您的用户目录（例如：/Users/betta/）。
4. 编辑 "known_hosts" 文件，找到并删除其中对应远程主机的行。根据错误信息，该行应该是第三行，并包含了之前的 ECDSA 密钥信息。在 Linux/macOS 上，可以使用以下命令来编辑该文件：

```javascript
nano ~/.ssh/known_hosts
```

使用 arrow keys 来导航到有关 gitlab.ftsview.com 的那一行，然后按 "Ctrl + K" 来删除该行，最后按 "Ctrl + X" 保存并退出。如果您使用的是 Windows，请使用合适的文本编辑器（如 Notepad++）打开 "known_hosts" 文件并删除相关行。

1. 确保您正确地将新的主机密钥添加到 "known_hosts" 文件中，以便以后再次连接时不会出现此问题。重新连接到远程主机时，您将收到类似于以下消息：

```vbnet
The authenticity of host 'gitlab.ftsview.com (IP_ADDRESS)' can't be established.
ED25519 key fingerprint is SHA256:jst2EVsPmHO81FhvLS8vsKNltl91eRTsRGK1jJTyhZ0.
Are you sure you want to continue connecting (yes/no)?
```

输入 "yes" 并按下 Enter 键，这将把新的主机密钥添加到您的 "known_hosts" 文件中。

1. 重新尝试连接到远程仓库。

请注意，执行这些步骤之前，请确保您理解了正在做的操作，并且您确实信任远程主机，以免出现安全问题。如果您不确定如何处理此问题，请联系系统管理员或负责服务器的人员寻求帮助。