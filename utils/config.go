package utils

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Configs struct {
	config map[string]string
	node   string
}

const MidStr = "-_-!"

var Conf *Configs

func init() {
	Conf = new(Configs)
	Conf.LoadConfig("config/app.ini")
}

func (conf *Configs) LoadConfig(path string) {
	conf.config = make(map[string]string)
	file, err := os.Open(path)
	// fmt.Print(file, err)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	buf := bufio.NewReader(file)

	for {
		lines, _, err := buf.ReadLine()
		line := strings.TrimSpace(string(lines))
		if err != nil {
			//读完最后一行退出
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		// 处理注释
		if strings.Index(line, "#") == 0 {
			continue
		}
		//如果是[xxx]
		n := strings.Index(line, "[")
		nl := strings.LastIndex(line, "]")

		if n > -1 && nl > -1 && nl > n+1 {
			conf.node = strings.TrimSpace(line[n+1 : nl])
			continue
		}
		if len(conf.node) == 0 || len(line) == 0 {
			continue
		}
		arr := strings.Split(line, "=")
		key := strings.TrimSpace(arr[0])
		value := strings.TrimSpace(arr[1])
		newKey := conf.node + MidStr + key
		conf.config[newKey] = value
	}

	// fmt.Println(conf.config)
}

func (conf *Configs) ReadStr(node, key string) string {
	key = node + MidStr + key
	if v, ok := conf.config[key]; !ok {
		return ""
	} else {
		return v
	}
}

func (conf *Configs) ReadInt32(node, key string) int {
	key = node + MidStr + key
	if v, ok := conf.config[key]; !ok {
		return 0
	} else {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0
		}
		return i
	}
}

func (conf *Configs) ReadInt64(node, key string) int64 {
	key = node + MidStr + key
	if v, ok := conf.config[key]; !ok {
		return 0
	} else {
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0
		}
		return i
	}
}
