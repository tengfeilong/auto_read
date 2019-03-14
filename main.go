package main

import (
	"auto_read/auto"
	"bufio"
	"log"
	"os/exec"
	"regexp"
)

var regAdbInfo = regexp.MustCompile(`^(.*)	device$`)

func getDevicesAll() (devices []string) {
	devices = make([]string, 0)
	cmd := exec.Command("adb", "devices")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println("out err", err)
	}
	err = cmd.Start()
	if err != nil {
		log.Println("run err", err)
	}
	rd := bufio.NewReader(stdout)
	for {
		line, _, err := rd.ReadLine()
		if err != nil || line == nil {
			break
		}
		submatch := regAdbInfo.FindStringSubmatch(string(line))
		if len(submatch) == 2 {
			log.Println("adb name", submatch[1])
			devices = append(devices, submatch[1])
		}
	}
	cmd.Wait()
	return devices
}

func main() {
	devices := getDevicesAll()
	log.Println(devices[0])

	//时间单位 1小时
	//var appTimeBase int64 = 3600
	// 单篇阅读时长 30秒
	//var readTimeBase = 30

	//刷宝对模拟器限制（多开模拟器）
	sb := auto.ShuaBaoInit(3600 * 5)
	sb.ShuabaoRead(devices)

	//波波视频  优看点
	//bb := auto.BoboInit(100, 120)
	//bb.BoboRead(devices)
	//种子视频
	//ziInit := auto.ZhongZiInit(100, 60)
	//ziInit.ZhongZiRead(devices)

	//init := auto.YouKanDianInit(100, 40)
	//init.YouKanDianRead(devices)
}
