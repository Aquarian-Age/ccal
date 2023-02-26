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
  CustomClass({required this.times, required this.stargn, required this.doorgn});
  CustomClass.fromJson(Map<String, dynamic> json)
      : times = json['times'],
        stargn = json['stargn'],
        doorgn = json['doorgn'];

  static Map<String, dynamic> toJson(CustomClass data) => {'times': data.times, 'stargn': data.stargn, 'doorgn': data.doorgn};
}

//
class PageJQHfw extends StatefulWidget {
  const PageJQHfw({super.key});

  @override
  State<StatefulWidget> createState() => _PateJqhState();
}

class _PateJqhState extends State<PageJQHfw> {
  // // GET 跳转到这个页面之前处理GET 可以处理JqhData空的问题
  // void getServe() async {
  //   const url = "http://localhost:6714";
  //   final client = http.Client();
  //   final request = http.Request('GET', Uri.parse(url))..followRedirects = false;
  //   final response = await client.send(request);

  //   if (response.statusCode == 200) {
  //     final resp = await http.get(Uri.parse('$url/jqh'));
  //     var jqhData = JqhData.fromJson(jsonDecode(resp.body));

  //     // info的json数组 转换为List<String>类型
  //     Map<String, dynamic> jqhMap = jsonDecode(resp.body);

  //     // 基础信息
  //     List<String> listInfo = List<String>.from(jqhMap['info'] as List);
  //     print("listInfo Type: " + listInfo.runtimeType.toString());
  //     setState(() {
  //       jqhData;
  //     });
  //   }
  //   client.close();
  // }

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
      title: const Text("date"),
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
  //late JqhData jqhData; //接受服务端返回的数据
  JqhData? jqhData; //接受服务端返回的数据
  String infoStr = '';

  // POST
  void getJqh() async {
    //doJsonString();
    // print("===============================");
    final CustomClass cc = CustomClass(times: dateStr, stargn: stargn, doorgn: doorgn);
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
        toEncodable: (Object? value) => value is CustomClass ? CustomClass.toJson(value) : throw UnsupportedError('Cannot convert to JSON $value'),
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

//---------
      //JqhData misc = doJqhMap(jqhMap);
      //print(misc.runtimeType.toString()); // flutter: JqhData
      // print(misc.kanStar);
      //---------
      jqhData = doJqhMap(jqhMap);
      // jqhData ??= doNull() as JqhData?;
      // print("-----------------");
      // print(jqhData!.kanBaShen.toString());
      //print(jqhData.kanBaShen.runtimeType.toString());
      //-------------
      setState(() {
        infoStr = sprintf("%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n%s\n", infoA);
        //misc;
        jqhData;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    jqhData ??= doNull() as JqhData?;
    return Scaffold(
      // appBar: AppBar(
      //   title: const Text('JQH'),
      //   titleTextStyle: const TextStyle(fontSize: 21),
      //   centerTitle: true, //标题居中显示
      //   // backgroundColor: Colors.blue, //标题背景颜色
      //   toolbarHeight: 40, //标题高度
      // ),
      body: ListView(
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
                child: const Text("寄宫"),
              ),
              // <- 返回主页
              const SizedBox(
                width: 12,
              ),
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
            ],
          ),
          Text(
            infoStr,
            style: const TextStyle(fontSize: 14),
          ),
          //-----------------
          FittedBox(
            child: Column(
              children: [
                const SizedBox(
                  height: 10,
                ),
                _rowA(jqhData!),
                _rowB(jqhData!),
                _rowC(jqhData!),
              ],
            ),
          ),
          //     //--------------
        ],
      ),
    );
  }

  //-----------九宫布局 4 9 2 -----------
  Widget _rowA(JqhData jqh) {
    return Row(
      children: [
        //------巽宫-----
        Container(
          width: 140,
          height: 120,
          alignment: Alignment.centerLeft,
          // margin: const EdgeInsets.all(3), //外间距
          padding: const EdgeInsets.all(6),
          //color: Colors.white,
          decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(3), //圆角
            color: Colors.white, //Color.fromARGB(226, 248, 245, 245),
            //边框
            border: Border.all(
              width: 1,
              color: Colors.red,
            ),
          ),
          child: Column(
            children: [
              TextButton(onPressed: () {}, child: Text('')),
              Text(
                jqh.xunBaShen.toString(),
                style: const TextStyle(fontSize: 13),
              ),
              Text(
                "${jqh.xunStar} ${jqh.xunTianQi}",
                style: const TextStyle(fontSize: 13),
              ),
              Text(
                "${jqh.xunDoor} ${jqh.xunDiQi}",
                style: const TextStyle(fontSize: 13),
              ),
              Text(
                "${jqh.xunAgz}",
                style: const TextStyle(fontSize: 13),
              ),
            ],
          ),
        ),

        // ------ 离宫 -----
        Container(
          width: 140,
          height: 120,
          alignment: Alignment.centerLeft,
          // margin: const EdgeInsets.all(3),
          padding: const EdgeInsets.all(6),
          //color: Colors.white,
          decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(3), //圆角
            color: Colors.white, //Color.fromARGB(226, 248, 245, 245),
            //边框
            border: Border.all(
              width: 1,
              color: Colors.red,
            ),
          ),
          child: Column(
            children: [
              TextButton(onPressed: () {}, child: Text('')),
              Text(
                jqh.liBaShen.toString(),
                style: const TextStyle(fontSize: 13),
              ),
              Text(
                "${jqh.liStar} ${jqh.liTianQi}",
                style: const TextStyle(fontSize: 13),
              ),
              Text(
                "${jqh.liDoor} ${jqh.liDiQi}",
                style: const TextStyle(fontSize: 13),
              ),
              Text(
                "${jqh.liAgz}",
                style: const TextStyle(fontSize: 13),
              ),
            ],
          ),
        ),

        //  ---- 坤宫 -----
        Container(
          width: 140,
          height: 120,
          alignment: Alignment.centerLeft,
          // margin: const EdgeInsets.all(3),
          padding: const EdgeInsets.all(6),
          //color: Colors.white,
          decoration: BoxDecoration(
            borderRadius: BorderRadius.circular(3), //圆角
            color: Colors.white, //Color.fromARGB(226, 248, 245, 245),
            //边框
            border: Border.all(
              width: 1,
              color: Colors.red,
            ),
          ),
          child: Column(
            children: [
              TextButton(onPressed: () {}, child: Text('')),
              Text(
                jqh.kunBaShen.toString(),
                style: const TextStyle(fontSize: 13),
              ),
              Text(
                "${jqh.kunStar} ${jqh.kunTianQi}",
                style: const TextStyle(fontSize: 13),
              ),
              Text(
                "${jqh.kunDoor} ${jqh.kunDiQi}",
                style: const TextStyle(fontSize: 13),
              ),
              Text(
                "${jqh.kunAgz}",
                style: const TextStyle(fontSize: 13),
              ),
            ],
          ),
        ),
      ],
    );
  }
}

