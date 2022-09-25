package lastfm

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/shkh/lastfm-go"
	"go-musicfox/pkg/constants"
	"go-musicfox/utils"
)

type AuthInvalid struct {
	error
}

type Client struct {
	api *lastfm_go.Api
}

func NewClient() *Client {
	return &Client{
		api: lastfm_go.New(constants.LastfmKey, constants.LastfmSecret),
	}
}

func (c *Client) errorHandle(e error) (bool, error) {
	if e == nil {
		return false, nil
	}
	if lastfmErr, ok := e.(*lastfm_go.LastfmError); ok {
		switch lastfmErr.Code {
		case 9: // invalid session key
			return false, AuthInvalid{lastfmErr}
		case 11, 16: // server error
			return true, e
		default:
			utils.Logger().Printf("[ERROR] Lastfm request err: %+v", lastfmErr)
			return false, e
		}
	}
	utils.Logger().Printf("[ERROR] Lastfm other err: %+v", e)
	return false, e
}

func (c *Client) GetAuthUrlWithToken() (token, url string, err error) {
	if constants.LastfmKey == "" || constants.LastfmSecret == "" {
		return "", "", errors.New("lastfm key或secret为空")
	}
	token, err = c.api.GetToken()
	if _, err = c.errorHandle(err); err != nil {
		return
	}

	url = fmt.Sprintf(constants.LastfmAuthUrl, constants.LastfmKey, token)
	return
}

func (c *Client) SetSession(session string) {
	c.api.SetSession(session)
}

func (c *Client) GetSession(token string) (sessionKey string, err error) {
	err = c.api.LoginWithToken(token)
	if _, err = c.errorHandle(err); err != nil {
		return
	}
	sessionKey = c.api.GetSessionKey()
	return
}

func (c *Client) UpdateNowPlaying(args map[string]interface{}) error {
	if c.api.GetSessionKey() == "" {
		_, err := c.errorHandle(errors.New("empty session key"))
		return err
	}
	_, err := c.api.Track.UpdateNowPlaying(args)

	var retry bool
	if retry, err = c.errorHandle(err); retry {
		return c.UpdateNowPlaying(args)
	}
	return err
}

func (c *Client) Scrobble(args map[string]interface{}) error {
	if c.api.GetSessionKey() == "" {
		_, err := c.errorHandle(errors.New("empty session key"))
		return err
	}
	_, err := c.api.Track.Scrobble(args)

	var retry bool
	if retry, err = c.errorHandle(err); retry {
		return c.Scrobble(args)
	}
	return err
}

func (c *Client) GetUserInfo(args map[string]interface{}) (lastfm_go.UserGetInfo, error) {
	if c.api.GetSessionKey() == "" {
		_, err := c.errorHandle(errors.New("empty session key"))
		return lastfm_go.UserGetInfo{}, err
	}
	userInfo, err := c.api.User.GetInfo(args)

	var retry bool
	if retry, err = c.errorHandle(err); retry {
		return c.GetUserInfo(args)
	}
	return userInfo, err
}