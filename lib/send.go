package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

const addr = "127.0.0.1:8082"
const bufferSize = 1024 * 5 // 5k
const headerSize = 128

type FileHeader struct {
	FileName     string
	FileSize     int64
	PackageIndex int
}

func Send(path string) {
	listener, err := net.Listen("tcp", addr) //使用协议是tcp，监听的地址是addr
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close() //关闭监听的端口
	for {
		conn, err := listener.Accept() //用conn接收链接
		if err != nil {
			log.Fatal(err)
		}
		/// 写文件 start
		////////////////
		fileInfo, _ := os.Stat(path)
		fileHeader := &FileHeader{}
		fileHeader.FileName = fileInfo.Name()
		fileHeader.FileSize = fileInfo.Size()

		////////////////path
		file, _ := os.Open(path)
		defer file.Close()
		buf := make([]byte, bufferSize-headerSize)
		for {
			n, err := file.Read(buf)
			if err != nil && err != io.EOF {
				panic(err)
			} //
			if 0 == n {
				break
			}
			// time.Sleep(time.Second / 1000)
			fileHeader.PackageIndex++
			fileHeaderBytes, _ := json.Marshal(fileHeader)

			headers, _ := makeHeaderBytes(string(fileHeaderBytes), headerSize, " ")
			headers = append(headers, buf[:n]...)

			fmt.Println("n: ", n, "  ", "headers:", len(headers), string(headers[:headerSize]))
			// conn.Write(buf[:n])
			conn.Write(headers)
		}
		/// 写文件 end
		conn.Close() //与客户端断开连接。
	}
}

func makeHeaderBytes(str string, size int, fill string) ([]byte, error) {
	bytes := []byte(str)
	length := len(bytes)
	if length > size {
		return nil, errors.New("内容过长: length=" + strconv.Itoa(length))
	} else {
		emptySize := size - length
		for i := 0; i < emptySize; i++ {
			bytes = append(bytes, ([]byte(fill)[0]))
		}
		return bytes, nil
	}
}
