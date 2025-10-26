package main

import (
	"bytes"
	"fmt"
	"unicode"
)

func main() {
	// 1. 字节切片转换函数（大小写/标题格式）
	fmt.Println("=== 1. 字节切片转换函数 ===")
	byteConvertDemo()
	fmt.Println()

	// 2. 字节切片比较函数（相等/大小/忽略大小写）
	fmt.Println("=== 2. 字节切片比较函数 ===")
	byteCompareDemo()
	fmt.Println()

	// 3. 字节切片清理函数（trim前缀/后缀/空白）
	fmt.Println("=== 3. 字节切片清理函数 ===")
	byteTrimDemo()
	fmt.Println()

	// 4. 字节切片拆合函数（分割/连接/重复）
	fmt.Println("=== 4. 字节切片拆合函数 ===")
	byteSplitJoinDemo()
	fmt.Println()

	// 5. 字节切片子串判断函数（包含/前缀/后缀/计数）
	fmt.Println("=== 5. 字节切片子串判断函数 ===")
	byteSubstringDemo()
	fmt.Println()

	// 6. 字节切片替换函数（替换指定内容/字符映射）
	fmt.Println("=== 6. 字节切片替换函数 ===")
	byteReplaceDemo()
	fmt.Println()

	// 7. Buffer类型核心操作（声明/读写/容量管理）
	fmt.Println("=== 7. Buffer类型核心操作 ===")
	bufferDemo()
	fmt.Println()

	// 8. Reader类型核心操作（读取/定位/重置）
	fmt.Println("=== 8. Reader类型核心操作 ===")
	readerDemo()
	fmt.Println()
}

// 1. 字节切片转换函数：ToUpper/ToLower/ToTitle
func byteConvertDemo() {
	original := []byte("Hello World! 123 世界")
	fmt.Printf("原始内容: %q\n", original)

	// 转大写
	toUpper := bytes.ToUpper(original)
	fmt.Printf("转大写:   %q\n", toUpper) // "HELLO WORLD! 123 世界"（中文不受影响）

	// 转小写
	toLower := bytes.ToLower(original)
	fmt.Printf("转小写:   %q\n", toLower) // "hello world! 123 世界"

	// 转标题格式（每个单词首字母大写）
	toTitle := bytes.ToTitle(original)
	fmt.Printf("转标题:   %q\n", toTitle) // "HELLO WORLD! 123 世界"

	// 原切片不会被修改（转换函数返回新切片）
	// 修正：使用 bytes.Equal() 比较两个字节切片的内容，
	// 在 Go 中不能直接用 == 比较两个字节切片
	fmt.Printf("原切片内容是否与转大写后相同: %t\n", bytes.Equal(original, toUpper)) // false
	// 额外演示：原切片自身内容未被修改
	fmt.Printf("原始切片内容是否保持不变: %q\n", original) // 仍为 "Hello World! 123 世界"
}

// 2. 字节切片比较函数：Compare/Equal/EqualFold
func byteCompareDemo() {
	b1 := []byte("abc123")
	b2 := []byte("abc123")
	b3 := []byte("abc456")
	b4 := []byte("ABC123")
	b5 := []byte("Φφϕ") // 希腊字母，大小写形态不同但属于同类字符
	b6 := []byte("ϕΦφ")

	// Equal：严格比较（大小写、顺序完全一致才相等）
	fmt.Printf("b1 == b2 (Equal): %t\n", bytes.Equal(b1, b2)) // true
	fmt.Printf("b1 == b4 (Equal): %t\n", bytes.Equal(b1, b4)) // false

	// EqualFold：忽略大小写，判断字符"语义相似"
	fmt.Printf("b1 == b4 (EqualFold): %t\n", bytes.EqualFold(b1, b4)) // true
	fmt.Printf("b5 == b6 (EqualFold): %t\n", bytes.EqualFold(b5, b6)) // true

	// Compare：按字节ASCII码比较大小（返回-1/0/1）
	fmt.Printf("b1 vs b3 (Compare): %d\n", bytes.Compare(b1, b3)) // -1（b1的'1' < b3的'4'）
	fmt.Printf("b3 vs b1 (Compare): %d\n", bytes.Compare(b3, b1)) // 1
	fmt.Printf("b1 vs b2 (Compare): %d\n", bytes.Compare(b1, b2)) // 0
}

