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
	// Comma-separated list of creator IDs.
	Ids OptString
	// Page.
	Page OptInt
	// Page limit.
	Limit OptInt
}

func unpackCreatorsGetParams(packed middleware.Parameters) (params CreatorsGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "ids",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Ids = v.(OptString)
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
	// Decode query: ids.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "ids",
			Style:   uri.QueryStyleForm,
			Explode: false,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotIdsVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotIdsVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Ids.SetTo(paramsDotIdsVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "ids",
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
	// Unique identifier of the creator.
	CreatorID OptString
	// Comma-separated list of videos IDs.
	Ids OptString
	// Start Date.
	StartDate OptString
	// End Date.
	EndDate OptString
	// Page.
	Page OptInt
	// Page limit.
	Limit OptInt
}

func unpackVideosGetParams(packed middleware.Parameters) (params VideosGetParams) {
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
			Name: "ids",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Ids = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "start_date",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.StartDate = v.(OptString)
		}
	}
	{
		key := middleware.ParameterKey{
			Name: "end_date",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.EndDate = v.(OptString)
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
	// Decode query: ids.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "ids",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotIdsVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotIdsVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Ids.SetTo(paramsDotIdsVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "ids",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: start_date.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "start_date",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotStartDateVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotStartDateVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.StartDate.SetTo(paramsDotStartDateVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "start_date",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: end_date.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "end_date",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotEndDateVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotEndDateVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.EndDate.SetTo(paramsDotEndDateVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "end_date",
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
