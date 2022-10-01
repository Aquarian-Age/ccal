class Info {
  String t;
  String lu;
  String gzs;
  String nayins;
  String zqs;
  String yjs;
  String jianchus;
  String huangheis;
  String huangheiHs;
  String xk;
  String moon;
  String riqins;
  String solar;
  String jueliris;
  String jitanbings;
  Info({
    required this.t,
    required this.lu,
    required this.gzs,
    required this.nayins,
    required this.zqs,
    required this.yjs,
    required this.jianchus,
    required this.huangheis,
    required this.huangheiHs,
    required this.xk,
    required this.moon,
    required this.riqins,
    required this.solar,
    required this.jueliris,
    required this.jitanbings,
  });
  Info.fromJson(Map<String, dynamic> json)
      : t = json['t'],
        lu = json['lu'],
        gzs = json['gan_zhi'],
        nayins = json['na_yin'],
        zqs = json['zhong_qi_string'],
        yjs = json['yue_jiang_string'],
        jianchus = json['jian_chu'],
        huangheis = json['huang_hei'],
        huangheiHs = json['huang_hei_h'],
        xk = json['xun_kong'],
        moon = json['moon'],
        riqins = json['ri_qin'],
        solar = json['solar'],
        jueliris = json['jue_li_ri'],
        jitanbings = json['ji_tan_bing'];
}
