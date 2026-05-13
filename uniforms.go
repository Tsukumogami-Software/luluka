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
		return make([]float32, length * 2)
	case shaderir.Vec3:
		return make([]float32, length * 3)
	case shaderir.Vec4:
		return make([]float32, length * 4)
	case shaderir.IVec2:
		return make([]int32, length * 2)
	case shaderir.IVec3:
		return make([]int32, length * 3)
	case shaderir.IVec4:
		return make([]int32, length * 4)
	case shaderir.Mat2:
		return make([]float64, length * 4)
	case shaderir.Mat3:
		return make([]float64, length * 9)
	case shaderir.Mat4:
		return make([]float64, length * 16)
	case shaderir.Array:
		log.Panicf("Array of array is forbidden")
	}

	log.Panicf("Unknown array type: %v", t)
	return make([]any, t.Length)
}

func parseUniformValue(t shaderir.Type, values map[int]string) any {
	switch t.Main {
	case shaderir.Bool:
		res, err := strconv.ParseBool(values[0])
		if err != nil {
			log.Printf("Failed to parse bool: %s", values[0])
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
			f, err := strconv.ParseInt(values[index], 10, 32)
			if err != nil {
				log.Printf("Failed to parse float: %s", values[index])
			}
			res[index] = int32(f)
		}
		return res
	case shaderir.IVec3:
		res := make([]int32, 3)
		for index := range 3 {
			f, err := strconv.ParseInt(values[index], 10, 32)
			if err != nil {
				log.Printf("Failed to parse float: %s", values[index])
			}
			res[index] = int32(f)
		}
		return res
	case shaderir.IVec4:
		res := make([]int32, 4)
		for index := range 4 {
			f, err := strconv.ParseInt(values[index], 10, 32)
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
	case shaderir.Array:
		return parseArrayValue(t.Sub[0], t.Length, values)
	}
	return 0
}

func parseArrayValue(t shaderir.Type, length int, values map[int]string) any {
	switch t.Main {
	case shaderir.Bool:
		res := make([]bool, length)
		for i := range length {
			b, err := strconv.ParseBool(values[i])
			if err != nil {
				log.Printf("Failed to parse bool: %s", values[i])
			}
			res[i] = b
		}
		return res
	case shaderir.Int:
		res := make([]int32, length)
		for i := range length {
			d, err := strconv.ParseInt(values[i], 10, 32)
			if err != nil {
				log.Printf("Failed to parse bool: %s", values[i])
			}
			res[i] = int32(d)
		}
		return res
	case shaderir.Float:
		res := make([]float32, length)
		for i := range length {
			f, err := strconv.ParseFloat(values[i], 32)
			if err != nil {
				log.Printf("Failed to parse float: %s", values[i])
			}
			res[i] = float32(f)
		}
		return res
	case shaderir.Vec2:
		res := make([]float32, length * 2)
		for i := range length * 2 {
			f, err := strconv.ParseFloat(values[i], 32)
			if err != nil {
				log.Printf("Failed to parse float: %s", values[i])
			}
			res[i] = float32(f)
		}
		return res
	case shaderir.Vec3:
		res := make([]float32, length * 3)
		for i := range length * 3 {
			f, err := strconv.ParseFloat(values[i], 32)
			if err != nil {
				log.Printf("Failed to parse float: %s", values[i])
			}
			res[i] = float32(f)
		}
		return res
	case shaderir.Vec4:
		res := make([]float32, length * 4)
		for i := range length * 4 {
			f, err := strconv.ParseFloat(values[i], 32)
			if err != nil {
				log.Printf("Failed to parse float: %s", values[i])
			}
			res[i] = float32(f)
		}
		return res
	case shaderir.IVec2:
		res := make([]int32, length * 2)
		for i := range length * 2 {
			d, err := strconv.ParseInt(values[i], 10, 32)
			if err != nil {
				log.Printf("Failed to parse bool: %s", values[i])
			}
			res[i] = int32(d)
		}
		return res
	case shaderir.IVec3:
		res := make([]int32, length * 3)
		for i := range length * 3 {
			d, err := strconv.ParseInt(values[i], 10, 32)
			if err != nil {
				log.Printf("Failed to parse bool: %s", values[i])
			}
			res[i] = int32(d)
		}
		return res
	case shaderir.IVec4:
		res := make([]int32, length * 4)
		for i := range length * 4 {
			d, err := strconv.ParseInt(values[i], 10, 32)
			if err != nil {
				log.Printf("Failed to parse bool: %s", values[i])
			}
			res[i] = int32(d)
		}
		return res
	case shaderir.Mat2:
		res := make([]float32, length * 4)
		for i := range length * 4 {
			f, err := strconv.ParseFloat(values[i], 32)
			if err != nil {
				log.Printf("Failed to parse float: %s", values[i])
			}
			res[i] = float32(f)
		}
		return res
	case shaderir.Mat3:
		res := make([]float32, length * 9)
		for i := range length * 9 {
			f, err := strconv.ParseFloat(values[i], 32)
			if err != nil {
				log.Printf("Failed to parse float: %s", values[i])
			}
			res[i] = float32(f)
		}
		return res
	case shaderir.Mat4:
		res := make([]float32, length * 16)
		for i := range length * 16 {
			f, err := strconv.ParseFloat(values[i], 32)
			if err != nil {
				log.Printf("Failed to parse float: %s", values[i])
			}
			res[i] = float32(f)
		}
		return res
	case shaderir.Array:
		log.Panicf("Array of array is forbidden")
	}

	log.Panicf("Unknown array type: %v", t)
	return make([]any, t.Length)
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
