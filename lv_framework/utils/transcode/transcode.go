package transcode

import "github.com/axgle/mahonia"

//src 字符串
//srcCode 字符串当前编码
//tagCode 要转换的编码
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
