package lv_conv

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

/******************************************************************************
 * 从字节数组中取出整数值
 *
 *******************************************************************************/

func GetNumValueBig(srcArr []byte, startIndex, num int, typeName string) any {
	data := srcArr[startIndex : startIndex+num] //高位在前，低位在后
	switch typeName {
	case "uint8":
		return uint8(data[len(data)-1]) //取最后一个字节上的值
	case "uint16":
		return binary.BigEndian.Uint16(data)
	case "uint32":
		return binary.BigEndian.Uint32(data)
	case "uint64":
		return binary.BigEndian.Uint64(data)
	case "int8":
		return int8(data[0]) //取第一个字节上的值
	case "int16":
		buf := bytes.NewReader(data)
		var num int16
		err := binary.Read(buf, binary.BigEndian, &num)
		if err != nil {
			fmt.Println("转换失败:", err)
		}
		return num
	case "int32":
		buf := bytes.NewReader(data)
		var num int32
		err := binary.Read(buf, binary.BigEndian, &num)
		if err != nil {
			fmt.Println("转换失败:", err)
		}
		return num
	case "int64":
		buf := bytes.NewReader(data)
		var num int32
		err := binary.Read(buf, binary.BigEndian, &num)
		if err != nil {
			fmt.Println("转换失败:", err)
		}
		return num
	}
	return 0
}

func GetNumValueLittle(srcArr []byte, startIndex, endIndex int, typeName string) any {
	data := srcArr[startIndex:endIndex] //高位在前，低位在后
	switch typeName {
	case "uint8":
		return uint8(data[len(data)-1]) //取最后一个字节上的值
	case "uint16":
		return binary.LittleEndian.Uint16(data)
	case "uint32":
		return binary.LittleEndian.Uint32(data)
	case "uint64":
		return binary.LittleEndian.Uint64(data)
	case "int8":
		return int8(data[0]) //取第一个字节上的值
	case "int16":
		buf := bytes.NewReader(data)
		var num int16
		err := binary.Read(buf, binary.LittleEndian, &num)
		if err != nil {
			fmt.Println("转换失败:", err)
		}
		return num
	case "int32":
		buf := bytes.NewReader(data)
		var num int32
		err := binary.Read(buf, binary.LittleEndian, &num)
		if err != nil {
			fmt.Println("转换失败:", err)
		}
		return num
	case "int64":
		buf := bytes.NewReader(data)
		var num int32
		err := binary.Read(buf, binary.LittleEndian, &num)
		if err != nil {
			fmt.Println("转换失败:", err)
		}
		return num
	}
	return 0
}

func Uint32ToByteArr(number uint32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, number) // 使用大端字节序
	return buf.Bytes()
}

func Int32ToByteArr(number int32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, number) // 使用大端字节序
	return buf.Bytes()
}

func Int16ToByteArr(number int16) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, number) // 使用大端字节序
	return buf.Bytes()
}
func Uint16ToByteArr(number uint16) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, number) // 使用大端字节序
	return buf.Bytes()
}
func Uint8ToByteArr(number uint8) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, number) // 使用大端字节序
	return buf.Bytes()
}

func ToBytes(data any) []byte {
	var buf = new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, data)
	return buf.Bytes()
}
