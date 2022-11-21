package main

import(
	"log"
	"customercenter/customercenter"
	"customercenter/config"
)

func main() {
	config.ConfigMode()
	options := config.Reader(config.GetMode())
	log.Printf("option : %v", options)
	customercenter.MyPageDBInit()
	go customercenter.InitProducer(config.GetMode())
	// policy 와 같이 사용시 InitConsumer 가 중복으로 호출될수 있으니, 하나는 삭제 필요
	go customercenter.InitConsumer(config.GetMode())
	e := customercenter.RouteInit()

	e.Logger.Fatal(e.Start(":8084"))
}
