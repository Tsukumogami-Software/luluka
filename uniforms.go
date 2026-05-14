package main

import (
	"log"

	"github.com/Tsukumogami-Software/luluka/graphics"
	"github.com/Tsukumogami-Software/luluka/shaderir"
)

func defaultUniformValue(t shaderir.Type) any {
	switch t.Main {
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
		return make([]float64, 4)
	case shaderir.Mat3:
		return make([]float64, 9)
	case shaderir.Mat4:
		return make([]float64, 16)
	case shaderir.Array:
		return defaultArrayValue(t.Sub[0], t.Length)
	}

	log.Panicf("Unknown uniform type: %v", t)
	return 0
}

func defaultArrayValue(t shaderir.Type, length int) any {
	switch t.Main {
	case shaderir.Bool:
		return make([]bool, length)
	case shaderir.Int:
		return make([]int32, length)
	case shaderir.Float:
		return make([]float32, length)
	case shaderir.Vec2:
		return make([]float32, length*2)
	case shaderir.Vec3:
		return make([]float32, length*3)
	case shaderir.Vec4:
		return make([]float32, length*4)
	case shaderir.IVec2:
		return make([]int32, length*2)
	case shaderir.IVec3:
		return make([]int32, length*3)
	case shaderir.IVec4:
		return make([]int32, length*4)
	case shaderir.Mat2:
		return make([]float64, length*4)
	case shaderir.Mat3:
		return make([]float64, length*9)
	case shaderir.Mat4:
		return make([]float64, length*16)
	case shaderir.Array:
		log.Panicf("Array of array is forbidden")
	}

	log.Panicf("Unknown array type: %v", t)
	return make([]any, t.Length)
}

func parseUniformValues(uniformFlags []string, valuesFile string, uniformsDeclarations map[string]shaderir.Type) map[string]any {
	result := make(map[string]any, len(uniformsDeclarations))
	for name, t := range uniformsDeclarations {
		result[name] = defaultUniformValue(t)
	}

	if valuesFile != "" {
		uniformFile := parseYaml(valuesFile)
		for key, value := range uniformFile {
			result[key] = parseUniformValueFromYAML(uniformsDeclarations[key], value)
		}
	}

	uniformFlagsMap := makeUniformFlagsMap(uniformFlags)
	for key, values := range uniformFlagsMap {
		result[key] = parseFlagsUniformValue(uniformsDeclarations[key], values)
	}

	return result
}

func parseUniformDeclarations(source []byte) map[string]shaderir.Type {
	program, err := graphics.CompileShader(source)
	if err != nil {
		log.Panicf("Failed to parse shader: %v", err)
	}

	result := make(map[string]shaderir.Type, len(program.Uniforms))
	for i, uniformType := range program.Uniforms {
		result[program.UniformNames[i]] = uniformType
	}
	return result
}
