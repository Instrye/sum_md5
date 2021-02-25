package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

var file = flag.Bool("f", false, "是否是文件")

func main() {
	flag.Parse()
	str := flag.Arg(0)
	hash := md5.New()
	if *file {
		files, err := os.Open(str)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		defer files.Close()
		fileBufio := bufio.NewReader(files)
		for {
			buffers := make([]byte, 512)
			n, err := fileBufio.Read(buffers)
			hash.Write(buffers[:n])
			if err != nil {
				if err != io.EOF {
					fmt.Println(err)
				}
				break
			}
		}
	} else {
		io.WriteString(hash, str)
	}
	md5Hash := hash.Sum(nil)
	fmt.Printf("32 : %s\r\n", hex.EncodeToString(md5Hash))
	fmt.Printf("16 : %s\r\n", hex.EncodeToString(md5Hash[4:12]))
}
