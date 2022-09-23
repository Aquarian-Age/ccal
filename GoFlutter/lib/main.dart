import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:sprintf/sprintf.dart';

import 'info.dart';

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
  String ygzs = "";
  void getServe() async {
    const url = "http://localhost:6714";
    final client = http.Client();
    final request = http.Request('GET', Uri.parse(url))
      ..followRedirects = false;
    final response = await client.send(request);
    // ignore: avoid_print
    //print(response.statusCode);

    if (response.statusCode == 200) {
      final resp = await http.get(Uri.parse(url)); //从服务器获取数据
      final jsonMap = json.decode(resp.body);
      Info info = Info.fromJson(jsonMap);

      // print("=======================");
      // print('${info.ygz}' '${info.mgz}' '${info.dgz}' '${info.hgz}');

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

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        appBar: AppBar(
          title: const Text("app title"),
        ),
        body: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
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
}
