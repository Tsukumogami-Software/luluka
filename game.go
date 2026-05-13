package main

import (
	"log"
	"os"
	"time"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

const (
	screenWidth  = 1024
	screenHeight = 512
)

var game *Game

// Game contains the compiled shader and keeps track of the time
type Game struct {
	images    [4]*ebiten.Image
	shader    *ebiten.Shader
	uniforms  map[string]any
	startTime time.Time
	ui        *ebitenui.UI
}

func (g *Game) width() int {
	if g.images[0] == nil {
		return screenWidth
	}
	return g.images[0].Bounds().Dx()
}

func (g *Game) height() int {
	if g.images[0] == nil {
		return screenHeight
	}
	return g.images[0].Bounds().Dy()
}

// Draw displays the shader on the entire screen
func (g *Game) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)

	geom := ebiten.GeoM{}
	geom.Translate(512, 0)
	g.uniforms["Time"] = time.Since(g.startTime).Seconds()
	screen.DrawRectShader(
		g.width(),
		g.height(),
		g.shader,
		&ebiten.DrawRectShaderOptions{
			Uniforms: g.uniforms,
			Images:   g.images,
			GeoM:     geom,
		},
	)
}

// Update does nothing here
func (g *Game) Update() error {
	g.ui.Update()
	return nil
}

// Layout returns a fixed width/height
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func Run(shaderPath string, uniformFlags []string, imageFlags []string) {
	shaderFile, err := os.ReadFile(shaderPath)
	if err != nil {
		log.Panicf("Failed to read shader file: %v", err)
	}

	uniformsDeclarations := parseUniformDeclarations(shaderFile)
	uniforms := parseUniformValues(uniformFlags, uniformsDeclarations)

	images := [4]*ebiten.Image{}
	for i, imageFlag := range imageFlags {
		images[i] = loadImage(imageFlag)
	}

	shader, err := ebiten.NewShader(shaderFile)
	if err != nil {
		log.Panicf("Failed to create shader: %v", err)
	}

	root := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(
			image.NewNineSliceColor(colornames.Darkslategray),
		),
	)
	game = &Game{
		shader:    shader,
		uniforms:  uniforms,
		startTime: time.Now(),
		images:    images,
		ui:        &ebitenui.UI{Container: root},
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
