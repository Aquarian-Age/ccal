package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"liangzi.local/nongli/ganzhi"
	"liangzi.local/ts/jfj"
	"liangzi.local/ts/qimen"
	ts "liangzi.local/ts/tongshu"
	"liangzi.local/ts/xjbfs"
)

//干支查询页面
func selectlist(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("jz60.html")
		if err != nil {
			log.Fatal("t-err:", err)
		}
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		y, err := strconv.Atoi(r.Form["jzy"][0]) //这里的jzy 是select的name值
		if err != nil {
			log.Fatal("y-err:", err)
		}
		m, err := strconv.Atoi(r.Form["jzm"][0])
		if err != nil {
			log.Fatal("m-err:", err)
		}
		d, err := strconv.Atoi(r.Form["jzd"][0])
		if err != nil {
			log.Fatal("d-err:", err)
		}
		h, err := strconv.Atoi(r.Form["jzh"][0])
		if err != nil {
			log.Fatal("h-err", err)
		}
		///太岁出游日
		var ygz, mgz, dgz, hgz string
		for i := 0; i < 60; i++ {
			if i == y {
				ygz = jz60[i]
			}
			if i == m {
				mgz = jz60[i]
			}
			if i == d {
				dgz = jz60[i]
				//	break
			}
			if i == h {
				hgz = jz60[i]
			}
		}

		ymc := gztoYmc(mgz) //月名称 正 二 三...十一 十二
		//////
		日建除, 日黄黑, 时辰黄黑 := jchh(ymc, mgz, dgz, hgz)
		太岁出游 := xjbfs.XJBF太岁出游(dgz, jz60)
		阳贵人, 阴贵人 := xjbfs.XJBF太岁天乙贵人(ygz)
		天乙贵人 := xjbfs.XJBF天乙贵人(dgz, hgz)
		rilu := ganzhi.GZ禄(dgz)
		rilu = fmt.Sprintf("%s禄在 %s", dgz, rilu)
		孤辰, 寡宿 := xjbfs.XJBF孤辰寡宿(ygz, mgz, dgz, hgz)
		无禄 := xjbfs.XJBF无禄日(dgz)
		重日 := xjbfs.XJBF重日(dgz)
		复日 := xjbfs.XJBF复日(mgz, dgz)
		移居吉日 := ts.RSZL移居吉日(ymc, dgz)
		忌探病日 := ts.RSZL忌探病日(dgz)
		咸池 := xjbfs.XCTH咸池桃花(ygz, mgz, dgz, hgz)
		上朔日 := xjbfs.XJBF上朔(ygz, dgz)
		五不遇 := qimen.SJQM五不遇(dgz, hgz)
		旬孤虚信息, 孤虚map := qimen.SJQM孤虚法(dgz, hgz)
		var qmgx []string
		qmgx = append(qmgx, 旬孤虚信息)
		qmgx = append(qmgx, 孤虚map["孤"])
		qmgx = append(qmgx, 孤虚map["虚"])
		//方位
		guzhi := 孤虚map["孤"]
		xuzhi := 孤虚map["虚"]
		gzinfo := ganzhi.NewZHI(guzhi)
		xzinfo := ganzhi.NewZHI(xuzhi)
		gfx := gzinfo.FangXiang
		xfx := xzinfo.FangXiang
		qmgx = append(qmgx, gfx)
		qmgx = append(qmgx, xfx)
		liuchong := ganzhi.DZ六冲(dgz)
		liuchong = fmt.Sprintf("%s冲 %s", dgz, liuchong)
		sansha := ganzhi.GZ三煞(dgz)
		sansha = fmt.Sprintf("%s日煞在:%s", dgz, sansha)
		//太岁
		zhiTS := xjbfs.ZhiTaiSui(ygz)
		chTS := xjbfs.ChongTaiSui(ygz)
		haiTS := xjbfs.HaiTaiSui(ygz)
		xingTS := xjbfs.XingTaiSui(ygz)
		taisui := ygz + "年" + zhiTS + " " + chTS + " " + haiTS + " " + xingTS

		//金符九星
		jfs := jfj.JinFuJing(ymc, dgz)
		//金符经 金神七煞日
		_, qsd := jfj.QiShaDay(ygz, dgz)

		gzs := GZS{
			Ygz:       ygz,
			Mgz:       mgz,
			Dgz:       dgz,
			Hgz:       hgz,
			JC:        日建除,
			Dhh:       日黄黑,
			Hhh:       时辰黄黑,
			Tscy:      太岁出游,
			Tygr:      阳贵人 + " " + 阴贵人,
			TYDH:      天乙贵人,
			RiLu:      rilu,
			GuGua:     孤辰 + " " + 寡宿,
			WuLu:      无禄,
			ChongRi:   重日,
			FuRi:      复日,
			YiJu:      移居吉日,
			TanBing:   忌探病日,
			XianChi:   咸池,
			ShangShuo: 上朔日,
			Wbu:       五不遇,
			GuXu:      qmgx,
			LiuChong:  liuchong,
			SanSha:    sansha,
			TaiSui:    taisui,
			JinFu:     jfs,
			JinShen:   qsd,
		}
		json.NewEncoder(w).Encode(gzs)
	}
}