// 3. 字节切片清理函数：Trim/TrimPrefix/TrimSpace等
func byteTrimDemo() {
	original := []byte("  !!Hello World! 123..  ")
	fmt.Printf("原始内容: %q\n", original)

	// TrimSpace：清理前后空白（空格、制表符、换行等）
	trimSpace := bytes.TrimSpace(original)
	fmt.Printf("清理空白: %q\n", trimSpace) // "!!Hello World! 123.."

	// Trim：清理前后指定字符集（cutset为"!."）
	trimCutset := bytes.Trim(trimSpace, "!.")
	fmt.Printf("清理!和.: %q\n", trimCutset) // "Hello World! 123"（中间的!不清理）

	// TrimPrefix/TrimSuffix：清理指定前缀/后缀
	withPrefix := []byte("Prefix_Hello")
	withSuffix := []byte("Hello_Suffix")
	trimPre := bytes.TrimPrefix(withPrefix, []byte("Prefix_"))
	trimSuf := bytes.TrimSuffix(withSuffix, []byte("_Suffix"))
	fmt.Printf("清理前缀: %q\n", trimPre) // "Hello"
	fmt.Printf("清理后缀: %q\n", trimSuf) // "Hello"

	// TrimFunc：按自定义函数清理（清理非字母字符）
	trimFunc := bytes.TrimFunc(trimCutset, func(r rune) bool {
		return !unicode.IsLetter(r) // 返回true表示该字符需要被清理
	})
	fmt.Printf("清理非字母: %q\n", trimFunc) // "HelloWorld"
}

