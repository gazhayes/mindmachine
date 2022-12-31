package eventbucket

type EventBucket struct {
}

func (eb *EventBucket) CalculateMentions() bool {
	CalculateMentions()
	return true
}

func (eb *EventBucket) CurrentOrder() []Event {
	return getCurrentOrder()
}