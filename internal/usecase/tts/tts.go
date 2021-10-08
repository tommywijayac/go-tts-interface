package tts

type TTSEngine interface {
	CreateAudioFile(text, path string)
}

type TTSPlayer interface {
	Play(filepath string) error
}

type Usecase struct {
	engine TTSEngine
	player TTSPlayer
}

func New(engine TTSEngine, player TTSPlayer) *Usecase {
	return &Usecase{
		engine: engine,
		player: player,
	}
}

func (uc *Usecase) CreateAudioFile(text, path string) {
	uc.engine.CreateAudioFile(text, path)
}

func (uc *Usecase) Play(filepath string) {
	uc.player.Play(filepath)
}
