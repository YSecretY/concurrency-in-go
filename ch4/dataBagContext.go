package ch4

import (
	"context"
	"fmt"
)

/*
// ProcessRequest is an example of handling data with contexts
func ProcessRequest(userID, authToken string) {
	ctx := context.WithValue(context.Background(), "userID", userID)
	ctx = context.WithValue(ctx, "authToken", authToken)
	HandleResponse(ctx)
}

func HandleResponse(ctx context.Context) {
	fmt.Printf("handling response for %v (%v)", ctx.Value("userID"), ctx.Value("authToken"))
}
*/

// It is better to create specific type for keys
// in context to avoid collisions
type ctxKey int

const (
	ctxUserID ctxKey = iota
	ctxAuthToken
)

func UserID(c context.Context) string {
	return c.Value(ctxUserID).(string)
}

func AuthToken(c context.Context) string {
	return c.Value(ctxAuthToken).(string)
}

func ProcessRequest(userID, authToken string) {
	ctx := context.WithValue(context.Background(), ctxUserID, userID)
	ctx = context.WithValue(ctx, ctxAuthToken, authToken)
	HandleResponse(ctx)
}

func HandleResponse(ctx context.Context) {
	fmt.Printf("handling response for %v (auth: %v)", UserID(ctx), AuthToken(ctx))
}
