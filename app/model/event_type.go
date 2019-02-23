package model

//EventType event type
type EventType struct {
	value string
}

func (d EventType) String() string { return d.value }

// Fake enum, since go is wierd...
var (
	//SessionStarted ...
	SessionStarted = EventType{"sessionStarted"}
	//SessionFinished ...
	SessionFinished = EventType{"sessionFinished"}
	//ScreenPresented ...
	ScreenPresented = EventType{"screenPresented"}
	//ContentShared ...
	ContentShared = EventType{"contentShared"}
	//Other ...
	Custom = EventType{"custom"}

	//EventTypes ...
	EventTypes = []EventType{
		SessionStarted,
		SessionFinished,
		ScreenPresented,
		ContentShared,
		Custom}
)
