# game-xjson #
[**中文**](./README.md) | [**English**](./README_EN.md)

This project is an open-source initiative developed by the Zhuoshui Loutai development team, aimed at helping game developers quickly convert xlsx data into json data.

## Table of Contents ##
* [Introduction](#Introduction)
* [Version Implementation](#Version-Implementation)
* [Deployment](#Deployment)
* [Issue Reporting](#Issue Reporting)

## Introduction ##
#### What can game-xjson do?
* Assist game developers in adjusting game asset data in a visual environment.
* Help non-developers quickly convert xlsx data into json data.
#### So far, game-xjson offers the following support:
* Automatic detection of field types.
* Annotating fields in xlsx.
* Developer-friendly json formatted output.
* Adding auto-incrementing IDs to the generated json structure.
#### Data types supported by game-xjson:
* String type：string/text
  * test
* Integer type：int/integer
  * 314
* Floating-point type：float/double/number
  * 3.14
* Boolean type：bool/boolean
  * true/false
  * TRUE/FALSE
* JSON type：json/object
  * {"name":"test"}
* Array type：array/list
  * [1,2,3]
  * ["1","2","3"]
  * [1,"2",3]
  * [{"name":"test1"},{"name":"test2"}]

## Version Implementation ##
#### Development Version
* main
#### Stable Version
* version 1.0.0

## Deployment ##
```bash
go build -ldflags "-s -w -X github.com/ZSLTChenXiYin/game-xjson/cmd.Version=1.0.0"
```

## Issue Reporting ##
* Chen Xiyin reviews [Issues](https://github.com/ZSLTChenXiYin/game-xjson/issues) every Friday to Sunday and occasionally streams on Bilibili.
  * Chen Xiyin's email: imjfoy@163.com
  * Chen Xiyin's Bilibili UID: 352456302
