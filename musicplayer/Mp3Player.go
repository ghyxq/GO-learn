package musicplayer

import (
	"fmt"
	"time"
)

type MP3Player struct {
	stat int
	process int
}

func (m *MP3Player) Play(source string)  {

	fmt.Println("Play the mp3 music",source)

	m.process = 0

	for m.process <100 {
		time.Sleep(100 *time.Millisecond)
		fmt.Print(".")
		m.process +=10
	}

	fmt.Println("over3 playing ....",source)
}



type MP4Player struct {
	stat int
	process int
}

func (m *MP4Player) Play(source string)  {

	fmt.Println("Play the mp4 music",source)

	m.process = 0

	for m.process <100 {
		time.Sleep(100 *time.Millisecond)
		fmt.Print(".")
		m.process +=10
	}

	fmt.Println("over4 playing ....",source)
}