package model

type PhysCellId int

type ArfcnValueNr int

type PlmnId struct {
	Mcc Mcc
	Mnc Mnc
}

type Tai struct {
	PlmnId PlmnId
	Tac    Tac
	Nid    Nid
}

type TacInfo struct {
	TacList []Tac
}

type Ecgi struct {
	PlmnId      PlmnId
	EutraCellId EutraCellId
	Nid         Nid
}

type Ncgi struct {
	PlmnId   PlmnId
	NrCellId NrCellId
	Nid      Nid
}

type MbsfnArea struct {
	MbsfnAreaId      int
	CarrierFrequency int
}

type InterFreqTargetInfo struct {
	DlCarrierFreq ArfcnValueNr
	CellIdList    []PhysCellId
}

type JobType string

const (
	IMMEDIATE_MDT_ONLY      JobType = "IMMEDIATE_MDT_ONLY"
	LOGGED_MDT_ONLY         JobType = "LOGGED_MDT_ONLY"
	TRACE_ONLY              JobType = "TRACE_ONLY"
	IMMEDIATE_MDT_AND_TRACE JobType = "IMMEDIATE_MDT_AND_TRACE"
	RLF_REPORTS_ONLY        JobType = "RLF_REPORTS_ONLY"
	RCEF_REPORTS_ONLY       JobType = "RCEF_REPORTS_ONLY"
	LOGGED_MBSFN_MDT        JobType = "LOGGED_MBSFN_MDT"
)

type ReportTypeMdt string

const (
	PERIODICAL    ReportTypeMdt = "PERIODICAL"
	EVENT_TRIGGED ReportTypeMdt = "EVENT_TRIGGED"
)

type MeasurementLteForMdt string

const (
	M1Lte    MeasurementLteForMdt = "M1"
	M2Lte    MeasurementLteForMdt = "M2"
	M3Lte    MeasurementLteForMdt = "M3"
	M4_DLLte MeasurementLteForMdt = "M4_DL"
	M4_ULLte MeasurementLteForMdt = "M4_UL"
	M5_DLLte MeasurementLteForMdt = "M5_DL"
	M5_ULLte MeasurementLteForMdt = "M5_UL"
	M6_DLLte MeasurementLteForMdt = "M6_DL"
	M6_ULLte MeasurementLteForMdt = "M6_UL"
	M7_DLLte MeasurementLteForMdt = "M7_DL"
	M7_ULLte MeasurementLteForMdt = "M7_UL"
	M8Lte    MeasurementLteForMdt = "M8"
	M9Lte    MeasurementLteForMdt = "M9"
)

type MeasurementNrForMdt string

const (
	M1Nr    MeasurementNrForMdt = "M1"
	M2Nr    MeasurementNrForMdt = "M2"
	M3Nr    MeasurementNrForMdt = "M3"
	M4_DLNr MeasurementNrForMdt = "M4_DL"
	M4_ULNr MeasurementNrForMdt = "M4_UL"
	M5_DLNr MeasurementNrForMdt = "M5_DL"
	M5_ULNr MeasurementNrForMdt = "M5_UL"
	M6_DLNr MeasurementNrForMdt = "M6_DL"
	M6_ULNr MeasurementNrForMdt = "M6_UL"
	M7_DLNr MeasurementNrForMdt = "M7_DL"
	M7_ULNr MeasurementNrForMdt = "M7_UL"
	M8Nr    MeasurementNrForMdt = "M8"
	M9Nr    MeasurementNrForMdt = "M9"
)

type AreaScope struct {
	EutraCellIdList []EutraCellId
	NrCellIdList    []NrCellId
	TacList         []Tac
	TacInfoPerPlmn  map[string]TacInfo
}

type SensorMeasurement string

const (
	BAROMETRIC_PRESSURE SensorMeasurement = "BAROMETRIC_PRESSURE"
	UE_SPEED            SensorMeasurement = "UE_SPEED"
	UE_ORIENTATION      SensorMeasurement = "UE_ORIENTATION"
)

type ReportingTrigger string

const (
	PERIODICALTrigger      ReportingTrigger = "PERIODICAL"
	EVENTA2Trigger         ReportingTrigger = "EVENT_A2"
	EventA2PeriodicTrigger ReportingTrigger = "EVENT_A2_PERIODIC"
	AllRRMEventTriggers    ReportingTrigger = "ALL_RRM_EVENT_TRIGGERS"
)

type ReportIntervalMdt string

const (
	ReportIntervalMdt120ms   ReportIntervalMdt = "120"
	ReportIntervalMdt240ms   ReportIntervalMdt = "240"
	ReportIntervalMdt480ms   ReportIntervalMdt = "480"
	ReportIntervalMdt640ms   ReportIntervalMdt = "640"
	ReportIntervalMdt1024ms  ReportIntervalMdt = "1024"
	ReportIntervalMdt2048ms  ReportIntervalMdt = "2048"
	ReportIntervalMdt5120ms  ReportIntervalMdt = "5120"
	ReportIntervalMdt10240ms ReportIntervalMdt = "10240"
	ReportIntervalMdt1m      ReportIntervalMdt = "60000"
	ReportIntervalMdt6m      ReportIntervalMdt = "360000"
	ReportIntervalMdt12m     ReportIntervalMdt = "720000"
	ReportIntervalMdt30m     ReportIntervalMdt = "1800000"
	ReportIntervalMdt60m     ReportIntervalMdt = "3600000"
)

