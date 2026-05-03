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
	shader    *ebiten.Shader
	uniforms  map[string]any
	startTime time.Time
}

// Draw displays the shader on the entire screen
func (g *Game) Draw(screen *ebiten.Image) {
	g.uniforms["Time"] = time.Since(g.startTime).Seconds()
	screen.DrawRectShader(
		screenWidth,
		screenHeight,
		g.shader,
		&ebiten.DrawRectShaderOptions{
			Uniforms: g.uniforms,
		},
	)
}

// Update does nothing here
func (g *Game) Update() error {
	return nil
}

// Layout returns a fixed width/height
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func Run(shaderPath string, uniformFlags []string) {
	shaderFile, err := os.ReadFile(shaderPath)
	if err != nil {
		log.Panicf("Failed to read shader file: %v", err)
	}

	uniformsDeclarations := parseUniformDeclarations(shaderFile)
	uniforms := parseUniformValues(uniformFlags, uniformsDeclarations)

	shader, err := ebiten.NewShader(shaderFile)
	if err != nil {
		log.Panicf("Failed to create shader: %v", err)
	}

	game = &Game{
		shader:    shader,
		uniforms:  uniforms,
		startTime: time.Now(),
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
