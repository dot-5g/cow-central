package model

type Qfi int

type QfiRm int

type FiveQi int

type FiveQiRm int

type BitRate string

type BitRateRm string

type ArpPriorityLevel int

type ArpPriorityLevelRm int

type FiveQiPriorityLevel int

type FiveQiPriorityLevelRm int

type PacketDelBudget int

type PacketDelBudgetRm int

type PacketErrRate string

type PacketErrRateRm string

type PacketLossRate int

type PacketLossRateRm int

type AverWindow int

type AverWindowRm int

type MaxDataBurstVol int

type MaxDataBurstVolRm int

type SamplingRatio int

type SamplingRatioRm int

type RgWirelineCharacteristics byte

type RgWirelineCharacteristicsRm byte

type ExtMaxDataBurstVol int

type ExtMaxDataBurstVolRm int

type ExtPacketDelBudget int

type ExtPacketDelBudgetRm int

type AmbrRm struct {
	Uplink   BitRate
	Downlink BitRate
}
