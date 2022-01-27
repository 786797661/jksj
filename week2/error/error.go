package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {

	//v,err :=errout()
	//if err != nil {
	//	//fmt.Errorf("这个错误是%s",err)
	//	IsTemporary(err)
	//}else {
	//	fmt.Println(v)
	//}
	_, err := opfile("1")
	var another queryError
	errors.Is(err, &another)
	errors.As(err, &another)
	//fmt.Printf("original error : %T %v\n",errors.Cause(err),errors.Cause(err))
	//fmt.Printf("stack error: \n%+v\n",err )
}

func errout() (int, error) {
	err := errors.New("this is a error")
	return 0, err
}

//opaque errors
type temporary interface {
	Temporary() bool
}

func IsTemporary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.Temporary()
}

//func (p *queryError) Error() string {
//	return fmt.Sprintf("path=%s \nop=%s \ncreateTime=%s \nmessage=%s", p.path,
//		p.op, p.createTime, p.message)
//}

type errWriter struct {
	io.Writer
	err error
}

func (e *errWriter) Write(buf []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}
	var n int
	n, e.err = e.Writer.Write(buf)
	return n, e.err
}

type Header struct {
	Key, value string
}
type Status struct {
	Code   int
	Reason string
}

func WriteResponse(w io.Writer, st Status, header []Header, body io.Reader) error {
	ew := &errWriter{Writer: w}
	fmt.Fprintf(ew, "HTTP/1.1 %d %s\r\n", st.Code, st.Reason)
	for _, h := range header {
		fmt.Fprintf(ew, "%s:%s\r\n", h.Key, h.value)
	}
	fmt.Fprint(ew, "\r\n")
	io.Copy(ew, body)
	return ew.err
}

type queryError struct {
	query string
	err   error
}

func (e *queryError) UnWrap() error {
	return e.err
}
func (e *queryError) Error() string {
	return "error msg :" + e.query
}

func opfile(path string) ([]byte, error) {
	_, err := os.Open(path)
	_, te := err.(*queryError)
	if err != nil {
		err = fmt.Errorf("opfile faild:%w", te)
	}
	return nil, err
}
