package chacha

import (
	"encoding/hex"
	"testing"
)

var chacha *ChaCha
func init(){
	key,_:=hex.DecodeString("0000000000000000000000000000000000000000000000000000000000000000")
	nonce,_:=hex.DecodeString("0000000000000000")

	chacha=New(key,nonce)
}



func TestChaCha_EncodeAndDecode(t *testing.T) {

	plainText:=[]byte("Hello")
	cipherText,err:= chacha.Encode(plainText)
	if err != nil {
		t.Error(err)
	}
	t.Logf("cipherText:%X",cipherText)
	newText,err:= chacha.Decode(cipherText)
	if err != nil {
		t.Error(err)
	}
	if string(newText)!=string(plainText){
		t.Errorf("decode or encode error,plainText: %s,result: %s",string(plainText),string(newText))
	}
	t.Logf("cipherText:%s",newText)
}