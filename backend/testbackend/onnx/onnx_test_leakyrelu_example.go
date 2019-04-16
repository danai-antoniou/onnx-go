package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("LeakyRelu", "TestLeakyreluExample", NewTestLeakyreluExample)
}

// NewTestLeakyreluExample version: 3.
func NewTestLeakyreluExample() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "LeakyRelu",
		Title:  "TestLeakyreluExample",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x5e, 0xa, 0x22, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x9, 0x4c, 0x65, 0x61, 0x6b, 0x79, 0x52, 0x65, 0x6c, 0x75, 0x2a, 0xf, 0xa, 0x5, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x15, 0xcd, 0xcc, 0xcc, 0x3d, 0xa0, 0x1, 0x1, 0x12, 0x16, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x65, 0x61, 0x6b, 0x79, 0x72, 0x65, 0x6c, 0x75, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5a, 0xf, 0xa, 0x1, 0x78, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x62, 0xf, 0xa, 0x1, 0x79, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x42, 0x2, 0x10, 0x9},

		/*

		   &pb.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "LeakyRelu",
		     Attributes: ([]*pb.AttributeProto) (len=1 cap=1) {
		    (*pb.AttributeProto)(0xc00042ab00)(name:"alpha" type:FLOAT f:0.1 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]float32{-1, 0, 1}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]float32{-0.1, 0, 1}),
			),
		},
	}
}
