package task
import (
	"time"
	"github.com/shawncode/highlego-api/models"
	"fmt"
)
var ticker = time.NewTicker(time.Second * 60000)
var lChan = make(chan *models.Lottery, 50)

func Task() {
	time.Sleep(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			loadExpired()
		}
	}
}

func loadExpired() {
	if l, err := models.GetExpired(); err == nil {
		fmt.Println("load lotterys %v",l)
		for _,v:= range l {
			id,err:=models.UpdateExpired(&v)
			fmt.Println("id=%d,err=%v",id,err)
		}
	}else {
		fmt.Println(err.Error())
	}
}