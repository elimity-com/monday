package monday

// WebhooksService handles all the webhook related methods of the Monday API.
// A webhooks allows you to subscribe to events on your boards, and get notified by an HTTP post request to
// a specified URL with the event information as a payload.
//
// DOCS: https://monday.com/integrations/webhooks
type WebhooksService service

// Create returns a mutation that allows you to create a new webhook.
// After the mutation runs, the webhook will report the wanted event on the specific board given to the wanted URL.
// - id: the board's unique identifier.
// - url: the webhook URL.
// - event: the event to listen to (change_column_value / create_item / create_update).
//
// DOCS: https://monday.com/developers/v2#mutations-section-webhooks-create
func (*WebhooksService) Create(id int, url string, event WebhookEventType, webhookFields []WebhookField) Mutation {
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
			{"board_id", id},
			{"url", url},
			{"event", event.typ},
		},
	}
}

// Delete returns a mutation that deletes a webhook.
// After the mutation runs it will no longer report events to the URL given.
// - id: the webhook's unique identifier.
//
// DOCS: https://monday.com/developers/v2#mutations-section-webhooks-delete
func (*WebhooksService) Delete(id int, webhookFields []WebhookField) Mutation {
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

// The webhook's graphql field(s).
type WebhookField struct {
	field field
}

var (
	webhookBoardIDField = WebhookField{field{"board_id", nil}}
	webhookIDField      = WebhookField{field{"id", nil}}
)

// The webhook's board id.
func WebhookBoardIDField() WebhookField {
	return webhookBoardIDField
}

// The webhook's unique identifier.
func WebhookIDField() WebhookField {
	return webhookIDField
}

// The webhook's target type.
type WebhookEventType struct {
	typ string
}

var (
	webhookEventTypeChangeColumnValue = WebhookEventType{"change_column_value"}
	webhookEventTypeCreateItem        = WebhookEventType{"create_item"}
	webhookEventTypeCreateUpdate      = WebhookEventType{"create_update"}
)

// Column value changed on board.
func WebhookEventTypeChangeColumnValue() WebhookEventType {
	return webhookEventTypeChangeColumnValue
}

// An item was created on board.
func WebhookEventTypeCreateItem() WebhookEventType {
	return webhookEventTypeCreateItem
}

// An update was posted on board item.
func WebhookEventTypeCreateUpdate() WebhookEventType {
	return webhookEventTypeCreateUpdate
}
