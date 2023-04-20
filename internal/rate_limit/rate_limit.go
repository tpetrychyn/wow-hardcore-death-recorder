package rate_limit

import (
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

func NewInMemoryRateLimitMiddleware(rateFormatted string) (gin.HandlerFunc, error) {
	rate, err := limiter.NewRateFromFormatted(rateFormatted)
	if err != nil {
		return nil, err
	}

	// ideally this should use a distributed store like redis
	store := memory.NewStore()

	return mgin.NewMiddleware(limiter.New(store, rate)), nil
}
