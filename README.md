# A simple golang logger with level support and redis writer

This package is a simple golang logger wrapper, with four levels of logging (trace, info, warning and error). Level of logging is specified at initialization. 

Any os.Writer type can be used as the logger output. 

An redis writer is implemented in `redisWriter` package. User can create a handler which satisfies os.Writer interface, and writes log to redis. like this

```
import log "bitbucket.org/yanghu/levelLog"
import "bitbucket.org/yanghu/levelLog/redisWriter"

// NewRedisWriter initialization asks the redis server address, a list name to store logs, and the number of logs to store
rw, err := redisWriter.NewRedisWriter(redisWriter.ADDRESS, "test_log", 10000)
if err != nil {
    log.Fatal(err)
}
log.turnOnLogging(LevelWarning, rw)
log.Warninig("This warning will go to redis")
log.Error("So will this error")
```


