package server

type LobbyInfo struct {
	HostName   string `json:"HostName"`
	LobbyID    string `json:"lobby_id"`
	Players    int    `json:"players"`
	MaxPlayers int    `json:"max_players"`
}
