import 'dart:convert';

import 'package:http/http.dart' as http;

class JqhData {
  // String? times;
  // String? stargn;
  // String? doorgn;
  //
  String? kanBaShen;
  String? kanStar;
  String? kanTianQi;
  String? kanDoor;
  String? kanDiQi;
  String? kanAgz;
  String? genBaShen;
  String? genStar;
  String? genTianQi;
  String? genDoor;
  String? genDiQi;
  String? genAgz;
  String? zhenBaShen;
  String? zhenStar;
  String? zhenTianQi;
  String? zhenDoor;
  String? zhenDiQi;
  String? zhenAgz;
  String? xunBaShen;
  String? xunStar;
  String? xunTianQi;
  String? xunDoor;
  String? xunDiQi;
  String? xunAgz;
  String? liBaShen;
  String? liStar;
  String? liTianQi;
  String? liDoor;
  String? liDiQi;
  String? liAgz;
  String? kunBaShen;
  String? kunStar;
  String? kunTianQi;
  String? kunDoor;
  String? kunDiQi;
  String? kunAgz;
  String? duiBaShen;
  String? duiStar;
  String? duiTianQi;
  String? duiDoor;
  String? duiDiQi;
  String? duiAgz;
  String? qianBaShen;
  String? qianStar;
  String? qianTianQi;
  String? qianDoor;
  String? qianDiQi;
  String? qianAgz;
  String? zhongStar;
  String? zhongDiQi;
  String? zhongAgz;
  String? g1Hide;
  String? g8Hide;
  String? g3Hide;
  String? g4Hide;
  String? g9Hide;
  String? g2Hide;
  String? g7Hide;
  String? g6Hide;
  String? g5Hide;
  List<String>? info; //数组 13
  String? zhongBaShen;
  String? zhongTianQi;
  String? zhongDoor;
  String? g1s;
  String? g8s;
  String? g3s;
  String? g4s;
  String? g9s;
  String? g2s;
  String? g7s;
  String? g6s;
  String? g5s;

  JqhData({
    // required this.times,
    // required this.stargn,
    // required this.doorgn,
    //
    required this.kanBaShen,
    required this.kanStar,
    required this.kanTianQi,
    required this.kanDoor,
    required this.kanDiQi,
    required this.kanAgz,
    required this.genBaShen,
    required this.genStar,
    required this.genTianQi,
    required this.genDoor,
    required this.genDiQi,
    required this.genAgz,
    required this.zhenBaShen,
    required this.zhenStar,
    required this.zhenTianQi,
    required this.zhenDoor,
    required this.zhenDiQi,
    required this.zhenAgz,
    required this.xunBaShen,
    required this.xunStar,
    required this.xunTianQi,
    required this.xunDoor,
    required this.xunDiQi,
    required this.xunAgz,
    required this.liBaShen,
    required this.liStar,
    required this.liTianQi,
    required this.liDoor,
    required this.liDiQi,
    required this.liAgz,
    required this.kunBaShen,
    required this.kunStar,
    required this.kunTianQi,
    required this.kunDoor,
    required this.kunDiQi,
    required this.kunAgz,
    required this.duiBaShen,
    required this.duiStar,
    required this.duiTianQi,
    required this.duiDoor,
    required this.duiDiQi,
    required this.duiAgz,
    required this.qianBaShen,
    required this.qianStar,
    required this.qianTianQi,
    required this.qianDoor,
    required this.qianDiQi,
    required this.qianAgz,
    required this.zhongStar,
    required this.zhongDiQi,
    required this.zhongAgz,
    required this.g1Hide,
    required this.g8Hide,
    required this.g3Hide,
    required this.g4Hide,
    required this.g9Hide,
    required this.g2Hide,
    required this.g7Hide,
    required this.g6Hide,
    required this.g5Hide,
    required this.info, // 数组
    required this.zhongBaShen,
    required this.zhongTianQi,
    required this.zhongDoor,
    required this.g1s,
    required this.g8s,
    required this.g3s,
    required this.g4s,
    required this.g9s,
    required this.g2s,
    required this.g7s,
    required this.g6s,
    required this.g5s,
  });

