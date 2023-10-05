package goutils

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/5 19:30
  @describe :
*/

// StringTrim 根据指定的start跟count,裁剪出需要的字符串
// 支持汉字何emoji表情等多字节字符
func StringTrim(data string, start, count int) string {
	d := []rune(data)
	l := len(d)

	if start > l-1 {
		return ""
	}

	end := start + count

	if end <= 0 {
		return ""
	}

	if l <= end {
		end = l
	}

	return string(d[start:end])
}

// StringLen 返回字符串的长度,因为中文需要3个字节(byte),所以,我们使用rune来代替
func StringLen(data string) int {
	return len([]rune(data))
}
