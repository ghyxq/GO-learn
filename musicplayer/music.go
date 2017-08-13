package musicplayer

import "errors"

//音乐结构体基本定义
type Music struct {
	Id int
	Name string
}

type MusicManager struct {
	musics [] Music
}

func NewMusicManager() *MusicManager  {
	return &MusicManager{make([]Music,1)}
}


//add a music
func (m *MusicManager) Add(music *Music)  {

	m.musics = append(m.musics,*music)

}

//remove a music

func (m *MusicManager) Remove(index int) *Music  {

	if index <0 || index >len(m.musics){
		return nil
	}

	result := &m.musics[index]

	if index < len(m.musics) -1 {
		m.musics = append(m.musics[:index-1],m.musics[index+1:]...)
	} else if index == 0 {
       m.musics = make([]Music,0)
	} else {
		m.musics = append(m.musics[:index-1])

	}

	return result

}

//len
func (m *MusicManager) Len() int {

	return len(m.musics)

}

//get the elements

func (m *MusicManager) Get(index int)  (music *Music ,err error) {
	if index < 0 || index >= len(m.musics) {
		return nil , errors.New("index is not correct")

	}

	return &m.musics[index],nil

}


//find a element
func (m *MusicManager) Find(name string) *Music {
	if len(m.musics) ==0 {
       return nil
	}

	for _,result := range m.musics {

		if name == result.Name {

			return &result

		}
	}

	return  nil

}