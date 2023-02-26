class ClientData {
  String times;
  String gzs;
  String info;
  ClientData({required this.times, required this.gzs, required this.info});

  factory ClientData.fromJson(Map<String, dynamic> json) {
    return ClientData(
      times: json['times'], //传送时间到服务端
      gzs: json['gzs'], //获取服务端返回值
      info: json['info'], //获取服务端返回值
    );
  }
}

// clientDataFunc(String data) async {
//   const url = "http://localhost:6714";
//   final respones = await http.post(
//     Uri.parse(url),
//     headers: <String, String>{
//       'Content-Type': 'application/json; charset=UTF-8',
//     },
//     body: jsonEncode(<String, String>{
//       'times': data,
//     }),
//   );
//   if (respones.statusCode == 200) {
//     return ClientData.fromJson(jsonDecode(respones.body));
//   }
// }
