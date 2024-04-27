// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
)

// Encode encodes ChannelsPostBadRequest as json.
func (s *ChannelsPostBadRequest) Encode(e *jx.Encoder) {
	unwrapped := (*CronResponse)(s)

	unwrapped.Encode(e)
}

// Decode decodes ChannelsPostBadRequest from json.
func (s *ChannelsPostBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ChannelsPostBadRequest to nil")
	}
	var unwrapped CronResponse
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = ChannelsPostBadRequest(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ChannelsPostBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ChannelsPostBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ChannelsPostForbidden as json.
func (s *ChannelsPostForbidden) Encode(e *jx.Encoder) {
	unwrapped := (*CronResponse)(s)

	unwrapped.Encode(e)
}

// Decode decodes ChannelsPostForbidden from json.
func (s *ChannelsPostForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ChannelsPostForbidden to nil")
	}
	var unwrapped CronResponse
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = ChannelsPostForbidden(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ChannelsPostForbidden) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ChannelsPostForbidden) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ChannelsPostInternalServerError as json.
func (s *ChannelsPostInternalServerError) Encode(e *jx.Encoder) {
	unwrapped := (*CronResponse)(s)

	unwrapped.Encode(e)
}

// Decode decodes ChannelsPostInternalServerError from json.
func (s *ChannelsPostInternalServerError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ChannelsPostInternalServerError to nil")
	}
	var unwrapped CronResponse
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = ChannelsPostInternalServerError(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ChannelsPostInternalServerError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ChannelsPostInternalServerError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ChannelsPostNotFound as json.
func (s *ChannelsPostNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*CronResponse)(s)

	unwrapped.Encode(e)
}

// Decode decodes ChannelsPostNotFound from json.
func (s *ChannelsPostNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ChannelsPostNotFound to nil")
	}
	var unwrapped CronResponse
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = ChannelsPostNotFound(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ChannelsPostNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ChannelsPostNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ChannelsPostOK as json.
func (s *ChannelsPostOK) Encode(e *jx.Encoder) {
	unwrapped := (*CronResponse)(s)

	unwrapped.Encode(e)
}

// Decode decodes ChannelsPostOK from json.
func (s *ChannelsPostOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ChannelsPostOK to nil")
	}
	var unwrapped CronResponse
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = ChannelsPostOK(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ChannelsPostOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ChannelsPostOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *ChannelsPostReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ChannelsPostReq) encodeFields(e *jx.Encoder) {
	{
		if s.PlatformType.Set {
			e.FieldStart("platform_type")
			s.PlatformType.Encode(e)
		}
	}
	{
		if s.Period.Set {
			e.FieldStart("period")
			s.Period.Encode(e)
		}
	}
	{
		if s.ChannelType.Set {
			e.FieldStart("channel_type")
			s.ChannelType.Encode(e)
		}
	}
}

var jsonFieldsNameOfChannelsPostReq = [3]string{
	0: "platform_type",
	1: "period",
	2: "channel_type",
}

// Decode decodes ChannelsPostReq from json.
func (s *ChannelsPostReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ChannelsPostReq to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "platform_type":
			if err := func() error {
				s.PlatformType.Reset()
				if err := s.PlatformType.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"platform_type\"")
			}
		case "period":
			if err := func() error {
				s.Period.Reset()
				if err := s.Period.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"period\"")
			}
		case "channel_type":
			if err := func() error {
				s.ChannelType.Reset()
				if err := s.ChannelType.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"channel_type\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ChannelsPostReq")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ChannelsPostReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ChannelsPostReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ChannelsPostReqChannelType as json.
func (s ChannelsPostReqChannelType) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes ChannelsPostReqChannelType from json.
func (s *ChannelsPostReqChannelType) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ChannelsPostReqChannelType to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch ChannelsPostReqChannelType(v) {
	case ChannelsPostReqChannelTypeVspo:
		*s = ChannelsPostReqChannelTypeVspo
	case ChannelsPostReqChannelTypeGeneral:
		*s = ChannelsPostReqChannelTypeGeneral
	default:
		*s = ChannelsPostReqChannelType(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s ChannelsPostReqChannelType) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ChannelsPostReqChannelType) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ChannelsPostReqPeriod as json.
func (s ChannelsPostReqPeriod) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes ChannelsPostReqPeriod from json.
func (s *ChannelsPostReqPeriod) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ChannelsPostReqPeriod to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch ChannelsPostReqPeriod(v) {
	case ChannelsPostReqPeriodDay:
		*s = ChannelsPostReqPeriodDay
	case ChannelsPostReqPeriodWeek:
		*s = ChannelsPostReqPeriodWeek
	case ChannelsPostReqPeriodMonth:
		*s = ChannelsPostReqPeriodMonth
	case ChannelsPostReqPeriodYear:
		*s = ChannelsPostReqPeriodYear
	default:
		*s = ChannelsPostReqPeriod(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s ChannelsPostReqPeriod) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ChannelsPostReqPeriod) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ChannelsPostReqPlatformType as json.
