package main

import (
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 512
	screenHeight = 512
)

var game *Game

// Game contains the compiled shader and keeps track of the time
type Game struct {
	images    [4]*ebiten.Image
	shader    *ebiten.Shader
	uniforms  map[string]any
	startTime time.Time
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
	g.uniforms["Time"] = time.Since(g.startTime).Seconds()

	cx, cy := ebiten.CursorPosition()
	cx64 := min(max(float32(cx), 0.0), float32(g.width()))
	cy64 := min(max(float32(cy), 0.0), float32(g.height()))
	g.uniforms["Cursor"] = []float32{cx64 / float32(g.width()), cy64 / float32(g.height())}

	mouseButtons := 0b00
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mouseButtons += 0b10
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		mouseButtons += 0b01
	}
	g.uniforms["MouseButtons"] = mouseButtons

	screen.DrawRectShader(
		g.width(),
		g.height(),
		g.shader,
		&ebiten.DrawRectShaderOptions{
			Uniforms: g.uniforms,
			Images:   g.images,
		},
	)
}

// Update does nothing here
func (g *Game) Update() error {
	return nil
}

// Layout returns a fixed width/height
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.width(), g.height()
}

func Run(shaderPath string, uniformFlags []string, imageFlags []string, valuesFile string) {
	shaderFile, err := os.ReadFile(shaderPath)
	if err != nil {
		log.Panicf("Failed to read shader file: %v", err)
	}

	uniformsDeclarations := parseUniformDeclarations(shaderFile)
	uniforms := parseUniformValues(uniformFlags, valuesFile, uniformsDeclarations)

	images := [4]*ebiten.Image{}
	for i, imageFlag := range imageFlags {
		images[i] = loadImage(imageFlag)
	}

	shader, err := ebiten.NewShader(shaderFile)
	if err != nil {
		log.Panicf("Failed to create shader: %v", err)
	}

	game = &Game{
		shader:    shader,
		uniforms:  uniforms,
		startTime: time.Now(),
		images:    images,
	}

	ebiten.SetWindowSize(game.width(), game.height())
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
