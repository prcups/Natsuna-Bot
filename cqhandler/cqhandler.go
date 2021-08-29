package cqhandler

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
)


const (
	HTTP_ADDRESS = "http://127.0.0.1:5700"
	WEBSOCKET_ADDRESS = "ws://127.0.0.1:6700"
)


func BootHandler() {
	wsurl, err := url.Parse("ws://localhost:6700")
	if err != nil {
		log.Fatal(err.Error())
	}
	
	ws, _, err := websocket.DefaultDialer.Dial(wsurl.String(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer ws.Close()
	
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	
	for {
		select {
		case <- interrupt:
			log.Println("interrupt")
			err := ws.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Fatal(err)
			}
			return
		default: 
			_, resStr, err := ws.ReadMessage()
			if err != nil {
				log.Fatal(err)
			}
			var resMap map[string]interface{}
			err = json.Unmarshal(resStr, &resMap)
			if err != nil {
				log.Println(err)
			}

			log.Println(resMap)
			if resMap["notice_type"] == "group_increase" && int(resMap["group_id"].(float64)) == 247736999 {
				log.Println("来客人了")
				_, err := http.PostForm(
					HTTP_ADDRESS + "/send_group_msg",
					url.Values{
						"group_id": {"247736999"},
						"message": {"[CQ:at,qq=" + strconv.FormatInt(int64(resMap["user_id"].(float64)), 10) + "] 欢迎热爱技术的你加入兰州大学开源社区的大家庭！想了解更多请阅读群公告新生指南噢！"},
					},
				)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}
