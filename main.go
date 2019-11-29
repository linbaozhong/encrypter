package main

import (
	"encoding/hex"
	"encrypter/chacha"
	"encrypter/code"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

var (
	key     string = "1000000000000000000000000000000000000000000000000000000000000001" //密钥
	nonce   string = "1000000000000001"                                                 //向量
	target  string = "/Documents/Encrypted/"                                            //加密后的存放路径
	action  string
	example string = "命令格式:encypter [e|d|c] 文件路径,e：加密文件 ，d：解密文件, c: 查看加密文件明文"
)

func main() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	target = home + target
	flag.Parse()

	if len(flag.Args()) != 2 {
		fmt.Println(example)
		return
	}
	action = flag.Arg(0)
	fileName := flag.Arg(1)

	keyByte, err := hex.DecodeString(key)
	if err != nil {
		panic(err)
	}
	nonceByte, err := hex.DecodeString(nonce)
	if err != nil {
		panic(err)
	}
	encoder := chacha.New(keyByte, nonceByte)

	if _, err = os.Stat(fileName); err != nil {
		fmt.Printf("文件:%s不存在", fileName)
		return
	}

	var sourceFile *os.File
	if sourceFile, err = os.Open(fileName); err != nil {
		fmt.Printf("文件:%s打开失败", fileName)
		return
	}
	defer sourceFile.Close()
	switch action {
	case "e":
		var targetFileName = fmt.Sprintf("%s%s.enc", target, path.Base(fileName))

		targetFile, err := os.OpenFile(targetFileName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
		if err != nil {
			panic(err)
		}
		defer targetFile.Close()
		err = code.EncodeFile(encoder, sourceFile, targetFile)
		if err != nil {
			panic(err)
		}
		fmt.Printf("加密成功,压缩文件存储路径:%s \n", targetFileName)
	case "d":

		ext := path.Ext(fileName)
		if ext != ".enc" {
			fmt.Println("解密文件扩展名不为 .enc,请确认文件是否正确")
			return
		}

		var targetFileName = fmt.Sprintf("%s%c%s", pwd, os.PathSeparator, strings.TrimRight(path.Base(fileName), path.Ext(fileName)))

		targetFile, err := os.OpenFile(targetFileName, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
		if err != nil {
			panic(err)
		}
		defer targetFile.Close()
		err = code.DecodeFile(encoder, sourceFile, targetFile)
		if err != nil {
			panic(err)
		}
		fmt.Printf("解密成功,解压后文件存放位置：%s \n", targetFileName)
	case "c":
		ext := path.Ext(fileName)
		if ext != ".enc" {
			fmt.Println("解密文件扩展名不为 .enc,请确认文件是否正确")
			return
		}
		sExt := path.Ext(strings.TrimRight(fileName, ext))
		if sExt != ".txt" {
			fmt.Println("仅.txt 文件支持查看明文")
			return
		}

		var data []byte
		data, err = code.CheckText(encoder, sourceFile)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", data)

	default:
		fmt.Println(example)
	}

	return
}
