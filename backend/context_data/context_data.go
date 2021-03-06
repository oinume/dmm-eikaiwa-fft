package context_data

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/oinume/lekcije/backend/errors"
	"github.com/oinume/lekcije/backend/model"
)

type dbKey struct{}
type loggedInUserKey struct{}
type trackingIDKey struct{}
type apiTokenKey struct{}

func SetDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, dbKey{}, db)
}

func GetLoggedInUser(ctx context.Context) (*model.User, error) {
	value := ctx.Value(loggedInUserKey{})
	if user, ok := value.(*model.User); ok {
		return user, nil
	}
	return nil, errors.NewNotFoundError(errors.WithMessage("Failed to get loggedInUser from context"))
}

func MustLoggedInUser(ctx context.Context) *model.User {
	user, err := GetLoggedInUser(ctx)
	if err != nil {
		panic(err)
	}
	return user
}

func SetLoggedInUser(ctx context.Context, user *model.User) context.Context {
	return context.WithValue(ctx, loggedInUserKey{}, user)
}

func SetTrackingID(ctx context.Context, trackingID string) context.Context {
	return context.WithValue(ctx, trackingIDKey{}, trackingID)
}

func GetTrackingID(ctx context.Context) (string, error) {
	value := ctx.Value(trackingIDKey{})
	if trackingID, ok := value.(string); ok {
		return trackingID, nil
	}
	return "", errors.NewNotFoundError(errors.WithMessage("Failed to get trackingID from context"))
}

func MustTrackingID(ctx context.Context) string {
	trackingID, err := GetTrackingID(ctx)
	if err != nil {
		panic(err)
	}
	return trackingID
}

func WithAPIToken(ctx context.Context, apiToken string) context.Context {
	return context.WithValue(ctx, apiTokenKey{}, apiToken)
}

func GetAPIToken(ctx context.Context) (string, error) {
	value := ctx.Value(apiTokenKey{})
	if apiToken, ok := value.(string); ok {
		return apiToken, nil
	}
	return "", errors.NewNotFoundError(errors.WithMessage("failed to get api token from context"))
}

func MustAPIToken(ctx context.Context) string {
	v, err := GetAPIToken(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
