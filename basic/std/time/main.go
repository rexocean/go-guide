package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 1. Time类型基本用法演示
	fmt.Println("=== Time类型基本用法 ===")
	timeDemo()
	fmt.Println()

	// 2. 时间戳演示
	fmt.Println("=== 时间戳演示 ===")
	timestampDemo()
	fmt.Println()

	// 3. 时间解析演示
	fmt.Println("=== 时间解析演示 ===")
	timeParseDemo()
	fmt.Println()

	// 4. 时间格式化演示
	fmt.Println("=== 时间格式化演示 ===")
	timeFormatDemo()
	fmt.Println()

	// 5. 时间戳转Time演示
	fmt.Println("=== 时间戳转Time演示 ===")
	timestampDemo2(1627466760) // 2021-07-28 18:06:00的时间戳
	fmt.Println()

	// 6. 时间间隔与计算演示
	fmt.Println("=== 时间间隔与计算演示 ===")
	timeCalculationDemo()
	fmt.Println()

	// 7. 定时器演示
	fmt.Println("=== 定时器演示 ===")
	tickerAndTimerDemo()
}

// 1. Time类型基本用法
func timeDemo() {
	now := time.Now() // 获取当前时间
	fmt.Printf("当前时间: %v\n", now)

	// 获取时间的各个组成部分
	year := now.Year()     // 年
	month := now.Month()   // 月
	day := now.Day()       // 日
	hour := now.Hour()     // 小时
	minute := now.Minute() // 分钟
	second := now.Second() // 秒

	// 格式化输出，%02d表示不足两位补0
	fmt.Printf("格式化时间: %d-%02d-%02d %02d:%02d:%02d\n",
		year, month, day, hour, minute, second)

	// 时区转换示例
	fmt.Printf("UTC时间: %v\n", now.UTC())
	fmt.Printf("本地时间: %v\n", now.Local())

	// 检查时间是否为零值
	zeroTime := time.Time{}
	fmt.Printf("是否为零值时间: %v\n", zeroTime.IsZero())
}

// 2. 时间戳演示
func timestampDemo() {
	now := time.Now()               // 获取当前时间
	timestampSec := now.Unix()      // 秒级时间戳
	timestampNano := now.UnixNano() // 纳秒级时间戳

	fmt.Printf("当前秒级时间戳: %v\n", timestampSec)
	fmt.Printf("当前纳秒级时间戳: %v\n", timestampNano)
	fmt.Println("注: 时间戳以1970-01-01 00:00:00 UTC为起点")
}

// 3. 时间解析演示
func timeParseDemo() {
	// 使用Parse解析时间，默认时区为UTC
	t, err := time.Parse("2006-01-02 15:04:05", "2022-07-28 18:06:00")
	if err != nil {
		fmt.Printf("解析错误: %v\n", err)
		return
	}
	fmt.Printf("Parse解析结果(UTC): %v, 时区: %v\n", t, t.Location())

	// 加载上海时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("加载时区错误: %v\n", err)
		return
	}

	// 使用ParseInLocation解析指定时区的时间
	nowStr := time.Now().Format("2006/01/02 15:04:05")
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", nowStr, loc)
	if err != nil {
		fmt.Printf("ParseInLocation解析错误: %v\n", err)
		return
	}
	fmt.Printf("ParseInLocation解析结果(上海时区): %v, 时区: %v\n", timeObj, timeObj.Location())
}

// 4. 时间格式化演示
func timeFormatDemo() {
	now := time.Now()
	// 格式化模板必须使用Go诞生时间: 2006-01-02 15:04:05
	// 24小时制格式
	fmt.Println("24小时制带毫秒:", now.Format("2006-01-02 15:04:05.000"))
	// 12小时制格式
	fmt.Println("12小时制:", now.Format("2006-01-02 03:04:05 PM"))
	// 其他常用格式
	fmt.Println("年月日:", now.Format("2006/01/02"))
	fmt.Println("时分秒:", now.Format("15:04:05"))
	fmt.Println("自定义格式:", now.Format("02-01-2006 15时04分05秒"))
}

