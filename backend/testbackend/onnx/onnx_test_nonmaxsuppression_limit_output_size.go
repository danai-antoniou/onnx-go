package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("NonMaxSuppression", "TestNonmaxsuppressionLimitOutputSize", NewTestNonmaxsuppressionLimitOutputSize)
}

/*
&ir.ModelProto{
    IrVersion:   5,
    OpsetImport: {
        &ir.OperatorSetIdProto{Domain:"", Version:10},
    },
    ProducerName:    "backend-test",
    ProducerVersion: "",
    Domain:          "",
    ModelVersion:    0,
    DocString:       "",
    Graph:           &ir.GraphProto{
        Node: {
            &ir.NodeProto{
                Input:     {"boxes", "scores", "max_output_boxes_per_class", "iou_threshold", "score_threshold"},
                Output:    {"selected_indices"},
                Name:      "",
                OpType:    "NonMaxSuppression",
                Domain:    "",
                Attribute: nil,
                DocString: "",
            },
        },
        Name:        "test_nonmaxsuppression_limit_output_size",
        Initializer: nil,
        DocString:   "",
        Input:       {
            &ir.ValueInfoProto{
                Name: "boxes",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:1},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:6},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:4},
                                        Denotation: "",
                                    },
                                },
                            },
                        },
                    },
                    Denotation: "",
                },
                DocString: "",
            },
            &ir.ValueInfoProto{
                Name: "scores",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:1},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:1},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:6},
                                        Denotation: "",
                                    },
                                },
                            },
                        },
                    },
                    Denotation: "",
                },
                DocString: "",
            },
            &ir.ValueInfoProto{
                Name: "max_output_boxes_per_class",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 7,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:1},
                                        Denotation: "",
                                    },
                                },
                            },
                        },
                    },
                    Denotation: "",
                },
                DocString: "",
            },
            &ir.ValueInfoProto{
                Name: "iou_threshold",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:1},
                                        Denotation: "",
                                    },
                                },
                            },
                        },
                    },
                    Denotation: "",
                },
                DocString: "",
            },
            &ir.ValueInfoProto{
                Name: "score_threshold",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:1},
                                        Denotation: "",
                                    },
                                },
                            },
                        },
                    },
                    Denotation: "",
                },
                DocString: "",
            },
        },
        Output: {
            &ir.ValueInfoProto{
                Name: "selected_indices",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 7,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
                                        Denotation: "",
                                    },
                                },
                            },
                        },
                    },
                    Denotation: "",
                },
                DocString: "",
            },
        },
        ValueInfo:              nil,
        QuantizationAnnotation: nil,
    },
    MetadataProps: nil,
}
*/

// NewTestNonmaxsuppressionLimitOutputSize version: 5.
func NewTestNonmaxsuppressionLimitOutputSize() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "NonMaxSuppression",
		Title:  "TestNonmaxsuppressionLimitOutputSize",
		ModelB: []byte{0x8, 0x5, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xe1, 0x2, 0xa, 0x70, 0xa, 0x5, 0x62, 0x6f, 0x78, 0x65, 0x73, 0xa, 0x6, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0xa, 0x1a, 0x6d, 0x61, 0x78, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0xa, 0xd, 0x69, 0x6f, 0x75, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0xa, 0xf, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x10, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x22, 0x11, 0x4e, 0x6f, 0x6e, 0x4d, 0x61, 0x78, 0x53, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6e, 0x6f, 0x6e, 0x6d, 0x61, 0x78, 0x73, 0x75, 0x70, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5a, 0x1b, 0xa, 0x5, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x6, 0xa, 0x2, 0x8, 0x4, 0x5a, 0x1c, 0xa, 0x6, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x6, 0x5a, 0x28, 0xa, 0x1a, 0x6d, 0x61, 0x78, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x5f, 0x62, 0x6f, 0x78, 0x65, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x1, 0x5a, 0x1b, 0xa, 0xd, 0x69, 0x6f, 0x75, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x1, 0x5a, 0x1d, 0xa, 0xf, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x1, 0x62, 0x22, 0xa, 0x10, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x65, 0x73, 0x12, 0xe, 0xa, 0xc, 0x8, 0x7, 0x12, 0x8, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x3, 0x42, 0x2, 0x10, 0xa},

		/*

		   &ir.NodeProto{
		     Input:     []string{"boxes", "scores", "max_output_boxes_per_class", "iou_threshold", "score_threshold"},
		     Output:    []string{"selected_indices"},
		     Name:      "",
		     OpType:    "NonMaxSuppression",
		     Attributes: ([]*ir.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 6, 4),
				tensor.WithBacking([]float32{0, 0, 1, 1, 0, 0.1, 1, 1.1, 0, -0.1, 1, 0.9, 0, 10, 1, 11, 0, 10.1, 1, 11.1, 0, 100, 1, 101}),
			),

			tensor.New(
				tensor.WithShape(1, 1, 6),
				tensor.WithBacking([]float32{0.9, 0.75, 0.6, 0.95, 0.5, 0.3}),
			),

			tensor.New(
				tensor.WithShape(1),
				tensor.WithBacking([]float32{2}),
			),

			tensor.New(
				tensor.WithShape(1),
				tensor.WithBacking([]float32{0.5}),
			),

			tensor.New(
				tensor.WithShape(1),
				tensor.WithBacking([]float32{0}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 3),
				tensor.WithBacking([]int64{0, 0, 3, 0, 0, 0}),
			),
		},
	}
}
