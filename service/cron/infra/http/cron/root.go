package http

import (
	oas "github.com/sugar-cat7/vspo-portal/service/cron/infra/http/cron/internal/gen"
	"github.com/sugar-cat7/vspo-portal/service/cron/infra/http/cron/internal/handler/channel"
	"github.com/sugar-cat7/vspo-portal/service/cron/infra/http/cron/internal/handler/video"
	"github.com/sugar-cat7/vspo-portal/service/cron/usecase"
)

// Compile-time check for Handler.
var _ oas.Handler = (*RootHandler)(nil)

// RootHandler is a composite handler.
type RootHandler struct {
	channel.CH
	video.VH
}

// NewHandler creates a new instance of a RootHandler.
func NewHandler(channelInteractor usecase.ChannelInteractor, videoInteractor usecase.VideoInteractor) *RootHandler {
	return &RootHandler{
		CH: channel.NewHandler(channelInteractor),
		VH: video.NewHandler(videoInteractor),
	}
}