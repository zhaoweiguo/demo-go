package golang_socketio

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
	"sync"
)

const (
	SCREEN_WIDTH  = 1280
	SCREEN_HEIGHT = 720
)

var (
	quit     = make(chan bool)
	sdlMutex sync.Mutex
)

type VideoPlayer struct {
	Screen     *sdl.Window
	Renderer   *sdl.Renderer
	Texture    *sdl.Texture
	FrameWidth int
	FrameHeight int
}

func NewVideoPlayer(width, height int) (*VideoPlayer, error) {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		return nil, errors.Wrap(err, "初始化 SDL 失败")
	}

	screen, err := sdl.CreateWindow("Video Player", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, SCREEN_WIDTH, SCREEN_HEIGHT, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, errors.Wrap(err, "创建窗口失败")
	}

	renderer, err := sdl.CreateRenderer(screen, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return nil, errors.Wrap(err, "创建渲染器失败")
	}

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_IYUV, sdl.TEXTUREACCESS_STREAMING, width, height)
	if err != nil {
		return nil, errors.Wrap(err, "创建纹理失败")
	}

	return &VideoPlayer{
		Screen:     screen,
		Renderer:   renderer,
		Texture:    texture,
		FrameWidth: width,
		FrameHeight: height,
	}, nil
}

func (p *VideoPlayer) Play(data []byte) error {
	sdlMutex.Lock()
	defer sdlMutex.Unlock()

	if len(data) != p.FrameWidth * p.FrameHeight * 3 / 2 {
		return fmt.Errorf("数据长度不正确")
	}

	if err := p.Texture.UpdateYUV(nil, data[:p.FrameWidth], data[p.FrameWidth:p.FrameWidth*5/4], data[p.FrameWidth*5/4:], p.FrameWidth); err != nil {
		return errors.Wrap(err, "更新纹理失败")
	}

	p.Renderer.Clear()
	p.Renderer.Copy(p.Texture, nil, nil)
	p.Renderer.Present()

	return nil
}

func (p *VideoPlayer) Destroy() {
	sdlMutex.Lock()
	defer sdlMutex.Unlock()

	p.Texture.Destroy()
	p.Renderer.Destroy()
	p.Screen.Destroy()

	sdl.Quit()
}

type VideoDecoder struct {
	CodecContext *ffmpeg.CodecContext
	Frame        *ffmpeg.Frame
}

func NewVideoDecoder(codecpar *ffmpeg.CodecParameters) (*VideoDecoder, error) {
	codec, err := ffmpeg.FindDecoder(codecpar.CodecID())
	if err != nil {
		return nil, errors
