package consumer

const DoctorDeactiveteScheduleEventType = "schedule.deactivate"
const TimeSlotDeletedEventType = "schedule.time_slots_deleted"

type Event[T any] struct {
	EventType string
	Timestamp int64
	Data      T
}

type DoctorDeactiveteScheduleEvent = Event[DoctorDeactiveteScheduleData]

type DoctorDeactiveteScheduleData struct {
	TimeSlotsID []int `json:"time_slots_id"`
}

type TimeSlotDeletedEvent = Event[TimeSlotDeletedData]

type TimeSlotDeletedData struct {
	TimeSlotsID []int `json:"time_slots_id"`
}