type ReportAmountMdt string

const (
	ReportAmountMdt1        ReportAmountMdt = "1"
	ReportAmountMdt2        ReportAmountMdt = "2"
	ReportAmountMdt4        ReportAmountMdt = "4"
	ReportAmountMdt8        ReportAmountMdt = "8"
	ReportAmountMdt16       ReportAmountMdt = "16"
	ReportAmountMdt32       ReportAmountMdt = "32"
	ReportAmountMdt64       ReportAmountMdt = "64"
	ReportAmountMdtInfinity ReportAmountMdt = "infinity"
)

type EventForMdt string

const (
	OutOfCoverageEventForMtd EventForMdt = "OUT_OF_COVERAGE"
	A2EventForMtd            EventForMdt = "A2_EVENT"
)

type LoggingIntervalMdt string

const (
	LoggingIntervalMdt128ms  LoggingIntervalMdt = "128"
	LoggingIntervalMdt256ms  LoggingIntervalMdt = "256"
	LoggingIntervalMdt512ms  LoggingIntervalMdt = "512"
	LoggingIntervalMdt1024ms LoggingIntervalMdt = "1024"
	LoggingIntervalMdt2048ms LoggingIntervalMdt = "2048"
	LoggingIntervalMdt3072ms LoggingIntervalMdt = "3072"
	LoggingIntervalMdt4096ms LoggingIntervalMdt = "4096"
	LoggingIntervalMdt6144ms LoggingIntervalMdt = "6144"
)

type LoggingDurationMdt string

const (
	LoggingDurationMdt600  LoggingDurationMdt = "600"
	LoggingDurationMdt1200 LoggingDurationMdt = "1200"
	LoggingDurationMdt2400 LoggingDurationMdt = "2400"
	LoggingDurationMdt3600 LoggingDurationMdt = "3600"
	LoggingDurationMdt5400 LoggingDurationMdt = "5400"
	LoggingDurationMdt7200 LoggingDurationMdt = "7200"
)

type PositioningMethodMdt string

const (
	PositioningMethodMdtGNSS    PositioningMethodMdt = "GNSS"
	PositioningMethodMdtECellID PositioningMethodMdt = "E_CELL_ID"
)

type CollectionPeriodRmmLteMdt string

const (
	CollectionPeriodRmmLteMdt1024ms  CollectionPeriodRmmLteMdt = "1024"
	CollectionPeriodRmmLteMdt1280ms  CollectionPeriodRmmLteMdt = "1280"
	CollectionPeriodRmmLteMdt2048ms  CollectionPeriodRmmLteMdt = "2048"
	CollectionPeriodRmmLteMdt2560ms  CollectionPeriodRmmLteMdt = "2560"
	CollectionPeriodRmmLteMdt5120ms  CollectionPeriodRmmLteMdt = "5120"
	CollectionPeriodRmmLteMdt10240ms CollectionPeriodRmmLteMdt = "10240"
	CollectionPeriodRmmLteMdt1m      CollectionPeriodRmmLteMdt = "60000"
)

type MeasurementPeriodLteMdt string

const (
	MeasurementPeriodLteMdt1024ms  MeasurementPeriodLteMdt = "1024"
	MeasurementPeriodLteMdt1280ms  MeasurementPeriodLteMdt = "1280"
	MeasurementPeriodLteMdt2048ms  MeasurementPeriodLteMdt = "2048"
	MeasurementPeriodLteMdt2560ms  MeasurementPeriodLteMdt = "2560"
	MeasurementPeriodLteMdt5120ms  MeasurementPeriodLteMdt = "5120"
	MeasurementPeriodLteMdt10240ms MeasurementPeriodLteMdt = "10240"
	MeasurementPeriodLteMdt1m      MeasurementPeriodLteMdt = "60000"
)

type ReportIntervalNrMdt string

const (
	ReportIntervalNrMdt120ms   ReportIntervalNrMdt = "120"
	ReportIntervalNrMdt240ms   ReportIntervalNrMdt = "240"
	ReportIntervalNrMdt480ms   ReportIntervalNrMdt = "480"
	ReportIntervalNrMdt640ms   ReportIntervalNrMdt = "640"
	ReportIntervalNrMdt1024ms  ReportIntervalNrMdt = "1024"
	ReportIntervalNrMdt2048ms  ReportIntervalNrMdt = "2048"
	ReportIntervalNrMdt5120ms  ReportIntervalNrMdt = "5120"
	ReportIntervalNrMdt10240ms ReportIntervalNrMdt = "10240"
	ReportIntervalNrMdt20480ms ReportIntervalNrMdt = "20480"
	ReportIntervalNrMdt40960ms ReportIntervalNrMdt = "40960"
	ReportIntervalNrMdt1m      ReportIntervalNrMdt = "60000"
	ReportIntervalNrMdt6m      ReportIntervalNrMdt = "360000"
	ReportIntervalNrMdt12m     ReportIntervalNrMdt = "720000"
	ReportIntervalNrMdt30m     ReportIntervalNrMdt = "1800000"
	ReportIntervalNrMdt60m     ReportIntervalNrMdt = "3600000"
)

