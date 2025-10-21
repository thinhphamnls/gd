package gdentity

type CustomerPhoneType int32

const (
	CustomerPhoneTypeMobile  CustomerPhoneType = 1
	CustomerPhoneTypeHome    CustomerPhoneType = 2
	CustomerPhoneTypeWork    CustomerPhoneType = 3
	CustomerPhoneTypeMain    CustomerPhoneType = 4
	CustomerPhoneTypeHomeFax CustomerPhoneType = 5
	CustomerPhoneTypeWorkFax CustomerPhoneType = 6
)