// 5. 时间戳转Time演示
func timestampDemo2(timestamp int64) {
	// 将时间戳转为Time类型（秒级时间戳）
	timeObj := time.Unix(timestamp, 0)
	fmt.Printf("时间戳 %d 对应的时间: %v\n", timestamp, timeObj)

	// 获取转换后时间的各个组成部分
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()

	fmt.Printf("分解后: %d-%02d-%02d %02d:%02d:%02d\n",
		year, month, day, hour, minute, second)

	// 带时区的转换示例
	loc, _ := time.LoadLocation("Asia/Shanghai")
	chinaTime := time.Unix(timestamp, 0).In(loc)
	fmt.Printf("上海时区时间: %v\n", chinaTime)
}

// 6. 时间间隔与计算演示
func timeCalculationDemo() {
	now := time.Now()
	fmt.Printf("当前时间: %v\n", now)

	// Add: 时间加法
	oneHourLater := now.Add(1 * time.Hour)
	fmt.Printf("1小时后: %v\n", oneHourLater)

	thirtyMinutesAgo := now.Add(-30 * time.Minute)
	fmt.Printf("30分钟前: %v\n", thirtyMinutesAgo)

	// Sub: 计算时间差
	diff := oneHourLater.Sub(now)
	fmt.Printf("时间差: %v\n", diff)
	fmt.Printf("时间差(秒): %v\n", diff.Seconds())

	// Equal: 判断时间是否相等（考虑时区）
	time1 := time.Date(2023, 10, 1, 0, 0, 0, 0, time.UTC)
	time2 := time.Date(2023, 10, 1, 8, 0, 0, 0, time.FixedZone("CST", 8*3600))
	fmt.Printf("两个时间是否表示同一时刻: %v\n", time1.Equal(time2))

	// Before 和 After
	fmt.Printf("当前时间是否在1小时后之前: %v\n", now.Before(oneHourLater))
	fmt.Printf("当前时间是否在30分钟前之后: %v\n", now.After(thirtyMinutesAgo))
}

// 7. 定时器演示
func tickerAndTimerDemo() {
	var wg sync.WaitGroup
	wg.Add(3) // 等待3个goroutine完成

	// 1. 基本定时器 (Tick)
	go func() {
		defer wg.Done()
		fmt.Println("=== 基本定时器开始 (3秒后结束) ===")
		ticker := time.Tick(1 * time.Second) // 1秒间隔的定时器
		count := 0
		for t := range ticker {
			fmt.Printf("Tick触发: %v\n", t.Format("15:04:05"))
			count++
			if count >= 3 { // 触发3次后退出
				break
			}
		}
	}()

	// 2. AfterFunc 定时器
	go func() {
		defer wg.Done()
		fmt.Println("\n=== AfterFunc定时器开始 ===")
		fmt.Println("等待2秒后执行...")
		// 2秒后执行匿名函数
		time.AfterFunc(2*time.Second, func() {
			fmt.Println("AfterFunc: 2秒时间到!")
		})
		// 等待足够长的时间让AfterFunc执行
		time.Sleep(3 * time.Second)
	}()

	// 3. 对比Timer和Ticker
	go func() {
		defer wg.Done()
		fmt.Println("\n=== Timer与Ticker对比 (5秒后结束) ===")
		timer := time.NewTimer(1 * time.Second)
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop() // 停止ticker释放资源

		count := 0
		for {
			select {
			case t := <-timer.C:
				fmt.Printf("Timer触发: %v\n", t.Format("15:04:05"))
				timer.Reset(1 * time.Second) // 重置timer实现循环触发
			case t := <-ticker.C:
				fmt.Printf("Ticker触发: %v\n", t.Format("15:04:05"))
			}

			count++
			if count >= 5 {
				break
			}
		}
	}()

	wg.Wait()
	fmt.Println("\n所有定时器演示结束")
}
