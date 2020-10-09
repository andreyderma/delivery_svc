package config

const (
	UNASSIGNED_STATUS = "UNASSIGNED"
	TAKEN_STATUS      = "TAKEN"
	SUCCESS_STATUS    = "SUCCESS"
)

var (
	LatitudeRegexString  = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
	LongitudeRegexString = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
)
