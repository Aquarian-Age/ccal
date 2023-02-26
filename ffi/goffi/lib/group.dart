class Group {
  String id;
  String name;
  String time;

  Group({required this.id, required this.name, required this.time});

  Group.fromJson(Map<String, dynamic> json)
      : id = json["id"],
        name = json["name"],
        time = json["time"];
}
