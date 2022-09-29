import 'dart:convert';

import 'package:datepicker/client.dart';
import 'package:flutter/material.dart';
import 'package:flutter_my_picker_null_safety/common/date.dart';
import 'package:flutter_my_picker_null_safety/flutter_my_picker.dart';
import 'package:http/http.dart' as http;

class DemoPage extends StatefulWidget {
  @override
  _DemoPageState createState() => _DemoPageState();
}

class _DemoPageState extends State<DemoPage> {
  DateTime date = DateTime.now();
  String dateStr = "";

  @override
  void initState() {
    super.initState();
    // DateTime date = DateTime.now();
    // print(MyDate.format('yyyy-MM-dd HH:mm:ss', date));
    setState(() {
      date; //= date;
      dateStr = MyDate.format('yyyy-MM-dd HH:mm:ss', date);
    });
  }

  _change(formatString) {
    return (date) {
      // print("change: " + MyDate.format(formatString, date));
      setState(() {
        date; //= date;
        dateStr = MyDate.format(formatString, date);
      });
    };
  }

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
    return Column(
      crossAxisAlignment: CrossAxisAlignment.center,
      children: <Widget>[
        Container(
            // padding: const EdgeInsets.only(top: 8.0),
            // width: MediaQuery.of(context).size.width,
            // child: Text(
            //   '当前时间： $dateStr',
            //   textAlign: TextAlign.center,
            // ),
            ),
        Row(
          crossAxisAlignment: CrossAxisAlignment.center,
          children: [
            _Button('select date', () {
              MyPicker.showDateTimePicker(
                context: context,
                background: Colors.black,
                color: Colors.white,
                current: date,
                magnification: 1.2,
                squeeze: 1.45,
                offAxisFraction: 0.2,
                onChange: _change('yyyy-MM-dd HH:mm'),
              );
            }),
            const Text(' '), //一个间隔
            OutlinedButton(
              onPressed: getInfo,
              child: const Text(
                'Info',
                style: TextStyle(fontSize: 21),
              ),
            ),
          ],
        ),
        Text(infoStr),
      ],
    );
  }
}

typedef OnChangeButton = Function();

class _Button extends StatelessWidget {
  final OnChangeButton onPressed;
  final String text;

  const _Button(
    this.text,
    this.onPressed,
  );

  @override
  Widget build(BuildContext context) {
    return ElevatedButton(
      onPressed: onPressed,
      child: Text(text),
    );
  }
}
