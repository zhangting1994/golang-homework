package main

import (
	"fmt"
	"time"
)

// 参考 Hystrix 实现一个滑动窗口计数器
// 这里只实现一个滑动窗口计数器，Hystrix略复杂，手动滑稽。

type rolling struct {
	gwindowSize int64 //窗口大小，毫秒为单位
	glimit int64 //窗口内限流大小
	gsplitNum int64 //切分小窗口的数目大小
	counters []int64 //每个小窗口的计数数组
	gindex int64 //当前小窗口计数器的索引
	startTime time.Time //窗口开始时间
}

func (r *rolling) init(windowSize int64, limit int64, splitNum int64) {
	r.gwindowSize = windowSize
	r.glimit = limit
	r.gsplitNum = splitNum
	r.counters = make([]int64, splitNum)
	r.gindex = 0
	r.startTime = time.Now()
}

func (r *rolling) slideWindow(windowsNum int64) {
	if(windowsNum != 0) {
		slideNum := min(windowsNum, r.gsplitNum)
		for i:=0; i<int(slideNum); i++ {
			r.gindex = (r.gindex + 1) % r.gsplitNum
			r.counters[r.gindex] = 0
		}
	
		tmp := windowsNum * r.gwindowSize / r.gsplitNum
		unix := r.startTime.Unix() + int64(tmp)
		r.startTime = time.Unix(unix, 0)
	}
}

func (r *rolling) IsValid() int {
	now := max(time.Now().Unix() - r.startTime.Unix() - r.gwindowSize, 0)
	tmp := r.gwindowSize / r.gsplitNum
	windowsNum := now / tmp
	r.slideWindow(windowsNum)

	var count int64
	count = 0

	for i:=0;i<int(r.gsplitNum);i++ {
		count += r.counters[i];
	}

	//超过限制则限流
    if(count >= r.glimit) {
		return 0
	} else {
        r.counters[r.gindex]++;
        return 1
    }
}

func min(a int64, b int64) int64 {
	if (a >= b) {
		return b
	} else {
		return a
	}
}

func max(a int64, b int64) int64 {
	if (a >= b) {
		return a
	} else {
		return b
	}
}

func main() {
	r := rolling{}
	r.init(1000, 10, 10) //1秒一个窗口，开10个，每个限制10个请求
	fmt.Println(r.IsValid()) //这里应该开多个携程验证
	// fmt.Println(r.startTime)
	// fmt.Println(time.Unix(0, 1603546715761482000))
	// fmt.Println(time.Unix(1603546715, 0))
}

