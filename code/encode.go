package code

import (
	"io"
	"os"
)


type Encoder interface {
	Encode(plainData []byte) (cipherData []byte, err error)
}

func EncodeFile(encoder Encoder,sourefile *os.File,targetFile *os.File)(err error){
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
		cipher,err=encoder.Encode(buf[:n])
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