func (s ChannelsPostReqPlatformType) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes ChannelsPostReqPlatformType from json.
func (s *ChannelsPostReqPlatformType) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ChannelsPostReqPlatformType to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch ChannelsPostReqPlatformType(v) {
	case ChannelsPostReqPlatformTypeYoutube:
		*s = ChannelsPostReqPlatformTypeYoutube
	case ChannelsPostReqPlatformTypeTwitch:
		*s = ChannelsPostReqPlatformTypeTwitch
	case ChannelsPostReqPlatformTypeTwitcasting:
		*s = ChannelsPostReqPlatformTypeTwitcasting
	case ChannelsPostReqPlatformTypeNiconico:
		*s = ChannelsPostReqPlatformTypeNiconico
	default:
		*s = ChannelsPostReqPlatformType(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s ChannelsPostReqPlatformType) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ChannelsPostReqPlatformType) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ChannelsPostUnauthorized as json.
func (s *ChannelsPostUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*CronResponse)(s)

	unwrapped.Encode(e)
}

// Decode decodes ChannelsPostUnauthorized from json.
func (s *ChannelsPostUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ChannelsPostUnauthorized to nil")
	}
	var unwrapped CronResponse
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = ChannelsPostUnauthorized(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ChannelsPostUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ChannelsPostUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *CronResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *CronResponse) encodeFields(e *jx.Encoder) {
	{
		if s.Message.Set {
			e.FieldStart("message")
			s.Message.Encode(e)
		}
	}
}

var jsonFieldsNameOfCronResponse = [1]string{
	0: "message",
}

// Decode decodes CronResponse from json.
func (s *CronResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode CronResponse to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "message":
			if err := func() error {
				s.Message.Reset()
				if err := s.Message.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode CronResponse")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *CronResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *CronResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ChannelsPostReqChannelType as json.
func (o OptChannelsPostReqChannelType) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes ChannelsPostReqChannelType from json.
func (o *OptChannelsPostReqChannelType) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptChannelsPostReqChannelType to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptChannelsPostReqChannelType) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptChannelsPostReqChannelType) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ChannelsPostReqPeriod as json.
func (o OptChannelsPostReqPeriod) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes ChannelsPostReqPeriod from json.
func (o *OptChannelsPostReqPeriod) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptChannelsPostReqPeriod to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptChannelsPostReqPeriod) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptChannelsPostReqPeriod) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ChannelsPostReqPlatformType as json.
func (o OptChannelsPostReqPlatformType) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes ChannelsPostReqPlatformType from json.
func (o *OptChannelsPostReqPlatformType) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptChannelsPostReqPlatformType to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptChannelsPostReqPlatformType) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptChannelsPostReqPlatformType) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes string as json.
func (o OptString) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes string from json.
func (o *OptString) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptString to nil")
	}
	o.Set = true
	v, err := d.Str()
	if err != nil {
		return err
	}
	o.Value = string(v)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptString) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptString) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes VideosPostReqPeriod as json.
func (o OptVideosPostReqPeriod) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes VideosPostReqPeriod from json.
func (o *OptVideosPostReqPeriod) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptVideosPostReqPeriod to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptVideosPostReqPeriod) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptVideosPostReqPeriod) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes VideosPostReqVideoType as json.
func (o OptVideosPostReqVideoType) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes VideosPostReqVideoType from json.
func (o *OptVideosPostReqVideoType) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptVideosPostReqVideoType to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptVideosPostReqVideoType) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptVideosPostReqVideoType) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes VideosPostBadRequest as json.
func (s *VideosPostBadRequest) Encode(e *jx.Encoder) {
	unwrapped := (*CronResponse)(s)

	unwrapped.Encode(e)
}

// Decode decodes VideosPostBadRequest from json.
func (s *VideosPostBadRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode VideosPostBadRequest to nil")
	}
	var unwrapped CronResponse
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = VideosPostBadRequest(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *VideosPostBadRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *VideosPostBadRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes VideosPostForbidden as json.
func (s *VideosPostForbidden) Encode(e *jx.Encoder) {
	unwrapped := (*CronResponse)(s)

	unwrapped.Encode(e)
}

// Decode decodes VideosPostForbidden from json.
func (s *VideosPostForbidden) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode VideosPostForbidden to nil")
	}
	var unwrapped CronResponse
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = VideosPostForbidden(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *VideosPostForbidden) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *VideosPostForbidden) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes VideosPostInternalServerError as json.
func (s *VideosPostInternalServerError) Encode(e *jx.Encoder) {
	unwrapped := (*CronResponse)(s)

	unwrapped.Encode(e)
}

// Decode decodes VideosPostInternalServerError from json.
func (s *VideosPostInternalServerError) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode VideosPostInternalServerError to nil")
	}
	var unwrapped CronResponse
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = VideosPostInternalServerError(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *VideosPostInternalServerError) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *VideosPostInternalServerError) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes VideosPostNotFound as json.
func (s *VideosPostNotFound) Encode(e *jx.Encoder) {
	unwrapped := (*CronResponse)(s)

	unwrapped.Encode(e)
}

