package chilliflake

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type IID int64

type Idgen struct {
	epoch  time.Time
	nodeId int64
}

func New() *Idgen {
	nid := int64(rand.Int31n(64))
	nid = nid << 16

	return &Idgen{
		epoch:  time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
		nodeId: nid,
	}
}

func (i *Idgen) New() IID {
	now := time.Since(i.epoch).Nanoseconds() / 1000000
	now = now << 22

	rid := int64(rand.Int31n(math.MaxUint16))

	return IID(now | i.nodeId | rid)
}

func PrintBinary(msg string, val int64) {
	fmt.Printf("%s => %d [%64b] \n", msg, val, val)
}
