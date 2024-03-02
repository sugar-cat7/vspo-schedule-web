// Code generated by ogen, DO NOT EDIT.

package api

import (
	"github.com/go-faster/errors"
)

type ApiKeyAuth struct {
	APIKey string
}

// GetAPIKey returns the value of APIKey.
func (s *ApiKeyAuth) GetAPIKey() string {
	return s.APIKey
}

// SetAPIKey sets the value of APIKey.
func (s *ApiKeyAuth) SetAPIKey(val string) {
	s.APIKey = val
}

type ChannelsPostBadRequest CronResponse

func (*ChannelsPostBadRequest) channelsPostRes() {}

type ChannelsPostForbidden CronResponse

func (*ChannelsPostForbidden) channelsPostRes() {}

type ChannelsPostInternalServerError CronResponse

func (*ChannelsPostInternalServerError) channelsPostRes() {}

type ChannelsPostNotFound CronResponse

func (*ChannelsPostNotFound) channelsPostRes() {}

type ChannelsPostOK CronResponse

func (*ChannelsPostOK) channelsPostRes() {}

type ChannelsPostReq struct {
	// Video's platform type.
	PlatformType OptChannelsPostReqPlatformType `json:"platform_type"`
	// Period for performing updates.
	Period OptChannelsPostReqPeriod `json:"period"`
	// Period for performing updates.
	ChannelType OptChannelsPostReqChannelType `json:"channel_type"`
}

// GetPlatformType returns the value of PlatformType.
func (s *ChannelsPostReq) GetPlatformType() OptChannelsPostReqPlatformType {
	return s.PlatformType
}

// GetPeriod returns the value of Period.
func (s *ChannelsPostReq) GetPeriod() OptChannelsPostReqPeriod {
	return s.Period
}

// GetChannelType returns the value of ChannelType.
func (s *ChannelsPostReq) GetChannelType() OptChannelsPostReqChannelType {
	return s.ChannelType
}

// SetPlatformType sets the value of PlatformType.
func (s *ChannelsPostReq) SetPlatformType(val OptChannelsPostReqPlatformType) {
	s.PlatformType = val
}

// SetPeriod sets the value of Period.
func (s *ChannelsPostReq) SetPeriod(val OptChannelsPostReqPeriod) {
	s.Period = val
}

// SetChannelType sets the value of ChannelType.
func (s *ChannelsPostReq) SetChannelType(val OptChannelsPostReqChannelType) {
	s.ChannelType = val
}

// Period for performing updates.
type ChannelsPostReqChannelType string

const (
	ChannelsPostReqChannelTypeVspo ChannelsPostReqChannelType = "vspo"
	ChannelsPostReqChannelTypeAll  ChannelsPostReqChannelType = "all"
)

