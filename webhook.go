package monday

func CreateWebhook(boardID int, url string, event WebhookEventType, webhookFields []WebhookField) Mutation {
	if len(webhookFields) == 0 {
		webhookFields = append(webhookFields, webhookIDField)
	}

	var fields []field
	for _, wf := range webhookFields {
		fields = append(fields, wf.field)
	}
	return Mutation{
		name:   "create_webhook",
		fields: fields,
		args: []argument{
			{"board_id", boardID},
			{"url", url},
			{"event", event.typ},
		},
	}
}

func DeleteWebhook(id int, webhookFields []WebhookField) Mutation {
	if len(webhookFields) == 0 {
		webhookFields = append(webhookFields, webhookIDField)
	}

	var fields []field
	for _, wf := range webhookFields {
		fields = append(fields, wf.field)
	}
	return Mutation{
		name:   "delete_webhook",
		fields: fields,
		args: []argument{
			{"id", id},
		},
	}
}

type WebhookField struct {
	field field
}

var (
	webhookBoardIDField = WebhookField{field{"board_id", nil}}
	webhookIDField      = WebhookField{field{"id", nil}}
)

// The webhooks's board id.
func WebhookBoardIDField() WebhookField {
	return webhookBoardIDField
}

// The webhooks's unique identifier.
func WebhookIDField() WebhookField {
	return webhookIDField
}

type WebhookEventType struct {
	typ string
}

var (
	webhookEventTypeChangeColumnValue = WebhookEventType{"change_column_value"}
	webhookEventTypeCreateItem        = WebhookEventType{"create_item"}
	webhookEventTypeCreateUpdate      = WebhookEventType{"create_update"}
)

func WebhookEventTypeChangeColumnValue() WebhookEventType {
	return webhookEventTypeChangeColumnValue
}

func WebhookEventTypeCreateItem() WebhookEventType {
	return webhookEventTypeCreateItem
}

func WebhookEventTypeCreateUpdate() WebhookEventType {
	return webhookEventTypeCreateUpdate
}
