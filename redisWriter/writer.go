// Package redisWriter provides ...
package redisWriter

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

const (
	ADDRESS = "127.0.0.1:6379"
	NETWORK = "tcp"
)

type RedisWriter struct {
	address    string
	Logname    string
	Conn       redis.Conn
	EntryLimit int
}

func NewRedisWriter(address, logname string, entryLimit int) (rw *RedisWriter, err error) {
	rw = &RedisWriter{address: address, Logname: logname, EntryLimit: entryLimit}
	rw.Conn, err = redis.Dial(NETWORK, address)
	if err != nil {
		return nil, err
	}
	return
}

func (rw *RedisWriter) Write(p []byte) (n int, err error) {
	rw.Conn.Send("LPUSH", rw.Logname, string(p[:]))
	rw.Conn.Send("LTRIM", rw.Logname, 0, rw.EntryLimit)
	if err := rw.Conn.Flush(); err != nil {
		return 0, err
	}
	return 1, nil
}

// can also use the following code to deal with sring reply
// 	for len(reply) > 0 {
// 		var logString string
// 		reply, err = redis.Scan(reply, &logString)
// 		log.Println(logString)
// 	}

func ReadLog(conn redis.Conn, logName string, limit int) (logs []string) {
	reply, err := redis.Values(conn.Do("LRANGE", logName, 0, limit))
	if err != nil {
		log.Fatal(err)
	}
	if err := redis.ScanSlice(reply, &logs); err != nil {
		log.Fatal(err)
	}
	return
}
