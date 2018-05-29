package lib

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func Receive(listenPort string) {
	fmt.Println("Server listening:", listenPort)
	if _, err := os.Stat("./downloads"); os.IsNotExist(err) { // 如果downloads目录不存在则创建
		os.Mkdir("./downloads", 0700)
	}
	listener, err := net.Listen("tcp", "0.0.0.0:"+listenPort) //使用协议是tcp，监听的地址是addr
	if err != nil {
		panic(err)
	}
	defer listener.Close() //关闭监听的端口
	for {
		conn, err := listener.Accept() //用conn接收链接
		if err != nil {
			log.Fatal(err)
		}
		buf := make([]byte, bufferSize)
		for {
			n, err := conn.Read(buf)
			if err == io.EOF {
				conn.Close()
				break
			}
			/// 读取Header
			fileHeader := &FileHeader{}
			json.Unmarshal(buf[:headerSize], fileHeader)
			if fileHeader.PackageIndex == 1 {
				_, err1 := os.Create("./downloads/" + fileHeader.FileName) //创建文件
				if err1 != nil {
					panic(err1)
				}
			}
			///
			file, errF := os.OpenFile("./downloads/"+fileHeader.FileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
			if errF != nil {
				fmt.Println("出错了！", errF)
			}
			writer := bufio.NewWriterSize(file, bufferSize)
			// fmt.Println(fileHeader, len(buf), "n->", n)
			///
			nn, ew := writer.Write(buf[headerSize:n]) // panic: runtime error: slice bounds out of range
			fmt.Println("len(buf):", len(buf), " headerSize:", headerSize, "n: ", n)

			writer.Flush()
			file.Close()
			if ew != nil {
				fmt.Println("ERR:", ew, nn)
			} else {
				// fmt.Println("nn:", nn)
			}
			///
		}
		conn.Close() //与客户端断开连接。
	}
}
