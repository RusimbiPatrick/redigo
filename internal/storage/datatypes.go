type DataType int

const (
	STRING DataType = iota
	LIST
	HASH
	SET
	ZSET
)

type RedisObject struct {
	Type DataType
	Encoding interface{} // Could be string, []string, map[string]string etc
	TTL time.Time
}


