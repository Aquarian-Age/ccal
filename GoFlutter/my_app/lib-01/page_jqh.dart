import 'package:flutter/material.dart';

class PageJQHfw extends StatefulWidget {
  const PageJQHfw({super.key});

  @override
  State<StatefulWidget> createState() => _PateJQHState();
}

class _PateJQHState extends State<PageJQHfw> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('JQH'),
        titleTextStyle: const TextStyle(fontSize: 21),
        centerTitle: true, //标题居中显示
        backgroundColor: Colors.blue, //标题背景颜色
        toolbarHeight: 26, //标题高度
      ),
      body: Column(
        children: const [
          Text('data'),
        ],
      ),
    );
  }
}
