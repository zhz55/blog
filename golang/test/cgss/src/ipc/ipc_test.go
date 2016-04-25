package ipc

import (
    "testing"
)

type EchoServer struct {
}

func （server *EchoServer) Handle(request string) string {
    return "ECHO:" + request
}

func (server *EchoServer) Name() string {
    return "EchoServer"
}

func TestIpc(t *testing.T) {
    
}
