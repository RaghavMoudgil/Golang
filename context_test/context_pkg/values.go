package context_pkg

import (
	"context"
	"fmt"
)

func PrintCtx(c context.Context) {
	userId := c.Value("val")
	fmt.Println("User ID:", userId)
}
