package web

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

//func Run() {
//	app := httpd.New()
//	app.Get("/", IndexHandlerFunc)
//	app.Run()
//}
//func IndexHandlerFunc(c *httpd.Context) {
//	c.Write.Write([]byte("Hello ant-go"))
//}

// ==================

const (
	readBufferSize  = 1024
	writeBufferSize = 1024
)

type Client struct {
	conn     *websocket.Conn
	messages chan []byte
}

var clients map[*Client]bool
var broadcast chan []byte
var addClients chan *Client

func (c *Client) readPump() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		fmt.Printf("receive message is :%s\n", message)
		broadcast <- message
	}
}

func (c *Client) writePump() {
	for {
		select {
		case message := <-c.messages:
			fmt.Printf("send message is :%s\n", message)
			c.conn.WriteMessage(1, message)
		}
	}
}

func Push() {
	var addr = "localhost:3010"
	//u := url.URL{Scheme: "ws", Host: addr}

	u := url.URL{Scheme: "ws", Host: addr, Path: "/ws"}

	fmt.Println("接口地址:" + u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		fmt.Printf("连接错误: %v \n", err)

		return
	}

	werr := conn.WriteMessage(websocket.TextMessage, []byte("11111111111111"))

	fmt.Printf("发送失败: %v", werr)
}

func (c *Client) pushPump() {
	for {
		select {
		case message := <-c.messages:
			fmt.Printf("send message is :%s\n", message)
			c.conn.WriteMessage(1, message)
		}
	}
}

func manager() {

	clients = make(map[*Client]bool)
	broadcast = make(chan []byte, 10)
	addClients = make(chan *Client)

	for {
		select {
		case message := <-broadcast:
			for client := range clients {
				select {
				case client.messages <- message:
				default:
					close(client.messages)
					delete(clients, client)
				}
			}
		case itemClient := <-addClients:
			clients[itemClient] = true
		}
	}
}

func Run() {

	var homeTemplate = template.Must(template.ParseFiles("./web/index.html"))

	m := martini.Classic()

	m.Get("/", func(res http.ResponseWriter, req *http.Request) {

		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		homeTemplate.Execute(res, req.Host)
	})

	m.Get("/ws", func(res http.ResponseWriter, req *http.Request) { // res and req are injected by Martini
		conn, err := websocket.Upgrade(res, req, nil, readBufferSize, writeBufferSize)
		if err != nil {
			log.Println(err)
			return
		}
		client := &Client{conn: conn, messages: make(chan []byte, 5)}
		addClients <- client
		go client.writePump()
		client.readPump()
	})

	go manager()

	m.RunOnAddr("localhost:3010")
	//m.Run()

}

// ==========
