package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

func main() {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)

	fmt.Printf("Type: %T, Value: %v \n", id, id)
	fmt.Printf("Type: %T, Value: %v \n", id.String(), id.String())
	fmt.Printf("Type: %T, Value: %v \n", id.Time(), id.Time())
	fmt.Printf("Type: %T, Value: %v \n", ulid.Time(id.Time()), ulid.Time(id.Time()))
}
