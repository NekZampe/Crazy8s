package server

type PlayerMessage struct {
	PlayerID byte
	MsgType  byte
	Payload  string
}

func (s *Server) processPlayRequest(playerID byte, payload string) {

}

func (s *Server) processSkipRequest(playerID byte) {

}

func (s *Server) processCrazy8SuitRequest(playerID byte, payload string) {

}
