# game-xjson #
[**中文**](./README.md) | [**English**](./README_EN.md)

本项目为由浊水楼台开发团队开源的，希望帮助游戏开发者快速将xlsx数据转为json数据的项目。

## 内容导引 ##
* [介绍](#介绍)
* [版本实现](#版本实现)
* [部署](#部署)
* [问题反馈](#问题反馈)

## 介绍 ##
#### game-xjson能做些什么？
* 帮助游戏开发者在可视化的环境下调整游戏资产数据
* 帮助非开发者快速将xlsx数据转为json数据
#### 目前为止，game-xjson可以提供以下支持：
* 自动检测字段类型
* 在xlsx中对字段进行注释
* 对开发者友好的json格式化输出
* 在生成的json结构中添加自增id
#### game-xjson支持的数据类型：
* 字符串类型：string/text
  * test
* 整数类型：int/integer
  * 314
* 浮点数类型：float/double/number
  * 3.14
* 布尔类型：bool/boolean
  * true/false
  * TRUE/FALSE
* JSON类型：json/object
  * {"name":"test"}
* 数组类型：array/list
  * [1,2,3]
  * ["1","2","3"]
  * [1,"2",3]
  * [{"name":"test1"},{"name":"test2"}]

## 版本实现 ##
#### 开发版
* main
#### 稳定版
* version 1.0.0

## 构建 ##
```bash
go build -ldflags "-s -w -X github.com/ZSLTChenXiYin/game-xjson/cmd.Version=1.0.0"
```

## 问题反馈 ##
* 陈汐胤会在每周五至周日查看 [Issues](https://github.com/ZSLTChenXiYin/game-xjson/issues)，还会不定期地在bilibili直播。
  * 陈汐胤的e-mail: imjfoy@163.com
  * 陈汐胤的bilibili UID: 352456302
