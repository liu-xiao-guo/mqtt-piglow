package main

import (
  	"fmt"
 	 //import the Paho Go MQTT library
 	 MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
 	 "os"
  	"time"
)

const TOPIC = "testubuntucore/counter"

//define a function for the default message handler
var f MQTT.MessageHandler = func(client *MQTT.Client, msg MQTT.Message) {
  	fmt.Printf("Recived TOPIC: %s\n", msg.Topic())	
  	fmt.Printf("Received MSG: %s\n", msg.Payload())
	
	s := string(msg.Payload()[:])
	fmt.Printf("check: %t\n",  (s == "on"))	
	
	if ( s == "on" ) {
		fmt.Println("on is received!")
		TurnAllOn()
	} else if ( s == "off" )  {
		fmt.Println("off is received!")
		GlowOff()
	}
}

func main() {
	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	opts := MQTT.NewClientOptions().AddBroker("tcp://iot.eclipse.org:1883")

	opts.SetClientID("go-simple")
	opts.SetDefaultPublishHandler(f)
	
	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
	 	panic(token.Error())
	}
	
	//subscribe to the topic and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	if token := c.Subscribe(TOPIC, 0, nil); token.Wait() && token.Error() != nil {
	 	fmt.Println(token.Error())
	 	os.Exit(1)
	}
	
	// Pubish messages to TOPIC at qos 1 and wait for the receipt
	//from the server after sending each message
	i := 0;
	for true {
	 	text := fmt.Sprintf("this is msg #%d! from MQTT piglow", i)
	 	token := c.Publish(TOPIC, 0, false, text)
		token.Wait()
		time.Sleep(5 *time.Second)
		i++;
	}
	
	time.Sleep(3 * time.Second)
	
	//unsubscribe from /go-mqtt/sample
	if token := c.Unsubscribe(TOPIC); token.Wait() && token.Error() != nil {
	 	fmt.Println(token.Error())
		os.Exit(1)
	}
	
	c.Disconnect(250)
}
