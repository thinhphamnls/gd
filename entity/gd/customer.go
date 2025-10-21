package gdentity

type CustomerActiveType int32

const (
	CustomerInActive CustomerActiveType = 0
	CustomerActive   CustomerActiveType = 1
	CustomerLead     CustomerActiveType = 2
)

type CustomerDeletedType int32

const (
	CustomerNotDelete CustomerDeletedType = 0
	CustomerDeleted   CustomerDeletedType = 1
)

type CustomerSyncQbType int32

const (
	CustomerSynQbNotSync      CustomerSyncQbType = 0
	CustomerSyncQbSyncSuccess CustomerSyncQbType = 1
	CustomerSyncQbSyncFail    CustomerSyncQbType = 2
)
