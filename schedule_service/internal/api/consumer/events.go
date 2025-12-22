package consumer

const DoctorCreatedEventType = "doctor.created"
const DoctorActiveChangedEventType = "doctor.active_changed"
const DoctorDeletedEventType = "doctor.deleted"

type Event[T any] struct {
	EventType string
	Timestamp int64
	Data      T
}

type DoctorCreatedEvent = Event[DoctorCreatedData]

type DoctorCreatedData struct {
	DoctorID int `json:"doctor_id"`
}

type DoctorActiveChangedEvent = Event[DoctorActiveChangedData]

type DoctorActiveChangedData struct {
	DoctorID int  `json:"doctor_id"`
	Active   bool `json:"active"`
}

type DoctorDeletedEvent = Event[DoctorDeletedData]

type DoctorDeletedData struct {
	DoctorID int `json:"doctor_id"`
}