  JqhData.fromJson(dynamic json) {
    // times = json['times'];
    // stargn = json['stargn'];
    // doorgn = json['doorgn'];
    //
    kanBaShen = json['kan_ba_shen'];
    kanStar = json['kan_star'];
    kanTianQi = json['kan_tian_qi'];
    kanDoor = json['kan_door'];
    kanDiQi = json['kan_di_qi'];
    kanAgz = json['kan_agz'];
    genBaShen = json['gen_ba_shen'];
    genStar = json['gen_star'];
    genTianQi = json['gen_tian_qi'];
    genDoor = json['gen_door'];
    genDiQi = json['gen_di_qi'];
    genAgz = json['gen_agz'];
    zhenBaShen = json['zhen_ba_shen'];
    zhenStar = json['zhen_star'];
    zhenTianQi = json['zhen_tian_qi'];
    zhenDoor = json['zhen_door'];
    zhenDiQi = json['zhen_di_qi'];
    zhenAgz = json['zhen_agz'];
    xunBaShen = json['xun_ba_shen'];
    xunStar = json['xun_star'];
    xunTianQi = json['xun_tian_qi'];
    xunDoor = json['xun_door'];
    xunDiQi = json['xun_di_qi'];
    xunAgz = json['xun_agz'];
    liBaShen = json['li_ba_shen'];
    liStar = json['li_star'];
    liTianQi = json['li_tian_qi'];
    liDoor = json['li_door'];
    liDiQi = json['li_di_qi'];
    liAgz = json['li_agz'];
    kunBaShen = json['kun_ba_shen'];
    kunStar = json['kun_star'];
    kunTianQi = json['kun_tian_qi'];
    kunDoor = json['kun_door'];
    kunDiQi = json['kun_di_qi'];
    kunAgz = json['kun_agz'];
    duiBaShen = json['dui_ba_shen'];
    duiStar = json['dui_star'];
    duiTianQi = json['dui_tian_qi'];
    duiDoor = json['dui_door'];
    duiDiQi = json['dui_di_qi'];
    duiAgz = json['dui_agz'];
    qianBaShen = json['qian_ba_shen'];
    qianStar = json['qian_star'];
    qianTianQi = json['qian_tian_qi'];
    qianDoor = json['qian_door'];
    qianDiQi = json['qian_di_qi'];
    qianAgz = json['qian_agz'];
    zhongStar = json['zhong_star'];
    zhongDiQi = json['zhong_di_qi'];
    zhongAgz = json['zhong_agz'];
    g1Hide = json['g1_hide'];
    g8Hide = json['g8_hide'];
    g3Hide = json['g3_hide'];
    g4Hide = json['g4_hide'];
    g9Hide = json['g9_hide'];
    g2Hide = json['g2_hide'];
    g7Hide = json['g7_hide'];
    g6Hide = json['g6_hide'];
    g5Hide = json['g5_hide'];
    info = json['info'] != null ? json['info'].cast<String>() : []; //数组
    zhongBaShen = json['zhong_ba_shen'];
    zhongTianQi = json['zhong_tian_qi'];
    zhongDoor = json['zhong_door'];
    g1s = json['g1s'];
    g8s = json['g8s'];
    g3s = json['g3s'];
    g4s = json['g4s'];
    g9s = json['g9s'];
    g2s = json['g2s'];
    g7s = json['g7s'];
    g6s = json['G6s'];
    g5s = json['g5s'];
  }

