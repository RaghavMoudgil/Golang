package main

import (
	ctx "context"
	v "context_test/context_pkg"
	"time"
)

func main() {
	//elow code is to simply send the values via context
	c := ctx.Background()
	c = ctx.WithValue(c, "val", 2)
	v.PrintCtx(c)

	// below is the code to send the context with cancel
	cc := ctx.Background()
	cc, cancel := ctx.WithCancel(cc)
	v.CancelCtx(cc)
	time.Sleep(1 * time.Second) //chage the time and you will handle the cancelation of the context

	cancel()

	// below is the code to send the context with timeout
	ct := ctx.Background()
	ct, cancel = ctx.WithTimeout(ct, 4*time.Second) //change the time to change the timeout
	//bwloe is the code for deadline its same as timeout but the time strts from now means you are given sokme deadline to complete the task
	ct, cancel = ctx.WithDeadline(ct, time.Now().Add(2*time.Second)) // change the time to change the deadline

	v.TimeoutCtx(ct)
	defer cancel() // ensure the context is cancelled to avoid resource leaks
}
