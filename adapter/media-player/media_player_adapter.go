package media_player

import (
	"errors"
)

type AdvancedMediaPlayerAdapter struct {
	Amp AdvancedMediaPlayer
}

func (a *AdvancedMediaPlayerAdapter) Play(mediaType MediaType, fileName string) (string, error) {
	var res string
	var err error
	if mediaType == MP4 {
		res, err = a.Amp.PlayMp4(fileName)
	} else if mediaType == VLC {
		res, err = a.Amp.PlayVlc(fileName)
	} else {
		err = errors.New("media type is not supported")
	}

	return res, err
}