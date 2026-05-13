// Copyright 2019 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package graphics

const (
	ShaderSrcImageCount = 4

	// PreservedUniformVariablesCount represents the number of preserved uniform variables.
	// Any shaders in Ebitengine must have these uniform variables.
	PreservedUniformVariablesCount = 1 + // the destination texture size
		1 + // the source texture sizes array
		1 + // the destination image region origin
		1 + // the destination image region size
		1 + // the source image region origins
		1 + // the source image region sizes array
		1 // the projection matrix

	ProjectionMatrixUniformVariableIndex = 6

	PreservedUniformDwordCount = 2 + // the destination texture size
		2*ShaderSrcImageCount + // the source texture sizes array
		2 + // the destination image region origin
		2 + // the destination image region size
		2*ShaderSrcImageCount + // the source image region origins array
		2*ShaderSrcImageCount + // the source image region sizes array
		16 // the projection matrix

	ProjectionMatrixUniformDwordIndex = 2 +
		2*ShaderSrcImageCount +
		2 +
		2 +
		2*ShaderSrcImageCount +
		2*ShaderSrcImageCount
)

const (
	VertexFloatCount = 12
)

var (
	quadIndices = []uint32{0, 1, 2, 1, 2, 3}
)
