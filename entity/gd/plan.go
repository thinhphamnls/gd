package gdentity

type PlanType int32

const (
	PlanTypeBasic        PlanType = 1
	PlanTypePro          PlanType = 20
	PlanTypeAnnualBasic  PlanType = 30
	PlanTypeAnnualPro    PlanType = 50
	PlanTypeAnnualFreeze PlanType = 40
	PlanTypeFreeze       PlanType = 10
	PlanTypeGrowth       PlanType = 60
	PlanTypeAnnualGrowth PlanType = 70
)

type PlanEnableType int32

const (
	PlanDisable PlanEnableType = 0
	PlanEnable  PlanEnableType = 1
)
