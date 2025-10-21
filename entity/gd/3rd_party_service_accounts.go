package gdentity

type ThirdPartyServiceType int32

const (
	ThirdPartyServicePostalMail     ThirdPartyServiceType = 0
	ThirdPartyServiceDoorMamba      ThirdPartyServiceType = 1
	ThirdPartyServiceWebsiteBuilder ThirdPartyServiceType = 2
	ThirdPartyServiceMasterAI       ThirdPartyServiceType = 3
)

type ThirdPartyServiceStatus int32

const (
	ThirdPartyServiceInActive ThirdPartyServiceStatus = 0
	ThirdPartyServiceActive   ThirdPartyServiceStatus = 1
)
