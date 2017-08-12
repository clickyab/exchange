package static

import (
	"fmt"
	"math/rand"
	"time"
)

func slotKeyGen(imp, slot string) string {
	return fmt.Sprintf(`%s_%s_%s`, prefixSlot, imp, slot)
}

func inRange(min, max int) int {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Intn(max-min) + min
}
