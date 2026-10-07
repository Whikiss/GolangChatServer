package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

func (this *Server) Handler(conn net.Conn) {
	//....连接当前业务
	fmt.Print("连接建立成功")
}

// 启动服务器的接口
func (this *Server) Start() {
	//stock listener
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("netListen err:", err)
		return
	}
	//close listen stock
	defer listener.Close()

	for {
		//Accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			continue
		}

		//do handler
		go this.Handler(conn)
	}

}