  Map<String, dynamic> toJson() {
    final map = <String, dynamic>{};
    // map['times'] = times;
    // map['stargn'] = stargn;
    // map['doorgn'] = doorgn;
    //
    map['kan_ba_shen'] = kanBaShen;
    map['kan_star'] = kanStar;
    map['kan_tian_qi'] = kanTianQi;
    map['kan_door'] = kanDoor;
    map['kan_di_qi'] = kanDiQi;
    map['kan_agz'] = kanAgz;
    map['gen_ba_shen'] = genBaShen;
    map['gen_star'] = genStar;
    map['gen_tian_qi'] = genTianQi;
    map['gen_door'] = genDoor;
    map['gen_di_qi'] = genDiQi;
    map['gen_agz'] = genAgz;
    map['zhen_ba_shen'] = zhenBaShen;
    map['zhen_star'] = zhenStar;
    map['zhen_tian_qi'] = zhenTianQi;
    map['zhen_door'] = zhenDoor;
    map['zhen_di_qi'] = zhenDiQi;
    map['zhen_agz'] = zhenAgz;
    map['xun_ba_shen'] = xunBaShen;
    map['xun_star'] = xunStar;
    map['xun_tian_qi'] = xunTianQi;
    map['xun_door'] = xunDoor;
    map['xun_di_qi'] = xunDiQi;
    map['xun_agz'] = xunAgz;
    map['li_ba_shen'] = liBaShen;
    map['li_star'] = liStar;
    map['li_tian_qi'] = liTianQi;
    map['li_door'] = liDoor;
    map['li_di_qi'] = liDiQi;
    map['li_agz'] = liAgz;
    map['kun_ba_shen'] = kunBaShen;
    map['kun_star'] = kunStar;
    map['kun_tian_qi'] = kunTianQi;
    map['kun_door'] = kunDoor;
    map['kun_di_qi'] = kunDiQi;
    map['kun_agz'] = kunAgz;
    map['dui_ba_shen'] = duiBaShen;
    map['dui_star'] = duiStar;
    map['dui_tian_qi'] = duiTianQi;
    map['dui_door'] = duiDoor;
    map['dui_di_qi'] = duiDiQi;
    map['dui_agz'] = duiAgz;
    map['qian_ba_shen'] = qianBaShen;
    map['qian_star'] = qianStar;
    map['qian_tian_qi'] = qianTianQi;
    map['qian_door'] = qianDoor;
    map['qian_di_qi'] = qianDiQi;
    map['qian_agz'] = qianAgz;
    map['zhong_star'] = zhongStar;
    map['zhong_di_qi'] = zhongDiQi;
    map['zhong_agz'] = zhongAgz;
    map['g1_hide'] = g1Hide;
    map['g8_hide'] = g8Hide;
    map['g3_hide'] = g3Hide;
    map['g4_hide'] = g4Hide;
    map['g9_hide'] = g9Hide;
    map['g2_hide'] = g2Hide;
    map['g7_hide'] = g7Hide;
    map['g6_hide'] = g6Hide;
    map['g5_hide'] = g5Hide;
    map['info'] = info;
    map['zhong_ba_shen'] = zhongBaShen;
    map['zhong_tian_qi'] = zhongTianQi;
    map['zhong_door'] = zhongDoor;
    map['g1s'] = g1s;
    map['g8s'] = g8s;
    map['g3s'] = g3s;
    map['g4s'] = g4s;
    map['g9s'] = g9s;
    map['g2s'] = g2s;
    map['g7s'] = g7s;
    map['G6s'] = g6s;
    map['g5s'] = g5s;
    return map;
  }
}

// 客户端输入 以POST json格式 传送给服务端
Future<JqhData> createData(String data, String stargn, String doorgn) async {
  const url = "http://localhost:6714";
  final respones = await http.post(
    Uri.parse(url),
    headers: <String, String>{
      'Content-Type': 'application/json; charset=UTF-8',
    },
    body: jsonEncode(<String, String>{
      'times': data,
      'stargn': stargn,
      'doorgn': doorgn,
    }),
  );
  //服务器返回来的数据
  if (respones.statusCode == 200) {
    return JqhData.fromJson(jsonDecode(respones.body));
  } else {
    throw Exception('Failed to create clientData');
  }
}

