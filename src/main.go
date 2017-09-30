package main

import (
	"fmt"
	"os"
	"io"
	"encoding/base64"
	"bufio"
)

func main() {
	args := os.Args
	if args == nil || len(args) < 4 {
		fmt.Println("please input file path!")
		os.Exit(1)
	}
	runCmd(args[1], args[2], args[3])
}

func runCmd(flag, in_file, out_file string) {
	switch (flag) {
	case "encode":
		fmt.Println("encode file to base64.")
		working(in_file, out_file, true)
	case "decode":
		fmt.Println("decode file to txt.")
		working(in_file, out_file, false)
	default:
		fmt.Println("error cmd!")
	}
}

func working(in_file, out_file string, code bool) {
	fi, err := os.Open(in_file)
	if err != nil {
		fmt.Printf("open file %s failed!\n", in_file)
		os.Exit(1)
	}
	defer fi.Close()
	fo, err := os.Create(out_file)
	if err != nil {
		fmt.Printf("create file %s failed!\n", out_file)
		os.Exit(2)
	}
	defer fo.Close()
	switch code {
	case true:
		encode_base64(fi, fo)
	case false:
		decode_base64(fi, fo)
	default:
	}
}

func encode_base64(fi, fo *os.File) {
	buf := bufio.NewReader(fi)
	const base64_line_len int = 48  // (txt)48Bytes => (base64)6bit => (base64)64Bytes
	for flag := true; flag ; {  // break for while
		data := [base64_line_len]byte{0}
		i := 0
		for ; i < base64_line_len; {  // break for while
			c, err := buf.ReadByte()
			if err != nil {
				if err != io.EOF {
					fmt.Printf("Read Error: %s\n", err.Error())
					os.Exit(1)
				}
				flag = false
				break
			}
			if c == '\r' {  // delete \r char
				continue
			}
			data[i] = c
			i++
		}
		if i == 0 {  // i == 0
			continue
		}
		encodeString := base64.StdEncoding.EncodeToString(data[:i])
		if flag == true {
			encodeString += "\n"
		}
		fo.Write([]byte(encodeString))
	}
}

func decode_base64(fi, fo *os.File) {
	buf := bufio.NewReader(fi)
	const base64_line_len int = 64  // (txt)48Bytes => (base64)6bit => (base64)64Bytes
	for flag := true; flag ; {
		data := [base64_line_len]byte{0}
		i := 0
		for ; i < base64_line_len; {
			c, err := buf.ReadByte()
			if err != nil {
				if err != io.EOF {
					fmt.Printf("Read Error: %s\n", err.Error())
					os.Exit(1)
				}
				flag = false
				break
			}
			if c == '\r' || c == '\n' {  // delete \r\n char
				continue
			}
			data[i] = c
			i++
		}
		if i == 0 {  // i == 0
			continue
		}
		decodeBytes, err := base64.StdEncoding.DecodeString(string(data[:i]))
		if err != nil {
			fmt.Printf("Encode Base64 Error: %s\n", err.Error())
			os.Exit(2)
		}
		fo.Write(decodeBytes)
	}
}