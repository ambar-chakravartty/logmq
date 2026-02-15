
# LogMQ – A Disk-Backed Log-Based Message Broker (Kafka-lite)

LogMQ is a minimal, single-node, disk-backed message broker built in Go.

It implements:
- Append-only log storage
- Offset-based consumption
- TCP-based protocol
- Persistent durability via fsync
- Concurrent client handling

## Architecture Overview

Client (Producer / Consumer)
        |
        v
   TCP Server (Go)
        |
        v
   Storage Layer
        |
        v
 Append-Only Log Files

Each topic is stored as:

<topic>.log

Messages are stored in a binary format:

[4 bytes length][N bytes payload]

Offsets represent the byte position in the file.

---

## Features (MVP)

- Single-node broker
- Multiple topics (auto-created)
- Persistent disk-backed storage
- Offset-based message replay
- Concurrent client handling
- Manual durability via fsync
- Plain text TCP protocol

---

## Protocol

The broker speaks a simple TCP text protocol.

### Produce

Request:
PRODUCE <topic> <message>

Example:
PRODUCE orders hello-world

Response:
OK <offset>

Example:
OK 128

The returned offset is the byte position at which the message was written.

---

### Consume

Request:
CONSUME <topic> <offset>

Example:
CONSUME orders 128

Response:
MESSAGE <offset> <length> <payload>

Example:
MESSAGE 128 11 hello-world

If offset is beyond end of file:
EOF

---

## Storage Format

Each message is stored as:

- 4-byte big-endian integer (payload length)
- Raw payload bytes

Example (hex):

00 00 00 0B | 68 65 6C 6C 6F 2D 77 6F 72 6C 64

Advantages of length-prefix encoding:
- Fast sequential reads
- Easy parsing
- Industry-standard binary framing technique

---

## Concurrency Model

- One goroutine per client connection
- Mutex per topic file
- Storage layer isolated from networking layer

This ensures:
- Safe concurrent appends
- Clean separation of concerns
- Scalable client handling

---

## How to Run

1. Build:
   go build 

2. Start server:
   ./logmq

Default port: 8080

---

## Manual Testing

Using netcat:

nc localhost 8080

Then:

PRODUCE orders hello
CONSUME orders 0

---

## Example Workflow

1. Producer sends:
   PRODUCE orders item-1

2. Server responds:
   OK 0

3. Consumer reads:
   CONSUME orders 0

4. Server responds:
   MESSAGE 0 6 item-1

---

