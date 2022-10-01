import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter_my_picker_null_safety/common/date.dart';
import 'package:flutter_my_picker_null_safety/flutter_my_picker.dart';
import 'package:http/http.dart' as http;

// https://docs.flutter.dev/cookbook/networking/send-data
// 客户端输入 以POST json格式 传送给服务端
Future<ClientData> createData(String data) async {
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
    return ClientData.fromJson(jsonDecode(respones.body));
  } else {
    throw Exception('Failed to create clientData');
  }
}

class ClientData {
  String times;
  String gzs;
  String info;
  ClientData({required this.times, required this.gzs, required this.info});

  // ClientData.fromJson(Map<String, dynamic> json)
  //     : times = json['times'],
  //       gzs = json['gzs'],
  //       info = json['info'];
  factory ClientData.fromJson(Map<String, dynamic> json) {
    return ClientData(
      times: json['times'],
      gzs: json['gzs'],
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
  String dateStr = '';
  void _datePicker(context) {
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
  late ClientData clientData;
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
      clientData = ClientData.fromJson(jsonDecode(respones.body));
      setState(() {
        infoStr = clientData.info;
      });
    } else {
      infoStr = '连接服务端失败';
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Info"), //style: TextStyle(color: Colors.blue)),
        titleTextStyle: const TextStyle(fontSize: 21),
        centerTitle: true, //标题居中显示
        backgroundColor: Colors.blue, //标题背景颜色
        toolbarHeight: 26, //标题高度
      ),
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
