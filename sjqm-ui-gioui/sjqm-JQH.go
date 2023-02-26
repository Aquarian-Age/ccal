/*
 * Created by GoLand
 * User: Amrta
 * Mail: liangzi2021@yandex.com
 * Date:  2021年 7月 26日
 *
 */

package main

import (
	"fmt"
	"github.com/Aquarian-Age/xa/pkg/pub"
	"log"
	"strconv"
	"strings"
	"time"

	"gioui.org/widget"
	"github.com/Aquarian-Age/sjqm/v3"
	"github.com/Aquarian-Age/xa/pkg/gz"
)

var (
	t = time.Now().Local()
)

func convEdit(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Println("convEdit:", err)
	}
	return i
}

var number = map[int]string{
	1: `1,6 1,6`,
	8: `8,7,5,10`,
	3: `3,4,3,8`,
	4: `4,5,3,8`,
	9: `9,3,2,7`,
	2: `2,8,5,10`,
	7: `7,2,4,9`,
	6: `6,1,4,9`,
	5: `5`,
}

type Res struct {
	ed1, ed2, ed3, ed4, ed5 widget.Editor
	btn1                    widget.Clickable
	btna                    widget.Clickable

	qm   *sjqm.Sjqm
	info string

	baShen  []string
	jiuXing []string
	baMen   []string
	tianQi  []string
	diQi    []string
	anGan   []string
	g5s     string

	bg1, bg8, bg3, bg4, bg9, bg2, bg7, bg6 widget.Clickable
	btnGxs                                 string
}

func SJQM(tx time.Time) *Res {
	res := &Res{}

	res.ed1 = widget.Editor{}
	res.ed2 = widget.Editor{}
	res.ed3 = widget.Editor{}
	res.ed4 = widget.Editor{}
	res.ed5 = widget.Editor{}
	res.btn1 = widget.Clickable{}
	res.btna = widget.Clickable{}
	//
	res.bg1 = widget.Clickable{}
	res.bg8 = widget.Clickable{}
	res.bg3 = widget.Clickable{}
	res.bg4 = widget.Clickable{}
	res.bg9 = widget.Clickable{}
	res.bg2 = widget.Clickable{}
	res.bg7 = widget.Clickable{}
	res.bg6 = widget.Clickable{}

	qm := sjqm.NewJieQiHourSjqm(tx)
	res.qm = qm

	yjo := qm.Gzo.YueJiangStruct()
	dan, mu := yjo.GuiDengTianMen(qm.Dgz)
	var tms string
	if dan != "" {
		tms = "天门:" + dan
	}
	if mu != "" {
		tms = "天门:" + mu
	}
	if dan != "" && mu != "" {
		tms = "天门:" + dan + mu
	}

	solars := fmt.Sprintf("%s ", tx.String()[:19])
	lunar, _ := qm.Gzo.GetLunar()
	solars += lunar + " JQH" + "\n"
	gzs := fmt.Sprintf("%s %s %s %s ", qm.Ygz, qm.Mgz, qm.Dgz, qm.Hgz)
	xk1, xk2, xk3, xk4 := qm.XunKong()
	xks := fmt.Sprintf("%s %s %s %s ", xk1, xk2, xk3, xk4)

	jus := qm.JuString[7:] + " " + qm.XunShouDunString[7:] + "\n"
	xunshous := fmt.Sprintf("值符:%s%d宫 值使:%s%d宫", qm.ZhiFu, qm.ZhiFuGn, qm.ZhiShi, qm.ZhiShiGn)
	//
	var others string
	shiGanRuMus := qm.ShiGanRuMu()
	wubuyus := qm.WuBuYu()
	tianxians := qm.TianXianShiGe()
	others = "\n" + shiGanRuMus + wubuyus + tianxians
	infos := solars + gzs + jus + xks + xunshous + others

	res.info = infos

	var bsarr, stararr, tianqiarr, bamenarr, diqiarr, angarr []string
	bsarr = nil
	stararr = nil
	tianqiarr = nil
	bamenarr = nil
	diqiarr = nil
	angarr = nil

	yjs := "月将:" + yjo.Zhi + "\n"
	tianmas := yjo.TaiChongTianMa(qm.Hgz)
	tianms := "天马:" + tianmas
	res.g5s = yjs + tianms

	res.baShen = bsarr
	res.jiuXing = stararr
	res.baMen = bamenarr
	res.tianQi = tianqiarr
	res.diQi = diqiarr
	res.anGan = angarr

	return res
}
func (res *Res) G1s() string {
	return "g1s"
}
func (res *Res) G8s() string {

	return "g8s"
}
func (res *Res) G3s() string {

	return "g3s"
}
func (res *Res) G4s() string {

	return "g4s"
}
func (res *Res) G9s() string {

	return "g9s"
}
func (res *Res) G2s() string {

	return "g2s"
}
func (res *Res) G7s() string {

	return "g7s"
}
func (res *Res) G6s() string {

	return "g6s"
}
func (res *Res) G5s() string {

	return "g5s"
}
