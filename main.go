package main

import (
	"flag"
	"fmt"
	"net"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	host string
	port int
)

/*
ec2 pulic ip: 3.236.42.248 port: 4000, 4001
ga anycast ip: 13.248.235.115\76.223.107.21 port: 26458, 26459
*/

func main() {
	flag.StringVar(&host, "h", "13.248.235.115", "host")
	flag.IntVar(&port, "p", 26458, "connect port")
	flag.Parse()
	logrus.Infof("host: %s port: %d", host, port)

	// 1、与服务端建立连接
	t1 := time.Now().UnixMilli()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		fmt.Printf("conn server failed, err:%v\n", err)
		return
	}
	t2 := time.Now().UnixMilli()
	fmt.Printf("建立连接时间:%d\n", t2-t1)
	var count int
	var total int64
	// 2、使用 conn 连接进行数据的发送和接收
	for {
		f1 := time.Now().UnixMilli()
		// 发送消息
		_, err = conn.Write([]byte("hello"))
		if err != nil {
			fmt.Printf("send failed, err:%v\n", err)
			return
		}
		f2 := time.Now().UnixMilli()
		// 从服务端接收回复消息
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read failed:%v\n", err)
			return
		}
		f3 := time.Now().UnixMilli()
		fmt.Printf("第%d次，共耗时%d,写入%d,接收%d\n收到服务端回复:%v\n", count, f3-f1, f2-f1, f3-f2, string(buf[:n]))
		count++
		total += f3 - f1
		if count == 100 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("100次耗时:%d\n", total)
}
