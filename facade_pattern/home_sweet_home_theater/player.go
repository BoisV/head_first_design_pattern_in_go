package home_sweet_home_theater

import "fmt"

type IPlayer interface {
	pause()
	play(movie string)
	setSurroundAudio()
	setTwoChannelAudio()
	stop()
	on()
	off()
}

type StreamingPlayer struct {
	Amplifier IAmplifier
	movie     string
}

func NewStreamingPlayer(amplifier IAmplifier) *StreamingPlayer {
	return &StreamingPlayer{Amplifier: amplifier, movie: ""}
}

func (s *StreamingPlayer) on() {
	fmt.Println("Streaming Player on")
}

func (s *StreamingPlayer) off() {
	fmt.Println("Streaming Player off")
}

func (s *StreamingPlayer) play(movie string) {
	s.movie = movie
	fmt.Printf("Streaming Player playing \"%s\"\n", s.movie)
}

func (s *StreamingPlayer) pause() {
	fmt.Printf("Streaming Player paused \"%s\"\n", s.movie)
}

func (s *StreamingPlayer) stop() {
	fmt.Printf("Streaming Player stooped \"%s\"\n", s.movie)
}

func (s *StreamingPlayer) setSurroundAudio() {
	//TODO implement me
	panic("implement me")
}

func (s *StreamingPlayer) setTwoChannelAudio() {
	//TODO implement me
	panic("implement me")
}
