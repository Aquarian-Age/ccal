import 'dart:convert';
import 'dart:core';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:sprintf/sprintf.dart';

import 'info.dart';

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
    return ClientData.fromJson(jsonDecode(respones.body));
  } else {
    throw Exception('Failed to create clientData');
  }
}

class ClientData {
  // int id;
  String times;

  ClientData({required this.times});

  factory ClientData.fromJson(Map<String, dynamic> json) {
    return ClientData(
      // id: json['id'],
      times: json['times'],
    );
  }
}

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
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
  //
  String ygzs = "";
  void getServe() async {
    const url = "http://localhost:6714";
    final client = http.Client();
    final request = http.Request('GET', Uri.parse(url))
      ..followRedirects = false;
    final response = await client.send(request);

    if (response.statusCode == 200) {
      final resp = await http.get(Uri.parse(url)); //从服务器获取数据
      final jsonMap = json.decode(resp.body);
      Info info = Info.fromJson(jsonMap);
      setState(() {
        var solar = info.solar;
        var gzs = info.gzs;
        var lus = info.lu;
        var xks = info.xk;
        var nys = info.nayins;
        var moons = "阴历: ${info.moon}";
        var zqs = info.zqs;
        var yjs = info.yjs;
        var huangh = info.huangheis;
        var riqin = info.riqins;
        var hhs = info.huangheiHs;
        var jlr = info.jueliris;
        var jtb = info.jitanbings;
        var listInfo = [
          solar,
          lus,
          gzs,
          xks,
          nys,
          moons,
          zqs,
          yjs,
          huangh,
          riqin,
          hhs,
          jlr,
          jtb,
        ];
        ygzs = sprintf(
            "%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s%s\n", listInfo);
      });
    }

    // print(ygzs);

    client.close();
  }

//客户端
  final TextEditingController _controller = TextEditingController();
  Future<ClientData>? _futureClientData;

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: const Text("app create data"),
        ),
        body: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
              //客户端文本输入
              Container(
                alignment: Alignment.center,
                padding: const EdgeInsets.all(8.0),
                child: (_futureClientData == null)
                    ? buildColumn()
                    : buildFutureBuilder(),
              ),
              Text(
                ygzs,
                // style: Theme.of(context).textTheme.headline4,
              ),
            ],
          ),
        ),
        floatingActionButton: FloatingActionButton(
          onPressed: getServe,
          tooltip: 'Increment',
          child: const Icon(Icons.add_home),
        ),
      ),
    );
  }

//获取输入文本
  Column buildColumn() {
    return Column(
      mainAxisAlignment: MainAxisAlignment.center,
      children: <Widget>[
        TextField(
          controller: _controller,
          decoration: const InputDecoration(hintText: 'Enter Text'),
        ),
        ElevatedButton(
          onPressed: () {
            setState(() {
              _futureClientData = createData(_controller.text);
            });
          },
          child: const Text('Create Data'),
        ),
      ],
    );
  }

  FutureBuilder<ClientData> buildFutureBuilder() {
    return FutureBuilder<ClientData>(
      future: _futureClientData,
      builder: (context, snapshot) {
        if (snapshot.hasData) {
          return Text(snapshot.data!.times);
        } else if (snapshot.hasError) {
          return Text('${snapshot.error}');
        }
        return const CircularProgressIndicator();
      },
    );
  }
}
