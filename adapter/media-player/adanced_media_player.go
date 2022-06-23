package media_player

import (
	"errors"
	"fmt"
)

type AdvancedMediaPlayer interface {
	PlayMp4(fileName string) (string, error)
	PlayVlc(fileName string) (string, error)
}

type Mp4Player struct {

}

func (m *Mp4Player) PlayMp4(fileName string) (string, error) {
	return fmt.Sprintf("%s%s\n", fileName, ".mp4 is playing"), nil
}


func (m *Mp4Player) PlayVlc(fileName string) (string, error) {
	return "", errors.New("mp4 player can not play .vlc file")
}

type VlcPlayer struct {

}

func (m *VlcPlayer) PlayMp4(fileName string) (string, error) {
	return "", errors.New("vlc player can not play .mp4 file")
}


func (m *VlcPlayer) PlayVlc(fileName string) (string, error) {
	return fmt.Sprintf("%s%s\n", fileName, ".vlc is playing"), nil
}