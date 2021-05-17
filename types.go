package go_tgbot

import "encoding/json"

type BotAPI struct {
	Token string
	Debug bool
}

type Error struct {
	Code int
	Message string
	ResponseParameters
}

func (e Error) Error() string {
	return e.Message
}


type APIResponse struct {
	Ok          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result"`
	ErrorCode   int                 `json:"error_code"`
	Description string              `json:"description"`
	Parameters  *ResponseParameters `json:"parameters"`
}

type ResponseParameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id"`
	RetryAfter      int   `json:"retry_after"`
}

// Chat contains information about the place a message was sent.
type Chat struct {
	ID int64 `json:"id"`
	Type string `json:"type"`
	Title string `json:"title"`
	UserName string `json:"username"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Photo *ChatPhoto `json:"photo"`
	AllMembersAreAdmins bool `json:"all_members_are_administrators"`
	Bio string `json:"bio"`
	Description string `json:"description,omitempty"`
	InviteLink string `json:"invite_link,omitempty"`
	//PinnedMessage *Message `json:"pinned_message"`
	Permissions *ChatPermissions `json:"permissions"`
	SlowModeDelay int `json:"slow_mode_delay"`
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
	StickerSetName string `json:"sticker_set_name"`
	CanSetStickerSet bool `json:"can_set_sticker_set"`
	LinkedChatId int `json:"linked_chat_id"`
	Location *ChatLocation
}

type ChatPhoto struct {
	SmallFileId string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
	BigFileId string `json:"big_file_id"`
	BigFileUniqueId string `json:"big_file_unique_id"`
}

type ChatLocation struct {
	Location *Location `json:"location"`
	Address string     `json:"address"`
}

type ChatPermissions struct {
	CanSendMessages bool `json:"can_send_messages"`
	CanSendMediaMessages bool `json:"can_send_media_messages"`
	CanSendPolls bool `json:"can_send_polls"`
	CanSendOtherMessages bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	CanChangeInfo bool `json:"can_change_info"`
	CanInviteUsers bool `json:"can_invite_users"`
	CanPinMessages bool `json:"can_pin_messages"`
}

type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
	HorizontalAccuracy float64 `json:"horizontal_accuracy"`
	LivePeriod int `json:"live_period"`
	Heading int `json:"heading"`
	ProximityAlertRadius int `json:"proximity_alert_radius"`
}

type User struct {
	Id int `json:"id"`
	IsBot bool `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	LanguageCode string `json:"language_code"`
	CanJoinGroups bool `json:"can_join_groups"`
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages"`
	SupportsInlineQueries bool `json:"supports_inline_queries"`
}