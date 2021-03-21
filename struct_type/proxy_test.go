package struct_type

import (
	"fmt"
	"testing"
)

func TestProxy(t *testing.T)  {
	nginx := newNginxServer()
	code , msg := nginx.handlerRequest("create", "post")
	fmt.Printf("response code:%d, msg:%s\n", code, msg)
	code , msg = nginx.handlerRequest("create", "post")
	fmt.Printf("response code:%d, msg:%s\n", code, msg)
	code , msg = nginx.handlerRequest("create", "post")
	fmt.Printf("response code:%d, msg:%s\n", code, msg)
}