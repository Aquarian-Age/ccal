import 'dart:convert';
import 'dart:ffi';

import 'package:ffi/ffi.dart';
import 'package:flutter/material.dart';
import 'package:goffi/group.dart';

void main() {
  runApp(const MyApp());
}

// linux desktop
typedef NativeFunc = Pointer<Utf8> Function();
typedef DarkFunc = Pointer<Utf8> Function();

class MyApp extends StatelessWidget {
  const MyApp({super.key});
//
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    // linux desktop
    final gets = DynamicLibrary.open("./linux/libgets.dll");
    final fns = gets.lookupFunction<NativeFunc, DarkFunc>('gets');

    // final DynamicLibrary nativeAddLib = Platform.isAndroid
    //     ? DynamicLibrary.open('libgets.so')
    //     : DynamicLibrary.process();
    // final armlib = nativeAddLib.lookupFunction<NativeFunc, DarkFunc>("gets");

// json 解析方法(简单)
    // String str = fns().toDartString();
    // Map<String, dynamic> grp = json.decode(str);
//或 json 解析(class)
    String str = fns().toDartString();
    final strjson = json.decode(str);
    Group group = Group.fromJson(strjson);

//
    return MaterialApp(
      title: 'Flutter dark:ffi Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      //home: const MyHomePage(title: 'Flutter Demo Home Page'),
      home: Scaffold(
        body: Center(
          child: Text(
            // grp['time'] + "\n" + fns().toDartString(), //json 简单解析
            group.id + group.name + group.time, //json class 方式解析
            // armlib.toString(),
            style: const TextStyle(fontSize: 26),
          ),
        ),
      ),
    );
  }
}

// class MyHomePage extends StatefulWidget {
//   const MyHomePage({super.key, required this.title});

//   final String title;

//   @override
//   State<MyHomePage> createState() => _MyHomePageState();
// }

// class _MyHomePageState extends State<MyHomePage> {
//   int _counter = 0;

//   void _incrementCounter() {
//     setState(() {
//       _counter++;
//     });
//   }

//   @override
//   Widget build(BuildContext context) {
//     return Scaffold(
//       appBar: AppBar(
//         title: Text(widget.title),
//       ),
//       body: Center(
//         child: Column(
//           mainAxisAlignment: MainAxisAlignment.center,
//           children: <Widget>[
//             const Text(
//               'You have pushed the button this many times:',
//             ),
//             Text(
//               '$_counter',
//               style: Theme.of(context).textTheme.headline4,
//             ),
//           ],
//         ),
//       ),
//       floatingActionButton: FloatingActionButton(
//         onPressed: _incrementCounter,
//         tooltip: 'Increment',
//         child: const Icon(Icons.add),
//       ), // This trailing comma makes auto-formatting nicer for build methods.
//     );
//   }
// }
