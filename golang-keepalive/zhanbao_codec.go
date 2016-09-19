package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
)


func Encode(message string) ([]byte, error) {
	//读取消息的长度
	var length int32 = int32(len(message))
	var pkg *bytes.Buffer = new(bytes.Buffer)

	//写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}

	//写入消息实体
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))

	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}


func Decode(reader *bufio.Reader) (string, error) {
	//读取消息的长度
	lengthByte, _ := reader.Peek(4)			//读取前4位， 也就是 int32 的长度
	lengthBuff := bytes.NewBuffer(lengthByte)

	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}

	//读取消息真正的内容
	pack := make([]byte, int(4 + length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}






