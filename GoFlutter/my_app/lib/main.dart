import 'dart:convert';
import 'dart:core';

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

  if (respones.statusCode == 200) {
    //var js = jsonDecode(respones.body);
    // print("decode: " + js['times']); //服务器返回来的数据
    // print("infos:" + js['info']);
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

  factory ClientData.fromJson(Map<String, dynamic> json) {
    return ClientData(
      times: json['times'],
      gzs: json['gzs'],
      info: json['info'],
    );
  }
}

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  //
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      //title: '',
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
      ClientData clientData = ClientData.fromJson(jsonDecode(respones.body));
      setState(() {
        infoStr = clientData.info;
      });
    }
  }

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
          backgroundColor: Colors.white, //标题背景颜色
          toolbarHeight: 24, //标题高度
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
                  TextButton(onPressed: getInfo, child: const Text('Info')),
                ],
              ),
              Text(infoStr),
            ],
          ),
        ),
      ),
    );
  }
}
