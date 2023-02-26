import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter_my_picker_null_safety/common/date.dart';
import 'package:flutter_my_picker_null_safety/flutter_my_picker.dart';
import 'package:http/http.dart' as http;

// https://docs.flutter.dev/cookbook/networking/send-data
// 客户端输入 以POST json格式 传送给服务端
Future<InfoData> createData(String data) async {
  const url = "http://localhost:6714";
  final respones = await http.post(
    Uri.parse(url),
    headers: <String, String>{
      'Content-Type': 'application/json; charset=UTF-8',
    },
    body: jsonEncode(<String, String>{
      'times': data,
    }),
  );
  //服务器返回来的数据
  if (respones.statusCode == 200) {
    return InfoData.fromJson(jsonDecode(respones.body));
  } else {
    throw Exception('Failed to create infoData');
  }
}

class InfoData {
  String times;
  String info;
  InfoData({required this.times, required this.info});

  factory InfoData.fromJson(Map<String, dynamic> json) {
    return InfoData(
      times: json['times'],
      info: json['info'],
    );
  }
}

class PageInfo extends StatefulWidget {
  const PageInfo({super.key});

  @override
  State<PageInfo> createState() => _PageInfoState();
}

class _PageInfoState extends State<PageInfo> {
  //时间选择器组件
  //DateTime time = DateTime.now();
  String dateStr = '';
  void _datePicker(context) {
    //  print("context Type: " + context.runtimeType.toString());// Type: StatefulElement
    MyPicker.showDateTimePicker(
      context: context,
      title: const Text("myPicker"),
      onConfirm: (time) {
        setState(() {
          dateStr = MyDate.format('yyyy-MM-dd HH:mm:ss', time);
        });
      },
    );
  }

//传送时间到服务端
  late InfoData infoData;
  String infoStr = '';
  void getInfo() async {
    const url = "http://localhost:6714";
    final respones = await http.post(
      Uri.parse("$url/info"),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: jsonEncode(<String, String>{
        'times': dateStr, //传送时间到服务端
      }),
    );

    if (respones.statusCode == 200) {
      Map<String, dynamic> clid = jsonDecode(respones.body);
      //  print("clid Type: " + clid.runtimeType.toString());//_InternalLinkedHashMap<String, dynamic>
      //infoData = InfoData.fromJson(jsonDecode(respones.body));
      setState(() {
        infoStr = clid['info'];
        // infoStr = infoData.info;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      // appBar: AppBar(
      //   title: const Text("Info"), //style: TextStyle(color: Colors.blue)),
      //   titleTextStyle: const TextStyle(fontSize: 21),
      //   centerTitle: true, //标题居中显示
      //   //backgroundColor: Colors.blue, //标题背景颜色
      //   toolbarHeight: 40, //标题高度
      // ),
      body: Center(
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.center,
          children: <Widget>[
            Row(
              children: [
                OutlinedButton(
                  onPressed: () {
                    _datePicker(context);
                  },
                  child: const Text("select Date"),
                ),
                TextButton(onPressed: getInfo, child: const Text('show')),
//
                const SizedBox(
                  width: 12,
                ),
                // <- 返回主页
                TextButton.icon(
                    onPressed: () {
                      Navigator.pop(context);
                    },
                    icon: const Icon(
                      Icons.keyboard_backspace,
                      color: Colors.green,
                      size: 30.0,
                    ),
                    label: const Text('')),
                // 另外一个页面
                // TextButton(
                //   onPressed: () {
                //     Navigator.push(
                //         context,
                //         MaterialPageRoute(
                //             builder: ((context) => Home())));
                //   },
                //   child: const Text('JQH'),
                // ),
              ],
            ),
            Text(infoStr),
          ],
        ),
      ),
    );
  }
}
