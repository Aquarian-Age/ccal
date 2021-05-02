package main

import "liangzi.local/ts/xjbfs"

//日建除
func jchh(ymc, mgz, dgz, hgz string) (djc, dhh, hhh string) {
	yjc := xjbfs.JC本月建除(ymc)
	djc = xjbfs.JC本日建除(dgz, yjc) //日建除
	zhimap := xjbfs.MakeHHMap()
	dhh = xjbfs.XJBF黄黑起法(mgz, dgz, zhimap) //本日黄黑神煞
	hhh = xjbfs.XJBF黄黑起法(dgz, hgz, zhimap) //时辰黄黑神煞
	return
}
