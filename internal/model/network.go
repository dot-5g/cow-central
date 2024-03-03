package model

type ApplicationId string

type ApplicationIdRm string

type PduSessionId int

type Mcc string

type MccRm string

type Mnc string

type MncRm string

type Tac string

type TacRm string

type EutraCellId string

type EutraCellIdRm string

type NrCellId string

type NrCellIdRm string

type Dnai string

type DnaiRm string

type FiveGMmCause uint

type AreaCodeRm string

type AmfName string

type AreaCode string

type N3IwfId string

type WAgfId string

type TngfId string

type NgeNbId string

type Nid string

type NidRm string

type NfSetId string

type NfServiceSetId string

type PlmnAssiUeRadioCapId byte

type TypeAllocationCode string

type HfcNId string

type HfcNIdRm string

type ENbId string

type Gli byte

type Gci string

type Snssai struct {
	Sst uint8
	Sd  string
}

type RatType string

const (
	NR             RatType = "NR"
	EUTRA          RatType = "EUTRA"
	WLAN           RatType = "WLAN"
	VIRTUAL        RatType = "VIRTUAL"
	NBIOT          RatType = "NBIOT"
	WIRELINE       RatType = "WIRELINE"
	WIRELINE_CABLE RatType = "WIRELINE_CABLE"
	WIRELINE_BBF   RatType = "WIRELINE_BBF"
	LTE_M          RatType = "LTE-M"
	NR_U           RatType = "NR_U"
	EUTRA_U        RatType = "EUTRA_U"
	TRUSTED_N3GA   RatType = "TRUSTED_N3GA"
	TRUSTED_WLAN   RatType = "TRUSTED_WLAN"
	UTRA           RatType = "UTRA"
	GERA           RatType = "GERA"
)

type Area struct {
	Tacs     []Tac
	AreaCode AreaCode
}

type RestrictionType string

const (
	ALLOWED_AREAS     RestrictionType = "ALLOWED_AREAS"
	NOT_ALLOWED_AREAS RestrictionType = "NOT_ALLOWED_AREAS"
)

type ServiceAreaRestriction struct {
	RestrictionType               RestrictionType
	Areas                         []Area
	MaxNumOfTAs                   Uinteger
	MaxNumOfTAsForNotAllowedAreas Uinteger
}

type CoreNetworkType string

const (
	FIVEGC CoreNetworkType = "5GC"
	EPC    CoreNetworkType = "EPC"
)

type StationaryIndication string

const (
	StationaryIndicationSTATIONARY StationaryIndication = "STATIONARY"
	StationaryIndicationMOBILE     StationaryIndication = "MOBILE"
)

type ScheduledCommunicationTime struct {
	DaysOfWeek     []DayOfWeek
	TimeOfDayStart TimeOfDay
	TimeOfDayEnd   TimeOfDay
}

type GNbId struct {
	BitLength int
	GNBValue  string
}

type GlobalRanNodeId struct {
	PlmnId  PlmnId
	N3IwfId N3IwfId
	GNbId   GNbId
	NgeNbId NgeNbId
	WagfId  WAgfId
	TngfId  TngfId
	Nid     Nid
	ENbId   ENbId
}

type ScheduledCommunicationType string

const (
	ScheduledCommunicationTypeDownlinkOnly  ScheduledCommunicationType = "DOWNLINK_ONLY"
	ScheduledCommunicationTypeUplinkOnly    ScheduledCommunicationType = "UPLINK_ONLY"
	ScheduledCommunicationTypeBidirectional ScheduledCommunicationType = "BIDIRECTIONAL"
)

type TrafficProfile string

const (
	TrafficProfileSINGLE_TRANS_UL     TrafficProfile = "SINGLE_TRANS_UL"
	TrafficProfileSINGLE_TRANS_DL     TrafficProfile = "SINGLE_TRANS_DL"
	TrafficProfileDUAL_TRANS_UL_FIRST TrafficProfile = "DUAL_TRANS_UL_FIRST"
	TrafficProfileDUAL_TRANS_DL_FIRST TrafficProfile = "DUAL_TRANS_DL_FIRST"
	TrafficProfileMULTI_TRANS         TrafficProfile = "MULTI_TRANS"
)

type BatteryIndication struct {
	BatteryInd      bool
	ReplaceableInd  bool
	RechargeableInd bool
}
