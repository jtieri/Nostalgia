package mediaserver

import (
	"github.com/jtieri/Nostalgia/mediaserver/queue"
)

const MediaPath string = "C:\\Users\\justi\\go\\src\\github.com\\jtieri\\Nostalgia\\media\\"

type MediaServer struct {
	port        int
	address     string
	StreamQueue *queue.StreamQueue
}

/*
Media WebServer starts and loads three videos into queue.
The first video is started with the command below.
VLC will listen at the address and serve the video via HLS.
Once the video ends, start the next video.
Once the next video has begun adjust the queue to ensure there are still 3 videos in it.

Web WebServer starts and listens for connections.
Handles Site navigation, blog, forum, etc.

*/
// Command to
// vlc C:\Users\justi\go\src\github.com\jtieri\Nostalgia\media\now_playing\prog_index.m3u8 --sout #http{mux=ffmpeg{mux=flv},dst=:8081/stream} :no-sout-all :sout-keep
func New(port int, address string) *MediaServer {
	return &MediaServer{
		port:        port,
		address:     address,
		StreamQueue: queue.New(),
	}
}

func (ms *MediaServer) Start() (err error) {
	err = ms.StreamQueue.Load(MediaPath)
	if err != nil {
		return err
	}

	return nil
}
