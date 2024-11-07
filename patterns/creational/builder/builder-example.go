// source: https://www.linkedin.com/learning/go-design-patterns/
package builder

import "fmt"

type Notification struct {
	title    string
	message  string
	image    string
	icon     string
	priority int
	notType  string
}

// The NotificationBuilder has fields exported as well as a few methods
// to demonstrate
type NotificationBuilder struct {
	Title    string
	Subtitle string
	Message  string
	Image    string
	Icon     string
	Priority int
	NotType  string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.Title = title
}

func (nb *NotificationBuilder) SetMessage(message string) {
	nb.Message = message
}

func (nb *NotificationBuilder) SetImage(image string) {
	nb.Image = image
}

func (nb *NotificationBuilder) SetIcon(icon string) {
	nb.Icon = icon
}

func (nb *NotificationBuilder) SetPriority(pri int) {
	nb.Priority = pri
}

func (nb *NotificationBuilder) SetType(notType string) {
	nb.NotType = notType
}

func (nb *NotificationBuilder) Build() (Notification, error) {
	var emptyNotification Notification
	if nb.Icon != "" && nb.Subtitle == "" {
		return emptyNotification, fmt.Errorf("You need to specify a subtitle when using an icon")
	}
	if nb.Priority > 5 {
		return emptyNotification, fmt.Errorf("Priority must be 0 to 5")
	}

	return Notification{
		title:    nb.Title,
		message:  nb.Message,
		image:    nb.Image,
		icon:     nb.Icon,
		priority: nb.Priority,
		notType:  nb.NotType,
	}, nil
}


func main() {
	nb := newNotificationBuilder()

	// setters
	nb.SetTitle("Hello")
	nb.SetMessage("World")
	nb.SetIcon("icon.png")
	nb.SetPriority(5)
	nb.SetType("Reminder")

	// build action
	notification, err := nb.Build()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(notification) // just a placeholder so that the linter stops complaining that notifcation is unused
}
