# calendar

- [calendar_todo](https://github.com/taihangg/calendar_todo)修改版

日历 农历算法源自许剑伟先生的寿星万年历。
万年历的代码移植自HongchenMeng先生的c#移植项目，地址：https://github.com/HongchenMeng/SharpSxwnl
感谢许剑伟先生！感谢HongchenMeng先生!


## todo部分删除不用

- calendar部分剥离出来单独使用

- 节日内容字数过多将会滚动显示。


# Flutter 

```bash
Flutter 3.3.9 • channel stable • ssh://git@github.com/flutter/flutter.git
Framework • revision b8f7f1f986 (2 周前) • 2022-11-23 06:43:51 +0900
Engine • revision 8f2221fbef
Tools • Dart 2.18.5 • DevTools 2.15.0

flutter build apk --obfuscate --split-debug-info=build --release --build-name=1.1.9 --build-number=8

flutter build apk --debug --obfuscate --split-debug-info=build 

flutter build linux --release

flutter pub run msix:create

```
## 添加内容

1. 日干支
2. 日建除
3. 二十八宿(日禽)
4. 宗教节日

