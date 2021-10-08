package oto

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/hajimehoshi/oto"
	"github.com/tosone/minimp3"
)

type Oto struct {
	ctx    *oto.Context
	player *oto.Player
}

func New() *Oto {
	//move these values to config
	ctx, err := oto.NewContext(24000, 1, 2, 1024)
	if err != nil {
		// panic here.
		return nil
	}

	return &Oto{
		ctx:    ctx,
		player: ctx.NewPlayer(),
	}
}

func (oto *Oto) Play(filepath string) error {
	if _, err := os.Stat(filepath); err == nil {
		audio, err := os.Open(filepath)
		if err != nil {
			return err
		}

		defer audio.Close()

		b, err := ioutil.ReadAll(audio)
		if err != nil {
			return err
		}

		_, data, _ := minimp3.DecodeFull(b)
		// fmt.Println(dec.SampleRate)
		// fmt.Println(dec.Channels)
		_, err = oto.player.Write(data)
		return nil
	} else if os.IsNotExist(err) {
		return err
	} else {
		return errors.New("Schrodinger state: file may or may not exist..")
	}
}
