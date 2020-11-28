package structs

type Notification struct {
	Title       string
	Description string
	FkUser      int64
}

type NotificationUser struct {
	NotificationId int64
	FkUser         int64
}
