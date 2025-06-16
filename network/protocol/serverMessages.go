package protocol

const (
	MsgStartGame   byte = 0x06
	MsgRequestTurn byte = 0x07
	MsgRequestSuit byte = 0x08
	MsgPickup2     byte = 0x09
	MsgMissTurn    byte = 0x10
	MsgLeaderBoard byte = 0x11
	MsgLastCard    byte = 0x12
	MsgWinner      byte = 0x13
)

func BuildStartGame(msg string) []byte {
	return buildPacket(MsgStartGame, 0xFF, msg)
}

func BuildRequestTurn(playerID byte, msg string) []byte {
	return buildPacket(MsgRequestTurn, playerID, msg)
}

func BuildRequestSuit(playerID byte) []byte {
	return buildPacket(MsgRequestSuit, playerID, "Please select a suit:")
}

func BuildPickup2(playerID byte) []byte {
	return buildPacket(MsgPickup2, playerID, "You pick up 2 cards!")
}

func BuildMissTurn(playerID byte) []byte {
	return buildPacket(MsgMissTurn, playerID, "You miss your turn :( ")
}

func BuildLeaderboard(msg string) []byte {
	return buildPacket(MsgLeaderBoard, 0xFF, msg)
}

func BuildLastCard(msg string) []byte {
	return buildPacket(MsgLastCard, 0xFF, msg)
}

func BuildWinner(msg string) []byte {
	return buildPacket(MsgWinner, 0xFF, msg)
}
