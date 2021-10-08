//Interface to htgo package
package htgo

import (
	htgotts "github.com/hegedustibor/htgo-tts"
	"github.com/hegedustibor/htgo-tts/voices"
)

type HTGO struct{}

func New() *HTGO {
	return &HTGO{}
}

func (htgo *HTGO) CreateAudioFile(text, path string) {
	speech := htgotts.Speech{Folder: "audio", Language: voices.Indonesian}
	speech.Speak(text)
}
