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
	m1 := storage.ChatMessage{Channel: "channel", ID: "1"}
	m2 := storage.ChatMessage{Channel: "channel2", ID: "2"}
	assert.NoError(t, s.Add(m1))
	assert.NoError(t, s.Add(m2))
	opt := storage.QueryOptions{Channel: "channel"}
	res, err := s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m1}, res)
}

func TestName(t *testing.T) {
	s := testClient(t)
	defer s.Close()
	m1 := storage.ChatMessage{Name: "channel", DisplayName: "channel", ID: "1"}
	m2 := storage.ChatMessage{Name: "channel2", DisplayName: "channel2", ID: "2"}
	assert.NoError(t, s.Add(m1))
	assert.NoError(t, s.Add(m2))
	opt := storage.QueryOptions{Name: "channel"}
	res, err := s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m1}, res)
}

func TestChannelName(t *testing.T) {
	s := testClient(t)
	defer s.Close()
	m1 := storage.ChatMessage{Name: "name", Channel: "channel", ID: "1"}
	m2 := storage.ChatMessage{Name: "name", Channel: "channel", ID: "2"}
	m3 := storage.ChatMessage{Name: "name1", Channel: "channel", ID: "3"}
	m4 := storage.ChatMessage{Name: "name", Channel: "channel1", ID: "4"}
	assert.NoError(t, s.Add(m1))
	assert.NoError(t, s.Add(m2))
	assert.NoError(t, s.Add(m3))
	assert.NoError(t, s.Add(m4))
	opt := storage.QueryOptions{Name: "name", Channel: "channel"}
	res, err := s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m1, m2}, res)
}

func TestDate(t *testing.T) {
	s := testClient(t)
	defer s.Close()
	m1 := storage.ChatMessage{Name: "channel", Time: time.Now().Unix(), ID: "1"}
	m2 := storage.ChatMessage{Name: "channel2", Time: time.Now().Unix(), ID: "2"}
	assert.NoError(t, s.Add(m1))
	assert.NoError(t, s.Add(m2))
	opt := storage.QueryOptions{Date: time.Now()}
	res, err := s.Query(opt)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []storage.ChatMessage{m1, m2}, res)

	opt = storage.QueryOptions{Date: time.Now().AddDate(0, 0, -2)}
	res, err = s.Query(opt)
	assert.NoError(t, err)
	assert.Len(t, res, 0)
}

func TestUniqueID(t *testing.T) {
	s := testClient(t)
	defer s.Close()
	m1 := storage.ChatMessage{Channel: "channel", ID: "1"}
	m2 := storage.ChatMessage{Channel: "channel2", ID: "1"}
	assert.NoError(t, s.Add(m1))
	assert.Error(t, s.Add(m2))
}
