package main

import "kafka/players"

func main() {
	go players.Producer()

	consumer := players.Consumer{}
	consumer.Consume()
}
