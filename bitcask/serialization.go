package main

import (
	"bytes"
	"encoding/binary"
)


var encoding = binary.LittleEndian

func serializeI16(i uint16) []byte{
	return encoding.AppendUint16([]byte{}, i)
}

func deserializeI16(b []byte) uint16{
	return encoding.Uint16(b)
}

func serializeI32(i uint32) []byte{
	return encoding.AppendUint32([]byte{}, i)
}

func deserializeI32(b []byte) uint32{
	return encoding.Uint32(b)
}

func serializeLog(l LogEntry) []byte {
	var b bytes.Buffer
	b.Write(serializeI32(l.crc))
	b.Write(serializeI32(l.ts))
	b.Write(serializeI16(l.keySize))
	b.Write(serializeI32(l.valSize))
	b.Write(l.key)
	b.Write(l.value)

	return b.Bytes()
}

func deserializeToLog(b *[]byte) (LogEntry, error) {
	advanceBuf := func(b *[]byte, size int) {
		*b = (*b)[size:]
	}
	
	crc := deserializeI32(*b)
	advanceBuf(b, 4)
	
	ts := deserializeI32(*b)
	advanceBuf(b, 4)
	
	keySize := deserializeI16(*b)
	advanceBuf(b, 2)
	
	valSize := deserializeI32(*b)
	advanceBuf(b, 4)

	key := (*b)[:keySize]
	advanceBuf(b, int(keySize))

	val := (*b)[:valSize]
	advanceBuf(b, int(valSize))

	return LogEntry{
		crc: crc,
		ts: ts,
		keySize: keySize,
		valSize: valSize,
		key: key,
		value: val,
	}, nil
}