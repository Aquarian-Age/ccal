import 'package:flutter/material.dart';
import 'package:marquee_flutter/marquee_flutter.dart';

class TitleDay extends StatelessWidget {
  final double screenWidth;

  TitleDay(this.num, this.screenWidth)
      : assert(1 <= num),
        assert(num <= 7);
  final int num;

  final List<String> _weekDayName = ["一", "二", "三", "四", "五", "六", "日"];

  @override
  Widget build(BuildContext context) {
    return Container(
      width: screenWidth / 8,
      height: screenWidth / 8 / 10 * 6,
      decoration: BoxDecoration(
        color: Colors.blue[300],
        border: Border.all(width: 0.5, color: Colors.black38),
        borderRadius: BorderRadius.all(Radius.circular(6.0)),
      ),
      child: Center(
        child: FittedBox(
          child: Text(
            _weekDayName[num - 1],
            style: TextStyle(fontSize: screenWidth / 20, color: Colors.black),
          ),
        ),
      ),
    );
  }
}

class DayBox extends StatelessWidget {
  //传入宗教节日
  final String zongJiao;
  //
  final double font;
  final double screenWidth;
  final DateTime date;
  final bool showNoteIcon;
  final bool noteActive;
  final bool selected;
  final bool baskgroundGrey;
  final bool isToday;
  final List<TextSpan> gregorianStrs;
  final List<TextSpan> lunarStrs;
  final Function(DateTime, bool) onSelectCallback;

  ///双击返回宗教节日
  final String Function(String zongJiao) onDoubleTapCallback;

  DayBox(
    this.zongJiao,
    this.font,
    this.date,
    this.screenWidth, {
    this.showNoteIcon = false,
    this.noteActive = true,
    this.selected = false,
    this.baskgroundGrey = false,
    this.isToday = false,
    this.gregorianStrs,
    this.lunarStrs,
    this.onSelectCallback,
    //
    this.onDoubleTapCallback,
  });

  Widget _buildText(List<TextSpan> strs, AlignmentGeometry alignment) {
    int length = 0;
    String str;
    strs.forEach((var e) {
      length += e.text.length;
      if (length >= 5) {
        str = "${strs[0].text} ${strs[1].text} ${strs[2].text} ${strs[3].text} ${e.text}"; // ${e.text}
        // print("str $str ---> ${str.length}");
      }
    });
    final richText = RichText(text: TextSpan(children: strs));

    return Container(
      alignment: alignment,
      child: (length < 5)
          ? richText
          : Container(
              width: screenWidth / 8,
              height: screenWidth / 8 / 3, //阴历节日高度位置
              child: MarqueeWidget(
                text: str,
                textStyle: TextStyle(fontSize: font, color: Colors.redAccent),
                scrollAxis: Axis.horizontal,
                ratioOfBlankToScreen: 0.05,
              ),
            ),
    );
  }

  @override
  Widget build(BuildContext context) {
    Color backgroundColor;
    if (true == isToday) {
      backgroundColor = selected ? Colors.yellowAccent : Colors.yellow;
    } else if (baskgroundGrey) {
      backgroundColor = Colors.grey[300];
    }

    List<Widget> stackChildren = [];

    // 用一个单独的Container来处理选中时
    ///双击返回宗教节日候的效果
    // 如果直接在显示层处理选中效果，点击选中的时候显示内容会有细微的大小变化
    stackChildren.add(Container(
        decoration: BoxDecoration(
      color: backgroundColor,
      border: Border.all(width: selected ? 2.0 : 0.1, color: selected ? Colors.red : Colors.black38),
      borderRadius: BorderRadius.all(Radius.circular(8.0)),
    )));

    // 日期数字
    stackChildren.add(
      Container(
        alignment: Alignment.center,
        child: Text("${date.day}", style: TextStyle(fontSize: font, color: Colors.black)), //screenWidth / 25
      ),
    );

    // 需要显示月份的情况
    if (null != gregorianStrs) {
      stackChildren.add(_buildText(gregorianStrs, Alignment.topCenter));
    }

    if (null != lunarStrs) {
      stackChildren.add(_buildText(lunarStrs, Alignment.bottomCenter));
    }

    return GestureDetector(
      onTap: () {
        if (null != onSelectCallback) {
          onSelectCallback(date, !selected);
        }
      },
      //返回宗教节日到月视图页面
      onDoubleTap: () {
        if (onDoubleTapCallback != null) {
          onDoubleTapCallback(zongJiao);
        }
      },
      child: Container(
        width: screenWidth / 8,
        height: screenWidth / 7, //8// 日历子部件高度
        child: Stack(children: stackChildren),
      ),
    );
  }
}
