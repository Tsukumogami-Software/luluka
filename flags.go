package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/Tsukumogami-Software/luluka/shaderir"
)

func parseFlagsUniformValue(t shaderir.Type, values map[int]string) any {
	switch t.Main {
	case shaderir.Bool:
		return parseFlagsBool(values[0])
	case shaderir.Int:
		return parseFlagsInt(values[0])
	case shaderir.Float:
		return parseFlagsFloat(values[0])
	case shaderir.Vec2:
		res := make([]float32, 2)
		for i := range 2 {
			res[i] = parseFlagsFloat(values[i])
		}
		return res
	case shaderir.Vec3:
		res := make([]float32, 3)
		for i := range 3 {
			res[i] = parseFlagsFloat(values[i])
		}
		return res
	case shaderir.Vec4:
		res := make([]float32, 4)
		for i := range 4 {
			res[i] = parseFlagsFloat(values[i])
		}
		return res
	case shaderir.IVec2:
		res := make([]int32, 2)
		for i := range 2 {
			res[i] = parseFlagsInt(values[i])
		}
		return res
	case shaderir.IVec3:
		res := make([]int32, 3)
		for i := range 3 {
			res[i] = parseFlagsInt(values[i])
		}
		return res
	case shaderir.IVec4:
		res := make([]int32, 4)
		for i := range 4 {
			res[i] = parseFlagsInt(values[i])
		}
		return res
	case shaderir.Mat2:
		res := make([]float32, 4)
		for i := range 4 {
			res[i] = parseFlagsFloat(values[i])
		}
		return res
	case shaderir.Mat3:
		res := make([]float32, 9)
		for i := range 9 {
			res[i] = parseFlagsFloat(values[i])
		}
		return res
	case shaderir.Mat4:
		res := make([]float32, 16)
		for i := range 16 {
			res[i] = parseFlagsFloat(values[i])
		}
		return res
	case shaderir.Array:
		return parseFlagsArrayValue(t.Sub[0], t.Length, values)
	}
	return 0
}

func parseFlagsBool(value string) bool {
	b, err := strconv.ParseBool(value)
	if err != nil {
		log.Printf("Failed to parse bool: %s", value)
	}
	return b
}

func parseFlagsInt(value string) int32 {
	i, err := strconv.ParseInt(value, 10, 32)
	if err != nil {
		log.Printf("Failed to parse bool: %s", value)
	}
	return int32(i)
}

func parseFlagsFloat(value string) float32 {
	f, err := strconv.ParseFloat(value, 32)
	if err != nil {
		log.Printf("Failed to parse float: %s", value)
	}
	return float32(f)
}

func parseFlagsArrayValue(t shaderir.Type, length int, values map[int]string) any {
	switch t.Main {
	case shaderir.Bool:
		res := make([]bool, length)
		for i := range length {
			res[i] = parseFlagsBool(values[i])
		}
		return res
	case shaderir.Int:
		res := make([]int32, length)
		for i := range length {
			res[i] = parseFlagsInt(values[i])
		}
		return res
	case shaderir.Float:
		res := make([]float32, length)
		for i := range length {
			res[i] = parseFlagsFloat(values[i])
		}
		return res
	case shaderir.Vec2:
		res := make([]float32, length*2)
		for i := range length * 2 {
			res[i] = parseFlagsFloat(values[i])
		}
		return res
	case shaderir.Vec3:
		res := make([]float32, length*3)
		for i := range length * 3 {
			res[i] = parseFlagsFloat(values[i])
		}
		return res
	case shaderir.Vec4:
		res := make([]float32, length*4)
		for i := range length * 4 {
			res[i] = parseFlagsFloat(values[i])
		}
		return res
	case shaderir.IVec2:
		res := make([]int32, length*2)
		for i := range length * 2 {
			res[i] = parseFlagsInt(values[i])
		}
		return res
	case shaderir.IVec3:
		res := make([]int32, length*3)
		for i := range length * 3 {
			res[i] = parseFlagsInt(values[i])
		}
		return res
	case shaderir.IVec4:
		res := make([]int32, length*4)
		for i := range length * 4 {
			res[i] = parseFlagsInt(values[i])
		}
		return res
	case shaderir.Mat2:
		res := make([]float32, length*4)
		for i := range length * 4 {
			res[i] = parseFlagsFloat(values[i])
		}
		return res
	case shaderir.Mat3:
		res := make([]float32, length*9)
		for i := range length * 9 {
			res[i] = parseFlagsFloat(values[i])
		}
		return res
	case shaderir.Mat4:
		res := make([]float32, length*16)
		for i := range length * 16 {
			res[i] = parseFlagsFloat(values[i])
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
