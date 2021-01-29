package managers

import (
	"github.com/karimabedrabbo/eyo/api/apputils"

	"github.com/microcosm-cc/bluemonday"
)

type Sanitize struct {
	SanitizeClient *bluemonday.Policy
}
var sanitizer *Sanitize

func SetupSanitize() *Sanitize {
	return &Sanitize {
		SanitizeClient: bluemonday.StrictPolicy(),
	}
}

func InitSanitize() {
	sanitizer = SetupSanitize()
}

func GetSanitize() *Sanitize {
	return sanitizer
}

func (s *Sanitize) SanitizeString(str string) string {
	return s.SanitizeClient.Sanitize(str)
}

func (s *Sanitize) SanitizeJsonObj(obj interface{}) (interface{}, error) {
	var err error
	var objBytes []byte
	if objBytes, err = apputils.Marshal(obj); err != nil {
		return nil, err
	}
	return s.SanitizeJsonBytes(objBytes)
}

func (s *Sanitize) SanitizeJsonBytes(b []byte) (interface{}, error) {
	var err error
	var i interface{}
	if i, err = apputils.Unmarshal(b); err != nil {
		return nil, err
	}

	s.sanitize(i)

	return i, nil
}

func (s *Sanitize) sanitize(data interface{}) {
	switch d := data.(type) {
	case map[string]interface{}:
		for k, v := range d {
			switch tv := v.(type) {
			case string:
				d[k] = sanitizer.SanitizeClient.Sanitize(tv)
			case map[string]interface{}:
				s.sanitize(tv)
			case []interface{}:
				s.sanitize(tv)
			case nil:
				delete(d, k)
			}
		}
	case []interface{}:
		if len(d) > 0 {
			switch d[0].(type) {
			case string:
				for i, s := range d {
					d[i] = sanitizer.SanitizeClient.Sanitize(s.(string))
				}
			case map[string]interface{}:
				for _, t := range d {
					s.sanitize(t)
				}
			case []interface{}:
				for _, t := range d {
					s.sanitize(t)
				}
			}
		}
	}
}