package server

import (
	"fmt"
	"grpcui/go-snake-telnet/snake"
	"net"
	"time"

	"github.com/eolinker/eosc/log"
)

const (
	LEFT_TOP_ASCII = "\033[0;0H"
	CLEAR_ASCII    = "\033[2J"
)

type Server struct {
	addr string
}

//返回Server实例
func New(addr string) *Server {
	return &Server{addr: addr}
}

//运行服务
func (this *Server) Run() {
	listener, err := net.Listen("tcp", this.addr)
	if err != nil {
		log.Fatal("TCP establish faild", err.Error())
	}
	defer listener.Close()
	fmt.Printf("TCP established, listening on %s", this.addr)

	//Socket流程
	for {
		//接收TCP流
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("connection failed")
			continue
		}
		log.Printf("Client IP and Port:", conn.RemoteAddr().String())
		//处理流
		go this.doHandler(conn)
	}
}

//处理每个连接
func (this *Server) doHandler(conn net.Conn) {
	//启动

	//清屏初始化位置
	conn.Write([]byte(CLEAR_ASCII + LEFT_TOP_ASCII))
	conn.Write([]byte(LEFT_TOP_ASCII))

	//处理键盘指令输入
	this.HandleKeyboardInput(conn, nil)

	//Over检测
	tick := time.Tick(300 * time.Millisecond)
	for range tick {
		conn.Write([]byte(LEFT_TOP_ASCII + game.Render()))
		if game.IsOver {
			break
		}
	}

	conn.Close()
}

//读取键盘输入和指令
func (this *Server) HandleKeyboardInput(conn net.Conn, game *snake.Game) {

}
