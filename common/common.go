package common

import (
	"github.com/bwmarrin/discordgo"
	"github.com/fzzy/radix/extra/pool"
	"github.com/jonas747/yagpdb/common/pastebin"
)

var (
	RedisPool  *pool.Pool
	BotSession *discordgo.Session
	Conf       *Config
	Pastebin   *pastebin.Pastebin
)
