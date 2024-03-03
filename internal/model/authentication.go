package model

type SorMac string

type CounterSor string

type AckInd bool

type SecuredPacket string

type RoutingId string

type UpuMac string

type CounterUpu string

type UpuAckInd bool

type UpuData struct {
	SecPacket        SecuredPacket
	DefaultConfNssai []Snssai
	RoutingId        RoutingId
}
