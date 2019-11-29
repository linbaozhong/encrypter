package code

import (
	"io"
	"os"
)


type Decoder interface {
	Decode(plainData []byte) (cipherData []byte, err error)
}

func DecodeFile(encoder Decoder,sourefile *os.File,targetFile *os.File)(err error){
	buf:=make([]byte,1024)
	for {
		var n int
		n,err=sourefile.Read(buf)
		if err==io.EOF{
			err=nil
			break
		}
		if err!=nil{
			return
		}
		var cipher []byte
		cipher,err=encoder.Decode(buf[:n])
		if err!=nil{
			return
		}
		_,err=targetFile.Write(cipher)
		if err!=nil{
			return
		}
	}
	return
}

func CheckText(encoder Decoder,sourefile *os.File)(plainData []byte,err error){
	buf:=make([]byte,1024)
	for {
		var n int
		n,err=sourefile.Read(buf)
		if err==io.EOF{
			err=nil
			break
		}
		if err!=nil{
			return
		}
		plainData,err=encoder.Decode(buf[:n])
		if err!=nil{
			return
		}
	}
	return
}