type LoggingIntervalNrMdt string

const (
	LoggingIntervalNrMdt1280ms   LoggingIntervalNrMdt = "1280"
	LoggingIntervalNrMdt2560ms   LoggingIntervalNrMdt = "2560"
	LoggingIntervalNrMdt5120ms   LoggingIntervalNrMdt = "5120"
	LoggingIntervalNrMdt10240ms  LoggingIntervalNrMdt = "10240"
	LoggingIntervalNrMdt20480ms  LoggingIntervalNrMdt = "20480"
	LoggingIntervalNrMdt30720ms  LoggingIntervalNrMdt = "30720"
	LoggingIntervalNrMdt40960ms  LoggingIntervalNrMdt = "40960"
	LoggingIntervalNrMdt61440ms  LoggingIntervalNrMdt = "61440"
	LoggingIntervalNrMdt320ms    LoggingIntervalNrMdt = "320"
	LoggingIntervalNrMdt640ms    LoggingIntervalNrMdt = "640"
	LoggingIntervalNrMdtInfinity LoggingIntervalNrMdt = "infinity"
)

type CollectionPeriodRmmNrMdt string

const (
	CollectionPeriodRmmNrMdt1024ms  CollectionPeriodRmmNrMdt = "1024"
	CollectionPeriodRmmNrMdt2048ms  CollectionPeriodRmmNrMdt = "2048"
	CollectionPeriodRmmNrMdt5120ms  CollectionPeriodRmmNrMdt = "5120"
	CollectionPeriodRmmNrMdt10240ms CollectionPeriodRmmNrMdt = "10240"
	CollectionPeriodRmmNrMdt1m      CollectionPeriodRmmNrMdt = "60000"
)

type LoggingDurationNrMdt string

const (
	LoggingDurationNrMdt600  LoggingDurationNrMdt = "600"
	LoggingDurationNrMdt1200 LoggingDurationNrMdt = "1200"
	LoggingDurationNrMdt2400 LoggingDurationNrMdt = "2400"
	LoggingDurationNrMdt3600 LoggingDurationNrMdt = "3600"
	LoggingDurationNrMdt5400 LoggingDurationNrMdt = "5400"
	LoggingDurationNrMdt7200 LoggingDurationNrMdt = "7200"
)

type TraceDepth string

const (
	TraceDepthMinimum                     TraceDepth = "MINIMUM"
	TraceDepthMedium                      TraceDepth = "MEDIUM"
	TraceDepthMaximum                     TraceDepth = "MAXIMUM"
	TraceDepthMINIMUM_WO_VENDOR_EXTENSION TraceDepth = "MINIMUM_WO_VENDOR_EXTENSION"
	TraceDepthMEDIUM_WO_VENDOR_EXTENSION  TraceDepth = "MEDIUM_WO_VENDOR_EXTENSION"
	TraceDepthMAXIMUM_WO_VENDOR_EXTENSION TraceDepth = "MAXIMUM_WO_VENDOR_EXTENSION"
)

type MdtConfiguration struct {
	JobType                JobType
	ReportType             ReportTypeMdt
	AreaScope              AreaScope
	MeasurementLteList     []MeasurementLteForMdt
	MeasurementNrList      []MeasurementNrForMdt
	SensorMeasurementList  []SensorMeasurement
	ReportingTriggerList   []ReportingTrigger
	ReportInterval         ReportIntervalMdt
	ReportIntervalNr       ReportIntervalNrMdt
	ReportAmount           ReportAmountMdt
	EventThresholdRsrp     int
	EeventThresholdRsrpNr  int
	EventThresholdRsrq     int
	EventThresholdRsrqNr   int
	EventList              []EventForMdt
	LoggingInterval        LoggingIntervalMdt
	LoggingIntervalNr      LoggingIntervalNrMdt
	LoggingDuration        LoggingDurationMdt
	LoggingDurationNr      LoggingDurationNrMdt
	PositioningMethod      PositioningMethodMdt
	CollectionPeriodRmmLte CollectionPeriodRmmLteMdt
	CollectionPeriodRmmNr  CollectionPeriodRmmNrMdt
	MeasurementPeriodLte   MeasurementPeriodLteMdt
	MdtAllowedPlmnIdList   []PlmnId
	MbsfnAreaList          []MbsfnArea
	InterFreqTargetList    []InterFreqTargetInfo
}

type TraceData struct {
	TraceRef                 string
	TraceDepth               TraceDepth
	NeTypeList               string
	EventList                string
	CollectionEntityIpv4Addr Ipv4Addr
	CollectionEntityIpv6Addr Ipv6Addr
	InterfaceList            string
}
