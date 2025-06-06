package utils

import (
	"context"
	"encoding/json"

	"github.com/heroiclabs/nakama-common/runtime"
	"github.com/rotisserie/eris"
	"google.golang.org/grpc/codes"
)

const (
	AdminAccountID = "00000000-0000-0000-0000-000000000000"
)

// GetUserID gets the Nakama UserID from the given context.
func GetUserID(ctx context.Context) (string, error) {
	userID, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		return "", eris.New("unable to get user id from context")
	}
	return userID, nil
}

// MarshalResult marshals the given result and converts any marshalling error into an "Internal" RPC error.
func MarshalResult(logger runtime.Logger, result any) (string, error) {
	bz, err := json.Marshal(result)
	if err != nil {
		return LogErrorWithMessageAndCode(logger, err, codes.FailedPrecondition, "unable to marshal response: %v", err)
	}
	return string(bz), nil
}
