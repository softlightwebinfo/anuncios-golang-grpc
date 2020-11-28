package services

import (
	"cientosdeanuncios.com/backend/structs"
	"database/sql"
	"errors"
)

func CreateNotification(DB *sql.DB, notification *structs.Notification) (int64, error) {
	var id int64 = 0
	err := DB.QueryRow("INSERT INTO notifications(fk_user, title, description) VALUES($1, $2, $3) returning id",
		&notification.FkUser,
		&notification.Title,
		&notification.Description,
	).Scan(&id)

	if err != nil {
		return 0, errors.New("No se ha podido crear la notificación")
	}
	return id, nil
}

func CreateUserNotification(DB *sql.DB, notification *structs.NotificationUser) {
	_, _ = DB.Exec("INSERT INTO users_notifications(fk_user, fk_notification) VALUES($1, $2)",
		&notification.FkUser,
		&notification.NotificationId,
	)
}

func SendNotifications(DB *sql.DB, users []int64, title string, description string) {
	notification, err := CreateNotification(DB, &structs.Notification{
		Title:       title,
		Description: description,
		FkUser:      users[0],
	})

	if err != nil {
		return
	}

	for _, userRow := range users {
		go CreateUserNotification(DB, &structs.NotificationUser{
			NotificationId: notification,
			FkUser:         userRow,
		})
	}
}
