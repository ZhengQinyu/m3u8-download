package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	url  string
	name string
)

/* 只初始化一次的方法，优先级比 main 还高 */
func init() {
	flag.StringVar(&url, "u", "", "m3u8链接")
	flag.StringVar(&name, "n", "", "文件名称")
}

func main() {
	flag.Parse()
	if url == "" {
		fmt.Println("-u 参数必传")
		return
	}
	if name == "" {
		fmt.Println("-n 参数必传")
		return
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("获取m3u8失败，请检查 -u 参数对应的链接")
		return
	}
	defer resp.Body.Close()
	// 获取链接前缀
	url = url[:strings.LastIndex(url, "/")]

	var lines []string

	var buffer = bufio.NewReader(resp.Body)
	f, _ := os.Create(name + ".ts")
	for {
		line, err := buffer.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "#") {
			lines = append(lines, line)
		}
	}
	for idx, line := range lines {
		r, e := http.Get(url + "/" + line)
		if e != nil {
			fmt.Printf("下载片段 %d/%d 失败：%s\n", idx+1, len(lines), line)
			return
		} else {
			_, _ = io.Copy(f, r.Body)
			fmt.Printf("下载片段 %d/%d 完成\n", idx+1, len(lines))
		}
		r.Body.Close()
	}
	f.Close()
}
