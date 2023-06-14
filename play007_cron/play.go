package play007_cron

import (
	"log"
	"time"

	"github.com/robfig/cron"
)

func TestLog() {
	log.Println("Run models.TestLog...")
}
func TestLog2() {
	log.Println("Run models.TestLog2...")
}

func TestCron() {
	log.Println("Starting...")

	c := cron.New()
	c.AddFunc("* * * * * *", TestLog)
	c.AddFunc("* * * * * *", TestLog2)
	c.Start()

	t1 := time.NewTimer(time.Second * 2)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
			log.Println("<-t1.C")
		}
	}
}