// AllValues returns all ChannelsPostReqChannelType values.
func (ChannelsPostReqChannelType) AllValues() []ChannelsPostReqChannelType {
	return []ChannelsPostReqChannelType{
		ChannelsPostReqChannelTypeVspo,
		ChannelsPostReqChannelTypeAll,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s ChannelsPostReqChannelType) MarshalText() ([]byte, error) {
	switch s {
	case ChannelsPostReqChannelTypeVspo:
		return []byte(s), nil
	case ChannelsPostReqChannelTypeAll:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *ChannelsPostReqChannelType) UnmarshalText(data []byte) error {
	switch ChannelsPostReqChannelType(data) {
	case ChannelsPostReqChannelTypeVspo:
		*s = ChannelsPostReqChannelTypeVspo
		return nil
	case ChannelsPostReqChannelTypeAll:
		*s = ChannelsPostReqChannelTypeAll
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Period for performing updates.
type ChannelsPostReqPeriod string

const (
	ChannelsPostReqPeriodAll   ChannelsPostReqPeriod = "all"
	ChannelsPostReqPeriodDay   ChannelsPostReqPeriod = "day"
	ChannelsPostReqPeriodMonth ChannelsPostReqPeriod = "month"
	ChannelsPostReqPeriodWeek  ChannelsPostReqPeriod = "week"
)

// AllValues returns all ChannelsPostReqPeriod values.
func (ChannelsPostReqPeriod) AllValues() []ChannelsPostReqPeriod {
	return []ChannelsPostReqPeriod{
		ChannelsPostReqPeriodAll,
		ChannelsPostReqPeriodDay,
		ChannelsPostReqPeriodMonth,
		ChannelsPostReqPeriodWeek,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s ChannelsPostReqPeriod) MarshalText() ([]byte, error) {
	switch s {
	case ChannelsPostReqPeriodAll:
		return []byte(s), nil
	case ChannelsPostReqPeriodDay:
		return []byte(s), nil
	case ChannelsPostReqPeriodMonth:
		return []byte(s), nil
	case ChannelsPostReqPeriodWeek:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *ChannelsPostReqPeriod) UnmarshalText(data []byte) error {
	switch ChannelsPostReqPeriod(data) {
	case ChannelsPostReqPeriodAll:
		*s = ChannelsPostReqPeriodAll
		return nil
	case ChannelsPostReqPeriodDay:
		*s = ChannelsPostReqPeriodDay
		return nil
	case ChannelsPostReqPeriodMonth:
		*s = ChannelsPostReqPeriodMonth
		return nil
	case ChannelsPostReqPeriodWeek:
		*s = ChannelsPostReqPeriodWeek
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Video's platform type.
type ChannelsPostReqPlatformType string

const (
	ChannelsPostReqPlatformTypeAll         ChannelsPostReqPlatformType = "all"
	ChannelsPostReqPlatformTypeYoutube     ChannelsPostReqPlatformType = "youtube"
	ChannelsPostReqPlatformTypeTwitch      ChannelsPostReqPlatformType = "twitch"
	ChannelsPostReqPlatformTypeTwitcasting ChannelsPostReqPlatformType = "twitcasting"
	ChannelsPostReqPlatformTypeNiconico    ChannelsPostReqPlatformType = "niconico"
)

// AllValues returns all ChannelsPostReqPlatformType values.
func (ChannelsPostReqPlatformType) AllValues() []ChannelsPostReqPlatformType {
	return []ChannelsPostReqPlatformType{
		ChannelsPostReqPlatformTypeAll,
		ChannelsPostReqPlatformTypeYoutube,
		ChannelsPostReqPlatformTypeTwitch,
		ChannelsPostReqPlatformTypeTwitcasting,
		ChannelsPostReqPlatformTypeNiconico,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s ChannelsPostReqPlatformType) MarshalText() ([]byte, error) {
	switch s {
	case ChannelsPostReqPlatformTypeAll:
		return []byte(s), nil
	case ChannelsPostReqPlatformTypeYoutube:
		return []byte(s), nil
	case ChannelsPostReqPlatformTypeTwitch:
		return []byte(s), nil
	case ChannelsPostReqPlatformTypeTwitcasting:
		return []byte(s), nil
	case ChannelsPostReqPlatformTypeNiconico:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *ChannelsPostReqPlatformType) UnmarshalText(data []byte) error {
	switch ChannelsPostReqPlatformType(data) {
	case ChannelsPostReqPlatformTypeAll:
		*s = ChannelsPostReqPlatformTypeAll
		return nil
	case ChannelsPostReqPlatformTypeYoutube:
		*s = ChannelsPostReqPlatformTypeYoutube
		return nil
	case ChannelsPostReqPlatformTypeTwitch:
		*s = ChannelsPostReqPlatformTypeTwitch
		return nil
	case ChannelsPostReqPlatformTypeTwitcasting:
		*s = ChannelsPostReqPlatformTypeTwitcasting
		return nil
	case ChannelsPostReqPlatformTypeNiconico:
		*s = ChannelsPostReqPlatformTypeNiconico
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type ChannelsPostUnauthorized CronResponse

func (*ChannelsPostUnauthorized) channelsPostRes() {}

// Ref: #/components/schemas/CronResponse
type CronResponse struct {
	Message OptString `json:"message"`
}

// GetMessage returns the value of Message.
func (s *CronResponse) GetMessage() OptString {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *CronResponse) SetMessage(val OptString) {
	s.Message = val
}

// NewOptChannelsPostReqChannelType returns new OptChannelsPostReqChannelType with value set to v.
func NewOptChannelsPostReqChannelType(v ChannelsPostReqChannelType) OptChannelsPostReqChannelType {
	return OptChannelsPostReqChannelType{
		Value: v,
		Set:   true,
	}
}

// OptChannelsPostReqChannelType is optional ChannelsPostReqChannelType.
type OptChannelsPostReqChannelType struct {
	Value ChannelsPostReqChannelType
	Set   bool
}

// IsSet returns true if OptChannelsPostReqChannelType was set.
func (o OptChannelsPostReqChannelType) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptChannelsPostReqChannelType) Reset() {
	var v ChannelsPostReqChannelType
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptChannelsPostReqChannelType) SetTo(v ChannelsPostReqChannelType) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptChannelsPostReqChannelType) Get() (v ChannelsPostReqChannelType, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptChannelsPostReqChannelType) Or(d ChannelsPostReqChannelType) ChannelsPostReqChannelType {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptChannelsPostReqPeriod returns new OptChannelsPostReqPeriod with value set to v.
func NewOptChannelsPostReqPeriod(v ChannelsPostReqPeriod) OptChannelsPostReqPeriod {
	return OptChannelsPostReqPeriod{
		Value: v,
		Set:   true,
	}
}

// OptChannelsPostReqPeriod is optional ChannelsPostReqPeriod.
type OptChannelsPostReqPeriod struct {
	Value ChannelsPostReqPeriod
	Set   bool
}

// IsSet returns true if OptChannelsPostReqPeriod was set.
func (o OptChannelsPostReqPeriod) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptChannelsPostReqPeriod) Reset() {
	var v ChannelsPostReqPeriod
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptChannelsPostReqPeriod) SetTo(v ChannelsPostReqPeriod) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptChannelsPostReqPeriod) Get() (v ChannelsPostReqPeriod, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptChannelsPostReqPeriod) Or(d ChannelsPostReqPeriod) ChannelsPostReqPeriod {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptChannelsPostReqPlatformType returns new OptChannelsPostReqPlatformType with value set to v.
func NewOptChannelsPostReqPlatformType(v ChannelsPostReqPlatformType) OptChannelsPostReqPlatformType {
	return OptChannelsPostReqPlatformType{
		Value: v,
		Set:   true,
	}
}

// OptChannelsPostReqPlatformType is optional ChannelsPostReqPlatformType.
type OptChannelsPostReqPlatformType struct {
	Value ChannelsPostReqPlatformType
	Set   bool
}

// IsSet returns true if OptChannelsPostReqPlatformType was set.
func (o OptChannelsPostReqPlatformType) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptChannelsPostReqPlatformType) Reset() {
	var v ChannelsPostReqPlatformType
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptChannelsPostReqPlatformType) SetTo(v ChannelsPostReqPlatformType) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptChannelsPostReqPlatformType) Get() (v ChannelsPostReqPlatformType, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptChannelsPostReqPlatformType) Or(d ChannelsPostReqPlatformType) ChannelsPostReqPlatformType {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptVideosPostReqPeriod returns new OptVideosPostReqPeriod with value set to v.
func NewOptVideosPostReqPeriod(v VideosPostReqPeriod) OptVideosPostReqPeriod {
	return OptVideosPostReqPeriod{
		Value: v,
		Set:   true,
	}
}

// OptVideosPostReqPeriod is optional VideosPostReqPeriod.
type OptVideosPostReqPeriod struct {
	Value VideosPostReqPeriod
	Set   bool
}

// IsSet returns true if OptVideosPostReqPeriod was set.
func (o OptVideosPostReqPeriod) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptVideosPostReqPeriod) Reset() {
	var v VideosPostReqPeriod
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptVideosPostReqPeriod) SetTo(v VideosPostReqPeriod) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptVideosPostReqPeriod) Get() (v VideosPostReqPeriod, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptVideosPostReqPeriod) Or(d VideosPostReqPeriod) VideosPostReqPeriod {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptVideosPostReqPlatformType returns new OptVideosPostReqPlatformType with value set to v.
func NewOptVideosPostReqPlatformType(v VideosPostReqPlatformType) OptVideosPostReqPlatformType {
	return OptVideosPostReqPlatformType{
		Value: v,
		Set:   true,
	}
}

// OptVideosPostReqPlatformType is optional VideosPostReqPlatformType.
type OptVideosPostReqPlatformType struct {
	Value VideosPostReqPlatformType
	Set   bool
}

// IsSet returns true if OptVideosPostReqPlatformType was set.
func (o OptVideosPostReqPlatformType) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptVideosPostReqPlatformType) Reset() {
	var v VideosPostReqPlatformType
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptVideosPostReqPlatformType) SetTo(v VideosPostReqPlatformType) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptVideosPostReqPlatformType) Get() (v VideosPostReqPlatformType, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptVideosPostReqPlatformType) Or(d VideosPostReqPlatformType) VideosPostReqPlatformType {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptVideosPostReqVideoType returns new OptVideosPostReqVideoType with value set to v.
func NewOptVideosPostReqVideoType(v VideosPostReqVideoType) OptVideosPostReqVideoType {
	return OptVideosPostReqVideoType{
		Value: v,
		Set:   true,
	}
}

// OptVideosPostReqVideoType is optional VideosPostReqVideoType.
type OptVideosPostReqVideoType struct {
	Value VideosPostReqVideoType
	Set   bool
}

// IsSet returns true if OptVideosPostReqVideoType was set.
func (o OptVideosPostReqVideoType) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptVideosPostReqVideoType) Reset() {
	var v VideosPostReqVideoType
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptVideosPostReqVideoType) SetTo(v VideosPostReqVideoType) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptVideosPostReqVideoType) Get() (v VideosPostReqVideoType, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptVideosPostReqVideoType) Or(d VideosPostReqVideoType) VideosPostReqVideoType {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

type VideosPostBadRequest CronResponse

func (*VideosPostBadRequest) videosPostRes() {}

type VideosPostForbidden CronResponse

func (*VideosPostForbidden) videosPostRes() {}

type VideosPostInternalServerError CronResponse

func (*VideosPostInternalServerError) videosPostRes() {}

type VideosPostNotFound CronResponse

func (*VideosPostNotFound) videosPostRes() {}

type VideosPostOK CronResponse

func (*VideosPostOK) videosPostRes() {}

type VideosPostReq struct {
	// Video's platform type.
	PlatformType OptVideosPostReqPlatformType `json:"platform_type"`
	// Period for performing updates.
	Period    OptVideosPostReqPeriod    `json:"period"`
	VideoType OptVideosPostReqVideoType `json:"video_type"`
}

// GetPlatformType returns the value of PlatformType.
func (s *VideosPostReq) GetPlatformType() OptVideosPostReqPlatformType {
	return s.PlatformType
}

// GetPeriod returns the value of Period.
func (s *VideosPostReq) GetPeriod() OptVideosPostReqPeriod {
	return s.Period
}

// GetVideoType returns the value of VideoType.
func (s *VideosPostReq) GetVideoType() OptVideosPostReqVideoType {
	return s.VideoType
}

// SetPlatformType sets the value of PlatformType.
func (s *VideosPostReq) SetPlatformType(val OptVideosPostReqPlatformType) {
	s.PlatformType = val
}

// SetPeriod sets the value of Period.
func (s *VideosPostReq) SetPeriod(val OptVideosPostReqPeriod) {
	s.Period = val
}

// SetVideoType sets the value of VideoType.
func (s *VideosPostReq) SetVideoType(val OptVideosPostReqVideoType) {
	s.VideoType = val
}

// Period for performing updates.
type VideosPostReqPeriod string

const (
	VideosPostReqPeriodAll   VideosPostReqPeriod = "all"
	VideosPostReqPeriodDay   VideosPostReqPeriod = "day"
	VideosPostReqPeriodMonth VideosPostReqPeriod = "month"
	VideosPostReqPeriodWeek  VideosPostReqPeriod = "week"
)

// AllValues returns all VideosPostReqPeriod values.
func (VideosPostReqPeriod) AllValues() []VideosPostReqPeriod {
	return []VideosPostReqPeriod{
		VideosPostReqPeriodAll,
		VideosPostReqPeriodDay,
		VideosPostReqPeriodMonth,
		VideosPostReqPeriodWeek,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s VideosPostReqPeriod) MarshalText() ([]byte, error) {
	switch s {
	case VideosPostReqPeriodAll:
		return []byte(s), nil
	case VideosPostReqPeriodDay:
		return []byte(s), nil
	case VideosPostReqPeriodMonth:
		return []byte(s), nil
	case VideosPostReqPeriodWeek:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *VideosPostReqPeriod) UnmarshalText(data []byte) error {
	switch VideosPostReqPeriod(data) {
	case VideosPostReqPeriodAll:
		*s = VideosPostReqPeriodAll
		return nil
	case VideosPostReqPeriodDay:
		*s = VideosPostReqPeriodDay
		return nil
	case VideosPostReqPeriodMonth:
		*s = VideosPostReqPeriodMonth
		return nil
	case VideosPostReqPeriodWeek:
		*s = VideosPostReqPeriodWeek
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Video's platform type.
type VideosPostReqPlatformType string

const (
	VideosPostReqPlatformTypeAll         VideosPostReqPlatformType = "all"
	VideosPostReqPlatformTypeYoutube     VideosPostReqPlatformType = "youtube"
	VideosPostReqPlatformTypeTwitch      VideosPostReqPlatformType = "twitch"
	VideosPostReqPlatformTypeTwitcasting VideosPostReqPlatformType = "twitcasting"
	VideosPostReqPlatformTypeNiconico    VideosPostReqPlatformType = "niconico"
)

// AllValues returns all VideosPostReqPlatformType values.
func (VideosPostReqPlatformType) AllValues() []VideosPostReqPlatformType {
	return []VideosPostReqPlatformType{
		VideosPostReqPlatformTypeAll,
		VideosPostReqPlatformTypeYoutube,
		VideosPostReqPlatformTypeTwitch,
		VideosPostReqPlatformTypeTwitcasting,
		VideosPostReqPlatformTypeNiconico,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s VideosPostReqPlatformType) MarshalText() ([]byte, error) {
	switch s {
	case VideosPostReqPlatformTypeAll:
		return []byte(s), nil
	case VideosPostReqPlatformTypeYoutube:
		return []byte(s), nil
	case VideosPostReqPlatformTypeTwitch:
		return []byte(s), nil
	case VideosPostReqPlatformTypeTwitcasting:
		return []byte(s), nil
	case VideosPostReqPlatformTypeNiconico:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *VideosPostReqPlatformType) UnmarshalText(data []byte) error {
	switch VideosPostReqPlatformType(data) {
	case VideosPostReqPlatformTypeAll:
		*s = VideosPostReqPlatformTypeAll
		return nil
	case VideosPostReqPlatformTypeYoutube:
		*s = VideosPostReqPlatformTypeYoutube
		return nil
	case VideosPostReqPlatformTypeTwitch:
		*s = VideosPostReqPlatformTypeTwitch
		return nil
	case VideosPostReqPlatformTypeTwitcasting:
		*s = VideosPostReqPlatformTypeTwitcasting
		return nil
	case VideosPostReqPlatformTypeNiconico:
		*s = VideosPostReqPlatformTypeNiconico
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type VideosPostReqVideoType string

const (
	VideosPostReqVideoTypeAll           VideosPostReqVideoType = "all"
	VideosPostReqVideoTypeVspoBroadcast VideosPostReqVideoType = "vspo_broadcast"
	VideosPostReqVideoTypeClip          VideosPostReqVideoType = "clip"
	VideosPostReqVideoTypeFreechat      VideosPostReqVideoType = "freechat"
)

// AllValues returns all VideosPostReqVideoType values.
func (VideosPostReqVideoType) AllValues() []VideosPostReqVideoType {
	return []VideosPostReqVideoType{
		VideosPostReqVideoTypeAll,
		VideosPostReqVideoTypeVspoBroadcast,
		VideosPostReqVideoTypeClip,
		VideosPostReqVideoTypeFreechat,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s VideosPostReqVideoType) MarshalText() ([]byte, error) {
	switch s {
	case VideosPostReqVideoTypeAll:
		return []byte(s), nil
	case VideosPostReqVideoTypeVspoBroadcast:
		return []byte(s), nil
	case VideosPostReqVideoTypeClip:
		return []byte(s), nil
	case VideosPostReqVideoTypeFreechat:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *VideosPostReqVideoType) UnmarshalText(data []byte) error {
	switch VideosPostReqVideoType(data) {
	case VideosPostReqVideoTypeAll:
		*s = VideosPostReqVideoTypeAll
		return nil
	case VideosPostReqVideoTypeVspoBroadcast:
		*s = VideosPostReqVideoTypeVspoBroadcast
		return nil
	case VideosPostReqVideoTypeClip:
		*s = VideosPostReqVideoTypeClip
		return nil
	case VideosPostReqVideoTypeFreechat:
		*s = VideosPostReqVideoTypeFreechat
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type VideosPostUnauthorized CronResponse

func (*VideosPostUnauthorized) videosPostRes() {}

type YoutubeApiKey struct {
	APIKey string
}

// GetAPIKey returns the value of APIKey.
func (s *YoutubeApiKey) GetAPIKey() string {
	return s.APIKey
}

// SetAPIKey sets the value of APIKey.
func (s *YoutubeApiKey) SetAPIKey(val string) {
	s.APIKey = val
}
