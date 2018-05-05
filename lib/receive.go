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

func Receive() {
	conn, err := net.Dial("tcp", addr) //拨号操作，需要指定协议。
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
			panic("File")
		}
		writer := bufio.NewWriterSize(file, bufferSize)
		fmt.Println(fileHeader, len(buf), "n->", n)
		///
		nn, ew := writer.Write(buf[headerSize:n])
		writer.Flush()
		file.Close()
		if ew != nil {
			fmt.Println("ERR:", ew, nn)
		} else {
			// fmt.Println("nn:", nn)
		}
		///
	}

}
