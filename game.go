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
	startTime time.Time
}

// Draw displays the shader on the entire screen
func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawRectShader(
		screenWidth,
		screenHeight,
		g.shader,
		&ebiten.DrawRectShaderOptions{
			Uniforms: map[string]any{
				"Center": []float32{
					float32(screenWidth) / 2,
					float32(screenHeight) / 2,
				},
				"Time": time.Now().Sub(g.startTime).Seconds(),
			},
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
	parseUniformValues(uniformFlags, uniformsDeclarations)

	shader, err := ebiten.NewShader(shaderFile)
	if err != nil {
		log.Panicf("Failed to create shader: %v", err)
	}

	game = &Game{
		shader:    shader,
		startTime: time.Now(),
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
