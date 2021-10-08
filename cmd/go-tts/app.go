package main

import (
	"log"
	"net/http"

	ttsuc "github.com/go-tts-interface/internal/usecase/tts"
	"github.com/go-tts-interface/pkg/htgo"
	"github.com/go-tts-interface/pkg/oto"
)

type TTSService struct {
	tts *ttsuc.Usecase
}

func New() TTSService {
	htgo := htgo.New()
	oto := oto.New()
	return TTSService{
		tts: ttsuc.New(htgo, oto),
	}
}

func (s *TTSService) createHandler(w http.ResponseWriter, r *http.Request) {
	texts, ok := r.URL.Query()["text"]
	if !ok || len(texts[0]) < 1 {
		log.Println("URL param 'text' is missing")
		return
	}

	text := texts[0]
	s.tts.CreateAudioFile(text, "audio")
	w.Write([]byte(text))
}

func (s *TTSService) playHandler(w http.ResponseWriter, r *http.Request) {
	s.tts.Play("/home/tommy/go/src/github.com/go-tts-interface/cmd/go-tts/audio/test.mp3")
	w.Write([]byte("played"))
}
