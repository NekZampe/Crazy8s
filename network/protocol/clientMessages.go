package protocol

const (
	MsgJoin   byte = 0x01
	MsgLeave  byte = 0x02
	MsgPlay   byte = 0x03
	MsgCrazy8 byte = 0x04
	MsgSkip   byte = 0x05
)

func BuildJoin(playerID byte) []byte {
	return buildPacket(MsgJoin, playerID, "")
}

func BuildLeave(playerID byte) []byte {
	return buildPacket(MsgLeave, playerID, "")
}

func BuildPlay(playerID byte, cards string) []byte {
	return buildPacket(MsgPlay, playerID, cards)
}

func BuildCrazy8(playerID byte, suit string) []byte {
	return buildPacket(MsgCrazy8, playerID, suit)
}

func BuildSkip(playerID byte) []byte {
	return buildPacket(MsgSkip, playerID, "")
}
