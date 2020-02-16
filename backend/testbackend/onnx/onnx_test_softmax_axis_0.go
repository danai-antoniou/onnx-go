package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Softmax", "TestSoftmaxAxis0", NewTestSoftmaxAxis0)
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
                OpType:    "Softmax",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:        "axis",
                        RefAttrName: "",
                        DocString:   "",
                        Type:        2,
                        F:           0,
                        I:           0,
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
        Name:        "test_softmax_axis_0",
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
        },
        ValueInfo:              nil,
        QuantizationAnnotation: nil,
    },
    MetadataProps: nil,
}
*/

// NewTestSoftmaxAxis0 version: 3.
func NewTestSoftmaxAxis0() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Softmax",
		Title:  "TestSoftmaxAxis0",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x65, 0xa, 0x1c, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x7, 0x53, 0x6f, 0x66, 0x74, 0x6d, 0x61, 0x78, 0x2a, 0xb, 0xa, 0x4, 0x61, 0x78, 0x69, 0x73, 0x18, 0x0, 0xa0, 0x1, 0x2, 0x12, 0x13, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x6f, 0x66, 0x74, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x5f, 0x30, 0x5a, 0x17, 0xa, 0x1, 0x78, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x62, 0x17, 0xa, 0x1, 0x79, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Softmax",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc0001762a0)(name:"axis" type:INT )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4, 5),
				tensor.WithBacking([]float32{1.7640524, 0.4001572, 0.978738, 2.2408931, 1.867558, 0.9772779, 0.95008844, 0.1513572, 0.10321885, 0.41059852, 0.14404356, 1.4542735, 0.7610377, 0.121675014, 0.44386324, 0.33367434, 1.4940791, 0.20515826, 0.3130677, 0.85409576, 2.5529897, 0.6536186, 0.8644362, 0.742165, 2.2697546, 1.4543657, 0.045758516, 0.18718386, 1.5327792, 1.4693588, 0.15494743, 0.37816253, 0.88778573, 1.9807965, 0.34791216, 0.15634897, 1.2302907, 1.2023798, 0.3873268, 0.30230275, 1.048553, 1.420018, 1.7062702, 1.9507754, 0.5096522, 0.4380743, 1.2527953, 0.7774904, 1.6138978, 0.21274029, 0.89546657, 0.3869025, 0.51080513, 1.1806322, 0.028182229, 0.42833188, 0.06651722, 0.3024719, 0.6343221, 0.36274117}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4, 5),
				tensor.WithBacking([]float32{0.033031557, 0.008444946, 0.015061609, 0.053213075, 0.036633715, 0.015039632, 0.014636223, 0.006584828, 0.0062753535, 0.008533584, 0.0065368456, 0.024232252, 0.012115052, 0.006392248, 0.008822225, 0.007901759, 0.025216287, 0.0069488017, 0.007740596, 0.013296577, 0.072704315, 0.010881124, 0.0134347845, 0.011888553, 0.054771263, 0.024234489, 0.0059249336, 0.0068250173, 0.02621129, 0.024600575, 0.006608512, 0.008261231, 0.013752171, 0.04102606, 0.008015066, 0.00661778, 0.019369539, 0.018836394, 0.008337286, 0.0076577165, 0.01615071, 0.023416221, 0.031177018, 0.039812714, 0.009422149, 0.008771302, 0.019810386, 0.012316027, 0.02842613, 0.0070016873, 0.013858207, 0.00833375, 0.009433018, 0.018431168, 0.0058217053, 0.008686263, 0.0060492125, 0.007659011, 0.010673171, 0.008134809}),
			),
		},
	}
}
