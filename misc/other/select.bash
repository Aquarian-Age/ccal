#!/bin/bash

### select 循环
PS3="输入要选择的数字显示相应的信息"

### 设置选择列表
select number in one two three four five six seven eight nine
do
	case $number in
		one)
			echo "111111111111"
			echo $number 是数字串中的第一个数字
		;;
		two)
			echo "222222222"
			echo $number 是数字串中的第二个数字
		;;
		three)
			echo "333-------"
			echo $number 是数字串中的第３个数字
		;;
		four)
			echo "444-----"
			echo $number 是数字串中的第４个数字
		;;
		five)
			echo "555------"
			echo $number 是数字串中的第５个数字
		;;
		six)
			echo "666---------"
			echo $number 是数字串中的第６个数字
		;;
		seven)
			echo "777---------"
			echo $number 是数字串中的第７个数字
		;;
		eight)
			echo "888----"
			echo "$number 是数字串的第８个数字"
		;;
		nine)
			echo "999----"
			echo "$number 是数字串的第9个数字"
		;;
		*)
			echo "Error try again(select 1..7)?"
		;;
	esac
done
	
