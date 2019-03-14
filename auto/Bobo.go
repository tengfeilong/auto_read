package auto

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

type BoBo struct {
	exeCount int64
	readTime int64
}

func BoboInit(exeCount, readTime int64) BoBo {
	return BoBo{exeCount: exeCount,
		readTime: readTime}
}

func (bb BoBo) BoboRead(devices []string) {
	count := 0
	i := 1
	for int64(count) < bb.exeCount {
		count++
		log.Printf("波波 阅读%d次", count)
		// 随机数转换
		i = -i
		//取随机数 阅读时长
		rt := int64(rand.Intn(5)) + bb.readTime
		x1 := 600 + rand.Intn(100)*i
		y1 := 500 - rand.Intn(40)*i
		for _, device := range devices {
			sprintf := fmt.Sprintf("-s %s shell input tap %d %d", device, x1, y1)
			split := strings.Split(sprintf, " ")
			runAdb(split...)
		}
		//开始时长
		start := time.Now().Unix()
		for time.Now().Unix()-start < rt {
			time.Sleep(time.Second * 3)
		}
		for _, device := range devices {
			x1 := RandInt64(300, 350)
			x2 := x1 + RandInt64(5, 10)*i

			y1 = 950 + rand.Intn(10)*i
			y2 := y1 - RandInt64(400, 450)
			log.Println("下拉刷新", x1, y1, x2, y2)
			sprintf := fmt.Sprintf("-s %s shell input swipe %d %d %d %d 1000", device, x1, y2, x2, y1)
			split := strings.Split(sprintf, " ")
			runAdb(split...)
		}
	}
	log.Println("阅读完成：波波")
}
