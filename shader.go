package main

import (
	"log"
	"os"
	"time"

	"github.com/Tsukumogami-Software/Luluka/shader"
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

func ParseUniforms(source []byte) {
	program, err := shader.Compile(source, "__vertex", "Fragment", 4)
	if err != nil {
		log.Panicf("Failed to parse shader: %v", err)
	}

	for i, uniformType := range program.Uniforms {
		log.Printf("Found uniform %s: %v\n", program.UniformNames[i], uniformType)
	}
}

func Run(shaderPath string) {
	shaderFile, err := os.ReadFile(shaderPath)
	if err != nil {
		log.Panicf("Failed to read shader file: %v", err)
	}

	ParseUniforms(shaderFile)

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