//------------ 九宫布局 3 5 7 -----------
Widget _rowB(jqh) {
  return Row(
    children: [
      // -------- 震宫 -------
      Container(
        width: 140,
        height: 120,
        alignment: Alignment.centerLeft,
        // margin: const EdgeInsets.all(3),
        padding: const EdgeInsets.all(6),
        //color: Colors.white,
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(3), //圆角
          color: Colors.white, //Color.fromARGB(226, 248, 245, 245),
          //边框
          border: Border.all(
            width: 1,
            color: Colors.red,
          ),
        ),
        child: Column(
          children: [
            TextButton(onPressed: () {}, child: Text('')),
            Text(
              jqh.zhenBaShen.toString(),
              style: const TextStyle(fontSize: 13),
            ),
            Text(
              "${jqh.zhenStar} ${jqh.zhenTianQi}",
              style: const TextStyle(fontSize: 13),
            ),
            Text(
              "${jqh.zhenDoor} ${jqh.zhenDiQi}",
              style: const TextStyle(fontSize: 13),
            ),
            Text(
              "${jqh.zhenAgz}",
              style: const TextStyle(fontSize: 13),
            ),
          ],
        ),
      ),

      // --------中5 -----------
      Container(
        width: 140,
        height: 120,
        alignment: Alignment.centerLeft,
        // margin: const EdgeInsets.all(3),
        padding: const EdgeInsets.all(6),
        //color: Colors.white,
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(3), //圆角
          color: Colors.white, //Color.fromARGB(226, 248, 245, 245),
          //边框
          border: Border.all(
            width: 1,
            color: Colors.red,
          ),
        ),
        child: Column(
          children: [
            TextButton(onPressed: () {}, child: Text('')),
            Text(
              jqh.zhongBaShen.toString(),
              style: const TextStyle(fontSize: 13),
            ),
            Text(
              "${jqh.zhongStar} ${jqh.zhongTianQi}",
              style: const TextStyle(fontSize: 13),
            ),
            Text(
              "${jqh.zhongDoor} ${jqh.zhongDiQi}",
              style: const TextStyle(fontSize: 13),
            ),
            Text(
              "${jqh.zhongAgz}",
              style: const TextStyle(fontSize: 13),
            ),
          ],
        ),
      ),

      // ------ 兑宫 ---------
      Container(
        width: 140,
        height: 120,
        alignment: Alignment.centerLeft,
        // margin: const EdgeInsets.all(3),
        padding: const EdgeInsets.all(6),
        //color: Colors.white,
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(3), //圆角
          color: Colors.white, //Color.fromARGB(226, 248, 245, 245),
          //边框
          border: Border.all(
            width: 1,
            color: Colors.red,
          ),
        ),
        child: Column(
          // crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            TextButton(onPressed: () {}, child: Text('')),
            Text(
              jqh.duiBaShen.toString(),
              style: const TextStyle(fontSize: 13),
            ),
            Text(
              "${jqh.duiStar} ${jqh.duiTianQi}",
              style: const TextStyle(fontSize: 13),
            ),
            Text(
              "${jqh.duiDoor} ${jqh.duiDiQi}",
              style: const TextStyle(fontSize: 13),
            ),
            Text(
              "${jqh.duiAgz}",
              style: const TextStyle(fontSize: 13),
            ),
          ],
        ),
      ),
    ],
  );
}

