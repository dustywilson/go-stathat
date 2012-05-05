package stathat

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type StatHat struct {
	ezkey string
}

func NewStatHat(ezkey string) *StatHat {
	sh := new(StatHat)
	sh.ezkey = ezkey
	return sh
}

func (sh *StatHat) Count(stat string, count float32) error {
	if sh.ezkey == "" {
		return errors.New("No StatHat ezkey was set.")
	}
	r, err := http.PostForm("https://api.stathat.com/ez", url.Values{
		"ezkey": {sh.ezkey},
		"stat":  {stat},
		"count": {fmt.Sprintf("%f", count)},
	})
	if err != nil {
		return err
	}
	r.Body.Close()
	return nil
}

func (sh *StatHat) Value(stat string, value float32) error {
	if sh.ezkey == "" {
		return errors.New("No StatHat ezkey was set.")
	}
	r, err := http.PostForm("https://api.stathat.com/ez", url.Values{
		"ezkey": {sh.ezkey},
		"stat":  {stat},
		"value": {fmt.Sprintf("%f", value)},
	})
	if err != nil {
		return err
	}
	r.Body.Close()
	return nil
}
