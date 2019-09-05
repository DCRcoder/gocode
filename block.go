package gocode

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"hash/adler32"
	"os"
)

const BlockSize = 128 // default block size unit byte

type Block struct {
	Adler32CheckSum uint32
	MD5CheckSum []byte
	Index int64
}

func RollingBlock(path string){
	blocks := make([]Block, 0)
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	i := int64(0)
	liner := make([]byte, 0)
	for scanner.Scan() {
		tmp := scanner.Bytes()
		liner = append(liner, tmp...)
		if len(liner) > BlockSize{
			h := md5.New()
			b := Block{
				Adler32CheckSum: adler32.Checksum(liner),
				MD5CheckSum: h.Sum(nil),
				Index: i,
			}
			blocks = append(blocks, b)
			i++
			liner = make([]byte, 0, 0)
		}
	}
	fmt.Printf("%v", blocks)

}