// 4. 字节切片拆合函数：Split/Join/Repeat/Fields
func byteSplitJoinDemo() {
	original := []byte("  Hello,World,Golang  ")
	fmt.Printf("原始内容: %q\n", original)

	// Split：按分隔符分割（不包含分隔符）
	splitComma := bytes.Split(original, []byte(","))
	fmt.Printf("按,分割:   %q\n", splitComma) // ["  Hello", "World", "Golang  "]

	// SplitN：指定分割次数（n=2表示只分割前2个分隔符）
	splitN := bytes.SplitN(original, []byte(","), 2)
	fmt.Printf("按,分割2次: %q\n", splitN) // ["  Hello", "World,Golang  "]

	// SplitAfter：分割后包含分隔符
	splitAfter := bytes.SplitAfter(original, []byte(","))
	fmt.Printf("分割后含,: %q\n", splitAfter) // ["  Hello,", "World,", "Golang  "]

	// Fields：按连续空白分割（自动清理前后空白）
	fields := bytes.Fields(original)
	fmt.Printf("按空白分割: %q\n", fields) // ["Hello,World,Golang"]

	// FieldsFunc：按自定义函数分割（按非字母分割）
	fieldsFunc := bytes.FieldsFunc(original, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	fmt.Printf("按非字母分割: %q\n", fieldsFunc) // ["Hello", "World", "Golang"]

	// Join：将切片用指定分隔符连接
	joined := bytes.Join(fieldsFunc, []byte("-"))
	fmt.Printf("连接结果:   %q\n", joined) // "Hello-World-Golang"

	// Repeat：重复字节切片（重复3次）
	repeated := bytes.Repeat([]byte("Go"), 3)
	fmt.Printf("重复3次:   %q\n", repeated) // "GoGoGo"
}

// 5. 字节切片子串判断函数：Contains/HasPrefix/HasSuffix/Count
func byteSubstringDemo() {
	original := []byte("Hello Golang! Go is good.")
	fmt.Printf("原始内容: %q\n", original)

	// HasPrefix/HasSuffix：判断前缀/后缀
	fmt.Printf("含前缀Hello: %t\n", bytes.HasPrefix(original, []byte("Hello"))) // true
	fmt.Printf("含后缀good.: %t\n", bytes.HasSuffix(original, []byte("good."))) // true

	// Contains：判断是否包含子串
	fmt.Printf("含子串Go:   %t\n", bytes.Contains(original, []byte("Go")))   // true
	fmt.Printf("含子串Java: %t\n", bytes.Contains(original, []byte("Java"))) // false

	// ContainsRune：判断是否包含指定字符（支持中文）
	fmt.Printf("含字符'好': %t\n", bytes.ContainsRune(original, '好')) // true

	// ContainsAny：判断是否包含字符集中任意一个字符
	fmt.Printf("含x/y/z/G: %t\n", bytes.ContainsAny(original, "xyzG")) // true（含G）

	// Index/LastIndex：查找子串第一次/最后一次出现位置
	fmt.Printf("Go第一次位置: %d\n", bytes.Index(original, []byte("Go")))      // 6
	fmt.Printf("Go最后一次位置: %d\n", bytes.LastIndex(original, []byte("Go"))) // 14

	// Count：统计子串出现次数（不重叠）
	fmt.Printf("Go出现次数: %d\n", bytes.Count(original, []byte("Go"))) // 2
}

// 6. 字节切片替换函数：Replace/Map/Runes
func byteReplaceDemo() {
	original := []byte("Hello World! World is big.")
	fmt.Printf("原始内容: %q\n", original)

	// Replace：替换子串（n=-1表示全部替换）
	replace1 := bytes.Replace(original, []byte("World"), []byte("Golang"), 1)
	replaceAll := bytes.Replace(original, []byte("World"), []byte("Golang"), -1)
	fmt.Printf("替换1次:   %q\n", replace1)   // "Hello Golang! World is big."
	fmt.Printf("替换全部:   %q\n", replaceAll) // "Hello Golang! Golang is big."

	// Map：按自定义函数映射字符（小写转大写）
	mapUpper := bytes.Map(func(r rune) rune {
		if unicode.IsLower(r) {
			return unicode.ToUpper(r)
		}
		return r
	}, original)
	fmt.Printf("小写转大写: %q\n", mapUpper) // "HELLO WORLD! WORLD IS BIG."

	// Runes：字节切片转rune切片（处理多字节字符）
	chinese := []byte("你好，世界！")
	runes := bytes.Runes(chinese)
	fmt.Printf("字节切片长度: %d\n", len(chinese)) // 13（每个中文字符占3字节，标点占1字节）
	fmt.Printf("rune切片长度: %d\n", len(runes)) // 5（每个字符对应1个rune）
	fmt.Printf("rune内容:     %q\n", runes)    // ['你','好',',','世','界','！']
}

// 7. Buffer类型核心操作：声明/读写/容量/重置
func bufferDemo() {
	// 1. Buffer的4种声明方式
	var buf1 bytes.Buffer // 直接声明（零值可用）
	//buf2 := new(bytes.Buffer)                // new创建指针
	buf3 := bytes.NewBuffer([]byte("Hello")) // 从字节切片创建
	buf4 := bytes.NewBufferString("World")   // 从字符串创建

	fmt.Printf("buf3初始: %q\n", buf3.String()) // "Hello"
	fmt.Printf("buf4初始: %q\n", buf4.String()) // "World"

	// 2. 往Buffer写入数据
	buf1.Write([]byte("Go "))                  // 写入字节切片
	buf1.WriteString("Buffer ")                // 写入字符串
	buf1.WriteByte('!')                        // 写入单个字节
	buf1.WriteRune('哈')                        // 写入单个rune（中文）
	fmt.Printf("buf1写入后: %q\n", buf1.String()) // "Go Buffer !哈"

	// 3. 从Buffer读取数据
	readBuf := make([]byte, 5)
	n, _ := buf1.Read(readBuf)                       // 读取5字节到切片
	fmt.Printf("Read读取: %q（长度%d）\n", readBuf[:n], n) // "Go Bu"

	readStr, _ := buf1.ReadString('!')        // 按分隔符'!'读取
	fmt.Printf("ReadString读取: %q\n", readStr) // "ffer !"

	r, _, _ := buf1.ReadRune()        // 读取单个rune
	fmt.Printf("ReadRune读取: %q\n", r) // "哈"

	// 4. 容量与重置
	fmt.Printf("buf3初始Len: %d, Cap: %d\n", buf3.Len(), buf3.Cap()) // 5, 5
	buf3.Grow(10)                                                  // 扩容（确保剩余容量≥10）
	fmt.Printf("buf3扩容后Cap: %d\n", buf3.Cap())                     // 16（Go按2的倍数扩容）
	buf3.Reset()                                                   // 清空Buffer
	fmt.Printf("buf3重置后Len: %d\n", buf3.Len())                     // 0
}

// 8. Reader类型核心操作：读取/定位/重置
func readerDemo() {
	data := []byte("123456789")
	reader := bytes.NewReader(data)

	// 基本信息（Size固定，Len随读取减少）
	fmt.Printf("Reader总长度(Size): %d\n", reader.Size()) // 9（底层数据总长度）
	fmt.Printf("初始剩余长度(Len): %d\n", reader.Len())      // 9

	// 1. 基础读取
	readBuf := make([]byte, 2)
	n, _ := reader.Read(readBuf)                     // 读取2字节
	fmt.Printf("Read读取: %q（长度%d）\n", readBuf[:n], n) // "12"
	fmt.Printf("读取后剩余Len: %d\n", reader.Len())       // 7

	b, _ := reader.ReadByte()                  // 读取单个字节
	fmt.Printf("ReadByte读取: %q\n", b)          // "3"
	fmt.Printf("读取后剩余Len: %d\n", reader.Len()) // 6

	// 2. 定位读取（Seek+ReadAt）
	reader.Seek(0, 0)                              // 重置读取位置到开头（offset=0, whence=0）
	fmt.Printf("Seek重置后剩余Len: %d\n", reader.Len()) // 9

	readAtBuf := make([]byte, 3)
	reader.ReadAt(readAtBuf, 3)                    // 从偏移量3的位置读取（不改变当前读取位置）
	fmt.Printf("ReadAt读取(偏移3): %q\n", readAtBuf)   // "456"
	fmt.Printf("ReadAt后剩余Len: %d\n", reader.Len()) // 9（位置未变）

	// 3. 循环读取与重置
	fmt.Println("=== 循环读取所有内容 ===")
	reader.Seek(0, 0) // 重置到开头
	for {
		b, err := reader.ReadByte()
		if err != nil {
			break // 读取完毕（EOF）
		}
		fmt.Printf("%q ", b) // "1" "2" "3" ... "9"
	}
	fmt.Println()

	// 4. 重置Reader（切换底层数据）
	newData := []byte("abcdef")
	reader.Reset(newData)                         // 重置为新的字节切片
	fmt.Printf("Reset后总长度: %d\n", reader.Size())  // 6
	fmt.Printf("Reset后剩余Len: %d\n", reader.Len()) // 6
}
