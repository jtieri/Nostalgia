package queue

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

/*VLC
vlc C:\Users\justi\go\src\github.com\jtieri\Nostalgia\media\now_playing\prog_index.m3u8 --sout #http{mux=ffmpeg{mux=flv},dst=:8081/stream} :no-sout-all :sout-keep

FFMPEG
ffmpeg -y -i C:\Users\justi\go\src\github.com\jtieri\Nostalgia\media\SLV.mkv -hls_time 4 -hls_base_url C:\Users\justi\go\src\github.com\jtieri\Nostalgia\media\now_playing\  -hls_segment_filename "C:\Users\justi\go\src\github.com\jtieri\Nostalgia\media\now_playing\fileSequence%d.ts" C:\Users\justi\go\src\github.com\jtieri\Nostalgia\media\now_playing\prog_index.m3u8
*/
type StreamQueue struct {
	FirstFileName  string
	SecondFileName string
	ThirdFileName  string
}

func New() *StreamQueue {
	return &StreamQueue{
		FirstFileName:  "",
		SecondFileName: "",
		ThirdFileName:  "",
	}
}

func (sq *StreamQueue) Load(mediaDir string) (err error) {
	dirContents, err := ioutil.ReadDir(mediaDir)
	if err != nil {
		return err
	}

	sq.FirstFileName = mediaDir + dirContents[0].Name()
	sq.SecondFileName = mediaDir + dirContents[1].Name()
	sq.ThirdFileName = mediaDir + dirContents[2].Name()

	return nil
}

func (sq *StreamQueue) Start(address string, port int) (err error) {
	fmt.Println(sq.FirstFileName)
	err = playVideoVLC("C:\\Users\\justi\\go\\src\\github.com\\jtieri\\Nostalgia\\media\\now_playing\\prog_index.m3u8", address, port)
	if err != nil {
		return err
	}

	return nil
}

func segmentVideoFFMPEG(fileLocation string) {
	/*
		ffmpeg -y \
		    -i sample.mov \
		    -hls_time 9 \
		    -hls_segment_filename "/Library/WebServer/Documents/vod/fileSequence%d.ts" \
		    -hls_playlist_type vod \
		    /Library/WebServer/Documents/vod/prog_index.m3u8
	*/
}

func playVideoVLC(filename string, address string, port int) (err error) {
	command := fmt.Sprintf("vlc %s --sout #http{mux=ffmpeg{mux=flv},dst=%s:%d/stream} :no-sout-all :sout-keep",
		filename,
		address,
		port)
	fmt.Println(command)

	err = exec.Command("C:\\Program Files\\VideoLAN\\VLC\\vlc.exe", command).Start()
	if err != nil {
		return err
	}

	return nil
}
