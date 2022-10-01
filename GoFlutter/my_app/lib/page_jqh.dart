import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter_my_picker_null_safety/common/date.dart';
import 'package:flutter_my_picker_null_safety/flutter_my_picker.dart';
import 'package:flutter_picker/flutter_picker.dart';
import 'package:http/http.dart' as http;
import 'package:sprintf/sprintf.dart';

import 'jqh_data.dart';

// 向服务端传送的数据
class CustomClass {
  final String times;
  final int stargn;
  final int doorgn;
  CustomClass(
      {required this.times, required this.stargn, required this.doorgn});
  CustomClass.fromJson(Map<String, dynamic> json)
      : times = json['times'],
        stargn = json['stargn'],
        doorgn = json['doorgn'];

  static Map<String, dynamic> toJson(CustomClass data) =>
      {'times': data.times, 'stargn': data.stargn, 'doorgn': data.doorgn};
}

//
class PageJQHfw extends StatefulWidget {
  const PageJQHfw({super.key});

  @override
  State<StatefulWidget> createState() => _PateJqhState();
}

class _PateJqhState extends State<PageJQHfw> {
  // 下拉寄宫数组 List<int>
  int stargn = 0;
  int doorgn = 0;
  void _showPickerNumber(context) {
    Picker(
        adapter: NumberPickerAdapter(
          data: [
            const NumberPickerColumn(items: [2, 8]),
            const NumberPickerColumn(items: [2, 8]),
          ],
        ),
        delimiter: [
          PickerDelimiter(
              child: Container(
            width: 20.0,
            alignment: Alignment.center,
            // child: Icon(Icons.more_vert),
          ))
        ],
        hideHeader: true,
        title: const Text("selectNumber"),
        selectedTextStyle: const TextStyle(color: Colors.blue),
        onConfirm: (Picker picker, List value) {
          //  print(value.toString());
          // print(picker.getSelectedValues());
          // print(picker.getSelectedValues().runtimeType.toString()); //List<int>
          setState(() {
            var list = picker.getSelectedValues();
            stargn = list[0];
            doorgn = list[1];
            // print(list[0].runtimeType.toString());
            // print(stargn.toString());
            //print(doorgn.toString());
          });
        }).showDialog(context);
  }

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

  // doJsonString() {
  //   final CustomClass cc =
  //       CustomClass(times: dateStr, stargn: stargn, doorgn: doorgn);
  //   final jsonStr = jsonEncode({"cliJqh": cc},
  //       toEncodable: (Object? value) => value is CustomClass
  //           ? CustomClass.toJson(value)
  //           : throw UnsupportedError('Cannot convert to JSON $value'));
  //   //  {"cliJqh":{"text":"2022-10-01 23:31:00","stargn":8,"doorgn":8}}
  //   // flutter: {"cliJqh":{"text":"2022-10-01 23:08:00","stargn":2,"doorgn":8}}
  //  //print(jsonStr);
  //   //print(jsonStr.runtimeType.toString()); //String
  // }

//
  late JqhData jqhData; //接受服务端返回的数据
  String infoStr = '';
  void getJqh() async {
    //doJsonString();
    // print("===============================");
    final CustomClass cc =
        CustomClass(times: dateStr, stargn: stargn, doorgn: doorgn);
    // print("cc ============== >" +
    //     cc.times +
    //     "  stargn: " +
    //     cc.stargn.toString() +
    //     "  doorgn: " +
    //     cc.doorgn.toString());
    const url = "http://localhost:6714";
    final respones = await http.post(
      Uri.parse("$url/jqh"),
      headers: <String, String>{
        'Content-Type': 'application/json; charset=UTF-8',
      },
      //传送时数据到服务端
      //body: jsonEncode(<String, String>{'times': dateStr}),
      body: jsonEncode(
        {'cliJqh': cc},
        toEncodable: (Object? value) => value is CustomClass
            ? CustomClass.toJson(value)
            : throw UnsupportedError('Cannot convert to JSON $value'),
      ),
    );

    if (respones.statusCode == 200) {
      jqhData = JqhData.fromJson(jsonDecode(respones.body));

      // info的json数组 转换为List<String>类型
      Map<String, dynamic> jqhMap = jsonDecode(respones.body);

      // 基础信息
      List<String> listInfo = List<String>.from(jqhMap['info'] as List);
      // print("listInfo Type: " + listInfo.runtimeType.toString());

// 简单的格式化输出
      var infoA = [
        "${listInfo[0]}            ${listInfo[4]}",
        "${listInfo[1]}    ${listInfo[5]}",
        "${listInfo[2]}    ${listInfo[6]}",
        "${listInfo[3]}          ${listInfo[7]}",
        listInfo[8],
        listInfo[9],
        listInfo[10],
        listInfo[11],
        listInfo[12],
      ];
      //九宫信息
      // String kanBaShen = jqhMap['kan_ba_shen'];
      // print("kanBanShen:" + kanBaShen);

      JqhData misc = doJqhMap(jqhMap);
      //print(misc.runtimeType.toString()); // flutter: JqhData
      // print(misc.kanStar);
      setState(() {
        infoStr = sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n", infoA);
        misc;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('JQH'),
        titleTextStyle: const TextStyle(fontSize: 21),
        centerTitle: true, //标题居中显示
        // backgroundColor: Colors.blue, //标题背景颜色
        toolbarHeight: 40, //标题高度
      ),
      body: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Row(
            children: [
              ElevatedButton(
                onPressed: () {
                  _datePicker(context);
                },
                child: const Text('select Data'),
              ),
              const SizedBox(
                width: 12,
              ),
              ElevatedButton(onPressed: getJqh, child: const Text('show')),
              const SizedBox(
                width: 12,
              ),
              ElevatedButton(
                  onPressed: () {
                    _showPickerNumber(context);
                  },
                  child: const Text("寄宫"))
            ],
          ),
          // 基础信息
          Text(
            infoStr,
            style: const TextStyle(fontSize: 14),
          ),
        ],
      ),
    );
  }
}

// 下拉寄宫数组 List<int>
// void _showPickerNumber(context) {
//   Picker(
//       adapter: NumberPickerAdapter(
//         data: [
//           const NumberPickerColumn(items: [2, 8]),
//           const NumberPickerColumn(items: [2, 8]),
//         ],
//       ),
//       delimiter: [
//         PickerDelimiter(
//             child: Container(
//           width: 20.0,
//           alignment: Alignment.center,
//           // child: Icon(Icons.more_vert),
//         ))
//       ],
//       hideHeader: true,
//       title: const Text("selectNumber"),
//       selectedTextStyle: const TextStyle(color: Colors.blue),
//       onConfirm: (Picker picker, List value) {
//         //  print(value.toString());
//         print(picker.getSelectedValues());
//         print(picker.getSelectedValues().runtimeType.toString()); //List<int>
//       }).showDialog(context);
// }

// flutter: [
//2022-10-01 14:11,
//干支: 壬寅 己酉 丁亥 丁未,
//旬空: 辰巳 寅卯 午未 寅卯,
//秋分:2022-09-23 09:03:00,

//起局: 阴遁中元1局,4
//值符: 天心 落4宫,
//值使: 开门 落3宫,
//旬首: 甲辰遁(壬),

//,五不遇 天显时 8
//三门: 太冲:午从魁:子小吉:戌,
//四户: 地四户: 除在:申 定在:亥 危在:寅 开在:巳,
//私门: 六合:卯 太阴:戌 太常:申,
//九月初六]
