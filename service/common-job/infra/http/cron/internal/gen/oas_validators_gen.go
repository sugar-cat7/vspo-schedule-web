// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"
)

func (s APICronChannelsGetPlatformType) Validate() error {
	switch s {
	case "youtube":
		return nil
	case "twitch":
		return nil
	case "twitcasting":
		return nil
	case "niconico":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s APICronCreatorsGetCreatorTypeItem) Validate() error {
	switch s {
	case "vspo":
		return nil
	case "general":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s APICronCreatorsGetPeriod) Validate() error {
	switch s {
	case "day":
		return nil
	case "week":
		return nil
	case "month":
		return nil
	case "year":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s APICronCreatorsGetPlatformTypeItem) Validate() error {
	switch s {
	case "youtube":
		return nil
	case "twitch":
		return nil
	case "twitcasting":
		return nil
	case "niconico":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s APICronVideosGetPeriod) Validate() error {
	switch s {
	case "day":
		return nil
	case "month":
		return nil
	case "week":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s APICronVideosGetPlatformTypeItem) Validate() error {
	switch s {
	case "youtube":
		return nil
	case "twitch":
		return nil
	case "twitcasting":
		return nil
	case "niconico":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s APICronVideosGetVideoType) Validate() error {
	switch s {
	case "vspo_stream":
		return nil
	case "clip":
		return nil
	case "freechat":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}
