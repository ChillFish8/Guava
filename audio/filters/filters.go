package filters

import "fmt"

//
// A base ffmpeg filter (Volume and Speed already included)
// anything can be given to this and will be formatted to
// ffmpeg as `name=k1=v1:k2=v2`

type Filter struct {
	Name   string            `json:"filter_name"`
	Values map[string]string `json:"filter_values"`
}

func (f *Filter) SetValue(name string, value string) (bool, string) {
	f.Values[name] = value
	return true, "OK"
}

func (f *Filter) ToString() string {
	base := f.Name + "="
	count := 0
	maximum := len(f.Values)
	for k, v := range f.Values {
		base += fmt.Sprintf("%s=%s", k, v)
		if count < maximum-1 {
			base += ":"
		}
	}
	return base
}

// Some helper functions for making a new filter manager.
func NewFFMpegFilters(vol int, speed float32) FFMpegFilters {
	return FFMpegFilters{
		Volume:       vol,
		Speed:        speed,
		OtherFilters: map[string]Filter{},
	}
}

// A struct that contains the basic volume and speed filters for ffmpeg
// as well as adding optional filters which is then formatted before starting.
type FFMpegFilters struct {
	Volume       int     `json:"volume"`
	Speed        float32 `json:"speed"`
	OtherFilters map[string]Filter
}

func (f *FFMpegFilters) SetVolume(v int) (bool, string) {
	if (v > 100) || (v < 0) {
		return false, "Volume not in range 0 - 100"
	}
	f.Volume = v
	return true, "OK"
}

func (f *FFMpegFilters) SetSpeed(s float32) (bool, string) {
	if (s > 10) || (s < 0) {
		return false, "Speed not in range 0 - 100"
	}
	f.Speed = s
	return true, "OK"
}

func (f *FFMpegFilters) SetFilter(filter Filter) (bool, string) {
	f.OtherFilters[filter.Name] = filter
	return true, "OK"
}

func (f *FFMpegFilters) ToString() string {
	base := ""
	for _, v := range f.OtherFilters {
		base += v.ToString() + ","
	}
	base += fmt.Sprintf("volume=%v,atempo=%v", f.Volume, f.Speed)

	return base
}
