package test

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestUUID(t *testing.T) {
	id := uuid.New()
	fmt.Println(id)
	sid := fmt.Sprintf("%X", id[10:])
	fmt.Println(sid)
}
