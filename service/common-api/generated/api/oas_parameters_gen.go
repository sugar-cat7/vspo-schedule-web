// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
)

// CreatorsGetParams is parameters of GET /creators operation.
type CreatorsGetParams struct {
	// A list of IDs that identify the creators you want to get. To get more than one creator, include
	// this parameter for each creator you want to get. The creator_id and creator_id parameters are
	// mutually exclusive.
	CreatorIds OptString
	// A filter used to filter the list of creators by the creator's type.(The default is "all".).
	CreatorType OptCreatorsGetCreatorType
	// Pagination object. Specify this parameter only if you don't specify the creator_ids query
	// parameter.
	Page OptInt
	// The maximum number of items to return per page(The default is 20.) Specify this parameter only if
	// you don't specify the creator_ids query parameter.
	Limit OptInt
}

func unpackCreatorsGetParams(packed middleware.Parameters) (params CreatorsGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "creator_ids",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.CreatorIds = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "creator_type",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.CreatorType = v.(OptCreatorsGetCreatorType)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "limit",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Limit = v.(OptInt)
		}
	}
	return params
}

func decodeCreatorsGetParams(args [0]string, argsEscaped bool, r *http.Request) (params CreatorsGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: creator_ids.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "creator_ids",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCreatorIdsVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotCreatorIdsVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.CreatorIds.SetTo(paramsDotCreatorIdsVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "creator_ids",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: creator_type.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "creator_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCreatorTypeVal CreatorsGetCreatorType
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotCreatorTypeVal = CreatorsGetCreatorType(c)
					return nil
				}(); err != nil {
					return err
				}
				params.CreatorType.SetTo(paramsDotCreatorTypeVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.CreatorType.Get(); ok {
					if err := func() error {
						if err := value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "creator_type",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: limit.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "limit",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotLimitVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotLimitVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Limit.SetTo(paramsDotLimitVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "limit",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// VideosGetParams is parameters of GET /videos operation.
type VideosGetParams struct {
	// A list of IDs that identify the videos you want to get. To get more than one video, include this
	// parameter for each video you want to get. The video_id and creator_id parameters are mutually
	// exclusive.
	VideoIds OptString
	// The ID of the creator whose list of videos you want to get. The video_id and creator_id parameters
	// are mutually exclusive.
	CreatorID OptString
	// The ISO 639-1 two-letter code for Japanese (i.e., ja).
	Language OptString
	// A filter used to filter the list of videos by the video's type.(The default is "vspo_broadcast".).
	VideoType OptVideosGetVideoType
	// A filter used to filter the list of videos by the video's vspo_broadcast type.(The default is
	// "all".).
	BroadcastType OptVideosGetBroadcastType
	// A filter used to filter the list of videos by when they were published.(The default is "all".).
	Period OptVideosGetPeriod
	// The order to sort the returned videos.(The default is "time".).
	Sort OptVideosGetSort
	// A filter used to filter the list of videos by the video's platform type.(The default is "all".).
	PlatformType OptVideosGetPlatformType
	// Pagination object. Specify this parameter only if you don't specify the creator_id or video_ids
	// query parameter.
	Page OptInt
	// The maximum number of items to return per page(The default is 20.) Specify this parameter only if
	// you don't specify the creator_id or video_ids query parameter.
	Limit OptInt
}

func unpackVideosGetParams(packed middleware.Parameters) (params VideosGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "video_ids",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.VideoIds = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "creator_id",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.CreatorID = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "language",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Language = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "video_type",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.VideoType = v.(OptVideosGetVideoType)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "broadcast_type",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.BroadcastType = v.(OptVideosGetBroadcastType)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "period",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Period = v.(OptVideosGetPeriod)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "sort",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Sort = v.(OptVideosGetSort)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "platform_type",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.PlatformType = v.(OptVideosGetPlatformType)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "page",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Page = v.(OptInt)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "limit",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Limit = v.(OptInt)
		}
	}
	return params
}

func decodeVideosGetParams(args [0]string, argsEscaped bool, r *http.Request) (params VideosGetParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: video_ids.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "video_ids",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotVideoIdsVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotVideoIdsVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.VideoIds.SetTo(paramsDotVideoIdsVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "video_ids",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: creator_id.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "creator_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotCreatorIDVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotCreatorIDVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.CreatorID.SetTo(paramsDotCreatorIDVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "creator_id",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: language.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "language",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotLanguageVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotLanguageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Language.SetTo(paramsDotLanguageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "language",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: video_type.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "video_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotVideoTypeVal VideosGetVideoType
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotVideoTypeVal = VideosGetVideoType(c)
					return nil
				}(); err != nil {
					return err
				}
				params.VideoType.SetTo(paramsDotVideoTypeVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.VideoType.Get(); ok {
					if err := func() error {
						if err := value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "video_type",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: broadcast_type.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "broadcast_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotBroadcastTypeVal VideosGetBroadcastType
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotBroadcastTypeVal = VideosGetBroadcastType(c)
					return nil
				}(); err != nil {
					return err
				}
				params.BroadcastType.SetTo(paramsDotBroadcastTypeVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.BroadcastType.Get(); ok {
					if err := func() error {
						if err := value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "broadcast_type",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: period.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "period",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPeriodVal VideosGetPeriod
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotPeriodVal = VideosGetPeriod(c)
					return nil
				}(); err != nil {
					return err
				}
				params.Period.SetTo(paramsDotPeriodVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Period.Get(); ok {
					if err := func() error {
						if err := value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "period",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: sort.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "sort",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotSortVal VideosGetSort
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotSortVal = VideosGetSort(c)
					return nil
				}(); err != nil {
					return err
				}
				params.Sort.SetTo(paramsDotSortVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.Sort.Get(); ok {
					if err := func() error {
						if err := value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "sort",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: platform_type.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "platform_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPlatformTypeVal VideosGetPlatformType
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotPlatformTypeVal = VideosGetPlatformType(c)
					return nil
				}(); err != nil {
					return err
				}
				params.PlatformType.SetTo(paramsDotPlatformTypeVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if value, ok := params.PlatformType.Get(); ok {
					if err := func() error {
						if err := value.Validate(); err != nil {
							return err
						}
						return nil
					}(); err != nil {
						return err
					}
				}
				return nil
			}(); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "platform_type",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: page.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotPageVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotPageVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Page.SetTo(paramsDotPageVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "page",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: limit.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "limit",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotLimitVal int
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToInt(val)
					if err != nil {
						return err
					}

					paramsDotLimitVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Limit.SetTo(paramsDotLimitVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "limit",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}
