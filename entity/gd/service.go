package gdentity

type ServiceType int32

const (
	ServiceCustom  ServiceType = 0
	ServiceDefault ServiceType = 2
	ServicePlan    ServiceType = 3
)

type ServiceDeleteType int32

const (
	ServiceNotDelete ServiceDeleteType = 0
	ServiceDeleted   ServiceDeleteType = 1
)

type ServiceArchiveType int32

const (
	ServiceNotArchive ServiceArchiveType = 0
	ServiceArchived   ServiceArchiveType = 1
)
