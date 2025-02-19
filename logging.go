package catalogue

import (
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// LoggingMiddleware logs method calls, parameters, results, and elapsed time.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

func (mw loggingMiddleware) List(tags []string, order string, pageNum, pageSize int) (socks []Sock, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "List",
			"tags", strings.Join(tags, ", "),
			"order", order,
			"pageNum", pageNum,
			"pageSize", pageSize,
			"result", len(socks),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	socks, err = mw.next.List(tags, order, pageNum, pageSize)
	if err != nil {
		level.Error(mw.logger).Log(
			"tags", strings.Join(tags, ", "),
			"order", order,
			"pageNum", pageNum,
			"pageSize", pageSize,
			"result", len(socks),
			"err", err.Error(),
		)
	}
	return socks, err
}

func (mw loggingMiddleware) Count(tags []string) (n int, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "Count",
			"tags", strings.Join(tags, ", "),
			"result", n,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	n, err = mw.next.Count(tags)
	if err != nil {
		level.Error(mw.logger).Log(
			"method", "Count",
			"tags", strings.Join(tags, ", "),
			"result", n,
			"err", err.Error(),
		)
	}
	return n, err
}

func (mw loggingMiddleware) Get(id string) (s Sock, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "Get",
			"id", id,
			"sock", s.ID,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	s, err = mw.next.Get(id)
	if err != nil {
		level.Error(mw.logger).Log(
			"method", "Get",
			"id", id,
			"sock", s.ID,
			"err", err.Error(),
		)
	}
	return s, err
}

func (mw loggingMiddleware) Tags() (tags []string, err error) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "Tags",
			"result", len(tags),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	tags, err = mw.next.Tags()
	if err != nil {
		level.Info(mw.logger).Log(
			"method", "Tags",
			"result", len(tags),
			"err", err.Error(),
		)
	}
	return tags, err
}

func (mw loggingMiddleware) Health() (health []Health) {
	defer func(begin time.Time) {
		level.Info(mw.logger).Log(
			"method", "Health",
			"result", len(health),
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.Health()
}
