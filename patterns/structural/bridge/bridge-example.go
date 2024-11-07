package bridge

import "fmt"

// Notification interface represents the abstraction
type Notification interface {
	SendNotification(string) // abstraction operation
}

// SMS is one implementation of Notification
type SMS struct{}

func (s *SMS) SendNotification(message string) {
	fmt.Println("Sending SMS notification:", message)
}

// Email is another implementation of Notification
type Email struct{}

func (e *Email) SendNotification(message string) {
	fmt.Println("Sending Email notification:", message)
}

// Push is another example of a notification type
type Push struct{}

func (p *Push) SendNotification(message string) {
	fmt.Println("Sending Push notification:", message)
}

// User represents the bridge that uses multiple Notification implementations
type User struct {
	Notifications []Notification // now holds a slice of Notification
}

// Notify sends a message using all available notification methods
func (u *User) Notify(message string) {
	for _, notifier := range u.Notifications {
		notifier.SendNotification(message)
	}
}

func main() {
	// Set up multiple notification methods for user1
	smsNotification := &SMS{}
	emailNotification := &Email{}
	pushNotification := &Push{}

	user := User{
		Notifications: []Notification{smsNotification, emailNotification, pushNotification},
	}
	
	// Notify user via all methods
	user.Notify("Your package is on its way!")
}