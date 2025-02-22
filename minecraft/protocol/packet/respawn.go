package packet

import (
	"bytes"
	"encoding/binary"
	"github.com/go-gl/mathgl/mgl32"
	"phoenixbuilder/minecraft/protocol"
)

const (
	// Respawn packets with these states are sent by the server.
	RespawnStateSearchingForSpawn = iota
	RespawnStateReadyToSpawn

	// A Respawn packet with this state is sent by the client.
	RespawnStateClientReadyToSpawn
)

// Respawn is sent by the server to make a player respawn client-side. It is sent in response to a
// PlayerAction packet with ActionType PlayerActionRespawn.
// As of 1.13, the server sends two of these packets with different states, and the client sends one of these
// back in order to complete the respawn.
type Respawn struct {
	// Position is the position on which the player should be respawned. The position might be in a different
	// dimension, in which case the client should first be sent a ChangeDimension packet.
	Position mgl32.Vec3
	// State is the 'state' of the respawn. It is one of the constants that may be found above, and the value
	// the packet contains depends on whether the server or client sends it.
	State byte
	// EntityRuntimeID is the entity runtime ID of the player that the respawn packet concerns. This is
	// apparently for the server to recognise which player sends this packet.
	EntityRuntimeID uint64
}

// ID ...
func (*Respawn) ID() uint32 {
	return IDRespawn
}

// Marshal ...
func (pk *Respawn) Marshal(buf *bytes.Buffer) {
	_ = protocol.WriteVec3(buf, pk.Position)
	_ = binary.Write(buf, binary.LittleEndian, pk.State)
	_ = protocol.WriteVaruint64(buf, pk.EntityRuntimeID)
}

// Unmarshal ...
func (pk *Respawn) Unmarshal(buf *bytes.Buffer) error {
	return chainErr(
		protocol.Vec3(buf, &pk.Position),
		binary.Read(buf, binary.LittleEndian, &pk.State),
		protocol.Varuint64(buf, &pk.EntityRuntimeID),
	)
}
