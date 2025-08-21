package context_pkg

import (
	"context"
	"fmt"
	"time"
)

func TimeoutCtx(c context.Context) {
	select {
	case <-c.Done():
		fmt.Println("Context timed out:", c.Err())
	case <-time.After(3 * time.Second):
		fmt.Println("Contextcomplete within 3 seconds")
	}
}
