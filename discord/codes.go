package discord

// Voice related op codes
const (
	StateUpdate  = 4
	ServerUpdate = 0
	Identify
)

type Gateway struct {
	Op int `json:"op"`
	D  interface{}
}

type VoiceStateUpdate struct {
	GuildId   int  `json:"guild_id"`
	ChannelId int  `json:"channel_id"`
	SelfMute  bool `json:"self_mute"`
	SelfDeaf  bool `json:"self_deaf"`
}
