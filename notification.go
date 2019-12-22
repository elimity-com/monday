package monday

// NotificationsService handles all the notification related methods of the Monday API.
// Notifications are sent in the platform in response to various triggers such as due dates, updates, and more.
type NotificationsService service

// Create returns a mutation that allows to trigger a notification within the platform.
// Keep in mind that notifications are async and you will not be able to query their ID's back after sending then out.
// - userID: the user's unique identifier.
// - targetID: the target's unique identifier.
// - text: the notification text.
// - targetType: the target's type (project/post).
//
// DOCS: https://monday.com/developers/v2#mutations-section-notifications
func (*NotificationsService) Create(userID, targetID int, text string, targetType NotificationType, notificationFields []NotificationsField) Mutation {
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
			{"target_type", targetType},
		},
	}
}

// Create returns a mutation that allows to trigger a notification within the platform.
// Keep in mind that notifications are async and you will not be able to query their ID's back after sending then out.
// - userID: the user's unique identifier.
// - targetID: the target's unique identifier.
// - text: the notification text.
// - payload: the notification payload (JSON).
// - targetType: the target's type (project/post).
//
// DOCS: https://monday.com/developers/v2#mutations-section-notifications
func CreateWithPayload(userID, targetID int, text, payload string, targetType NotificationType, notificationFields []NotificationsField) Mutation {
	notification := Notifications.Create(userID, targetID, text, targetType, notificationFields)
	notification.args = append(notification.args, argument{"payload", payload})
	return notification
}

// The notification's graphql field(s).
type NotificationsField struct {
	field field
}

var (
	notificationIDField   = NotificationsField{field{"id", nil}}
	notificationTextField = NotificationsField{field{"text", nil}}
)

// The notification's unique identifier.
func NotificationIDField() NotificationsField {
	return notificationIDField
}

// The notification text.
func NotificationTextField() NotificationsField {
	return notificationTextField
}

// The notification's graphql argument(s).
type NotificationArgument struct {
	arg argument
}

// The notification's target type.
type NotificationType struct {
	kind string
}

var (
	notificationTypeProject = NotificationType{"project"}
	notificationTypePost    = NotificationType{"post"}
)

// Pulse or Board.
func NotificationTypeProject() NotificationType {
	return notificationTypeProject
}

// Update.
func NotificationTypePost() NotificationType {
	return notificationTypePost
}
