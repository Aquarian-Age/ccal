import 'package:flutter/material.dart';
import 'month_view.dart';
import 'dart:ui';
void main() {
  if (window.physicalSize.isEmpty) {
    window.onMetricsChanged = () {
      if (!window.physicalSize.isEmpty) {
        window.onMetricsChanged = null;
        runApp(MyApp());
      }
    };
  } else {
    runApp(MyApp());
  }
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: '阴阳历',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: _HomePage(),
    );
  }
}

class _HomePage extends StatefulWidget {
  _HomePage({Key key}) : super(key: key);

  @override
  createState() {
    return _HomePageState();
  }
}

class _HomePageState extends State<_HomePage> {

  @override
  Widget build(BuildContext context) {

    return MonthView(
      onDateSelectedFn: (DateTime selectedDate) {},
      onMonthChangeFn: (DateTime showMonth) {
        ;
      },
      initDate: null,
    );
  }
}
