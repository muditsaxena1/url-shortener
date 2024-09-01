package utils

import (
	"sync"
	"time"
)

// Snowflake struct to hold instance data
type Snowflake struct {
	mu           sync.Mutex
	epoch        int64
	instanceID   int64
	sequence     int64
	lastTime     int64
	instanceBits uint8
	sequenceBits uint8
}

// NewSnowflake creates a new Snowflake instance
func NewSnowflake(instanceID int64) *Snowflake {
	if instanceID < 0 || instanceID > 15 { // 4 bits -> max value is 15
		panic("Instance ID must be between 0 and 15")
	}

	return &Snowflake{
		epoch:        time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), // valid for 69 years
		instanceID:   instanceID,
		instanceBits: 4,
		sequenceBits: 3,
	}
}

// GenerateID generates a new unique ID
func (sf *Snowflake) GenerateID() int64 {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	now := time.Now().UnixMilli()
	if now == sf.lastTime {
		sf.sequence = (sf.sequence + 1) & ((1 << sf.sequenceBits) - 1)
		if sf.sequence == 0 {
			for now <= sf.lastTime {
				now = time.Now().UnixMilli()
			}
		}
	} else {
		sf.sequence = 0
	}

	sf.lastTime = now

	// Calculate the 41-bit timestamp part
	timestamp := (now - sf.epoch) & ((1 << 41) - 1)

	// Shift bits to create the ID
	id := (timestamp << (sf.instanceBits + sf.sequenceBits)) | (sf.instanceID << sf.sequenceBits) | sf.sequence

	return id
}

func (sf *Snowflake) GenerateShortCode() string {
	return encodeIDToBase64(sf.GenerateID())
}
