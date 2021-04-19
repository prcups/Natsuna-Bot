package cqhandler

import (
	"log"
	"github.com/gorilla/websocket"
	"net/url"
	"os"
	"os/signal"
	"encoding/json"
	"net/http"
	"errors"
	_ "plugin"
	_ "filepath"
	"math/rand"
)

type naddr {
	qqnum string
	addr string
	to int
	isSent bool
}

var aa [105]naddr
var l int

const (
	HTTP_ADDRESS = "http://127.0.0.1:5700"
	WEBSOCKET_ADDRESS = "ws://127.0.0.1:6700"
)

func commandHandler(resStr []byte) {
		err = filepath.Walk("../recv_plugins", func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				log.Fatal(err.Error())
			}
			filepath.
		})
		
		if resMap["message_type"] == "private" {
			
		} else if resMap["message_type"] == "group" {
			
		}
	}
 }

func BootHandler() {
	url, err := url.Parse("ws://localhost:6700")
	if err != nil {
		log.Fatal(err.Error())
	}
	
	ws, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
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
				log.Fatal(err.Error)
			}
			return
		default: 
			id, resStr, err := ws.ReadMessage()
			if err != nil {
				log.Fatal(err)
			}
			commandHandler(message)
			var resMap, senderMap map
			err := json.Unmarshal(resStr, &resMap)
			if err != nil {
				log.Println(err.Error())
			}
			
			if resMap["message_type"] == "private" && resMap["message"][0] == ":" {
				l++
				aa[l] = {
					qqnum: resMap["user_id"],
					addr: resMap["message"][1:],
					to: l,
					isSent: false
				}
				ws.WriteMessage(id, "{ reply:/"记录成功！/"")
				//resp, err := http.PostForm({
					//"http://127.0.0.1:5700/send_private_msg",
					//url.Values{"user_id": resMap[user_id], "message": {"添加记录成功！"}}
				//})
				//resp.Body.Close()
			}	else if resMap["message"] == "分发地址" {
				for i := 1; i <= l; i++ {
					if aa[i].to == i {
						j := rand.Int() % l + 1
						for aa[i].to == j || aa[j].to == i {
							j := rand.Int() % l + 1
						}
						t := aa[i].to
						aa[i].to = aa[j].to
						aa[j].to = t
					}
				}
				
				for i := 1; i <= l; i++ {
					resp, _ := http.PostForm({
						"http://127.0.0.1:5700/send_private_msg",
						url.Values{"user_id": resMap[user_id], "group_id": "1022606672", message": aa[aa[i].to].addr}
					})
					resp.Body.Close()
				}
				l = 0
			}
		}
	}
}
