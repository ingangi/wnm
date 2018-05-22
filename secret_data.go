package main

import (
	"os"
	"fmt"
	"net"
	"strings"
	"encoding/binary"
)

//默认的服务器地址
var (
	server = "challenge.yuansuan.cn:7042"
)
//成功后，服务器开始发送数据包，每个数据包的格式如下：
//
//0        4        8      12
//+--------+--------+------+=============+
//|SEQUENCE|CHECKSUM|LENGTH| DATA        |
//+--------+--------+------+=============+
//
//SEQUENCE
//数据包序号（大端序）
//CHECKSUM (校验和)
//32位校验和，计算方式如下：首先将序号和数据拼接在一块，然后以32位为块，进行迭代异或操作。
//第一次迭代时，将第一块（数据包序号）与第二块进行异或，第二次迭代时，将第三块与上次迭代的结果进行异或，以此反复。
//如果 LENGTH 不是4的整数倍，需要用 0xAB 填充。
//LENGTH
//数据片段长度（大端序）
//DATA
//长度为 LENGTH 数据片段

const HEADLEN  = 12

func calcCheckSum(data []byte) uint32 {
	l := len(data)
	if l < 8 || l % 4 != 0 {
		fmt.Println("calcCheckSum len error:", len(data))
		return 0
	}

	res := binary.BigEndian.Uint32(data[0 : 4])
	for i:=4;i < l-4 ;i += 4  {
		tmp := binary.BigEndian.Uint32(data[i : i+4])
		res ^= tmp
	}

	return res
}


func main()  {
	//拿到服务器地址信息
	hawkServer,err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		fmt.Printf("hawk server [%s] resolve error: [%s]",server,err.Error())
		os.Exit(1)
	}
	//连接服务器
	connection,err := net.DialTCP("tcp",nil,hawkServer)
	if err != nil {
		fmt.Printf("connect to hawk server error: [%s]",err.Error())
		os.Exit(1)
	}
	defer connection.Close()

	buf := make([]byte, 1024, 1024)
	n, err2 := connection.Read(buf)
	if err2 != nil {
		fmt.Printf("Error: %v", err.Error())
		return
	}
	s := strings.Replace(string(buf[:n]), "\n", "", -1)
	id := strings.Split(s, ":")[1]
	fmt.Println(s, n, id)

	answer := fmt.Sprintf("IAM:%s:test\n", id)
	connection.Write([]byte(answer))
	fmt.Println("answer:", answer)


	buf = make([]byte, 1024, 1024)
	n, err2 = connection.Read(buf)
	if err2 != nil {
		fmt.Printf("Error: %v", err.Error())
		return
	}
	s = strings.Replace(string(buf[:n]), "\n", "", -1)
	fmt.Println(s, n)

	// 接下来的数据包全部写到文件里去
	buf = make([]byte, 1024*10, 1024*10)
	w := uint32(0)
	r := uint32(0)
	for {

		n, err2 = connection.Read(buf[w:])
		if err2 != nil {
			fmt.Printf("Error: %v", err2.Error())
			return
		}
		w += uint32(n)

		for {
			over := false
			if w > r && w - r > HEADLEN {
				seq := binary.BigEndian.Uint32(buf[r : r+4])
				len := binary.BigEndian.Uint32(buf[r+8 : r+8+4])
				fmt.Println("sqe", seq, len, w, r)
				return
				// 一帧
				if uint32(w) >= HEADLEN + len {
					//frame := buf[:HEADLEN + len]
					//data := buf[HEADLEN:HEADLEN + len]
					r += (HEADLEN + len)
					//fmt.Println("frame data len", len, w, r)
				} else {
					over = true
				}
			} else {
				over = true
			}

			if over {
				fmt.Println("move to head", w, r)
				w = uint32(copy(buf, buf[r:]))
				r = 0
				break
			}
			if r > w {
				r = 0
				w = 0
			}

		}
		//fmt.Println(n, buf[:n])
		//ioutil.WriteFile("./test.txt", buf, 0666)
	}


}
