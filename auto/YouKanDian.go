package auto

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type YouKanDian struct {
	exeCount int64
	readTime int64
}

func YouKanDianInit(exeCount, readTime int64) YouKanDian {
	return YouKanDian{exeCount: exeCount,
		readTime: readTime}
}

func (bb YouKanDian) YouKanDianRead(devices []string) {
	count := 0
	i := 1
	for int64(count) < bb.exeCount {
		count++
		log.Printf("优看点 阅读%d次", count)
		//取随机数 阅读时长
		i = -i
		rt := int64(rand.Intn(5)) + bb.readTime
		x1 := 600 + rand.Intn(100)*i
		y1 := 500 - rand.Intn(40)*i

		var griup sync.WaitGroup
		for _, device := range devices {
			griup.Add(1)
			go func(group sync.WaitGroup, device string) {
				sprintf := fmt.Sprintf("-s %s shell input tap %d %d", device, x1, y1)
				split := strings.Split(sprintf, " ")
				runAdb(split...)
				time.Sleep(time.Second)
				sprintf1 := fmt.Sprintf("-s %s shell input tap %d %d", device, 280, 220)
				split1 := strings.Split(sprintf1, " ")
				runAdb(split1...)

				//开始时长
				start := time.Now().Unix()
				for time.Now().Unix()-start < rt {
					time.Sleep(time.Second * 3)
				}

				sprintf2 := fmt.Sprintf("-s %s shell input tap %d %d", device, 20, 45)
				split2 := strings.Split(sprintf2, " ")
				runAdb(split2...)

				sprintf3 := fmt.Sprintf("-s %s shell input tap %d %d", device, 20, 45)
				split3 := strings.Split(sprintf3, " ")
				runAdb(split3...)

				x1 := RandInt64(300, 350)
				x2 := x1 + RandInt64(5, 10)

				y1 := 950 + rand.Intn(10)
				y2 := y1 - RandInt64(400, 450)

				sprintf4 := fmt.Sprintf("-s %s shell input swipe %d %d %d %d 1000", device, x1, y2, x2, y1)
				split4 := strings.Split(sprintf4, " ")
				runAdb(split4...)

				griup.Done()
			}(griup, device)
		}
		griup.Wait()

	}
	log.Println("阅读完成：优看点")
}
