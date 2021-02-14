package main

import "fmt"

type song struct {
	name   string
	artist string
	next   *song
}

type playlist struct {
	name       string
	head       *song
	nowPlaying *song
}

func createPlaylist(name string) *playlist {
	return &playlist{
		name: name,
	}
}

func (p *playlist) addSong(name, artist string) {
	s := &song{
		name:   name,
		artist: artist,
	}

	if p.head == nil {
		p.head = s
	} else {
		tail := p.head
		for tail.next != nil {
			tail = tail.next
		}
		tail.next = s
	}
}

func (p *playlist) showAllSogs() error {
	curSong := p.head
	if curSong == nil {
		fmt.Println("Playlist is empty")
		return nil
	}

	fmt.Printf("%+v\n", curSong)
	for curSong.next != nil {
		curSong = curSong.next
		fmt.Printf("%+v\n", curSong)
	}

	return nil
}

func (p *playlist) startPlaying() *song {
	p.nowPlaying = p.head
	return p.nowPlaying
}

func (p *playlist) playNext() *song {
	p.nowPlaying = p.nowPlaying.next
	return p.nowPlaying
}

func (p *playlist) showCurrentSong() {
	fmt.Printf("Playing %s by %s\n", p.nowPlaying.name, p.nowPlaying.artist)
}

func main() {
	p := createPlaylist("playlist")

	fmt.Printf("Adding songs to the playlist\n\n")
	p.addSong("Beat It", "Michael Jackson")
	p.addSong("Smooth Criminal", "Michael Jackson")
	p.addSong("Roar", "Katy Perry")
	p.addSong("Firework", "Katy Perry")
	p.showAllSogs()

	fmt.Println("\n<<< start playing >>>")
	p.startPlaying()
	p.showCurrentSong()

	fmt.Println("\nPlay next >>>")
	p.playNext()
	p.showCurrentSong()

	fmt.Println("\nPlay next >>>")
	p.playNext()
	p.showCurrentSong()
}
