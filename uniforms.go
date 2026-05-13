package main

import (
	"log"
	"strconv"
	"strings"

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
		result := make([]any, t.Length)
		for i := range t.Length {
			result[i] = defaultUniformValue(t.Sub[0])
		}
		return result
	}

	log.Panicf("Unknown uniform type: %v", t)
	return 0
}

func parseUniformValue(t shaderir.Type, values map[int]string) any {
	// TODO: parse Array
	switch t.Main {
	case shaderir.Bool:
		res, err := strconv.ParseBool(values[0])
		if err != nil {
			log.Panicf("Failed to parse bool: %s", values[0])
		}
		return res
	case shaderir.Int:
		res, err := strconv.ParseInt(values[0], 10, 32)
		if err != nil {
			log.Printf("Failed to parse int: %s", values[0])
		}
		return res
	case shaderir.Float:
		res, err := strconv.ParseFloat(values[0], 32)
		if err != nil {
			log.Printf("Failed to parse float: %s", values[0])
		}
		return res
	case shaderir.Vec2:
		res := make([]float32, 2)
		for index := range 2 {
			f, err := strconv.ParseFloat(values[index], 32)
			if err != nil {
				log.Printf("Failed to parse float: %s", values[index])
			}
			res[index] = float32(f)
		}
		return res
	case shaderir.Vec3:
		res := make([]float32, 3)
		for index := range 3 {
			f, err := strconv.ParseFloat(values[index], 32)
			if err != nil {
				log.Printf("Failed to parse float: %s", values[index])
			}
			res[index] = float32(f)
		}
		return res
	case shaderir.Vec4:
		res := make([]float32, 4)
		for index := range 4 {
			f, err := strconv.ParseFloat(values[index], 32)
			if err != nil {
				log.Printf("Failed to parse float: %s", values[index])
			}
			res[index] = float32(f)
		}
		return res
	case shaderir.IVec2:
		res := make([]int32, 2)
		for index := range 2 {
			f, err := strconv.ParseFloat(values[index], 32)
			if err != nil {
				log.Printf("Failed to parse float: %s", values[index])
			}
			res[index] = int32(f)
		}
		return res
	case shaderir.IVec3:
		res := make([]int32, 3)
		for index := range 3 {
			f, err := strconv.ParseFloat(values[index], 32)
			if err != nil {
				log.Printf("Failed to parse float: %s", values[index])
			}
			res[index] = int32(f)
		}
		return res
	case shaderir.IVec4:
		res := make([]int32, 4)
		for index := range 4 {
			f, err := strconv.ParseFloat(values[index], 32)
			if err != nil {
				log.Printf("Failed to parse float: %s", values[index])
			}
			res[index] = int32(f)
		}
		return res
	case shaderir.Mat2:
		res := make([]float32, 4)
		for index := range 4 {
			f, err := strconv.ParseFloat(values[index], 32)
			if err != nil {
				log.Printf("Failed to parse float: %s", values[index])
			}
			res[index] = float32(f)
		}
		return res
	case shaderir.Mat3:
		res := make([]float32, 9)
		for index := range 9 {
			f, err := strconv.ParseFloat(values[index], 32)
			if err != nil {
				log.Printf("Failed to parse float: %s", values[index])
			}
			res[index] = float32(f)
		}
		return res
	case shaderir.Mat4:
		res := make([]float32, 16)
		for index := range 16 {
			f, err := strconv.ParseFloat(values[index], 32)
			if err != nil {
				log.Printf("Failed to parse float: %s", values[index])
			}
			res[index] = float32(f)
		}
		return res
	}
	return 0
}

func makeUniformFlagsMap(uniformFlags []string) map[string]map[int]string {
	parsedFlags := make(map[string]map[int]string, len(uniformFlags))
	for _, flag := range uniformFlags {
		split := strings.Split(flag, ":")
		if len(split) != 2 {
			log.Panicf("Invalid uniform flag: %s", flag)
		}

		splitName := strings.Split(split[0], ".")
		if len(splitName) == 2 {
			//vector
			vecIndex, err := strconv.Atoi(splitName[1])
			if err != nil {
				log.Panicf("Failed to parse vector index: %s", flag)
			}
			_, exists := parsedFlags[splitName[0]]
			if exists {
				parsedFlags[splitName[0]][vecIndex] = split[1]
			} else {
				parsedFlags[splitName[0]] = map[int]string{
					vecIndex: split[1],
				}
			}
		} else {
			//scalar
			parsedFlags[split[0]] = map[int]string{
				0: split[1],
			}
		}
	}
	return parsedFlags
}

func parseUniformValues(uniformFlags []string, uniformsDeclarations map[string]shaderir.Type) map[string]any {
	result := make(map[string]any, len(uniformsDeclarations))
	for name, t := range uniformsDeclarations {
		result[name] = defaultUniformValue(t)
	}

	uniformFlagsMap := makeUniformFlagsMap(uniformFlags)
	for key, values := range uniformFlagsMap {
		result[key] = parseUniformValue(uniformsDeclarations[key], values)
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
