package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const bufferSize = 512 // 5k
const headerSize = 128

type FileHeader struct {
	FileName     string
	FileSize     int64
	PackageIndex int
}

func Send(path, addr string) {
	pathes, _ := getFilelist(path)
	for _, p := range pathes {
		SendSingleFile(p, addr)
	}
}

func SendSingleFile(path, addr string) {
	// conn, err := net.Dial("tcp", addr) //拨号操作，需要指定协议。
	conn, err := net.DialTimeout("tcp", addr, 4*time.Second) // 带超时机制的Dial
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}
	/////////////  写文件 start
	fileInfo, errFileInfo := os.Stat(path)
	if errFileInfo != nil {
		fmt.Println(errFileInfo)
	}
	fileHeader := &FileHeader{}
	fileHeader.FileName = fileInfo.Name()
	fileHeader.FileSize = fileInfo.Size()

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
		conn.Write(headers)
	}
	/// 写文件 end
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

func getFilelist(path string) ([]string, error) {
	var allPath []string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		allPath = append(allPath, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return allPath, nil
}
