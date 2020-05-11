package app

import (
	"github.com/4726/twitch-chat-logger/storage"
	twitch "github.com/gempir/go-twitch-irc/v2"
)

func privateMessageToStorageMessage(privmsg twitch.PrivateMessage) storage.ChatMessage {
	cm := storage.ChatMessage{
		Channel:     privmsg.Channel,
		Message:     privmsg.Message,
		RoomID:      privmsg.RoomID,
		ID:          privmsg.ID,
		Time:        privmsg.Time.Unix(),
		Bits:        privmsg.Bits,
		UserID:      privmsg.User.ID,
		Name:        privmsg.User.Name,
		DisplayName: privmsg.User.DisplayName,
		// SubscribMonths: privmsg.Tags["badge-info"],
	}
	if _, ok := privmsg.User.Badges["admin"]; ok {
		cm.Admin = true
	}
	if _, ok := privmsg.User.Badges["global_mod"]; ok {
		cm.GlobalMod = true
	}
	if _, ok := privmsg.User.Badges["moderator"]; ok {
		cm.Moderator = true
	}
	if _, ok := privmsg.User.Badges["staff"]; ok {
		cm.Staff = true
	}
	if _, ok := privmsg.User.Badges["turbo"]; ok {
		cm.Turbo = true
	}
	if _, ok := privmsg.User.Badges["subscriber"]; ok {
		cm.Subscriber = true
	}
	return cm
}
