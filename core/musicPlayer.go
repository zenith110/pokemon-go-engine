package core

import (
	"bytes"
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
)

type musicType int
type MusicPlayer struct {
	game               *Game
	audioContext       *audio.Context
	audioPlayer        *audio.Player
	current            time.Duration
	total              time.Duration
	soundEffectBytes   []byte
	soundEffectChannel chan []byte
	volume128          int
	musicType          musicType
	beginning          time.Duration
	end                time.Duration
	season             string
	mapLocation        string
}

func LoadMusicTrack(trackName string, trackExtention string) {
	switch trackExtention {
	case "ogg":
		var err error
		stream, err = vorbis.DecodeWithoutResampling(bytes.NewReader([]byte(trackName)))
		if err != nil {
			fmt.Printf("Error is %v", err)
		}
	}
	p, err := audiocontext.NewPlayer(stream)
}
