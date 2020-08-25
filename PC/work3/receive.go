/**
 * @Title  receiving
 * @description  #
 * @Author  沈来
 * @Update  2020/8/16 15:09
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

func main(){
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare("logs_topic", "topic", true, false, false, false, nil)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare("", false, false, true, false, nil, )
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(q.Name, "*.message.*", "logs_topic", false, nil)

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil, )
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func(){
		for d:= range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}