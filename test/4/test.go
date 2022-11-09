package main
import (
	"fmt"
	"unsafe"
	"reflect"
)
func main() {
	s :="wangkun"
	bytes2 := string2bytes2(s)
	byteS := string2bytes(s)
	ss := bytes2string(byteS)
	fmt.Println(bytes2,byteS,ss)
}
func string2bytes(s string) []byte{
	return *(*[]byte)(unsafe.Pointer(&s))
}
func string2bytes2(s string) []byte{
	return *(*[]byte)(unsafe.Pointer(&(*(*reflect.StringHeader)(unsafe.Pointer(&s)))))
}
func bytes2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

