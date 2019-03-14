package auto

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

type Shuabao struct {
	Apptime int64
}

func ShuaBaoInit(apptime int64) Shuabao {
	return Shuabao{Apptime: apptime}
}

func (sb Shuabao) ShuabaoRead(devices []string) {
	appstart := time.Now()
	count := 0
	as := time.Now().Unix() - appstart.Unix()
	for as < sb.Apptime {
		count++
		//取随机数5-10 阅读时长
		i := rand.Intn(6) + 15
		log.Printf("阅读%d次,阅读时长%d", count, i)
		start := time.Now()
		for time.Now().Unix()-start.Unix() < int64(i) {
			time.Sleep(time.Second * 3)
		}
		x := 400
		y1 := 825
		y2 := y1 - 605
		for _, device := range devices {
			sprintf := fmt.Sprintf("-s %s shell input swipe %d %d %d %d 1000 &", device, x, y1, x, y2)
			split := strings.Split(sprintf, " ")
			runAdb(split...)
		}
	}
	log.Println("阅读完成：刷宝")
}
