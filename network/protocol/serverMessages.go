package protocol

const (
	MsgStartGame           byte = 0x10 // msg to send to players show game is starting
	MsgRequestTurn         byte = 0x11 // Request player to take their turn, send data needed to do so
	MsgRequestSuit         byte = 0x12 // Request suit from player upon playing a crazy8
	MsgPickup2             byte = 0x13 // Notify player they are picking up a 2
	MsgMissTurn            byte = 0x14 // Notify player they missed their turn
	MsgLeaderBoard         byte = 0x15 // Send leaderboard data to user
	MsgLastCard            byte = 0x16 // Warn players one person is on their last card
	MsgWinner              byte = 0x17 // Notify players someone has won
	MsgBroadcastLobbyState byte = 0x18
	MsgBroadCastLatestPlay byte = 0x19
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

func BuildBroadcastLobbyState(msg string) []byte {
	return buildPacket(MsgBroadcastLobbyState, 0xFF, msg)
}

func BuildBroadcastLatestPlay(msg string) []byte {
	return buildPacket(MsgBroadCastLatestPlay, 0xFF, msg)
}
