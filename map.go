package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	LayerSky   = "sky"
	LayerCloud = "cloud"
	LayerTrees = "tree"
)

type Map struct {
	MaxWidth  int32
	MinWidth  int32
	MaxHeight int32
	MinHeight int32
	Clouds    []Cloud
}

type Cloud struct {
	ScrollX float32
	Tex     rl.Texture2D
	Src     rl.Rectangle
	Dst     rl.Rectangle
}

func NewCloud(
	src, dst rl.Rectangle,
	tex rl.Texture2D,
) Cloud {
	return Cloud{
		Tex: tex,
		Src: src,
		Dst: dst,
	}
}

func (m *Map) initMap() {
	m.initSky()
	m.initClouds()
}

func (m *Map) initSky() {

}

func (m *Map) updateSky() {
}

func (m *Map) initClouds() {
	cloudTex := getLoadedLayerTextures(LayerCloud)
	for _, c := range []Cloud{
		NewCloud(
			rl.NewRectangle(
				0,
				0,
				float32(cloudTex.Width),
				float32(cloudTex.Height),
			),
			rl.NewRectangle(
				40,
				10,
				float32(cloudTex.Width)/2,
				float32(cloudTex.Height)/2,
			),
			cloudTex,
		),
		NewCloud(
			rl.NewRectangle(
				0,
				0,
				float32(cloudTex.Width),
				float32(cloudTex.Height),
			),
			rl.NewRectangle(
				40+float32(cloudTex.Width)/2,
				60,
				float32(cloudTex.Width)/2,
				float32(cloudTex.Height)/2,
			),
			cloudTex,
		),
		// NewLayer(
		// 	"cloud3",
		// 	rl.NewRectangle(0, 0, float32(cloudTex.Width), float32(cloudTex.Height)),
		// 	rl.NewRectangle(40*8+float32(cloudTex.Width)/2, 30, float32(cloudTex.Width)/2, float32(cloudTex.Height)/2),
		// 	cloudTex,
		// ),
	} {
		m.Clouds = append(m.Clouds, c)
	}
}

func (m *Map) updateClouds(p Player, delta float32) {
	for i := range m.Clouds {
		m.Clouds[i].ScrollX += p.HorSpeed * -0.21 * delta
	}
}

func (m *Map) updateMap(p Player, delta float32) {
	m.updateClouds(p, delta)
}

func (m *Map) drawMap() {
	for i := range m.Clouds {
		c := &m.Clouds[i]
		rl.DrawTextureEx(c.Tex, rl.NewVector2(c.Dst.X+c.ScrollX, c.Dst.Y), 0, 1.3, rl.White)
	}
}
