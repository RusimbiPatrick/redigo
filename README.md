
# RediGo - Redis Clone in Go

A Redis-compatible in-memory datastore implementation in Go, supporting core data structures and the RESP protocol.

## Features

- **Data Types**: Strings, Hashes (partial)
- **Commands**: 15+ Redis-compatible commands
- **Persistence**: TBD (Planned)
- **Concurrency**: Thread-safe with RW locks
- **Protocol**: RESP (Redis Serialisation Protocol)
- **TTL**: Time-based expiration with heap

## Quick Start

# Build and run
make && ./redigo

# Connect with telnet
telnet localhost 6389
## Supported Commands

| Command   | Example                      | Description                  |
|-----------|------------------------------|------------------------------|
| SET       | SET key value                | Store string value           |
| GET       | GET key                      | Retrieve string value        |
| HSET      | HSET key field1 value1 ...   | Set hash fields              |
| HGET      | HGET key field               | Get hash field value         |
| HGETALL   | HGETALL key                  | Get all hash fields+values   |
| HDEL      | HDEL key field1 field2 ...   | Delete hash fields           |
| HEXISTS   | HEXISTS key field            | Check field existence        |
| EXPIRE    | EXPIRE key seconds           | Set key expiration           |
| TTL       | TTL key                      | Get time to live             |
| PING      | PING                         | Server liveness check        |


## Command Reference

### Hash Commands

#### HSET

```
HSET key field1 value1 [field2 value2 ...]
```

Stores multiple hash fields.

**Example:**

```bash
127.0.0.1:6389> HSET user:1000 name "John" age 30
(integer) 2
```

#### HGETALL

```redis
HGETALL key
```

Returns all fields and values in a hash.

**Response:**

```redis
1) "name"
2) "John"
3) "age"
4) "30"
```

#### HEXISTS

```redis
HEXISTS key field
```

Returns if field exists in hash.

**Return:**

- `1` if exists
- `0` if not exists

---

## Troubleshooting Guide

### Common Issues

#### Connection Refused

1. Check if the server is running:

```bash
ps aux | grep redigo
```

2. Verify listening port:

```bash
lsof -i :6389
```

#### Command Not Found

- Verify the command is implemented in:
  `internal/server/handler.go`
- Check command spelling (Redis commands are uppercase)

---

## Redis CLI Telnet Update

To interact with your server using `telnet`, connect with:

```bash
telnet localhost 6389
```

Once connected, you can send raw RESP commands, for example:

```bash
SET key value
```

But to interact with your server, it's generally better to use `telnet` for testing as it allows you to send RESP commands directly.
