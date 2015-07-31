# A simple golang logger with level support and redis writer

This package is a simple golang logger wrapper, with four levels of logging (trace, info, warning and error). Level of logging is specified at initialization. 

Click [here](https://godoc.org/bitbucket.org/yanghu/logger) to view docs.

Any os.Writer type can be used as the logger output. 

An redis writer is implemented in `logger/redis` package. User can create a handler which satisfies os.Writer interface, and writes log to redis. like this

```go
import log "bitbucket.org/yanghu/logger"
import "bitbucket.org/yanghu/logger/redis"

// NewRedisWriter initialization asks the redis server address, a list name to store logs, and the number of logs to store
rw, err := redis.NewWriter(redis.ADDRESS, "test_log", 10000)
if err != nil {
    log.Fatal(err)
}
log.turnOnLogging(LevelWarning, rw)
log.Warninig("This warning will go to redis")
log.Error("So will this error")

// Read logs in go. provide a redis connection(not necessarily the one our writer uses though)
// ReadLogs returns a slice of strings
readSize = 10 //read 10 entries
logs := redis.ReadLogs(rw.Conn, rw.Logname, readSize)

```


