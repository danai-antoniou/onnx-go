package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("ReduceL1", "TestReduceL1NegativeAxesKeepDimsExample", NewTestReduceL1NegativeAxesKeepDimsExample)
}

/*
&ir.ModelProto{
    IrVersion:   6,
    OpsetImport: {
        &ir.OperatorSetIdProto{Domain:"", Version:11},
    },
    ProducerName:    "backend-test",
    ProducerVersion: "",
    Domain:          "",
    ModelVersion:    0,
    DocString:       "",
    Graph:           &ir.GraphProto{
        Node: {
            &ir.NodeProto{
                Input:     {"data"},
                Output:    {"reduced"},
                Name:      "",
                OpType:    "ReduceL1",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:        "axes",
                        RefAttrName: "",
                        DocString:   "",
                        Type:        7,
                        F:           0,
                        I:           0,
                        S:           nil,
                        T:           (*ir.TensorProto)(nil),
                        G:           (*ir.GraphProto)(nil),
                        Floats:      nil,
                        Ints:        {-1},
                        Strings:     nil,
                        Tensors:     nil,
                        Graphs:      nil,
                    },
                    &ir.AttributeProto{
                        Name:        "keepdims",
                        RefAttrName: "",
                        DocString:   "",
                        Type:        2,
                        F:           0,
                        I:           1,
                        S:           nil,
                        T:           (*ir.TensorProto)(nil),
                        G:           (*ir.GraphProto)(nil),
                        Floats:      nil,
                        Ints:        nil,
                        Strings:     nil,
                        Tensors:     nil,
                        Graphs:      nil,
                    },
                },
                DocString: "",
            },
        },
        Name:        "test_reduce_l1_negative_axes_keep_dims_example",
        Initializer: nil,
        DocString:   "",
        Input:       {
            &ir.ValueInfoProto{
                Name: "data",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
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
                Name: "reduced",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
                                        Denotation: "",
                                    },
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
        ValueInfo:              nil,
        QuantizationAnnotation: nil,
    },
    MetadataProps: nil,
}
*/

// NewTestReduceL1NegativeAxesKeepDimsExample version: 6.
func NewTestReduceL1NegativeAxesKeepDimsExample() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "ReduceL1",
		Title:  "TestReduceL1NegativeAxesKeepDimsExample",
		ModelB: []byte{0x8, 0x6, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xad, 0x1, 0xa, 0x40, 0xa, 0x4, 0x64, 0x61, 0x74, 0x61, 0x12, 0x7, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x64, 0x22, 0x8, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x4c, 0x31, 0x2a, 0x14, 0xa, 0x4, 0x61, 0x78, 0x65, 0x73, 0x40, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x1, 0xa0, 0x1, 0x7, 0x2a, 0xf, 0xa, 0x8, 0x6b, 0x65, 0x65, 0x70, 0x64, 0x69, 0x6d, 0x73, 0x18, 0x1, 0xa0, 0x1, 0x2, 0x12, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x5f, 0x6c, 0x31, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x78, 0x65, 0x73, 0x5f, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x64, 0x69, 0x6d, 0x73, 0x5f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5a, 0x1a, 0xa, 0x4, 0x64, 0x61, 0x74, 0x61, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x2, 0x62, 0x1d, 0xa, 0x7, 0x72, 0x65, 0x64, 0x75, 0x63, 0x65, 0x64, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x1, 0x42, 0x2, 0x10, 0xb},

		/*

		   &ir.NodeProto{
		     Input:     []string{"data"},
		     Output:    []string{"reduced"},
		     Name:      "",
		     OpType:    "ReduceL1",
		     Attributes: ([]*ir.AttributeProto) (len=2 cap=2) {
		    (*ir.AttributeProto)(0xc00046e000)(name:"axes" type:INTS ints:-1 ),
		    (*ir.AttributeProto)(0xc00046e0e0)(name:"keepdims" type:INT i:1 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 2, 2),
				tensor.WithBacking([]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 2, 1),
				tensor.WithBacking([]float32{3, 7, 11, 15, 19, 23}),
			),
		},
	}
}
