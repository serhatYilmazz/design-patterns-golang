package media_player

import (
	"fmt"
)

type MediaType int

const (
	MP3 = iota
	VLC
	MP4
)

type MediaPlayer interface {
	Play(mediaType MediaType, fileName string) (string, error)
}

type AudioPlayer struct {
	Adapter *AdvancedMediaPlayerAdapter
}

func (a *AudioPlayer) Play(mediaType MediaType, fileName string) (string, error) {
	if mediaType == MP3 {
		return fmt.Sprintf("%s%s\n", fileName, ".mp3 is playing"), nil
	} else {
		return a.Adapter.Play(mediaType, fileName)
	}

	//return "", errors.New("not supported type") // No needed since adapter is handle this side.
}