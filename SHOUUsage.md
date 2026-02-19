假定阅读者具有基本的编程经验和项目构建基础

## 1. 前置准备
1. 使用学号和密码登录你的`Easy Connect VPN`
2. 确保你有`Linux`或者`wsl`、`termux`等任意一个`Linux`系统
3. 确保自己的系统上安装有`go`编译器
4. 欢迎使用`Windows`的同学提供纯`Windows`版本的编译过程

> 注意: 如果你是`termux`平台
> 那么`apt`所安装的`golang`应该是`1.25.6`版本
> 此时你需要修改[./go.mod](./go.mod)文件
> 将开头的`go 1.25.7`替换为`go 1.25.6`

## 2. 编译项目
执行以下命令: 
```shell
git clone https://github.com/starwindv/picoclaw-shou
cd picoclaw-shou

# 以下为可选项, 任选其一即可
make build     # 为当前 Linux 平台构建
make build-win # 为同 CPU 架构的 Windows 构建
make build-all # 为所有平台进行构建

# 原项目里只有 build 和 build-all
# 这个 build-win 是我新增的
# 你也可以按照同样的格式自己加一个
```

这样你就拥有了一个适合自己平台的`picoclaw-shou`  
将它放到合适的位置, 并加入环境变量里

`wsl`同学建议使用`build-win`命令得到`Windows`版本的可执行文件  
因为`Easy Connect`走不了`TUN`模式  
`wsl`里的原生程序走不了学校流量

但是！你可以在`pwsh`里`ssh`连接`wsl`, 在里面调用`tmux`,   
内部使用`Windows`的`pwsh`,   
然后你不仅可以实现程序后台保活, 并且保留全部的日志, 而且可以走学校流量

## 3. 初始化配置
执行以下命令进行初始化: 
```shell
picoclaw-{platform}-{architecture} onboard
```

接下来导航到配置目录: 
- Linux
```shell
cd ~/.picoclaw/
```
- Windows PowerShell
```shell
cd $env:USERPROFILE/.picoclaw
```
- Windows CMD
```shell
cd /D %USERPROFILE%/.picoclaw
```

使用你喜欢的文本编辑器打开当前目录下的`config.json`, 找到如下字段: 
```json
{
    "shou": {
        "student_id": "",
        "password": "",
        "api_base": ""
    }
}
```

- `api_base`: 填`https://chat.shou.edu.cn`
- `student_id`: 填你的学号
- `password`: 填你的密码, 可以去`api_base`的链接里登录一下看看你的密码是多少, 好像是和`urp`一致的  
  (不过学校的一堆密码不互通也不是少见的事情了, 多试试总会想起来的)

## 4. 申请开放应用
之后便按照[原项目教程](https://github.com/sipeed/picoclaw/README.zh.md)申请一个开放应用, 比如`QQ`、`钉钉`  
我不建议你直接用`telegram`或者`discord`  
因为学校 VPN 和出海 VPN 不兼容, 会掉 AI 服务

## 5. 启动网关
上述配置完成后, 你就可以使用`picoclaw`的子命令`gateway`运行网关, 进行助手的使用了: 
```shell
picoclaw gateway
```

## 6. Windows 后台运行脚本 (补充) 
如果你是Windows用户且不需要任何日志, 那么可以把以下脚本保存到一个.ps1脚本里: 
```shell
<#
.SYNOPSIS
/path/to/thisScript.ps1 (background exec) - 在后台隐藏窗口执行指定命令

.DESCRIPTION
通过 Start-Process 以隐藏窗口的方式执行指定命令及参数
用法: /path/to/thisScript.ps1 command arg1 arg2 ... argN
#>

if ($args.Count -eq 0) {
    Write-Host "用法: /path/to/thisScript.ps1 command [arg1] [arg2] ... [argN]"
    Write-Host "示例: /path/to/thisScript.ps1 picoclaw gateway"
    exit 1
}

# 第一个参数作为要执行的程序路径
$command = $args[0]

# 剩余参数作为命令参数列表
$arguments = if ($args.Count -gt 1) { $args[1..($args.Count - 1)] } else { @() }

# 以隐藏窗口方式启动进程
Start-Process -FilePath $command -ArgumentList $arguments -WindowStyle Hidden
```

之后使用它来启动`picoclaw gateway`

## 7. 注意事项
最后, 决定服务在线时长的也就只有`Easy Connect`的在线时长(它会自动掉线)和电脑网络状态了

其它关于`picoclaw`本身的教程, 请参见[原项目教程](https://github.com/sipeed/picoclaw/README.zh.md)
