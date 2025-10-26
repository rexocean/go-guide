package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 1. 演示Print系列函数（单纯打印日志）
	fmt.Println("=== 1. Print系列函数演示 ===")
	printDemo()
	fmt.Println()

	// 2. 演示Panic系列函数（打印日志+抛panic）
	fmt.Println("=== 2. Panic系列函数演示 ===")
	//panicDemo()
	fmt.Println() // 此句不会执行，因为panic会终止程序，仅作示例标记

	// 3. 演示Fatal系列函数（打印日志+强制退出）
	// 注：Fatal会直接终止程序，为避免影响后续演示，建议单独注释测试
	fmt.Println("=== 3. Fatal系列函数演示 ===")
	//fatalDemo()
	fmt.Println() // 此句不会执行，因为Fatal会强制退出

	// 4. 演示日志配置（Flags）
	fmt.Println("=== 4. 日志Flags配置演示 ===")
	logFlagsDemo()
	fmt.Println()

	// 5. 演示日志前缀配置（Prefix）
	fmt.Println("=== 5. 日志前缀配置演示 ===")
	logPrefixDemo()
	fmt.Println()

	// 6. 演示日志输出到文件
	fmt.Println("=== 6. 日志输出到文件演示 ===")
	logToFileDemo()
	fmt.Println()

	// 7. 演示自定义Logger
	fmt.Println("=== 7. 自定义Logger演示 ===")
	customLoggerDemo()
	fmt.Println()
}

// 1. Print系列函数：单纯打印日志，不影响程序执行
func printDemo() {
	// Print：直接打印字符串
	log.Print("Print函数：这是一条基础日志\n")

	// Printf：支持格式化输出（类似fmt.Printf）
	userId := 1001
	userName := "Alice"
	log.Printf("Printf函数：用户ID=%d，用户名=%s\n", userId, userName)

	// Println：打印后自动换行，支持多参数
	log.Println("Println函数：参数1", "参数2", "参数3（自动换行）")
}

// 2. Panic系列函数：打印日志后抛出panic，后续代码不执行（defer会执行）
func panicDemo() {
	// defer语句在panic前声明，会在panic触发时执行
	defer fmt.Println("panic触发后，defer语句执行：程序即将终止")

	// Panic：打印日志后抛panic
	log.Panic("Panic函数：这是一条panic日志，程序将终止")

	// 以下代码不会执行，因为panic已终止程序
	fmt.Println("panic后：这条语句永远不会被执行")
}

// 3. Fatal系列函数：打印日志后强制退出（os.Exit(1)），defer不执行
func fatalDemo() {
	// defer语句不会执行，因为Fatal会直接调用os.Exit
	defer fmt.Println("Fatal触发后：这条defer语句不会执行")

	// Fatal：打印日志后强制退出程序，退出码为1
	log.Fatal("Fatal函数：这是一条fatal日志，程序将强制退出")

	// 以下代码不会执行，因为Fatal已强制终止程序
	fmt.Println("Fatal后：这条语句永远不会被执行")
}

// 4. 日志Flags配置：控制日志输出的细节（日期、时间、文件路径等）
func logFlagsDemo() {
	// 1. 获取当前日志的Flags配置（默认是LstdFlags = Ldate | Ltime）
	currentFlags := log.Flags()
	fmt.Printf("默认Flags值：%d（对应Ldate | Ltime）\n", currentFlags)

	// 2. 配置1：仅显示日期（Ldate）
	log.SetFlags(log.Ldate)
	log.Print("Flags=Ldate：仅显示日期的日志")

	// 3. 配置2：日期+时间+微秒（Ldate | Ltime | Lmicroseconds）
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Print("Flags=日期+时间+微秒：带微秒精度的日志")

	// 4. 配置3：日期+时间+短文件名+行号（Ldate | Ltime | Lshortfile）
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Print("Flags=日期+时间+短文件名+行号：带代码位置的日志")

	// 5. 配置4：日期+时间+长文件名+行号（Ldate | Ltime | Llongfile）
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Print("Flags=日期+时间+长文件名+行号：带完整路径的日志")

	// 恢复默认配置，避免影响后续演示
	log.SetFlags(log.LstdFlags)
}

// 5. 日志前缀配置：给所有日志添加固定前缀，方便筛选
func logPrefixDemo() {
	// 1. 获取当前日志前缀（默认是空字符串）
	currentPrefix := log.Prefix()
	fmt.Printf("默认日志前缀：【%s】（空字符串）\n", currentPrefix)

	// 2. 设置自定义前缀（例如添加"[INFO]"标记）
	log.SetPrefix("[INFO] ")
	log.Print("前缀为[INFO]：这是一条普通信息日志")

	// 3. 切换前缀（例如添加"[ERROR]"标记）
	log.SetPrefix("[ERROR] ")
	log.Print("前缀为[ERROR]：这是一条错误信息日志")

	// 4. 清空前缀（恢复默认）
	log.SetPrefix("")
	log.Print("清空前缀：恢复默认无前缀的日志")
}

// 6. 日志输出到文件：将日志从控制台转向文件存储
func logToFileDemo() {
	// 1. 打开日志文件：不存在则创建，存在则追加写入
	// os.O_CREATE：创建文件；os.O_WRONLY：只写；os.O_APPEND：追加模式
	// 0644：文件权限（所有者可读写，其他用户只读）
	logFile, err := os.OpenFile(
		"app.log", // 日志文件名
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		// 若文件打开失败，用panic日志提示（此处也可改用Print，但panic更醒目）
		log.Panicf("打开日志文件失败：%v", err)
	}

	// 2. 延迟关闭文件（确保程序退出前关闭，避免数据丢失）
	defer logFile.Close()

	// 3. 设置日志输出到文件（替代默认的控制台输出）
	log.SetOutput(logFile)

	// 4. 打印日志（此时日志会写入app.log，而非控制台）
	log.SetPrefix("[FILE_LOG] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Print("这是一条写入文件的日志：包含前缀、时间和代码位置")
	log.Printf("这是一条格式化文件日志：用户操作类型=%s，操作结果=%s", "登录", "成功")

	// 恢复默认输出（控制台），避免影响后续演示
	log.SetOutput(os.Stdout)
	log.SetPrefix("")
	log.SetFlags(log.LstdFlags)
	fmt.Println("日志已写入app.log文件，可在当前目录查看")
}

// 7. 自定义Logger：创建独立的Logger实例，不影响全局配置
func customLoggerDemo() {
	// 1. 打开自定义日志文件（单独的文件，与全局日志分离）
	customFile, err := os.OpenFile(
		"custom.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0644,
	)
	if err != nil {
		log.Panicf("打开自定义日志文件失败：%v", err)
	}
	defer customFile.Close()

	// 2. 创建自定义Logger：参数依次为（输出目标、前缀、Flags）
	// 这里配置：输出到custom.log + 前缀"[CUSTOM]" + 时间+短文件名
	customLogger := log.New(
		customFile,
		"[CUSTOM] ",
		log.Ltime|log.Lshortfile,
	)

	// 3. 使用自定义Logger打印日志（全局日志配置不受影响）
	customLogger.Println("自定义Logger：这是一条独立日志，写入custom.log")
	customLogger.Printf("自定义Logger：格式化日志，数值=%d，字符串=%s", 123, "test")

	// 4. 验证全局日志是否正常（不受自定义Logger影响）
	log.Println("全局Logger：自定义Logger不影响全局配置，此日志仍输出到控制台")

	fmt.Println("自定义日志已写入custom.log文件，可在当前目录查看")
}
