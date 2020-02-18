package oauth

type Scope string

var ProfileScope Scope = "profile"
var ChatsScope Scope = "chats"
var FriendsScope Scope = "friends"
var RoomsScope Scope = "rooms"
var CommentsScope Scope = "comments"
var ScoresScope Scope = "scores"
var NotifsScope Scope = "notifs"

var Scopes = []Scope{ProfileScope, ChatsScope, FriendsScope, RoomsScope, CommentsScope, ScoresScope, NotifsScope}
