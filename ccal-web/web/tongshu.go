package main

import ts "liangzi.local/ts/tongshu"

//通书 日时总览方法
func Rszl(aliasmonth, dgz, rmc, aliasHour string) []string {
	return ts.RSZL日时总览(aliasmonth, dgz, rmc, aliasHour)
}

//通书择时要览
func Zsyl(dgz, hgz string) (tssj, tssx string) {
	return ts.ZSYL时辰吉凶(dgz, hgz)
}
