package context_pkg

import (
	"context"
	"fmt"
	"time"
)

func CancelCtx(c context.Context) {

	select {
	case <-c.Done():
		fmt.Println("Context cancelled:", c.Err())
	case <-time.After(2 * time.Second):
		fmt.Println("Context not cancelled within 2 seconds")
	}

}
