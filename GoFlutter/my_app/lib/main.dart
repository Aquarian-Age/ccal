import 'dart:core';

import 'package:flutter/material.dart';
import 'package:my_app/page_info.dart';
import 'package:my_app/page_jqh.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: const MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key});

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: const Text(
            "my_app",
            style: TextStyle(color: Colors.blue),
          ),
          centerTitle: true, //标题居中显示
          //backgroundColor: Colors.white, //标题背景颜色
          toolbarHeight: 26, //标题高度
        ),
        body: Center(
          // textDirection: TextDirection.rtl,
          // crossAxisAlignment: CrossAxisAlignment.center,
          child: Column(
            children: [
              Container(
                alignment: Alignment.center,
                width: 100,
                height: 60,
                margin: const EdgeInsets.fromLTRB(0, 20, 0, 0),
                decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(10),
                    color: Colors.blue),
                child: TextButton(
                  onPressed: () {
                    Navigator.push(
                        context,
                        MaterialPageRoute(
                            builder: (context) => const PageInfo()));
                  },
                  child: const Text(
                    'INFO',
                    style: TextStyle(color: Colors.white, fontSize: 26),
                  ),
                ),
              ),
              Container(
                alignment: Alignment.center,
                width: 100,
                height: 60,
                margin: const EdgeInsets.fromLTRB(0, 20, 0, 0),
                decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(10),
                    color: Colors.blue),
                child: TextButton(
                  onPressed: () {
                    Navigator.push(
                        context,
                        MaterialPageRoute(
                            builder: ((context) => const PageJQHfw())));
                  },
                  child: const Text(
                    'JQH',
                    style: TextStyle(color: Colors.white, fontSize: 26),
                  ),
                ),
              ),
              // OutlinedButton(
              //   onPressed: () {
              //     Navigator.push(context,
              //         MaterialPageRoute(builder: (context) => const PageInfo()));
              //   },
              //   child: const Text('INFO'),
              // ),

              // const SizedBox(height: 16), // 上下组件间距的高度
              // OutlinedButton(
              //   onPressed: () {
              //     Navigator.push(
              //         context,
              //         MaterialPageRoute(
              //             builder: ((context) => const PageJQHfw())));
              //   },
              //   child: const Text('JQH'),
              // ),
            ],
          ),
        ),
      ),
    );
  }
}
