package mongodb

import (
	"context"
	"testing"
	"time"

	"github.com/4726/twitch-chat-logger/storage"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func testClient(t *testing.T) *Storage {
	addr := "mongodb://192.168.1.232:27017"
	db := "db_test"
	coll := "messages"
	s := New(addr, db, coll)
	assert.NoError(t, s.Connect())
	c := s.client.Database(db).Collection(coll)
	c.DeleteMany(context.Background(), bson.M{})
	return s
}

func TestChannel(t *testing.T) {
	s := testClient(t)
	defer s.Close()
	m1 := storage.ChatMessage{Channel: "channel"}
	m2 := storage.ChatMessage{Channel: "channel2"}
	s.Add(m1)
	s.Add(m2)
	opt := storage.QueryOptions{Channel: "channel"}
	res, err := s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m1}, res)
}

func TestTerm(t *testing.T) {
	s := testClient(t)
	defer s.Close()
	m1 := storage.ChatMessage{Message: "hello world"}
	m2 := storage.ChatMessage{Message: "helloooo"}
	m3 := storage.ChatMessage{Message: "hel lo"}
	s.Add(m1)
	s.Add(m2)
	s.Add(m3)
	opt := storage.QueryOptions{Term: "world"}
	res, err := s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m1}, res)

	opt = storage.QueryOptions{Term: "hello world"}
	res, err = s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m1}, res)
}

func TestName(t *testing.T) {
	s := testClient(t)
	defer s.Close()
	m1 := storage.ChatMessage{Name: "channel", DisplayName: "channel"}
	m2 := storage.ChatMessage{Name: "channel2", DisplayName: "channel2"}
	s.Add(m1)
	s.Add(m2)
	opt := storage.QueryOptions{Name: "channel"}
	res, err := s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m1}, res)
}

func TestDate(t *testing.T) {
	s := testClient(t)
	defer s.Close()
	m1 := storage.ChatMessage{Name: "channel", Time: time.Now().Unix()}
	m2 := storage.ChatMessage{Name: "channel2", Time: time.Now().Unix()}
	s.Add(m1)
	s.Add(m2)
	opt := storage.QueryOptions{Date: time.Now()}
	res, err := s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m1, m2}, res)

	opt = storage.QueryOptions{Date: time.Now().AddDate(0, 0, -2)}
	res, err = s.Query(opt)
	assert.NoError(t, err)
	assert.Len(t, res, 0)
}

func TestSubscribeMin(t *testing.T) {
	s := testClient(t)
	defer s.Close()
	m1 := storage.ChatMessage{SubscribeMonths: 0}
	m2 := storage.ChatMessage{SubscribeMonths: 5}
	s.Add(m1)
	s.Add(m2)
	opt := storage.QueryOptions{SubscribeMin: 4}
	res, err := s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m2}, res)

	opt = storage.QueryOptions{SubscribeMin: 5}
	res, err = s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m2}, res)

	opt = storage.QueryOptions{SubscribeMin: 0}
	res, err = s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m1, m2}, res)

	opt = storage.QueryOptions{SubscribeMin: 6}
	res, err = s.Query(opt)
	assert.NoError(t, err)
	assert.Len(t, res, 0)
}

func TestAdmin(t *testing.T) {
	s := testClient(t)
	defer s.Close()
	m1 := storage.ChatMessage{Admin: false}
	m2 := storage.ChatMessage{Admin: true}
	s.Add(m1)
	s.Add(m2)
	opt := storage.QueryOptions{Admin: true}
	res, err := s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m2}, res)

	opt = storage.QueryOptions{Admin: false}
	res, err = s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m1, m2}, res)
}

func TestGlobalMod(t *testing.T) {
	s := testClient(t)
	defer s.Close()
	m1 := storage.ChatMessage{GlobalMod: false}
	m2 := storage.ChatMessage{GlobalMod: true}
	s.Add(m1)
	s.Add(m2)
	opt := storage.QueryOptions{GlobalMod: true}
	res, err := s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m2}, res)

	opt = storage.QueryOptions{GlobalMod: false}
	res, err = s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m1, m2}, res)
}

func TestModerator(t *testing.T) {
	s := testClient(t)
	defer s.Close()
	m1 := storage.ChatMessage{Moderator: false}
	m2 := storage.ChatMessage{Moderator: true}
	s.Add(m1)
	s.Add(m2)
	opt := storage.QueryOptions{Moderator: true}
	res, err := s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m2}, res)

	opt = storage.QueryOptions{Moderator: false}
	res, err = s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m1, m2}, res)
}

func TestStaff(t *testing.T) {
	s := testClient(t)
	defer s.Close()
	m1 := storage.ChatMessage{Staff: false}
	m2 := storage.ChatMessage{Staff: true}
	s.Add(m1)
	s.Add(m2)
	opt := storage.QueryOptions{Staff: true}
	res, err := s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m2}, res)

	opt = storage.QueryOptions{Staff: false}
	res, err = s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m1, m2}, res)
}

func TestTurbo(t *testing.T) {
	s := testClient(t)
	defer s.Close()
	m1 := storage.ChatMessage{Turbo: false}
	m2 := storage.ChatMessage{Turbo: true}
	s.Add(m1)
	s.Add(m2)
	opt := storage.QueryOptions{Turbo: true}
	res, err := s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m2}, res)

	opt = storage.QueryOptions{Turbo: false}
	res, err = s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m1, m2}, res)
}

func TestBits(t *testing.T) {
	s := testClient(t)
	defer s.Close()
	m1 := storage.ChatMessage{Bits: 0}
	m2 := storage.ChatMessage{Bits: 1000}
	s.Add(m1)
	s.Add(m2)
	opt := storage.QueryOptions{BitsMin: 100, BitsMax: 500}
	res, err := s.Query(opt)
	assert.NoError(t, err)
	assert.Len(t, res, 0)

	opt = storage.QueryOptions{BitsMin: 0, BitsMax: 1000}
	res, err = s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m1, m2}, res)

	opt = storage.QueryOptions{BitsMin: 10000, BitsMax: 20000}
	res, err = s.Query(opt)
	assert.NoError(t, err)
	assert.Len(t, res, 0)

	opt = storage.QueryOptions{BitsMin: 0, BitsMax: 0}
	res, err = s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m1, m2}, res)
}
