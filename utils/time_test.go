package utils

import (
	"fmt"  
	"time"
	"sort"
	"testing"
)
func TestTime(tttt *testing.T) {

	fmt.Println("获取当前时间字符串 => ",GetCurrTstr(DATETIME_FORMAT))
	fmt.Println("获取当前时间戳 => ", GetCurrTs())

	t := time.Now()
	fmt.Println("时间转字符串 => ",TimeToStr(t,DATETIME_FORMAT))
	fmt.Println("时间转字符串 => ",TimeToStr(t,DATE_FORMAT))
	fmt.Println("时间戳转字符串 => ", TimestampToStr(GetCurrTs(), DATETIME_FORMAT))

	//返回当前是一年中的第几天  
	//select to_char(sysdate,'ddd'),sysdate from dual;  
	yd := t.YearDay();  
	fmt.Println("一年中的第几天: ",yd)         

	//一年中的第几周  
	year,week := t.ISOWeek()  
	fmt.Println("一年中的第几周: ",year," | ",week)         

	//当前是周几  
	//select to_char(sysdate,'day') from dual;  
	//select to_char(sysdate,'day','NLS_DATE_LANGUAGE = American') from dual;     
	fmt.Println("当前是周几: ",t.Weekday().String())  

	// 字符串转成time.Time     
	tt,er := time.Parse(GoStdTime(),"2018-01-01 01:01:01")  
	if(er != nil){  
			fmt.Println("字符串转时间: parse error!")  
	}else{  
			fmt.Println("字符串转时间: ",tt.String())  
	}

	fmt.Println("\n演示时间 => ",TimeToStr(t,"y-m-d h:i:s"))  
		
	ta := t.AddDate(1,0,0)     
	fmt.Println("增加一年 => ",TimeToStr(ta,"y-m-d"))  
	
	ta = t.AddDate(0,1,0)  
	fmt.Println("增加一月 => ",TimeToStr(ta,"y-m-d"))  

	//select sysdate,sysdate + interval '1' day from dual;  
	ta = t.AddDate(0,0,1) //18  
	fmt.Println("增加一天 => ",TimeToStr(ta,"y-m-d"))  

	durdm,_ := time.ParseDuration("432h")  
	ta = t.Add(durdm)  
	fmt.Println("增加18天(18*24=432h) => ",TimeToStr(ta,"y-m-d"))  
	
	//select sysdate,sysdate - interval '7' hour from dual;  
	dur,_ := time.ParseDuration("-2h")  
	ta = t.Add(dur)  
	fmt.Println("减去二小时 => ",TimeToStr(ta,"y-m-d h:i:s"))  

	//select sysdate,sysdate - interval '7' MINUTE from dual;  
	durmi,_ := time.ParseDuration("-7m")  
	ta = t.Add(durmi)  
	fmt.Println("减去7分钟 => ",TimeToStr(ta,"y-m-d h:i:s"))  

	//select sysdate,sysdate - interval '10' second from dual;  
	durs,_ := time.ParseDuration("-10s")  
	ta = t.Add(durs)  
	fmt.Println("减去10秒 => ",TimeToStr(ta,"y-m-d h:i:s"))  

	ttr,er := time.Parse(GoStdTime(),"2014-06-09 16:58:06")  
	if(er != nil){  
			fmt.Println("字符串转时间: 转换失败!")  
	}else{  
			fmt.Println("字符串转时间: ",ttr.String())  
	}  

	//alter session set nls_date_format='yyyy-mm-dd hh24:mi:ss';  
	//select trunc(to_date('2014-06-09 16:58:06','yyyy-mm-dd hh24:mi:ss'),'mi') as dt from dual;   
	// SQL => 2014-06-09 16:58:00  
	// Truncate =>  2014-06-09 16:50:00  
	durtr,_ := time.ParseDuration("10m")  
	ta = ttr.Truncate(durtr)  
	fmt.Println("Truncate => ",TimeToStr(ta,"y-m-d H:i:s"))  

	//select round(to_date('2014-06-09 16:58:06','yyyy-mm-dd hh24:mi:ss'),'mi') as dt from dual;   
	// SQL => 2014-06-09 16:58:00  
	// Round =>  2014-06-09 17:00:00  
	ta = ttr.Round(durtr)  
	fmt.Println("Round => ",TimeToStr(ta,"y-m-d H:i:s"))  

	//日期比较  
	tar1,_ := time.Parse(GoStdTime(),"2014-06-09 19:38:36")  
	tar2,_ := time.Parse(GoStdTime(),"2015-01-14 17:08:26")  
	if tar1.After(tar2) {  
			fmt.Println("tar1 > tar2")  
	}else if tar1.Before(tar2) {  
			fmt.Println("tar1 < tar2")  
	}else{  
			fmt.Println("tar1 = tar2")  
	}  
	tar3,_ := time.Parse(GoStdTime(),"2000-07-19 15:58:16")  

	//日期列表中最晚日期  
	// select greatest('2014-06-09','2015-01-14','2000-07-19') from dual;      
	var arr TimeSlice  
	arr = []time.Time{tar1,tar2,tar3}  
	temp := Greatest(arr)  
	fmt.Println("日期列表中最晚日期 => ",TimeToStr(temp,"y-m-d"))      

	//日期数组从早至晚排序  
	fmt.Println("\n日期数组从早至晚排序")  
	sort.Sort(arr)  
	for _,at := range arr {  
				fmt.Println("Sort => ",TimeToStr(at,"y-m-d H:i:s"))  
	}
}
