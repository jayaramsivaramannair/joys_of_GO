package main

import (
	"fmt"
)

//Any struct or user defined type will be implementing the notifier interface as long as it has the notify method defined
type notifier interface {
	notify()
}


//user defined "user" struct
type user struct {
	name string
	email string
}

//notify method has been made available on user struct
// to implement the notifier interface with a pointer receiver
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

//user defined "admin" struct
type admin struct {
	name string
	email string
}


//notify method is also made available on admin struct
//to implement the notifier interface with a pointer receiver
func(a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}



//sendNotification accepts values that implements the notifier
//interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}

func main() {
	bill := user{"Bill", "bill@email.com"}
	//Since user struct impleements the notifier interface using a pointer receiver
	//SendNotification method must be invoked by passing a reference to an instance of user and not the instance itself
	sendNotification(&bill)


	lisa := admin{"Lisa", "lisa@email.com"}
	sendNotification(&lisa)
}