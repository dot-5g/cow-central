package NudrDataRepository

import "github.com/dot-5g/cow-central/internal/model"

type DefaultDnnIndicator bool

type LboRoamingAllowed bool

type UeUsageType int

type MpsPriorityIndicator bool

type McsPriorityIndicator bool

type ThreeGPPChargingCharacteristics string

type DlPacketCount int

type MicoAllowed bool

type SmsSubscribed bool

type SharedDataId string

type IwkEpsInd bool

type SecuredPacket string

type UpuRegInd bool

type ExtGroupId string

type NbIoTUePriority int

type CodeWord string

type AfId string

type LcsClientId string

type EventType string

const (
	LOSS_OF_CONNECTIVITY           EventType = "LOSS_OF_CONNECTIVITY"
	UE_REACHABILITY_FOR_DATA       EventType = "UE_REACHABILITY_FOR_DATA"
	UE_REACHABILITY_FOR_SMS        EventType = "UE_REACHABILITY_FOR_SMS"
	LOCATION_REPORTING             EventType = "LOCATION_REPORTING"
	CHANGE_OF_SUPI_PEI_ASSOCIATION EventType = "CHANGE_OF_SUPI_PEI_ASSOCIATION"
	ROAMING_STATUS                 EventType = "ROAMING_STATUS"
	COMMUNICATION_FAILURE          EventType = "COMMUNICATION_FAILURE"
	AVAILABILITY_AFTER_DDN_FAILURE EventType = "AVAILABILITY_AFTER_DDN_FAILURE"
	CN_TYPE_CHANGE                 EventType = "CN_TYPE_CHANGE"
	DL_DATA_DELIVERY_STATUS        EventType = "DL_DATA_DELIVERY_STATUS"
	PDN_CONNECTIVITY_STATUS        EventType = "PDN_CONNECTIVITY_STATUS"
	UE_CONNECTION_MANAGEMENT_STATE EventType = "UE_CONNECTION_MANAGEMENT_STATE"
)

type SorInfo struct {
	AckInd           model.AckInd
	SorMacIausf      model.SorMac
	CounterSor       model.CounterSor
	ProvisioningTime model.DateTime
}

type SorUpdateIndicator string

const (
	INITIAL_REGISTRATION   SorUpdateIndicator = "INITIAL_REGISTRATION"
	EMERGENCY_REGISTRATION SorUpdateIndicator = "EMERGENCY_REGISTRATION"
)

type UpuInfo struct {
	UpuDataList []model.UpuData
	UpuRegInd   UpuRegInd
	UpuAckInd   model.UpuAckInd
}

type MdtUserConsent string

const (
	CONSENT_NOT_GIVEN MdtUserConsent = "CONSENT_NOT_GIVEN"
	CONSENT_GIVEN     MdtUserConsent = "CONSENT_GIVEN"
)

type CagInfo struct {
	AllowedCagList   []model.CagId
	CagOnlyIndicator bool
}

type CagData struct {
	CagInfos         map[string]CagInfo
	ProvisioningTime model.DateTime
}

type EcRestrictionDataWb struct {
	EcModeARestricted bool
	EcModeBRestricted bool
}

type Nssai struct {
	SupportedFeatures   model.SupportedFeatures
	DefaultSingleNssais []model.Snssai
	ProvisioningTime    model.DateTime
}

type NetworkAreaInfo struct {
	Ecgis       []model.Ecgi
	Ncgis       []model.Ncgi
	GRanNodeIds []model.GlobalRanNodeId
	Tais        []model.Tai
}
type LocationArea struct {
	CivicAddresses []model.CivicAddress
	NwAreaInfo     NetworkAreaInfo
}

type ExpectedUeBehaviourData struct {
	StationaryIndication       model.StationaryIndication
	CommunicationDurationTime  model.DurationSec
	PeriodicTime               model.DurationSec
	ScheduledCommunicationTime model.ScheduledCommunicationTime
	ScheduledCommunicationType model.ScheduledCommunicationType
	ExpectedUmts               []LocationArea
	TrafficProfile             model.TrafficProfile
	BatteryIndication          model.BatteryIndication
	ValidityTime               model.DateTime
}

type MaximumResponseTime struct {
	MaximumResponseTime model.DurationSec
	Snssai              model.Snssai
	Dnn                 model.Dnn
	ValidityTime        model.DateTime
}

type MaximumLatency struct {
	MaximumLatency model.DurationSec
	Snssai         model.Snssai
	Dnn            model.Dnn
	ValidityTime   model.DateTime
}

type EdrxParameters struct {
	RatType   model.RatType
	EdrxValue string
}

type OperationMode string

const (
	OperationModeWB_S1 OperationMode = "WB_S1"
	OperationModeNB_S1 OperationMode = "NB_S1"
	OperationModeWB_N1 OperationMode = "WB_N1"
	OperationModeNB_N1 OperationMode = "NB_N1"
)

type PtwParameters struct {
	OperationMode OperationMode
	PtwValue      string
}

type AccessAndMobilitySubscriptionData struct {
	SupportedFeatures           model.SupportedFeatures
	Gpsis                       []model.Gpsi
	InternalGroupIds            []model.GroupId
	SubscribedUeAmbr            model.AmbrRm
	Nssai                       Nssai
	RatRestrictions             []model.RatType
	ForbiddenAreas              []model.Area
	ServiceAreaRestriction      model.ServiceAreaRestriction
	CoreNetworkTypeRestrictions []model.CoreNetworkType
	RfspIndex                   model.RfspIndexRm
	SubsRegTimer                model.DurationSecRm
	UeUsageType                 UeUsageType
	MpsPriority                 MpsPriorityIndicator
	McsPriority                 McsPriorityIndicator
	ActiveTime                  model.DurationSecRm
	DlPacketCount               DlPacketCount
	SorInfo                     SorInfo
	SorInfoExpectInd            bool
	SorafRetrieval              bool
	SorUpdateIndicatorList      []SorUpdateIndicator
	UpuInfo                     UpuInfo
	MicoAllowed                 MicoAllowed
	SharedAmDataIds             []SharedDataId
	OdbPacketServices           model.OdbPacketServices
	SubscribedDnnList           []model.Dnn
	ServiceGapTime              model.DurationSec
	MdtUserConsent              MdtUserConsent
	MdtConfiguration            model.MdtConfiguration
	TraceData                   model.TraceData
	CagData                     CagData
	StnSr                       model.StnSr
	CMsisdn                     model.CMsisdn
	NbIoTUePriority             NbIoTUePriority
	NssaiInclusionAllowed       bool
	RgWirelineCharacteristics   model.RgWirelineCharacteristics
	EcRestrictionDataWb         EcRestrictionDataWb
	EcRestrictionDataNb         bool
	ExpectedUeBehaviour         ExpectedUeBehaviourData
	MaximumResponseTime         []MaximumResponseTime
	MaximumLatencyList          []MaximumLatency
	PrimaryRatRestrictions      []model.RatType
	SecondaryRatRestrictions    []model.RatType
	EdrxParametersList          []EdrxParameters
	PtwParametersList           []PtwParameters
	IabOperationAllowed         bool
}
