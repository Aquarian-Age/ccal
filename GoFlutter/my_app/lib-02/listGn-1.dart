import 'package:flutter/material.dart';

// 九宫布局
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
          // 间距
          const SizedBox(
            height: 20,
          ),
          const Text(
            ' infoStr Text Show\ninfoStr Text Show\ninfoStr Text Show\ninfoStr Text Show\ninfoStr Text Show\ninfoStr Text Show\n',
            style: const TextStyle(fontSize: 14, color: Colors.red),
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
          //     //--------------
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
        child: Container(
          margin: const EdgeInsets.all(10),
          //padding: EdgeInsets.all(10),
          child: Column(
            children: const [
              Text('八神 空 马'),
              Text('九星 干'),
              Text('八门 干'),
              Text('暗干 2,3,4,5,10'),
            ],
          ),
        ),
      ),
      OutlinedButton(
        onPressed: () {
          print('outlint btn cliekce');
        },
        child: Container(
          margin: const EdgeInsets.all(10),
          //padding: EdgeInsets.all(10),
          child: Column(
            children: const [
              Text('八神 空 马'),
              Text('九星 干'),
              Text('八门 干'),
              Text('暗干 2,3,4,5,10'),
            ],
          ),
        ),
      ),
      OutlinedButton(
        onPressed: () {
          print('outlint btn cliekce');
        },
        child: Container(
          margin: const EdgeInsets.all(10),
          //  padding: EdgeInsets.all(10),
          child: Column(
            // crossAxisAlignment: CrossAxisAlignment.start,
            children: const [
              Text('八神 空 马'),
              Text('九星 干'),
              Text('八门 干'),
              Text('暗干 2,3,4,5,10'),
            ],
          ),
        ),
      ),
    ],
  );
}
