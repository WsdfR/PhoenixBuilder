package packet

import (
	"bytes"
	"encoding/binary"
	"phoenixbuilder/minecraft/protocol"
)

const (
	AdventureFlagWorldImmutable = 1 << iota
	AdventureFlagNoPVP
	_
	_
	_
	AdventureFlagAutoJump
	AdventureFlagAllowFlight
	AdventureFlagNoClip
	AdventureFlagWorldBuilder
	AdventureFlagFlying
	AdventureFlagMuted
)

const (
	CommandPermissionLevelNormal = iota
	CommandPermissionLevelOperator
	CommandPermissionLevelHost
	CommandPermissionLevelAutomation
	CommandPermissionLevelAdmin
)

const (
	ActionPermissionBuildAndMine = 1 << iota
	ActionPermissionDoorsAndSwitched
	ActionPermissionOpenContainers
	ActionPermissionAttackPlayers
	ActionPermissionAttackMobs
	ActionPermissionOperator
	ActionPermissionTeleport
)

const (
	PermissionLevelVisitor = iota
	PermissionLevelMember
	PermissionLevelOperator
	PermissionLevelCustom
)

// AdventureSettings is sent by the server to update game-play related features, in particular permissions to
// access these features for the client. It includes allowing the player to fly, build and mine, and attack
// entities. Most of these flags should be checked server-side instead of using this packet only.
// The client may also send this packet to the server when it updates one of these settings through the
// in-game settings interface. The server should verify if the player actually has permission to update those
// settings.
type AdventureSettings struct {
	// Flags is a set of flags that specify certain properties of the player, such as whether or not it can
	// fly and/or move through blocks. It is one of the AdventureFlag constants above.
	Flags uint32
	// CommandPermissionLevel is a permission level that specifies the kind of commands that the player is
	// allowed to use. It is one of the CommandPermissionLevel constants above.
	CommandPermissionLevel uint32
	// ActionPermissions is, much like Flags, a set of flags that specify actions that the player is allowed
	// to undertake, such as whether it is allowed to edit blocks, open doors etc. It is a combination of the
	// ActionPermission constants above.
	ActionPermissions uint32
	// PermissionLevel is the permission level of the player as it shows up in the player list built up using
	// the PlayerList packet. It is one of the PermissionLevel constants above.
	PermissionLevel uint32
	// CustomStoredPermissions ...
	CustomStoredPermissions uint32
	// PlayerUniqueID is a unique identifier of the player. It appears it is not required to fill this field
	// out with a correct value. Simply writing 0 seems to work.
	PlayerUniqueID int64
}

// ID ...
func (*AdventureSettings) ID() uint32 {
	return IDAdventureSettings
}

// Marshal ...
func (pk *AdventureSettings) Marshal(buf *bytes.Buffer) {
	_ = protocol.WriteVaruint32(buf, pk.Flags)
	_ = protocol.WriteVaruint32(buf, pk.CommandPermissionLevel)
	_ = protocol.WriteVaruint32(buf, pk.ActionPermissions)
	_ = protocol.WriteVaruint32(buf, pk.PermissionLevel)
	_ = protocol.WriteVaruint32(buf, pk.CustomStoredPermissions)
	_ = binary.Write(buf, binary.LittleEndian, pk.PlayerUniqueID)
}

// Unmarshal ...
func (pk *AdventureSettings) Unmarshal(buf *bytes.Buffer) error {
	return chainErr(
		protocol.Varuint32(buf, &pk.Flags),
		protocol.Varuint32(buf, &pk.CommandPermissionLevel),
		protocol.Varuint32(buf, &pk.ActionPermissions),
		protocol.Varuint32(buf, &pk.PermissionLevel),
		protocol.Varuint32(buf, &pk.CustomStoredPermissions),
		binary.Read(buf, binary.LittleEndian, &pk.PlayerUniqueID),
	)
}
