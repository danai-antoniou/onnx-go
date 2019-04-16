package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Split", "TestSplitEqualParts1d", NewTestSplitEqualParts1d)
}

// NewTestSplitEqualParts1d version: 3.
func NewTestSplitEqualParts1d() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Split",
		Title:  "TestSplitEqualParts1d",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xb3, 0x1, 0xa, 0x39, 0xa, 0x5, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x8, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x31, 0x12, 0x8, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x32, 0x12, 0x8, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x33, 0x22, 0x5, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x2a, 0xb, 0xa, 0x4, 0x61, 0x78, 0x69, 0x73, 0x18, 0x0, 0xa0, 0x1, 0x2, 0x12, 0x19, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x70, 0x6c, 0x69, 0x74, 0x5f, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x73, 0x5f, 0x31, 0x64, 0x5a, 0x13, 0xa, 0x5, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x6, 0x62, 0x16, 0xa, 0x8, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x31, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x62, 0x16, 0xa, 0x8, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x32, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x62, 0x16, 0xa, 0x8, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x33, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x42, 0x2, 0x10, 0x9},

		/*

		   &pb.NodeProto{
		     Input:     []string{"input"},
		     Output:    []string{"output_1", "output_2", "output_3"},
		     Name:      "",
		     OpType:    "Split",
		     Attributes: ([]*pb.AttributeProto) (len=1 cap=1) {
		    (*pb.AttributeProto)(0xc00013d100)(name:"axis" type:INT )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(6),
				tensor.WithBacking([]float32{1, 2, 3, 4, 5, 6}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]float32{1, 2}),
			),

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]float32{3, 4}),
			),

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]float32{5, 6}),
			),
		},
	}
}
