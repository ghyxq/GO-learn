package musicplayer

type Player interface {
	Play(source string)
}


func Play(source , mtype string)  {

	var p Player
	switch mtype {
	case "mp3":
		p = &MP3Player{}
	case "mp4":
		p = &MP4Player{}
	default:
		p = &MP3Player{}

	}

	p.Play(source)
}
