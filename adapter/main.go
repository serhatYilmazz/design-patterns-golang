package main

import (
	"fmt"
	media_player "github.com/serhatYilmazz/design-patterns-golang/adapter/media-player"
	"log"
)

func main() {

	ap := media_player.AudioPlayer{
		Adapter: &media_player.AdvancedMediaPlayerAdapter{
			Amp: &media_player.Mp4Player{},
		},
	}

	play, err := ap.Play(media_player.VLC, "ghost")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(play)

}
