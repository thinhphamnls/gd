package gdentity

type ScheduleStatusType int32

const (
	ScheduleDeleted   ScheduleStatusType = -1
	ScheduleNotActive ScheduleStatusType = 0
	ScheduleActive    ScheduleStatusType = 1
)
