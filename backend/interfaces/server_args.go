package interfaces

import (
	"net/http"

	"github.com/oinume/lekcije/backend/ga_measurement"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"gopkg.in/redis.v4"

	"github.com/oinume/lekcije/backend/interfaces/http/flash_message"
)

type ServerArgs struct {
	AccessLogger        *zap.Logger
	AppLogger           *zap.Logger
	DB                  *gorm.DB
	FlashMessageStore   flash_message.Store
	Redis               *redis.Client
	SenderHTTPClient    *http.Client
	GAMeasurementClient ga_measurement.Client
}