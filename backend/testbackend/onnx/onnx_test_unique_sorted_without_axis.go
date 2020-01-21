package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Unique", "TestUniqueSortedWithoutAxis", NewTestUniqueSortedWithoutAxis)
}

// NewTestUniqueSortedWithoutAxis version: 5.
func NewTestUniqueSortedWithoutAxis() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Unique",
		Title:  "TestUniqueSortedWithoutAxis",
		ModelB: []byte{0x8, 0x5, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xc1, 0x1, 0xa, 0x30, 0xa, 0x1, 0x58, 0x12, 0x1, 0x59, 0x12, 0x7, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0xf, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0x6, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x6, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x12, 0x1f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x73, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x6f, 0x75, 0x74, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x5a, 0xf, 0xa, 0x1, 0x58, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x6, 0x62, 0xf, 0xa, 0x1, 0x59, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x4, 0x62, 0x15, 0xa, 0x7, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x4, 0x62, 0x1d, 0xa, 0xf, 0x69, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x6, 0x62, 0x14, 0xa, 0x6, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x4, 0x42, 0x2, 0x10, 0xb},

		/*

		   &ir.NodeProto{
		     Input:     []string{"X"},
		     Output:    []string{"Y", "indices", "inverse_indices", "counts"},
		     Name:      "",
		     OpType:    "Unique",
		     Attributes: ([]*ir.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(6),
				tensor.WithBacking([]float32{2, 1, 1, 3, 4, 3}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(4),
				tensor.WithBacking([]float32{1, 2, 3, 4}),
			),

			tensor.New(
				tensor.WithShape(4),
				tensor.WithBacking([]int64{1, 0, 3, 4}),
			),

			tensor.New(
				tensor.WithShape(6),
				tensor.WithBacking([]int64{1, 0, 0, 2, 3, 2}),
			),

			tensor.New(
				tensor.WithShape(4),
				tensor.WithBacking([]int64{2, 1, 2, 1}),
			),
		},
	}
}