//------------ 九宫布局 8 1 6 -----------
Widget _rowC(jqh) {
  return Row(
    children: [
      // ----- 艮宫 -----
      Container(
        width: 140,
        height: 120,
        alignment: Alignment.centerLeft,
        // margin: const EdgeInsets.all(3),
        padding: const EdgeInsets.all(6),
        //color: Colors.white,
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(3), //圆角
          color: Colors.white, //Color.fromARGB(226, 248, 245, 245),
          //边框
          border: Border.all(
            width: 1,
            color: Colors.red,
          ),
        ),
        child: Column(
          children: [
            TextButton(onPressed: () {}, child: const Text('')),
            Text(
              jqh.genBaShen.toString(),
              style: const TextStyle(fontSize: 13),
            ),
            Text(
              "${jqh.genStar} ${jqh.genTianQi}",
              style: const TextStyle(fontSize: 13),
            ),
            Text(
              "${jqh.genDoor} ${jqh.genDiQi}",
              style: const TextStyle(fontSize: 13),
            ),
            Text(
              "${jqh.genAgz}",
              style: const TextStyle(fontSize: 13),
            ),
          ],
        ),
      ),

      // ----- 坎宫 -----
      Container(
        width: 140,
        height: 120,
        alignment: Alignment.centerLeft,
        // margin: const EdgeInsets.all(3),
        padding: const EdgeInsets.all(6),
        //color: Colors.white,
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(3), //圆角
          color: Colors.white, //Color.fromARGB(226, 248, 245, 245),
          //边框
          border: Border.all(
            width: 1,
            color: Colors.red,
          ),
        ),
        child: Column(
          children: [
            TextButton(onPressed: () {}, child: Text('')),
            Text(
              jqh.kanBaShen.toString(),
              style: const TextStyle(fontSize: 13),
            ),
            Text(
              "${jqh.kanStar} ${jqh.kanTianQi}",
              style: const TextStyle(fontSize: 13),
            ),
            Text(
              "${jqh.kanDoor} ${jqh.kanDiQi}",
              style: const TextStyle(fontSize: 13),
            ),
            Text(
              "${jqh.kanAgz}",
              style: const TextStyle(fontSize: 13),
            ),
          ],
        ),
      ),

      // ----- 乾宫 ------
      Container(
        width: 140,
        height: 120,
        alignment: Alignment.centerLeft,
        // margin: const EdgeInsets.all(3),
        padding: const EdgeInsets.all(6),
        //color: Colors.white,
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(3), //圆角
          color: Colors.white, //Color.fromARGB(226, 248, 245, 245),
          //边框
          border: Border.all(
            width: 1,
            color: Colors.red,
          ),
        ),
        child: Column(
          // crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            TextButton(onPressed: () {}, child: Text('')),
            Text(
              jqh.qianBaShen.toString(),
              style: const TextStyle(fontSize: 13),
            ),
            Text(
              "${jqh.qianStar} ${jqh.qianTianQi}",
              style: const TextStyle(fontSize: 13),
            ),
            Text(
              "${jqh.qianDoor} ${jqh.qianDiQi}",
              style: const TextStyle(fontSize: 13),
            ),
            Text(
              "${jqh.qianAgz}",
              style: const TextStyle(fontSize: 13),
            ),
          ],
        ),
      ),
    ],
  );
}
// Widget _row() {
//   print("");
//   // var kan = [jqhData.kanBaShen.toString(), jqhData.kanStar.toString()];
//   // print(kan.runtimeType.toString());
//   return Row(
//     children: [
//       OutlinedButton(
//         onPressed: () {
//           print('btnX');
//         },
//         child: Text(
//           '九星 旬空\n八神\n天盘 干\n底盘 干\n暗干',
//           textAlign: TextAlign.start,
//           style: TextStyle(fontSize: 16, color: Color.fromARGB(255, 228, 27, 218)),
//         ),
//       ),
//       OutlinedButton(
//         onPressed: () {
//           print('btnX');
//         },
//         child: Text(
//           '九星 旬空\n八神\n天盘 干\n底盘 干\n暗干',
//           textAlign: TextAlign.start,
//           style: TextStyle(fontSize: 16, color: Color.fromARGB(255, 228, 27, 218)),
//         ),
//       ),
//       OutlinedButton(
//         onPressed: () {
//           print('btnX');
//         },
//         child: Text(
//           '九星 旬空\n八神\n天盘 干\n底盘 干\n暗干', // jqhData.kanBaShen.toString(),
//           textAlign: TextAlign.start,
//           style: TextStyle(fontSize: 16, color: Color.fromARGB(255, 228, 27, 218)),
//         ),
//       ),
//     ],
//   );
// }
//--------------------------------------------
// body: Column(
//   crossAxisAlignment: CrossAxisAlignment.start,
//   children: [
//     Row(
//       children: [
//         ElevatedButton(
//           onPressed: () {
//             _datePicker(context);
//           },
//           child: const Text('select Data'),
//         ),
//         const SizedBox(
//           width: 12,
//         ),
//         ElevatedButton(onPressed: getJqh, child: const Text('show')),
//         const SizedBox(
//           width: 12,
//         ),
//         ElevatedButton(
//             onPressed: () {
//               _showPickerNumber(context);
//             },
//             child: const Text("寄宫"))
//       ],
//     ),
//     // 基础信息
//     Text(
//       infoStr,
//       style: const TextStyle(fontSize: 14),
//     ),
//     //-----------------
//     FittedBox(
//       child: Column(
//         children: [
//           const SizedBox(
//             height: 80,
//           ),
//           _row(),
//           _row(),
//           _row(),
//         ],
//       ),
//     ),
//     //--------------
//   ],
// ),
//--------------------------------------------
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

 
      // ),
      // OutlinedButton(
      //   onPressed: () {
      //     print('btnX');
      //   },
      //   child: Container(
      //     // margin: const EdgeInsets.all(3),
      //     //  padding: EdgeInsets.all(3),
      //     child: Column(
      //       // crossAxisAlignment: CrossAxisAlignment.start,
      //       children: [
      //         Text(jqh.qianBaShen.toString()),
      //         Text("${jqh.qianStar} ${jqh.qianTianQi}"),
      //         Text("${jqh.qianDoor} ${jqh.qianDiQi}"),
      //         Text("${jqh.qianAgz}"),
      //       ],
      //     ),
      //   ),
      // ),
      // OutlinedButton(
      //   onPressed: () {
      //     print('btnX');
      //   },
      //   child: Container(
      //     // margin: const EdgeInsets.all(3),
      //     //padding: EdgeInsets.all(3),
      //     child: Column(
      //       children: [
      //         Text(jqh.kanBaShen.toString()),
      //         Text("${jqh.kanStar} ${jqh.kanTianQi}"),
      //         Text("${jqh.kanDoor} ${jqh.kanDiQi}"),
      //         Text("${jqh.kanAgz}"),
      //       ],
      //     ),
      //   ),
      // ),
      // OutlinedButton(
      //   onPressed: () {
      //     print('btnX');
      //   },
      //   child: Container(
      //     // margin: const EdgeInsets.all(3),
      //     //padding: EdgeInsets.all(3),
      //     child: Column(
      //       children: [
      //         Text(jqh.genBaShen.toString()),
      //         Text("${jqh.genStar} ${jqh.genTianQi}"),
      //         Text("${jqh.genDoor} ${jqh.genDiQi}"),
      //         Text("${jqh.genAgz}"),
      //       ],
      //     ),
      //   ),
      // ),
 
      // OutlinedButton(
      //   onPressed: () {
      //     print('btnX');
      //   },
      //   child: Container(
      //     // margin: const EdgeInsets.all(3),
      //     //  padding: EdgeInsets.all(3),
      //     child: Column(
      //       // crossAxisAlignment: CrossAxisAlignment.start,
      //       children: [
      //         Text(jqh.duiBaShen.toString()),
      //         Text("${jqh.duiStar} ${jqh.duiTianQi}"),
      //         Text("${jqh.duiDoor} ${jqh.duiDiQi}"),
      //         Text("${jqh.duiAgz}"),
      //       ],
      //     ),
      //   ),
      // ),
 
      // OutlinedButton(
      //   onPressed: () {
      //     print('btnX');
      //   },
      //   child: Container(
      //     // margin: const EdgeInsets.all(3),
      //     //padding: EdgeInsets.all(3),
      //     child: Column(
      //       children: [
      //         Text(jqh.zhongBaShen.toString()),
      //         Text("${jqh.zhongStar} ${jqh.zhongTianQi}"),
      //         Text("${jqh.zhongDoor} ${jqh.zhongDiQi}"),
      //         Text("${jqh.zhongAgz}"),
      //       ],
      //     ),
      //   ),
      // ),

      // OutlinedButton(
      //   onPressed: () {
      //     print('btnX');
      //   },
      //   child: Container(
      //     // margin: const EdgeInsets.all(3),
      //     //padding: EdgeInsets.all(3),
      //     child: Column(
      //       children: [
      //         Text(jqh.zhenBaShen.toString()),
      //         Text("${jqh.zhenStar} ${jqh.zhenTianQi}"),
      //         Text("${jqh.zhenDoor} ${jqh.zhenDiQi}"),
      //         Text("${jqh.zhenAgz}"),
      //       ],
      //     ),
      //   ),
      // ),



        // SizedBox(
        //   width: 120,
        //   height: 120,
        //   child: Container(
        //     // margin: const EdgeInsets.all(3),
        //     child: Column(
        //       children: [
        //         TextButton(onPressed: () {}, child: Text('')),
        //         Text(
        //           jqh.xunBaShen.toString(),
        //           style: const TextStyle(fontSize: 13),
        //         ),
        //         Text(
        //           "${jqh.xunStar} ${jqh.xunTianQi}",
        //           style: const TextStyle(fontSize: 13),
        //         ),
        //         Text(
        //           "${jqh.xunDoor} ${jqh.xunDiQi}",
        //           style: const TextStyle(fontSize: 13),
        //         ),
        //         Text(
        //           "${jqh.xunAgz}",
        //           style: const TextStyle(fontSize: 13),
        //         ),
        //       ],
        //     ),
        //   ),
        // ),

        // OutlinedButton(
        //   onPressed: () {
        //     print('btnX');
        //   },
        //   child: Container(
        //     // margin: const EdgeInsets.all(3),
        //     child: Column(
        //       children: [
        //         Text(jqh.kunBaShen.toString()),
        //         Text("${jqh.kunStar} ${jqh.kunTianQi}"),
        //         Text("${jqh.kunDoor} ${jqh.kunDiQi}"),
        //         Text("${jqh.kunAgz}"),
        //       ],
        //     ),
        //   ),
        // ),

        // OutlinedButton(
        //   onPressed: () {
        //     print('btnX');
        //   },
        //   child: Container(
        //     // margin: const EdgeInsets.all(3),
        //     child: Column(
        //       children: [
        //         Text(jqh.xunBaShen.toString()),
        //         Text("${jqh.xunStar} ${jqh.xunTianQi}"),
        //         Text("${jqh.xunDoor} ${jqh.xunDiQi}"),
        //         Text("${jqh.xunAgz}"),
        //       ],
        //     ),
        //   ),
        // ),


        // OutlinedButton(
        //   onPressed: () {
        //     print('btnX');
        //   },
        //   child: Container(
        //     // margin: const EdgeInsets.all(3),
        //     child: Column(
        //       children: [
        //         Text(jqh.liBaShen.toString()),
        //         Text("${jqh.liStar} ${jqh.liTianQi}"),
        //         Text("${jqh.liDoor} ${jqh.liDiQi}"),
        //         Text("${jqh.liAgz}"),
        //       ],
        //     ),
        //   ),
        // ),