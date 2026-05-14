package main

import (
	"log"
	"os"

	"github.com/Tsukumogami-Software/luluka/shaderir"
	"github.com/goccy/go-yaml"
)

func parseYaml(valuesFile string) map[string]any {
	data, err := os.ReadFile(valuesFile)
	if err != nil {
		log.Panicf("Failed to read values file: %v", err)
	}

	var res map[string]any
	err = yaml.Unmarshal(data, &res)
	if err != nil {
		log.Panicf("Failed to read values file: %v", err)
	}
	return res
}

func parseUniformValueFromYAML(t shaderir.Type, value any) any {
	switch t.Main {
	case shaderir.Bool:
		return parseYamlBool(value)
	case shaderir.Int:
		return parseYamlInt(value)
	case shaderir.Float:
		return parseYamlFloat(value)
	case shaderir.Vec2, shaderir.Vec3, shaderir.Vec4, shaderir.Mat2, shaderir.Mat3, shaderir.Mat4:
		return parseYamlSlice(shaderir.Float, value)
	case shaderir.IVec2, shaderir.IVec3, shaderir.IVec4:
		return parseYamlSlice(shaderir.Int, value)
	case shaderir.Array:
		return parseYamlSlice(t.Sub[0].Main, value)
	}

	log.Panicf("Unknown uniform type: %v", t)
	return 0
}

func parseYamlBool(value any) bool {
	res, ok := value.(bool)
	if !ok {
		log.Printf("Failed to parse bool from yaml: %v", value)
	}
	return res
}

func parseYamlInt(value any) int32 {
	res, ok := value.(uint64)
	if !ok {
		log.Printf("Failed to parse int from yaml: %v", value)
	}
	return int32(res)
}

func parseYamlFloat(value any) float32 {
	ui64, ok := value.(uint64)
	if ok {
		return float32(ui64)
	}
	f64, ok := value.(float64)
	if !ok {
		log.Printf("Failed to parse float from yaml: %v", value)
	}
	return float32(f64)
}

func parseYamlSlice(t shaderir.BasicType, value any) any {
	slice, ok := value.([]any)
	if !ok {
		log.Panicf("Failed to parse slice from yaml: %v", value)
	}

	switch t {
	case shaderir.Bool:
		res := make([]bool, len(slice))
		for i, v := range slice {
			res[i] = parseYamlBool(v)
		}
		return res
	case shaderir.Int, shaderir.IVec2, shaderir.IVec3, shaderir.IVec4:
		res := make([]int32, len(slice))
		for i, v := range slice {
			res[i] = parseYamlInt(v)
		}
		return res
	case shaderir.Float, shaderir.Vec2, shaderir.Vec3, shaderir.Vec4, shaderir.Mat2, shaderir.Mat3, shaderir.Mat4:
		res := make([]float32, len(slice))
		for i, v := range slice {
			res[i] = parseYamlFloat(v)
		}
		return res
	case shaderir.Array:
		log.Panicf("Array of array is forbidden")
	}

	log.Panicf("Unknown array type: %v", t)
	return []any{}
}
