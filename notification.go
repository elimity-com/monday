package monday

func CreateNotification(notificationFields []NotificationField, notificationArguments []NotificationArgument) Mutation {
	if len(notificationFields) == 0 {
		notificationFields = append(notificationFields, notificationTextField)
	}
	var fields []field
	for _, nf := range notificationFields {
		fields = append(fields, nf.field)
	}
	var args []argument
	for _, na := range notificationArguments {
		args = append(args, na.arg)
	}
	return Mutation{
		name:   "create_notification",
		fields: fields,
		args:   args,
	}
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

// The notification text.
func NewNotificationTextField(value string) NotificationArgument {
	return NotificationArgument{argument{"text", value}}
}

// The user's unique identifier.
func NewNotificationUserIDField(value int) NotificationArgument {
	return NotificationArgument{argument{"user_id", value}}
}

// The target's unique identifier.
func NewNotificationTargetIDField(value int) NotificationArgument {
	return NotificationArgument{argument{"target_id", nil}}
}

// The target's type (Project / Post).
func NewNotificationTargetTypeField(value NotificationType) NotificationArgument {
	return NotificationArgument{argument{"target_type", value.kind}}
}

// The notification payload.
func NewNotificationPayloadField(value string) NotificationArgument {
	return NotificationArgument{argument{"payload", value}}
}
