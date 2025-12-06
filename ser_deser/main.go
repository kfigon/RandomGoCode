package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type User struct {
	name  string
	age   int
	ready bool
}

func main() {
	fmt.Println("hello")
	u := User{
		name:  "john",
		age:   123,
		ready: true,
	}

	d, _ := Serialize(&u,
		intSerializer("age", func(u *User) int { return u.age }),
		strSerializer("name", func(u *User) string { return u.name }),
		boolSerializer("ready", func(u *User) bool { return u.ready }),
	)

	for i := 0; i < len(d); i++ {
		if i != 0 && i%4 == 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%x", d[i])
	}

	recovered, err := Deserialize(d,
		intDeserializer("age", func(t *User, i int) { t.age = i }),
		strDeserializer("name", func(t *User, i string) { t.name = i }),
		boolDeserializer("ready", func(t *User, i bool) { t.ready = i }),
	)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", recovered)
}

type serializeFn[T any] func(*T, *bytes.Buffer) (fieldName string, err error)
type getter[T any, K any] func(*T) K

func intSerializer[T any](name string, g getter[T, int]) serializeFn[T] {
	return func(obj *T, b *bytes.Buffer) (string, error) {
		i := g(obj)
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, uint32(i))
		b.Write(buf)
		return name, nil
	}
}

func strSerializer[T any](name string, g getter[T, string]) serializeFn[T] {
	return func(obj *T, b *bytes.Buffer) (string, error) {
		s := g(obj)
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, uint32(len(s)))
		b.Write(buf)
		b.WriteString(s)
		return name, nil
	}
}

func boolSerializer[T any](name string, g getter[T, bool]) serializeFn[T] {
	return func(obj *T, b *bytes.Buffer) (string, error) {
		boolean := g(obj)
		if boolean {
			b.WriteByte(1)
		} else {
			b.WriteByte(0)
		}
		return name, nil
	}
}

func Serialize[T any](obj *T, funs ...serializeFn[T]) ([]byte, error) {
	out := bytes.NewBuffer(nil)
	for _, fn := range funs {
		if name, err := fn(obj, out); err != nil {
			return nil, fmt.Errorf("error serializing field %s: %w", name, err)
		} else {
			fmt.Printf("%s serialized\n", name)
		}
	}

	return out.Bytes(), nil
}

type deserializeFn[T any] func(*T, *bytes.Reader) error
type setter[T any, K any] func(*T, K)

func intDeserializer[T any](name string, set setter[T, int]) deserializeFn[T] {
	return func(t *T, b *bytes.Reader) error {
		buf := make([]byte, 4)
		_, err := b.Read(buf)
		if err != nil {
			return fmt.Errorf("error deserializing %s: %w", name, err)
		}
		set(t, int(binary.BigEndian.Uint32(buf)))
		return nil
	}
}

func strDeserializer[T any](name string, set setter[T, string]) deserializeFn[T] {
	return func(t *T, b *bytes.Reader) error {
		buf := make([]byte, 4)
		if _, err := b.Read(buf); err != nil {
			return fmt.Errorf("error deserializing string lenght of %s: %w", name, err)
		}
		ln := binary.BigEndian.Uint32(buf)

		strBuf := make([]byte, ln)
		if x, err := b.Read(strBuf); err != nil || x != int(ln) {
			return fmt.Errorf("error deserializing string val of %s: %w", name, err)
		}
		set(t, string(strBuf))
		return nil
	}
}

func boolDeserializer[T any](name string, set setter[T, bool]) deserializeFn[T] {
	return func(t *T, b *bytes.Reader) error {
		d, err := b.ReadByte()
		if err != nil {
			return fmt.Errorf("error deserializing bool %s: %w", name, err)
		}
		set(t, d == 1)
		return nil
	}
}
func Deserialize[T any](data []byte, funs ...deserializeFn[T]) (*T, error) {
	var out T
	buf := bytes.NewReader(data)
	for _, fn := range funs {
		if err := fn(&out, buf); err != nil {
			return nil, fmt.Errorf("deserialization error: %w", err)
		}
	}

	return &out, nil
}
