package main

import (
	"log"
	"strings"

	"github.com/Tsukumogami-Software/luluka/shader"
	"github.com/Tsukumogami-Software/luluka/shaderir"
)

// TODO: handle Sub types
// TODO: handle Texture, Array, Struct
// TODO: double check, matrices may need to be flattened

func defaultUniformValue(t shaderir.Type) any {
	switch t.Main{
	case shaderir.Bool:
		return false
	case shaderir.Int:
		return 0
	case shaderir.Float:
		return 0.0
	case shaderir.Vec2:
		return make([]float32, 2)
	case shaderir.Vec3:
		return make([]float32, 3)
	case shaderir.Vec4:
		return make([]float32, 4)
	case shaderir.IVec2:
		return make([]int32, 2)
	case shaderir.IVec3:
		return make([]int32, 3)
	case shaderir.IVec4:
		return make([]int32, 4)
	case shaderir.Mat2:
		result := make([][]float64, 2)
		for i := range 2 {
			result[i] = make([]float64, 2)
		}
		return result
	case shaderir.Mat3:
		result := make([][]float64, 2)
		for i := range 2 {
			result[i] = make([]float64, 2)
		}
		return result
	case shaderir.Mat4:
		result := make([][]float64, 2)
		for i := range 2 {
			result[i] = make([]float64, 2)
		}
		return result
	}

	log.Panicf("Unknown uniform type: %v", t)
	return 0
}

func parseUniformValues(uniformFlags []string, uniformsDeclarations map[string]shaderir.Type) map[string]any {
	result := make(map[string]any, len(uniformsDeclarations))
	for name, t := range uniformsDeclarations {
		result[name] = defaultUniformValue(t)
	}

	for _, flag := range uniformFlags {
		split := strings.Split(flag, ":")
		if len(split) != 2 {
			log.Panicf("Invalid uniform flag: %s", flag)
		}

		//TODO: check that this matches the uniforms declarations
		existing, ok := result[split[0]]
		if ok {
			existing = append(existing.([]any), split[1])
		} else {
			result[split[0]] = split[1]
		}
	}
	return result
}

func parseUniformDeclarations(source []byte) map[string]shaderir.Type {
	program, err := shader.Compile(source, "__vertex", "Fragment", 4)
	if err != nil {
		log.Panicf("Failed to parse shader: %v", err)
	}

	result := make(map[string]shaderir.Type, len(program.Uniforms))
	for i, uniformType := range program.Uniforms {
		result[program.UniformNames[i]] = uniformType
	}
	return result
}
