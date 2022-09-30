import 'package:flutter/material.dart';

// ignore: must_be_immutable
class PageJQHfw extends StatefulWidget {
  dynamic map; // 第一页面传来的值
  PageJQHfw(this.map, {super.key});

  @override
  State<StatefulWidget> createState() => _PateJQHState();
}

class _PateJQHState extends State<PageJQHfw> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('JQH'), titleTextStyle: const TextStyle(fontSize: 21),
        centerTitle: true, //标题居中显示
        backgroundColor: Colors.blue, //标题背景颜色
        toolbarHeight: 26, //标题高度
      ),
      body: Column(
        children: [
          // OutlinedButton(
          //   onPressed: () {
          //     Navigator.pop(context);
          //   },
          //   child: const Text(
          //     '返回主页',
          //     style: TextStyle(fontSize: 21),
          //   ),
          // ),
          Text("data+${widget.map[1]}"),
        ],
      ),
    );
  }
}

// class PageJQH extends StatelessWidget {
//   dynamic map; //第一页面传来的值
//   PageJQH(this.map, {super.key});

//   // Widget build(BuildContext context) {
//   //   return Scaffold(
//   //     appBar: AppBar(
//   //       title: const Text('JQH'), titleTextStyle: const TextStyle(fontSize: 21),
//   //       centerTitle: true, //标题居中显示
//   //       backgroundColor: Colors.blue, //标题背景颜色
//   //       toolbarHeight: 26, //标题高度
//   //     ),
//   //     body: Column(
//   //       children: [
//   //         OutlinedButton(
//   //           onPressed: () {
//   //             Navigator.pop(context);
//   //           },
//   //           child: const Text(
//   //             '返回主页',
//   //             style: TextStyle(fontSize: 21),
//   //           ),
//   //         ),
//   //         Text('data'),
//   //       ],
//   //     ),
//   //   );
//   // }
// }
