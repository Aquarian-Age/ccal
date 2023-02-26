import 'package:flutter/material.dart';

void main() => runApp(const AppBarApp());

class AppBarApp extends StatelessWidget {
  const AppBarApp({super.key});

  @override
  Widget build(BuildContext context) {
    return const MaterialApp(
      home: AppBarExample(),
    );
  }
}

class AppBarExample extends StatefulWidget {
  const AppBarExample({super.key});

  @override
  State<StatefulWidget> createState() => _AppBarExampleState();
}

class _AppBarExampleState extends State<AppBarExample> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      // appBar: AppBar(
      //   title: const Text('title'),
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
                  //_datePicker(context);
                  print('ElevateButton clicked');
                },
                child: const Text('select Data'),
              ),
              const SizedBox(
                width: 12,
              ),
              ElevatedButton(
                  onPressed: () {
                    print('ElevateButton clicked');
                  },
                  child: const Text('show')),
              const SizedBox(
                width: 12,
              ),
              ElevatedButton(
                  onPressed: () {
                    // _showPickerNumber(context);
                    print('ElevateButton clicked');
                  },
                  child: const Text("寄宫"))
            ],
          ),
          Text(
            ' infoStr Text Show\ninfoStr Text Show\ninfoStr Text Show\ninfoStr Text Show\ninfoStr Text Show\ninfoStr Text Show\n',
            style: const TextStyle(fontSize: 14),
          ),
          //-----------------
          FittedBox(
            child: Column(
              children: [
                const SizedBox(
                  height: 10, //间距
                ),
                _row(),
                _row(),
                _row(),
              ],
            ),
          ),
		  //--------------
        ],
      ),
    );
  }

}

Widget _row() {
  return Row(
    children: [
      OutlinedButton(
        onPressed: () {
          print('outlint btn cliekce');
        },
        child: Text(
          '九星 旬空\n八神\n天盘 干\n底盘 干\n暗干',
          textAlign: TextAlign.start,
          style:
              TextStyle(fontSize: 16, color: Color.fromARGB(255, 228, 27, 218)),
        ),
      ),
      OutlinedButton(
        onPressed: () {
          print('outlint btn cliekce');
        },
        child: Text(
          '九星 旬空\n八神\n天盘 干\n底盘 干\n暗干',
          textAlign: TextAlign.start,
          style:
              TextStyle(fontSize: 16, color: Color.fromARGB(255, 228, 27, 218)),
        ),
      ),
      OutlinedButton(
        onPressed: () {
          print('outlint btn cliekce');
        },
        child: Text(
          '九星 旬空\n八神\n天盘 干\n底盘 干\n暗干',
          textAlign: TextAlign.start,
          style:
              TextStyle(fontSize: 16, color: Color.fromARGB(255, 228, 27, 218)),
        ),
      ),
    ],
  );
}