doJqhMap(jqhMap) {
  String kanBaShen = jqhMap['kan_ba_shen'];
  String kanStar = jqhMap['kan_star'];
  String kanTianQi = jqhMap['kan_tian_qi'];
  String kanDoor = jqhMap['kan_door'];
  String kanDiQi = jqhMap['kan_di_qi'];
  String kanAgz = jqhMap['kan_agz'];
  String genBaShen = jqhMap['gen_ba_shen'];
  String genStar = jqhMap['gen_star'];
  String genTianQi = jqhMap['gen_tian_qi'];
  String genDoor = jqhMap['gen_door'];
  String genDiQi = jqhMap['gen_di_qi'];
  String genAgz = jqhMap['gen_agz'];
  String zhenBaShen = jqhMap['zhen_ba_shen'];
  String zhenStar = jqhMap['zhen_star'];
  String zhenTianQi = jqhMap['zhen_tian_qi'];
  String zhenDoor = jqhMap['zhen_door'];
  String zhenDiQi = jqhMap['zhen_di_qi'];
  String zhenAgz = jqhMap['zhen_agz'];
  String xunBaShen = jqhMap['xun_ba_shen'];
  String xunStar = jqhMap['xun_star'];
  String xunTianQi = jqhMap['xun_tian_qi'];
  String xunDoor = jqhMap['xun_door'];
  String xunDiQi = jqhMap['xun_di_qi'];
  String xunAgz = jqhMap['xun_agz'];
  String liBaShen = jqhMap['li_ba_shen'];
  String liStar = jqhMap['li_star'];
  String liTianQi = jqhMap['li_tian_qi'];
  String liDoor = jqhMap['li_door'];
  String liDiQi = jqhMap['li_di_qi'];
  String liAgz = jqhMap['li_agz'];
  String kunBaShen = jqhMap['kun_ba_shen'];
  String kunStar = jqhMap['kun_star'];
  String kunTianQi = jqhMap['kun_tian_qi'];
  String kunDoor = jqhMap['kun_door'];
  String kunDiQi = jqhMap['kun_di_qi'];
  String kunAgz = jqhMap['kun_agz'];
  String duiBaShen = jqhMap['dui_ba_shen'];
  String duiStar = jqhMap['dui_star'];
  String duiTianQi = jqhMap['dui_tian_qi'];
  String duiDoor = jqhMap['dui_door'];
  String duiDiQi = jqhMap['dui_di_qi'];
  String duiAgz = jqhMap['dui_agz'];
  String qianBaShen = jqhMap['qian_ba_shen'];
  String qianStar = jqhMap['qian_star'];
  String qianTianQi = jqhMap['qian_tian_qi'];
  String qianDoor = jqhMap['qian_door'];
  String qianDiQi = jqhMap['qian_di_qi'];
  String qianAgz = jqhMap['qian_agz'];
  String zhongStar = jqhMap['zhong_star'];
  String zhongDiQi = jqhMap['zhong_di_qi'];
  String zhongAgz = jqhMap['zhong_agz'];
  String g1Hide = jqhMap['g1_hide'];
  String g8Hide = jqhMap['g8_hide'];
  String g3Hide = jqhMap['g3_hide'];
  String g4Hide = jqhMap['g4_hide'];
  String g9Hide = jqhMap['g9_hide'];
  String g2Hide = jqhMap['g2_hide'];
  String g7Hide = jqhMap['g7_hide'];
  String g6Hide = jqhMap['g6_hide'];
  String g5Hide = jqhMap['g5_hide'];

  String zhongBaShen = jqhMap['zhong_ba_shen'];
  String zhongTianQi = jqhMap['zhong_tian_qi'];
  String zhongDoor = jqhMap['zhong_door'];
  String g1s = jqhMap['g1s'];
  String g8s = jqhMap['g8s'];
  String g3s = jqhMap['g3s'];
  String g4s = jqhMap['g4s'];
  String g9s = jqhMap['g9s'];
  String g2s = jqhMap['g2s'];
  String g7s = jqhMap['g7s'];
  String g6s = jqhMap['G6s'];
  String g5s = jqhMap['g5s'];

  return JqhData(
    //times: '', stargn: '', doorgn: '', //
    kanBaShen: kanBaShen, kanStar: kanStar, kanTianQi: kanTianQi,
    kanDoor: kanDoor, kanDiQi: kanDiQi, kanAgz: kanAgz,
    genBaShen: genBaShen, genStar: genStar, genTianQi: genTianQi,
    genDoor: genDoor, genDiQi: genDiQi, genAgz: genAgz,
    zhenBaShen: zhenBaShen, zhenStar: zhenStar, zhenTianQi: zhenTianQi,
    zhenDoor: zhenDoor, zhenDiQi: zhenDiQi, zhenAgz: zhenAgz,
    xunBaShen: xunBaShen, xunStar: xunStar, xunTianQi: xunTianQi,
    xunDoor: xunDoor, xunDiQi: xunDiQi, xunAgz: xunAgz,
    liBaShen: liBaShen, liStar: liStar, liTianQi: liTianQi, liDoor: liDoor,
    liDiQi: liDiQi, liAgz: liAgz,
    kunBaShen: kunBaShen, kunStar: kunStar, kunTianQi: kunTianQi,
    kunDoor: kunDoor, kunDiQi: kunDiQi, kunAgz: kunAgz,
    duiBaShen: duiBaShen, duiStar: duiStar, duiTianQi: duiTianQi,
    duiDoor: duiDoor, duiDiQi: duiDiQi, duiAgz: duiAgz,
    qianBaShen: qianBaShen, qianStar: qianStar, qianTianQi: qianTianQi,
    qianDoor: qianDoor, qianDiQi: qianDiQi, qianAgz: qianAgz,
    zhongStar: zhongStar, zhongDiQi: zhongDiQi, zhongAgz: zhongAgz,
    g1Hide: g1Hide, g8Hide: g8Hide, g3Hide: g3Hide, g4Hide: g4Hide,
    g9Hide: g9Hide, g2Hide: g2Hide, g7Hide: g7Hide, g6Hide: g6Hide,
    g5Hide: g5Hide,
    info: [], //
    zhongBaShen: zhongBaShen, zhongTianQi: zhongTianQi, zhongDoor: zhongDoor,
    g1s: g1s, g8s: g8s, g3s: g3s, g4s: g4s, g9s: g9s, g2s: g2s, g7s: g7s,
    g6s: g6s, g5s: g5s,
  );
}
