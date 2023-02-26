import 'package:flutter/material.dart';

import './demo.dart';

void main() => runApp(MyApp());

class MyApp extends StatefulWidget {
  @override
  _MyAppState createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: const Text(
            'Demo',
            style: TextStyle(color: Colors.blue), // 标题字体颜色
          ),
          centerTitle: true, //标题居中显示
          backgroundColor: Colors.white, //标题背景颜色
          toolbarHeight: 24, //标题高度
          // toolbarTextStyle:
          //     TextStyle(fontSize: 16, color: Color.fromARGB(0, 97, 8, 212)),
        ),
        body: DemoPage(),
      ),
    );
  }
}
