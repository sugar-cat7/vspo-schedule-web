// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func encodeChannelsPostResponse(response ChannelsPostRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *ChannelsPostOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *ChannelsPostBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *ChannelsPostUnauthorized:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *ChannelsPostForbidden:
		w.WriteHeader(403)
		span.SetStatus(codes.Error, http.StatusText(403))

		return nil

	case *ChannelsPostNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	case *ChannelsPostInternalServerError:
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeVideosPostResponse(response VideosPostRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *VideosPostOKApplicationJSON:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *VideosPostBadRequest:
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		return nil

	case *VideosPostUnauthorized:
		w.WriteHeader(401)
		span.SetStatus(codes.Error, http.StatusText(401))

		return nil

	case *VideosPostForbidden:
		w.WriteHeader(403)
		span.SetStatus(codes.Error, http.StatusText(403))

		return nil

	case *VideosPostNotFound:
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		return nil

	case *VideosPostInternalServerError:
		w.WriteHeader(500)
		span.SetStatus(codes.Error, http.StatusText(500))

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}
