package transparent_proxy

import (
	"io"
	"net"
	"syscall"

	"tp/director"
)

const SO_ORIGINAL_DST = 80

// from pkg/net/parse.go
// Convert i to decimal string.
func itod(i uint) string {
	if i == 0 {
		return "0"
	}

	// Assemble decimal in reverse order.
	var b [32]byte
	bp := len(b)
	for ; i > 0; i /= 10 {
		bp--
		b[bp] = byte(i%10) + '0'
	}

	return string(b[bp:])
}

func RUN() {
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go func(conn net.Conn) {
			file, err := conn.(*net.TCPConn).File()
			if err != nil {
				panic(err)
			} else {
				conn.Close()
			}
			// 代理把 TCP 连接拦截下来之后，它需要知道原来的目的地址是什么
			addr, err := syscall.GetsockoptIPv6Mreq(int(file.Fd()), syscall.IPPROTO_IP, SO_ORIGINAL_DST)
			if err != nil {
				panic(err)
			}
			conn, err = net.FileConn(file)
			if err != nil {
				panic(err)
			}
			defer conn.Close()
			// 目的地址IP xxx.xxx.xxx.xxx
			dst_ip := itod(uint(addr.Multiaddr[4])) + "." +
				itod(uint(addr.Multiaddr[5])) + "." +
				itod(uint(addr.Multiaddr[6])) + "." +
				itod(uint(addr.Multiaddr[7]))
			// 目的地址端口
			// dport := uint16(addr.Multiaddr[2])<<8 + uint16(addr.Multiaddr[3])

			// 通過 目的地址IP 『dst_ip』 進行服務發現，得到真實的後端地址 『forward_addr』
			forward_addr := director.Select(dst_ip)
			forward, err := net.Dial("tcp", forward_addr)
			defer forward.Close()
			if err != nil {
				panic(err)
			}
			stop := make(chan int8, 1)
			go func() {
				_, _ = io.Copy(conn, forward)
				stop <- 1
			}()
			go func() {
				_, _ = io.Copy(forward, conn)
				stop <- 1
			}()
			<-stop
		}(conn)
	}
}
