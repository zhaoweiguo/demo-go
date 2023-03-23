package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

/**
把解码后的 YUV 数据渲染到屏幕上，可以使用像 SDL、OpenGL 等图形库来实现。以下是一个基于 SDL 的简单例子：

*/

func main() {
	// 初始化 SDL
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		fmt.Println("SDL 初始化失败：", err)
		return
	}
	defer sdl.Quit()

	// 创建窗口
	window, err := sdl.CreateWindow("Video Player", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 640, 480, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println("创建窗口失败：", err)
		return
	}
	defer window.Destroy()

	// 创建渲染器
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("创建渲染器失败：", err)
		return
	}
	defer renderer.Destroy()

	// 创建纹理
	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_IYUV, sdl.TEXTUREACCESS_STREAMING, 640, 480)
	if err != nil {
		fmt.Println("创建纹理失败：", err)
		return
	}
	defer texture.Destroy()

	// 模拟从网络接收到的数据
	yuvData := make([]byte, 640*480*3/2)
	for i := 0; i < len(yuvData); i++ {
		yuvData[i] = byte(i % 256)
	}

	for {
		// 更新纹理
		texture.UpdateYUV(nil, yuvData[:640], yuvData[640:640*3/2], yuvData[640*3/2:], 640)

		// 渲染
		renderer.Clear()
		renderer.Copy(texture, nil, nil)
		renderer.Present()

		// 处理事件
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
	}
}
