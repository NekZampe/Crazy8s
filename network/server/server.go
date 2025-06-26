package server

import (
	"Crazy8s/ilogger"
	"Crazy8s/network/protocol"
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Listener net.Listener  // TCP listener
	Lobby    *Lobby        // The game lobby
	QuitChan chan struct{} // To signal shutdown
	logger   ilogger.Logger
	idMu     sync.Mutex    // protects idAlloc
	idAlloc  map[byte]bool // true if ID is assigned
}

func (s *Server) Start(address string) error {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	log, err := ilogger.NewFileLogger("server.log")
	if err != nil {
		return err
	}

	s.logger = log
	s.Listener = ln
	s.QuitChan = make(chan struct{})
	s.Lobby = NewLobby()
	s.logger.Debug("Server Started Successfully...")
	s.logger.Info("Accepting new connections")
	go s.acceptLoop()

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			select {
			case <-s.QuitChan:
				return // server is shutting down
			default:
				// log the error, maybe retry
			}
		}

		// assign player ID, add to lobby, start handling
		playerID := s.generateUniqueID()
		s.Lobby.AddPlayer(playerID, conn)
		s.logger.Debug(fmt.Sprintf("[Lobby] Player %d joined", playerID))
	}
}

func (s *Server) Stop() {
	close(s.QuitChan)
	err := s.Listener.Close()
	if err != nil {
		return
	}
}

func (s *Server) generateUniqueID() byte {
	s.idMu.Lock()
	defer s.idMu.Unlock()

	// Initialize idAlloc if nil
	if s.idAlloc == nil {
		s.idAlloc = make(map[byte]bool)
	}

	for id := byte(1); id < 255; id++ { // skip 0 if you want
		if !s.idAlloc[id] {
			s.idAlloc[id] = true
			return id
		}
	}
	// No available ID
	return 0
}

func (s *Server) freeID(id byte) {
	s.idMu.Lock()
	defer s.idMu.Unlock()

	delete(s.idAlloc, id)
}

func (s *Server) handlePlayer(playerID byte, conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			s.logger.Debug(fmt.Sprintf("Error closing connection for player %d: %v", playerID, err))
		}
		s.Lobby.RemovePlayer(playerID)
		s.freeID(playerID)
	}()

	for {
		msgType, pid, payload, err := protocol.ReadMessage(conn)
		if err != nil {
			s.logger.Debug(fmt.Sprintf("Player %d disconnected or error: %v", playerID, err))
			return
		}

		s.logger.Debug(fmt.Sprintf("Received msg type %d from player %d: %s", msgType, pid, payload))

		switch msgType {
		case 0x03:
			//Handle Play Request
			s.processPlayRequest(pid, payload)
		case 0x04:
			//Handle Skip Request
			s.processSkipRequest(pid)
		case 0x05:
			//Handle Crazy8 Suit Request
			s.processCrazy8SuitRequest(pid, payload)
		default:

		}

	}
}
