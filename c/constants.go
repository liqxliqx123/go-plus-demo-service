package c

import "time"

// http code constants definition
const (
	HTTPSuccessCode         = "200"
	HTTPErrorCode           = "500"
	HTTPInvalidParamsCode   = "400"
	HTTPCodeNotFound        = "404"
)

const (
	// ProjectName ...
	ProjectName       = "Demo"
	ErrMessage        = "wrong"
	SpanSubtypePrefix = "xxx_"
)

// key
const (
	LangEn          = "en"
	LangZh          = "zh"
	LangCn          = "cn"
	LangEnCn        = "en-cn"
	LangEnZh        = "en-zh"
	DataVersion     = "data_version"

	CampaignType   = "campaign_type"
	CampaignStatus = "campaign_status"

	RequestID = "REQUEST_ID"

	CustomerID = "customer_id"
	GID        = "gid"
	UserID     = "user_id"

	LogKey       = "log_key"
	Entity       = "entity"
	Service      = "service"
	User         = "user_id"
	XRequest     = "X-Request-Id"
	Debug        = "debug"
	ErrorMessage = "error_message"
	ErrorTrace   = "error_trace"

	Equal   = "equal"
	Between = "between"
	In      = "in"
	NotIn   = "not_in"
	Like    = "like"

	Asc  = "asc"
	Desc = "desc"

	ApolloAppID     = "APOLLO.APP_ID"
	ApolloCLuster   = "APOLLO.CLUSTER"
	ApolloHost      = "APOLLO.HOST"
	ApolloNameSpace = "APOLLO.NAMESPACE"
	ApolloToken     = "APOLLO.TOKEN"
	ArgoTokenName   = "xxx-apollo-token"

	Daily   = "daily"
	Monthly = "monthly"

	Production = "production"
	Local      = "local"
)

// time
const (
	FormatToSecond           = "2006-01-02 15:04:05"
	FormatToMinute           = "2006-01-02 15:04:00"
	FormatToMicroSecond      = "2006-01-02 15:04:05.000"
	FormatToHour             = "2006-01-02 15"
	FormatToDay              = "2006-01-02"
	FormatToDayDot           = "2006.1.2"
	FormatToDayBeta          = "20060102"
	StatisticalPeriod        = 7 * 24 * time.Hour
	FullTimeStart            = "2021-07-01 00:00:00"
	FormatToMonth            = "2006-01"
	FormatToMonth2           = "200601"
	FormatToHourMinute       = "00:00"
	FormatToHourMinuteSecond = "15:04:05"
	TimeFormatYYYYMMDD       = "20060102"
)

const (
	// APIGateWayStorage prefix api gateway path
	APIGateWayStorage = "storage"
	// ExtPdf pdf
	ExtPdf = ".pdf"
	// ExtPng png
	ExtPng  = ".png"
	ExtJpg  = ".jpg"
	ExtJpeg = ".jpeg"
)

// internal user
const (
	InternalGID = "INTERTOOLS"
	InternalUID = "internal_user"
)

const (
	// KafkaClientID defines kafka producer client ID
	KafkaClientID = "xxx"
)

const (
	xxxResultStatusPending    = "pending"
	xxxResultStatusStarting   = "starting"
	xxxResultStatusProcessing = "processing"
	xxxResultStatusFailing    = "failing"
)

const (
	RealTimexxxType = "real-time"
	DailyxxxType    = "daily"
)

const (
	EndDateTypeYesterday = "yesterday"
	EndDateTypeToday     = "today"
)

// TaskCommand task command
const TaskCommand = "touch /result; /bin/workflow_task %d"
