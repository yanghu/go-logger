// Package redis provides a redis writer implementation for logger
package redis

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

const (
	ADDRESS = "127.0.0.1:6379" // convenient server address for user
	NETWORK = "tcp"
)

// Writer implements io.Writer. It writes strings into redis list using LPUSH
// the maximum size is limited by EntryLimit. It's implemented using LTRIM in redis.
// Conn is dialed at initialization and kept alive during the lifetime of the writer.
type Writer struct {
	address    string
	Logname    string
	Conn       redis.Conn
	EntryLimit int
}

// returns a new Writer object.
func NewWriter(address, logname string, entryLimit int) (rw *Writer, err error) {
	rw = &Writer{address: address, Logname: logname, EntryLimit: entryLimit}
	rw.Conn, err = redis.Dial(NETWORK, address)
	if err != nil {
		return nil, err
	}
	return
}

// Implements Writer interface. Use LPUSH and LTRIM to keep a log message list in redis
func (rw *Writer) Write(p []byte) (n int, err error) {
	rw.Conn.Send("LPUSH", rw.Logname, string(p[:]))
	rw.Conn.Send("LTRIM", rw.Logname, 0, rw.EntryLimit)
	if err := rw.Conn.Flush(); err != nil {
		return 0, err
	}
	return 1, nil
}

// Remove all logs in redis log list maintaned by Writer
func (rw *Writer) FlushLog() (err error) {
	_, err = rw.Conn.Do("DEL", rw.Logname)
	return
}

// Returns logs stored in the list specified by logName.
func ReadLogs(conn redis.Conn, logName string, limit int) (logs []string) {
	reply, err := redis.Values(conn.Do("LRANGE", logName, 0, limit))
	if err != nil {
		log.Fatal(err)
	}
	if err := redis.ScanSlice(reply, &logs); err != nil {
		log.Fatal(err)
	}
	// can also use the following code to deal with sring reply
	// 	for len(reply) > 0 {
	// 		var logString string
	// 		reply, err = redis.Scan(reply, &logString)
	// 		log.Println(logString)
	// 	}
	return
}
