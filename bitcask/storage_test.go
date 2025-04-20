package main

import (
	"encoding/json"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	t.Run("put get", func(t *testing.T) {
		s := newStore(NewMemBuffer())

		data := map[string]any{
			"foo": "bar",
			"asd": 123,
			"bool": true,
			"arr": []int{1,2,3},
		}

		serialized, err := json.Marshal(data)
		assert.NoError(t, err)

		s.Put("abc", serialized)
		
		got, err := s.Get("abc")
		assert.NoError(t, err)
		assert.JSONEq(t, string(serialized), string(got))
	})

	t.Run("check state of storage", func(t *testing.T) {
		f := NewMemBuffer()
		s := newStore(f)

		data := []byte("this is a long string")
		key := "abc"
		s.Put(key, data)

		expected := []byte{byte(len(key)), 0, byte(len(data)),0,0,0} // u16, u32 in little endian
		expected = append(expected, []byte(key)...)
		expected = append(expected, data...)

		assert.Equal(t, expected, f.buf.Bytes()[8:]) // skip crc and timestamp
	})

	t.Run("put get typed", func(t *testing.T) {
		s := newStore(NewMemBuffer())

		data := map[string]string{
			"foo": "bar",
			"asd": "123",
		}

		serialized, err := json.Marshal(data)
		assert.NoError(t, err)

		s.Put("abc", serialized)
		
		got, err := s.Get("abc")
		assert.NoError(t, err)

		unmarshalled := map[string]string{}
		err = json.Unmarshal(got, &unmarshalled)
		assert.NoError(t, err)

		assert.Equal(t, data, unmarshalled)
	})

	t.Run("get unknown key", func(t *testing.T) {
		s := newStore(NewMemBuffer())
		_, err := s.Get("unknown")
		assert.Error(t, err)
		assert.ErrorIs(t, err, errKeyNotFound)
	})

	t.Run("delete key", func(t *testing.T) {
		key := "abc"
		s := newStore(NewMemBuffer())

		s.Put(key, []byte("some data"))

		s.Delete(key)

		got, err := s.Get(key)
		assert.Empty(t, got)
		assert.Error(t, err)
		assert.ErrorIs(t, err, errKeyNotFound)
	})

	t.Run("multiple inserts", func(t *testing.T) {
		data := map[string]string {
			"foo": "123432432wea",
			"bar": "asdfasdfas234",
			"asd": "12231dfsjkhfdskjghjfdsa",
			"baz": "asdfasdffdasgsagfds",
		}

		s := newStore(NewMemBuffer())

		for k, v := range data {
			s.Put(k, []byte(v))
		}

		assert.ElementsMatch(t, []string{"foo","bar","asd","baz"}, s.Keys())

		for k, v := range data {
			got, err := s.Get(k)
			assert.NoError(t, err, "getting key %q", k)
			assert.Equal(t, v, string(got), "getting value for key %q", k)
		}
	})
}

func TestMerging(t *testing.T) {
	data := map[string]string {
		"foo": "123432432wea",
		"bar": "asdfasdfas234",
		"asd": "12231dfsjkhfdskjghjfdsa",
		"baz": "asdfasdffdasgsagfds",
	}

	oldFile := NewMemBuffer()
	oldStore := newStore(oldFile)

	for k, v := range data {
		oldStore.Put(k, []byte(v))
		time.Sleep(1*time.Millisecond)
	}

	oldStore.Put("foo", []byte("updated foo"))
	oldStore.Delete("asd")

	expectedData := map[string]string {
		"foo": "updated foo",
		"bar": "asdfasdfas234",
		"baz": "asdfasdffdasgsagfds",
	}

	oldFile.Seek(0, io.SeekStart)
	newStore, err := newStoreWithRebuild([]io.ReadWriteSeeker{oldFile}, NewMemBuffer())
	assert.NoError(t, err)

	assert.ElementsMatch(t, []string{"foo","bar","baz"}, newStore.Keys())
	
	for k, v := range expectedData {
		got, err := newStore.Get(k)
		assert.NoError(t, err, "getting key %q", k)
		assert.Equal(t, v, string(got), "getting value for key %q", k)
	}
}