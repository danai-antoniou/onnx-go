package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("AveragePool", "TestAveragepool2dPrecomputedPadsCountIncludePad", NewTestAveragepool2dPrecomputedPadsCountIncludePad)
}

/*
&ir.ModelProto{
    IrVersion:   3,
    OpsetImport: {
        &ir.OperatorSetIdProto{Domain:"", Version:9},
    },
    ProducerName:    "backend-test",
    ProducerVersion: "",
    Domain:          "",
    ModelVersion:    0,
    DocString:       "",
    Graph:           &ir.GraphProto{
        Node: {
            &ir.NodeProto{
                Input:     {"x"},
                Output:    {"y"},
                Name:      "",
                OpType:    "AveragePool",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:        "count_include_pad",
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
                    &ir.AttributeProto{
                        Name:        "kernel_shape",
                        RefAttrName: "",
                        DocString:   "",
                        Type:        7,
                        F:           0,
                        I:           0,
                        S:           nil,
                        T:           (*ir.TensorProto)(nil),
                        G:           (*ir.GraphProto)(nil),
                        Floats:      nil,
                        Ints:        {5, 5},
                        Strings:     nil,
                        Tensors:     nil,
                        Graphs:      nil,
                    },
                    &ir.AttributeProto{
                        Name:        "pads",
                        RefAttrName: "",
                        DocString:   "",
                        Type:        7,
                        F:           0,
                        I:           0,
                        S:           nil,
                        T:           (*ir.TensorProto)(nil),
                        G:           (*ir.GraphProto)(nil),
                        Floats:      nil,
                        Ints:        {2, 2, 2, 2},
                        Strings:     nil,
                        Tensors:     nil,
                        Graphs:      nil,
                    },
                },
                DocString: "",
            },
        },
        Name:        "test_averagepool_2d_precomputed_pads_count_include_pad",
        Initializer: nil,
        DocString:   "",
        Input:       {
            &ir.ValueInfoProto{
                Name: "x",
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:5},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:5},
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
                Name: "y",
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:5},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:5},
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

// NewTestAveragepool2dPrecomputedPadsCountIncludePad version: 3.
func NewTestAveragepool2dPrecomputedPadsCountIncludePad() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "AveragePool",
		Title:  "TestAveragepool2dPrecomputedPadsCountIncludePad",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xcb, 0x1, 0xa, 0x57, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0xb, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x2a, 0x18, 0xa, 0x11, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x70, 0x61, 0x64, 0x18, 0x1, 0xa0, 0x1, 0x2, 0x2a, 0x15, 0xa, 0xc, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x40, 0x5, 0x40, 0x5, 0xa0, 0x1, 0x7, 0x2a, 0x11, 0xa, 0x4, 0x70, 0x61, 0x64, 0x73, 0x40, 0x2, 0x40, 0x2, 0x40, 0x2, 0x40, 0x2, 0xa0, 0x1, 0x7, 0x12, 0x36, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x32, 0x64, 0x5f, 0x70, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x64, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x70, 0x61, 0x64, 0x5a, 0x1b, 0xa, 0x1, 0x78, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x5, 0x62, 0x1b, 0xa, 0x1, 0x79, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x5, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "AveragePool",
		     Attributes: ([]*ir.AttributeProto) (len=3 cap=4) {
		    (*ir.AttributeProto)(0xc000176460)(name:"count_include_pad" type:INT i:1 ),
		    (*ir.AttributeProto)(0xc000176540)(name:"kernel_shape" type:INTS ints:5 ints:5 ),
		    (*ir.AttributeProto)(0xc000176620)(name:"pads" type:INTS ints:2 ints:2 ints:2 ints:2 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 1, 5, 5),
				tensor.WithBacking([]float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 1, 5, 5),
				tensor.WithBacking([]float32{2.52, 3.6, 4.8, 4.08, 3.24, 4.56, 6.4, 8.4, 7.04, 5.52, 7.2, 10, 13, 10.8, 8.4, 6.96, 9.6, 12.4, 10.24, 7.92, 6.12, 8.4, 10.8, 8.88, 6.84}),
			),
		},
	}
}
