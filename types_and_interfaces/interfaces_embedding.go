package main 


import (
	"fmt"
)

//Notifier is an interface that defined notification type behavior
type notifier interface {
	notify()
}


type user struct {
	name string
	email string
}

//notify implements a method which can be called via a value of type user.
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}


type admin struct {
	user
	level string
}

//notify implements a method which can be called via a value of type admin.
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

func main() {
	//create an admin user
	ad := admin {
		user: user {
			name: "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	//Send the admin user a notification.
	//The embedded inner type's implementation of the interface is NOT
	//promoted to the outer type. 
	sendNotification(&ad)


	//We can access the inner type's method directly.
	ad.user.notify()

	//The inner type's method is NOT promoted.
	ad.notify()
}

//SendNotification accepts values that implements the notifier interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}