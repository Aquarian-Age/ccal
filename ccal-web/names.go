package main

import (
	"time"

	"liangzi.local/nongli/ganzhi"
	"liangzi.local/ts/xjbfs"
)

var (
	T    = time.Now().Local() //当前时间
	jz60 = ganzhi.MakeJZ60()  //六十甲子
	i    xjbfs.ZRYLImpl       //协纪辩方择日方法
	zryl *xjbfs.ZRYL
)

//应答数据
type Resp struct {
	TodayS   string   `json:"today"`        //今日农历日期
	JiNian   string   `json:"jinianInfo"`   //纪年信息
	Dmj      string   `json:"dmInfo"`       //地母经
	Jq       []string `json:"jqInfo"`       //24节气
	ListDay  []string `json:"listdayInfo"`  //农历月历表(农历初一开始)
	StarName string   `json:"starnameInfo"` //当日值宿名称(28宿)
	StarInfo string   `json:"starInfo"`     //值宿信息
	Zeji     string   `json:"zejiInfo"`     //择吉 本日星数字
	JiGan    string   `json:"jigan"`        //择吉 本月吉干
	JiRi     []string `json:"goodjg"`       //择吉 本月吉日数组
	QiSha    []string `json:"qisha"`        //择吉 本月七煞数组
	XJBF              //协纪辩方书
	TSInfo            //通书内容
	YJ                //月将信息
	SJQM     string   `json:"sjqm"` //时家奇门
	Gdtm     string   `json:"dtms"` //贵登天门(协纪辩方书义例六)
	Other
	About
}

//干支查询页面
type GZS struct {
	Ygz, Mgz, Dgz, Hgz string
	JC                 string   //日建除
	Dhh                string   //黄黑神煞(日)
	Hhh                string   //黄黑神煞(时辰)
	Tscy               string   //太岁出游日
	Tygr               string   //太岁天乙贵人
	TYDH               string   //天乙贵人 日时论
	RiLu               string   //日禄
	GuGua              string   //孤辰寡宿
	WuLu               string   //无禄日（十恶大败日）
	ChongRi, FuRi      string   //重日 复日
	YiJu               string   //移居吉日 含满成开日
	TanBing            string   //忌探病日
	XianChi            string   //咸池 桃花
	ShangShuo          string   //上朔日
	Wbu                string   //奇门五不遇
	GuXu               []string //奇门时孤虚法
	LiuChong           string   //六冲
	SanSha             string   //三煞
	TaiSui             string   //值年太岁
	JinFu              string   //金符九星
	JinShen            string   //金符经 本月金神七煞日
}

//择日信息
type ZR struct {
	syear  int       //阳历年
	smonth int       //阳历月
	sday   int       //阳历日
	sweek  string    //阳历周几
	stime  time.Time //当日阳历时间戳　精确到日

	lyear  int //农历年
	lmonth int //农历月
	lday   int //农历日

	aliasyg    string //农历年干
	aliasyz    string //年支
	aliasygz   string //年干支
	aliasmgz   string //月干支
	aliasday   string //日别名　一　二　廿二
	aliasmonth string //月别名 正　二　三
	dgz        string //日干支
	daygan     string //日干
	hourgz     string //时干支
	hourz      string //时支
	leapmb     bool   //闰月
	ydx        string //月大小数组
	aliasHour  string //时辰(地支)
	hour       int    //时子:1 丑:2．．．
	leapmonth  int    //闰月
}

//协纪辩方书
type XJBF struct {
	Ji          []string `json:"suiji"`    //岁吉
	TaiSuiWuGui []string `json:"taisui"`   //太岁五鬼
	Sha         []string `json:"suisha"`   //岁煞
	Xiong       []string `json:"suixiong"` //岁凶
	Yjh         []string `json:"xjbfsyjh"` //月总论 日建除 黄黑道
	Rj          []string `json:"xjbfsrj"`  //月日论 日吉
	Rx          []string `json:"xjbfsrx"`  //月日论 日凶
	Hgx         []string `json:"xjbfshgx"` //时辰黄黑 时辰孤 虚
	Hcj         []string `json:"xjbfshcj"` //日时论 时辰吉
	Hcx         []string `json:"xjbfshcx"` //日时论 时辰凶
	BW          string   `json:"xjbfsBW"`  //协纪辩方书 辩伪
}

//通书内容
type TSInfo struct {
	TSRs []string `json:"trszl"` //通书 日时总览
	TSsj string   `json:"tsj"`   //通书 择时要览时辰吉
	TSsx string   `json:"tsx"`   //通书 择时要览时辰凶
}

//关于...
type About struct {
	Ccal string `json:"accal"`
	Data string `json:"ada"`
	Xlr  string `json:"axlr"`
	Xjbf string `json:"axjbf"`
	Ck   string `json:"ack"`
	Me   string `json:"ame"`
}

//月将
type YJ struct {
	YjName   string `json:"yjName"`   //月将名称
	Zhi      string `json:"yjzhi"`    //月将地支
	StarName string `json:"starName"` //十二星宫
	JieQi    string `json:"jq"`       //节气
}

//其他内容
type Other struct {
	JinFu     string   `json:"jinfu"` //金符九星
	JinShenQS []string `json:"jsqs"`  //金神七煞 这里会有出现空值
}
