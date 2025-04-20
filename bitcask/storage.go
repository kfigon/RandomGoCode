package main

import (
	"errors"
	"fmt"
	"io"
	"time"
)


type Key []byte

type LogEntry struct {
	crc uint32
	ts uint32
	keySize uint16
	valSize uint32
	key Key
	value []byte
}

func now() uint32 {
	return uint32(time.Now().UnixMilli())
}

func newLog(key Key, val []byte) LogEntry {
	return LogEntry{
		crc: 0xdeadbeef, // todo
		ts: now(),
		keySize: uint16(len(key)),
		valSize: uint32(len(val)),
		key: key,
		value: val,
	}
}

type KeyDirEntry struct{
	fileId string
	valueSize uint32
	valuePos uint
	timestamp uint32
}

// todo: can't use []byte as map key, use string for now
type KeyDir map[string]KeyDirEntry

type Store struct {
	activeFile io.ReadWriteSeeker
	dir KeyDir
}

func newStore(f io.ReadWriteSeeker) *Store {
	return &Store{
		activeFile: f,
		dir: KeyDir{},
	}
}

func newStoreWithRebuild(files []io.ReadWriteSeeker, f io.ReadWriteSeeker) (*Store, error) {
	// skip hint files for now. For now we'll populate the new file
	type recoveredEntry struct {
		data []byte
		key string
		ts uint32
	}

	pendingChanges := map[string]recoveredEntry{}
	deletedEntries := map[string]struct{}{}

	newEntry := func(log LogEntry, data []byte) recoveredEntry {
		return recoveredEntry{
			data: data,
			key: string(log.key),
			ts: log.ts, 
		}
	}

	for i, file := range files {
		allData, err := io.ReadAll(file) // streaming might be more appropiate
		if err != nil {
			return nil, fmt.Errorf("error reading file %d: %w", i, err)
		} else if len(allData) == 0 {
			return nil, fmt.Errorf("empty file %d", i)
		}

		firstSize := len(allData)
		for offset := 0; offset < firstSize; {
			log, err := deserializeToLog(&allData)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				return nil, fmt.Errorf("error reading entries from file %d: %w", i, err)
			}

			if current, ok := pendingChanges[string(log.key)]; !ok {
				pendingChanges[string(log.key)] = newEntry(log, log.value)
			} else if log.ts == 0 {
				deletedEntries[string(log.key)] = struct{}{}
			} else if log.ts > current.ts { // updated entry
				pendingChanges[string(log.key)] = newEntry(log, log.value)
			}

			offset += 4 + 4 + 2 + 4 + int(log.keySize) + int(log.valSize)
		}

		// todo: create hint files for fast recovery
	}

	for k := range deletedEntries {
		delete(pendingChanges, k)
	}

	dir := KeyDir{}
	s := &Store{
		activeFile: f,
		dir: dir,
	}

	for k, v := range pendingChanges {
		newLog := newLog(Key(k), v.data)
		newLog.ts = v.ts
		
		data := serializeLog(newLog)
		lastOffset:= s.persist(data)
	
		s.dir[k] = KeyDirEntry{
			fileId: "", // todo
			valueSize: uint32(len(data)),
			valuePos: uint(lastOffset),
			timestamp: v.ts,
		}
	}

	return s, nil
}

var errKeyNotFound = fmt.Errorf("key not found")

func (s *Store) Get(key string) ([]byte, error) {
	entry, ok := s.dir[key]
	if !ok {
		return nil, errKeyNotFound
	}

	// todo: use it file id

	_, err := s.activeFile.Seek(int64(entry.valuePos), io.SeekStart)
	if err != nil {
		return nil, fmt.Errorf("error seeking file: %w", err)
	}

	b := make([]byte, entry.valueSize)
	_, err = s.activeFile.Read(b)
	if err != nil {
		return nil, fmt.Errorf("error reading file %w", err)
	}

	log, err := deserializeToLog(&b)
	if err != nil {
		return nil, fmt.Errorf("error deserializing log entry: %w", err)
	}
	return log.value, nil
}

func (s *Store) Put(key string, val []byte) {
	data := serializeLog(newLog(Key(key), val))

	lastOffset:= s.persist(data)

	s.dir[key] = KeyDirEntry{
		fileId: "", // todo
		valueSize: uint32(len(data)),
		valuePos: uint(lastOffset),
		timestamp: now(),
	}
}

func (s *Store) persist(b []byte) int {
	lastOffset, _ := s.activeFile.Seek(0, io.SeekEnd)

	if _, err := s.activeFile.Write(b); err != nil {
		fmt.Println(err)
	}

	return int(lastOffset)
}

func (s *Store) Delete(key string) {
	logEntry := newLog(Key(key), nil)
	logEntry.ts = 0 // special tombstone value. Do not recover this when merging
	
	s.persist(serializeLog(logEntry))
	delete(s.dir, key)
}

func (s *Store) Keys() []string {
	keys := make([]string, 0, len(s.dir))
	for key := range s.dir {
		keys = append(keys, key)
	}
	return keys
}