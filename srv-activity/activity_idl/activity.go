package activity_idl

import (
	"github.com/cloudwego/kitex/client"
	"log"
	activity "seckill_system/srv-activity/kitex_gen/activity/activityservice"
)

var ActivityClient activity.Client

func InitActivityClient() {
	a, err := activity.NewClient("Activity", client.WithHostPorts("localhost:8889"))
	if err != nil {
		log.Println("can not connect the activity error:", err)
		return
	}
	ActivityClient = a
}
