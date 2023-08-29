package main

import "fmt"

// This ping function takes two parameter channel and message.
// Sender is a channel we are using to send a value 
func sendToChannel(sender chan<- string, msg string) {
	sender <- msg  //sending value to channel 
}

func receiveFromChannel(sender <-chan string, receive chan<- string) {
	msg := <-sender   // receiving value from a channel 
	receive <- msg
}

func main() {

	// TODO - create channels.
	sender := make(chan string, 1)
    receive := make(chan string, 1)

	//send message to ping channel 
	sendToChannel(sender, " \n ping channel is used to pass message to pong channel. \n ")

	// pass both channel to function
	receiveFromChannel(sender, receive)

	// print messsage from 2nd channel.
	fmt.Println(<-receive) 

}
