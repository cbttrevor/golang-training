package slowcode

import (
	"fmt"
	"time"
)

func RunSlowly(waitTime uint8) {
	waitDuration, _ := time.ParseDuration(fmt.Sprintf("%vs", waitTime))
	time.Sleep(waitDuration)
}
