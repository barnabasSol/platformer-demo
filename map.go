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
	Trees     []Tree
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

type Tree struct {
	ScrollX float32
	Tex     rl.Texture2D
	Src     rl.Rectangle
	Dst     rl.Rectangle
}

func NewTree(
	src, dst rl.Rectangle,
	tex rl.Texture2D,
) Tree {
	return Tree{
		Tex: tex,
		Src: src,
		Dst: dst,
	}
}
func (m *Map) initMap() {
	m.initClouds()
	m.initTree()
}

func (m *Map) initTree() {
	treeTex := getLoadedLayerTextures(LayerTrees)
	posx, posy := 0, 140
	for range 7 {
		m.Trees = append(m.Trees, NewTree(
			rl.NewRectangle(
				0,
				0,
				float32(treeTex.Width),
				float32(treeTex.Height),
			),
			rl.NewRectangle(
				float32(posx),
				float32(posy),
				float32(treeTex.Width)/2,
				float32(treeTex.Height)/2,
			),
			treeTex,
		))
		posx += int(treeTex.Width) + 43
	}

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
		NewCloud(
			rl.NewRectangle(
				0,
				0,
				float32(cloudTex.Width),
				float32(cloudTex.Height),
			),
			rl.NewRectangle(
				40*8+float32(cloudTex.Width)/2,
				30,
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
				40*12+float32(cloudTex.Width)/2,
				30,
				float32(cloudTex.Width)/2,
				float32(cloudTex.Height)/2,
			),
			cloudTex,
		),
	} {
		m.Clouds = append(m.Clouds, c)
	}
}

func (m *Map) updateClouds(p Player, delta float32) {
	for i := range m.Clouds {
		m.Clouds[i].ScrollX += p.HorSpeed * -0.21 * delta
	}
}

func (m *Map) updateTrees(
	p Player,
	delta float32,
) {
	for _, t := range m.Trees {
		t.ScrollX += p.HorSpeed * -.4 * delta
	}
}
func (m *Map) updateMap(p Player, delta float32) {
	m.updateClouds(p, delta)
	m.updateTrees(p, delta)
}

func (m *Map) drawMap() {
	for i := range m.Clouds {
		c := &m.Clouds[i]
		rl.DrawTextureEx(
			c.Tex,
			rl.NewVector2(c.Dst.X+c.ScrollX, c.Dst.Y),
			0,
			1.3,
			rl.White,
		)
	}
	for i := range m.Trees {
		t := &m.Trees[i]
		rl.DrawTextureEx(
			t.Tex,
			rl.NewVector2(t.Dst.X+t.ScrollX, t.Dst.Y),
			0,
			1.3,
			rl.White,
		)
	}
}