// Decode decodes VideosPostNotFound from json.
func (s *VideosPostNotFound) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode VideosPostNotFound to nil")
	}
	var unwrapped CronResponse
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = VideosPostNotFound(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *VideosPostNotFound) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *VideosPostNotFound) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes VideosPostOK as json.
func (s *VideosPostOK) Encode(e *jx.Encoder) {
	unwrapped := (*CronResponse)(s)

	unwrapped.Encode(e)
}

// Decode decodes VideosPostOK from json.
func (s *VideosPostOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode VideosPostOK to nil")
	}
	var unwrapped CronResponse
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = VideosPostOK(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *VideosPostOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *VideosPostOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *VideosPostReq) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *VideosPostReq) encodeFields(e *jx.Encoder) {
	{
		if s.PlatformType != nil {
			e.FieldStart("platform_type")
			e.ArrStart()
			for _, elem := range s.PlatformType {
				elem.Encode(e)
			}
			e.ArrEnd()
		}
	}
	{
		if s.Period.Set {
			e.FieldStart("period")
			s.Period.Encode(e)
		}
	}
	{
		if s.VideoType.Set {
			e.FieldStart("video_type")
			s.VideoType.Encode(e)
		}
	}
}

var jsonFieldsNameOfVideosPostReq = [3]string{
	0: "platform_type",
	1: "period",
	2: "video_type",
}

// Decode decodes VideosPostReq from json.
func (s *VideosPostReq) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode VideosPostReq to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "platform_type":
			if err := func() error {
				s.PlatformType = make([]VideosPostReqPlatformTypeItem, 0)
				if err := d.Arr(func(d *jx.Decoder) error {
					var elem VideosPostReqPlatformTypeItem
					if err := elem.Decode(d); err != nil {
						return err
					}
					s.PlatformType = append(s.PlatformType, elem)
					return nil
				}); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"platform_type\"")
			}
		case "period":
			if err := func() error {
				s.Period.Reset()
				if err := s.Period.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"period\"")
			}
		case "video_type":
			if err := func() error {
				s.VideoType.Reset()
				if err := s.VideoType.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"video_type\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode VideosPostReq")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *VideosPostReq) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *VideosPostReq) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes VideosPostReqPeriod as json.
func (s VideosPostReqPeriod) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes VideosPostReqPeriod from json.
func (s *VideosPostReqPeriod) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode VideosPostReqPeriod to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch VideosPostReqPeriod(v) {
	case VideosPostReqPeriodDay:
		*s = VideosPostReqPeriodDay
	case VideosPostReqPeriodMonth:
		*s = VideosPostReqPeriodMonth
	case VideosPostReqPeriodWeek:
		*s = VideosPostReqPeriodWeek
	default:
		*s = VideosPostReqPeriod(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s VideosPostReqPeriod) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *VideosPostReqPeriod) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes VideosPostReqPlatformTypeItem as json.
func (s VideosPostReqPlatformTypeItem) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes VideosPostReqPlatformTypeItem from json.
func (s *VideosPostReqPlatformTypeItem) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode VideosPostReqPlatformTypeItem to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch VideosPostReqPlatformTypeItem(v) {
	case VideosPostReqPlatformTypeItemYoutube:
		*s = VideosPostReqPlatformTypeItemYoutube
	case VideosPostReqPlatformTypeItemTwitch:
		*s = VideosPostReqPlatformTypeItemTwitch
	case VideosPostReqPlatformTypeItemTwitcasting:
		*s = VideosPostReqPlatformTypeItemTwitcasting
	case VideosPostReqPlatformTypeItemNiconico:
		*s = VideosPostReqPlatformTypeItemNiconico
	default:
		*s = VideosPostReqPlatformTypeItem(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s VideosPostReqPlatformTypeItem) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *VideosPostReqPlatformTypeItem) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes VideosPostReqVideoType as json.
func (s VideosPostReqVideoType) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes VideosPostReqVideoType from json.
func (s *VideosPostReqVideoType) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode VideosPostReqVideoType to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch VideosPostReqVideoType(v) {
	case VideosPostReqVideoTypeVspoBroadcast:
		*s = VideosPostReqVideoTypeVspoBroadcast
	case VideosPostReqVideoTypeClip:
		*s = VideosPostReqVideoTypeClip
	case VideosPostReqVideoTypeFreechat:
		*s = VideosPostReqVideoTypeFreechat
	default:
		*s = VideosPostReqVideoType(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s VideosPostReqVideoType) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *VideosPostReqVideoType) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes VideosPostUnauthorized as json.
func (s *VideosPostUnauthorized) Encode(e *jx.Encoder) {
	unwrapped := (*CronResponse)(s)

	unwrapped.Encode(e)
}

// Decode decodes VideosPostUnauthorized from json.
func (s *VideosPostUnauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode VideosPostUnauthorized to nil")
	}
	var unwrapped CronResponse
	if err := func() error {
		if err := unwrapped.Decode(d); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = VideosPostUnauthorized(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *VideosPostUnauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *VideosPostUnauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}