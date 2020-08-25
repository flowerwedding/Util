/**
 * @Title  server
 * @description  tcp聊天室的服务端
 * @Author  沈来
 * @Update  2020/8/10 8:41
 **/
package main


import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

var (
	GenUserID = 0
	enteringChannel = make(chan *User)
	leavingChannel = make(chan *User)
	messageChannel = make(chan Message, 8)
)

type User struct {
	ID              int
	Addr            string
	EnterAt         time.Time
	MessageChannel  chan string
}

type Message struct{
	OwnerID int
	Content string
}

func main(){
	listener, err := net.Listen("tcp",":9000")
	if err != nil {
		panic(err)
	}

	go broadCaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	GenUserID ++
	user := &User{
		ID:             GenUserID,
		Addr:           conn.RemoteAddr().String(),
		EnterAt:        time.Now(),
		MessageChannel: make(chan string, 8),
	}

	go sendMessage(conn, user.MessageChannel)

	user.MessageChannel <- "Welcome, " + strconv.Itoa(user.ID)
	messageChannel <- Message{
		OwnerID: user.ID,
		Content: " has enter",
	}

	var userActive = make(chan struct{})
	go func() {
		d := 5 * time.Minute
		timer := time.NewTimer(d)
		for {
			select {
			case <-timer.C:
				conn.Close()
			case <-userActive:
				timer.Reset(d)
			}
		}
	}()

	enteringChannel <- user

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messageChannel <- Message{
			OwnerID: user.ID,
			Content: strconv.Itoa(user.ID) +":"+ input.Text(),
		}

		userActive <- struct{}{}
	}

	if err := input.Err(); err != nil {
		log.Println("读取错误：", err)
	}

	leavingChannel <- user
	messageChannel <- Message{
		OwnerID: user.ID,
		Content: " has left",
	}
}

func sendMessage(conn net.Conn, ch <- chan string) {
	for msg := range ch {
		fmt.Println(conn, msg)
	}
}

func broadCaster(){
	users := make(map[*User]struct{})

	for{
		select{
		case user := <-enteringChannel:
			users[user] = struct{}{}
		case user := <-leavingChannel:
			delete(users,user)
			close(user.MessageChannel)
		case msg := <-messageChannel:
			for user := range users{
				//user.MessageChannel <- msg
				if user.ID == msg.OwnerID {
					continue
				}
				user.MessageChannel <- msg.Content
			}
		}
	}
}