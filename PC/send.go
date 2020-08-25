/**
 * @Title  sending
 * @description  #
 * @Author  沈来
 * @Update  2020/8/16 15:01
 **/
package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)
func failOnError(err error, msg string){
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func sending(s string,v string){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open an channel")
	defer ch.Close()

	err = ch.ExchangeDeclare("logs_topic", "topic", true, false, false, false, nil, )
	failOnError(err, "Failed to declare an exchange")

	body := s//s  发送信息，v 绑定条件
	err = ch.Publish("logs_topic", v, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] sent %s", body)
}