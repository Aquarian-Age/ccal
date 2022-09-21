import 'dart:async';
import 'dart:convert';

import 'package:http/http.dart' as http;

Future<void> main() async {
  final url = "http://localhost:6714";
  final client = http.Client();
  final request = new http.Request('GET', Uri.parse(url))
    ..followRedirects = false;
  final response = await client.send(request);

  print(response.headers['content-type']); //application/json; charset=utf-8
  print(response.statusCode); //200
  if (response.statusCode == 200) {
    final resp = await http.get(Uri.parse(url)); //从服务器获取数据
    //print(resp.body);
    Map<String, dynamic> js = jsonDecode(resp.body);
    print(js['t']);

    final respInfo = await http.get(Uri.parse(url + '/info'));
    //print(respInfo.body);
    final jsonInfo = json.decode(respInfo.body);
    Info info = Info.fromJson(jsonInfo);
    print(info.t);
  } else {
    throw Exception("get $url error");
  }

  client.close();
}

class Info {
  String t;
  String ygz;
  String mgz;
  String dgz;
  String hgz;
  String moon;
  String solar;
  Info({this.t, this.ygz, this.mgz});
  Info.fromJson(Map<String, dynamic> json)
      : t = json['t'],
        ygz = json['ygz'],
        mgz = json['mgz'],
        dgz = json['dgz'],
        hgz = json['hgz'],
        moon = json['moon'],
        solar = json['solar'];
}
