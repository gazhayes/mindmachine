package eventbucket

import (
	"mindmachine/mindmachine"
)

func HandleEvent(event mindmachine.Event) (h mindmachine.HashSeq, b bool) {
	currentState.mutex.Lock()
	defer currentState.mutex.Unlock()
	currentState.upsert(
		Event{
			EventID:    event.ID,
			Event:      event.Nostr(),
			MentionMap: make(map[string]struct{}),
		})
	return
}