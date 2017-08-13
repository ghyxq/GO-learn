package main

import (
	"musicplayer"
	"fmt"
	"strconv"
	"bufio"
	"os"
	"strings"
)

var musiclib *musicplayer.MusicManager

func handLibCommand(token [] string)  {
	switch token[1] {
	case "list":
		for i := 0 ; i < musiclib.Len() ; i++ {
			e,_ := musiclib.Get(i)
			fmt.Println("i is:",i," and the music is:",e.Id," ",e.Name)
		}
	case "remove":
		if len(token) ==3 {
			b,_ := strconv.Atoi(token[2])
            musiclib.Remove(b)
		}
	case "add":
		music := musicplayer.Music{0,token[2]}
		musiclib.Add(&music)

	default:
		fmt.Println("default hand lib...")

	}
	
}


func handplayCommand(token [] string)  {
	if len(token) != 3 {
		fmt.Println("play commond error!")
		return
	}

	music := musiclib.Find(token[1])

	musicplayer.Play(music.Name,token[2])
	

}

func main()  {

	fmt.Println("start the music ......")
	musiclib = musicplayer.NewMusicManager();
	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter the Command:")

		rawLine,_ ,_:= r.ReadLine()

		line := string(rawLine)

		if line == "q" || line =="e" {
			break
			
		}

		token := strings.Split(line," ")

		if token[0] == "lib" {
			handLibCommand(token)
		} else {
			handplayCommand(token)
		}

	}

}