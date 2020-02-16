package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Pad", "TestConstantPad", NewTestConstantPad)
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
                Input:     {"x", "pads", "value"},
                Output:    {"y"},
                Name:      "",
                OpType:    "Pad",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:        "mode",
                        RefAttrName: "",
                        DocString:   "",
                        Type:        3,
                        F:           0,
                        I:           0,
                        S:           {0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74},
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
        Name:        "test_constant_pad",
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:4},
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
            &ir.ValueInfoProto{
                Name: "pads",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 7,
                            Shape:    &ir.TensorShapeProto{
                                Dim: {
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:8},
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
                Name: "value",
                Type: &ir.TypeProto{
                    Value: &ir.TypeProto_TensorType{
                        TensorType: &ir.TypeProto_Tensor{
                            ElemType: 1,
                            Shape:    &ir.TensorShapeProto{},
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:3},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:7},
                                        Denotation: "",
                                    },
                                    &ir.TensorShapeProto_Dimension{
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:12},
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

// NewTestConstantPad version: 6.
func NewTestConstantPad() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Pad",
		Title:  "TestConstantPad",
		ModelB: []byte{0x8, 0x6, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xa1, 0x1, 0xa, 0x2d, 0xa, 0x1, 0x78, 0xa, 0x4, 0x70, 0x61, 0x64, 0x73, 0xa, 0x5, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1, 0x79, 0x22, 0x3, 0x50, 0x61, 0x64, 0x2a, 0x13, 0xa, 0x4, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x8, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0xa0, 0x1, 0x3, 0x12, 0x11, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x64, 0x5a, 0x1b, 0xa, 0x1, 0x78, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x5a, 0x12, 0xa, 0x4, 0x70, 0x61, 0x64, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x7, 0x12, 0x4, 0xa, 0x2, 0x8, 0x8, 0x5a, 0xf, 0xa, 0x5, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x6, 0xa, 0x4, 0x8, 0x1, 0x12, 0x0, 0x62, 0x1b, 0xa, 0x1, 0x79, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x7, 0xa, 0x2, 0x8, 0xc, 0x42, 0x2, 0x10, 0xb},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x", "pads", "value"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Pad",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc000176700)(name:"mode" type:STRING s:"constant" )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 3, 4, 5),
				tensor.WithBacking([]float32{1.7640524, 0.4001572, 0.978738, 2.2408931, 1.867558, -0.9772779, 0.95008844, -0.1513572, -0.10321885, 0.41059852, 0.14404356, 1.4542735, 0.7610377, 0.121675014, 0.44386324, 0.33367434, 1.4940791, -0.20515826, 0.3130677, -0.85409576, -2.5529897, 0.6536186, 0.8644362, -0.742165, 2.2697546, -1.4543657, 0.045758516, -0.18718386, 1.5327792, 1.4693588, 0.15494743, 0.37816253, -0.88778573, -1.9807965, -0.34791216, 0.15634897, 1.2302907, 1.2023798, -0.3873268, -0.30230275, -1.048553, -1.420018, -1.7062702, 1.9507754, -0.5096522, -0.4380743, -1.2527953, 0.7774904, -1.6138978, -0.21274029, -0.89546657, 0.3869025, -0.51080513, -1.1806322, -0.028182229, 0.42833188, 0.06651722, 0.3024719, -0.6343221, -0.36274117}),
			),

			tensor.New(
				tensor.WithShape(8),
				tensor.WithBacking([]int64{0, 0, 1, 3, 0, 0, 2, 4}),
			),

			tensor.New(
				tensor.WithShape(1),
				tensor.WithBacking([]float32{1.2}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 3, 7, 12),
				tensor.WithBacking([]float32{1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.7640524, 0.4001572, 0.978738, 2.2408931, 1.867558, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, -0.9772779, 0.95008844, -0.1513572, -0.10321885, 0.41059852, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 0.14404356, 1.4542735, 0.7610377, 0.121675014, 0.44386324, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 0.33367434, 1.4940791, -0.20515826, 0.3130677, -0.85409576, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, -2.5529897, 0.6536186, 0.8644362, -0.742165, 2.2697546, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, -1.4543657, 0.045758516, -0.18718386, 1.5327792, 1.4693588, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 0.15494743, 0.37816253, -0.88778573, -1.9807965, -0.34791216, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 0.15634897, 1.2302907, 1.2023798, -0.3873268, -0.30230275, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, -1.048553, -1.420018, -1.7062702, 1.9507754, -0.5096522, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, -0.4380743, -1.2527953, 0.7774904, -1.6138978, -0.21274029, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, -0.89546657, 0.3869025, -0.51080513, -1.1806322, -0.028182229, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 0.42833188, 0.06651722, 0.3024719, -0.6343221, -0.36274117, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2, 1.2}),
			),
		},
	}
}
