package mock_store

import (
	"github.com/golang/mock/gomock"
	"github.com/rl-os/api/store"
)

// Interface assertion
var _ store.Store = (*MockedStore)(nil)

type MockedStore struct {
	ctrl *gomock.Controller

	beatmap    *MockBeatmap
	beatmapSet *MockBeatmapSet
	user       *MockUser
	oauth      *MockOAuth
	friend     *MockFriend
	chat       *MockChat
}

func InitStore(ctrl *gomock.Controller) MockedStore {
	return MockedStore{
		ctrl:       ctrl,
		beatmap:    NewMockBeatmap(ctrl),
		beatmapSet: NewMockBeatmapSet(ctrl),
	}
}

func (ss MockedStore) Close() {}

func (ss MockedStore) Beatmap() store.Beatmap       { return ss.beatmap }
func (ss MockedStore) BeatmapSet() store.BeatmapSet { return ss.beatmapSet }
func (ss MockedStore) User() store.User             { return ss.user }
func (ss MockedStore) OAuth() store.OAuth           { return ss.oauth }
func (ss MockedStore) Friend() store.Friend         { return ss.friend }
func (ss MockedStore) Chat() store.Chat             { return ss.chat }

func (ss MockedStore) BeatmapExpect() *MockBeatmapMockRecorder       { return ss.beatmap.EXPECT() }
func (ss MockedStore) BeatmapSetExpect() *MockBeatmapSetMockRecorder { return ss.beatmapSet.EXPECT() }
func (ss MockedStore) UserExpect() *MockUserMockRecorder             { return ss.user.EXPECT() }
func (ss MockedStore) OAuthExpect() *MockOAuthMockRecorder           { return ss.oauth.EXPECT() }
func (ss MockedStore) FriendExpect() *MockFriendMockRecorder         { return ss.friend.EXPECT() }
func (ss MockedStore) ChatExpect() *MockChatMockRecorder             { return ss.chat.EXPECT() }
