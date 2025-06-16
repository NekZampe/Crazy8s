package protocol

import (
	"encoding/binary"
	"io"
	"net"
)

func buildPacket(msgType byte, playerID byte, payload string) []byte {
	payloadBytes := []byte(payload)
	packet := make([]byte, 4+len(payloadBytes))
	packet[0] = msgType
	packet[1] = playerID
	binary.BigEndian.PutUint16(packet[2:4], uint16(len(payloadBytes)))
	copy(packet[4:], payloadBytes)
	return packet
}

func SendMessage(conn net.Conn, data []byte) error {
	_, err := conn.Write(data)
	return err
}

func ReadMessage(conn net.Conn) (msgType byte, playerID byte, payload string, err error) {
	header := make([]byte, 4)
	if _, err = io.ReadFull(conn, header); err != nil {
		return
	}
	msgType = header[0]
	playerID = header[1]
	length := binary.BigEndian.Uint16(header[2:4])
	payloadBytes := make([]byte, length)
	if _, err = io.ReadFull(conn, payloadBytes); err != nil {
		return
	}
	payload = string(payloadBytes)
	return
}
