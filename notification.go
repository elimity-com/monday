package monday

func CreateNotification(userID, targetID int, text string, targetType NotificationType, notificationFields []NotificationField) Mutation {
	if len(notificationFields) == 0 {
		notificationFields = append(notificationFields, notificationTextField)
	}
	var fields []field
	for _, nf := range notificationFields {
		fields = append(fields, nf.field)
	}
	return Mutation{
		name:   "create_notification",
		fields: fields,
		args: []argument{
			{"text", text},
			{"user_id", userID},
			{"target_id", targetID},
			{"target_type", targetType.kind},
		},
	}
}

func CreateNotificationWithPayload(userID, targetID int, text, payload string, targetType NotificationType, notificationFields []NotificationField) Mutation {
	notification := CreateNotification(userID, targetID, text, targetType, notificationFields)
	notification.args = append(notification.args, argument{"payload", payload})
	return notification
}

type NotificationField struct {
	field field
}

var (
	notificationIDField   = NotificationField{field{"id", nil}}
	notificationTextField = NotificationField{field{"text", nil}}
)

// The notification's unique identifier.
func NotificationIDField() NotificationField {
	return notificationIDField
}

// The notification text.
func NotificationTextField() NotificationField {
	return notificationTextField
}

type NotificationArgument struct {
	arg argument
}

type NotificationType struct {
	kind string
}

var (
	notificationTypeProject = NotificationType{"Project"}
	notificationTypePost    = NotificationType{"Post"}
)

// Pulse or Board.
func NotificationTypeProject() NotificationType {
	return notificationTypeProject
}

// Update.
func NotificationTypePost() NotificationType {
	return notificationTypePost
}
