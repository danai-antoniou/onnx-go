package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("LogSoftmax", "TestLogsoftmaxAxis0", NewTestLogsoftmaxAxis0)
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
                OpType:    "LogSoftmax",
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
        Name:        "test_logsoftmax_axis_0",
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

// NewTestLogsoftmaxAxis0 version: 3.
func NewTestLogsoftmaxAxis0() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "LogSoftmax",
		Title:  "TestLogsoftmaxAxis0",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x6b, 0xa, 0x1f, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0xa, 0x4c, 0x6f, 0x67, 0x53, 0x6f, 0x66, 0x74, 0x6d, 0x61, 0x78, 0x2a, 0xb, 0xa, 0x4, 0x61, 0x78, 0x69, 0x73, 0x18, 0x0, 0xa0, 0x1, 0x2, 0x12, 0x16, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x6f, 0x66, 0x74, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x78, 0x69, 0x73, 0x5f, 0x30, 0x5a, 0x17, 0xa, 0x1, 0x78, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x62, 0x17, 0xa, 0x1, 0x79, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "LogSoftmax",
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
				tensor.WithBacking([]float32{-3.410292, -4.774187, -4.195606, -2.9334512, -3.3067863, -4.1970663, -4.2242556, -5.0229874, -5.0711255, -4.763746, -5.0303006, -3.7200708, -4.4133067, -5.0526695, -4.730481, -4.8406696, -3.6802652, -4.969186, -4.8612766, -4.3202486, -2.6213546, -4.5207257, -4.309908, -4.4321795, -2.9045897, -3.7199786, -5.128586, -4.9871607, -3.641565, -3.7049856, -5.019397, -4.7961817, -4.2865586, -3.1935477, -4.826432, -5.0179954, -3.9440536, -3.9719644, -4.787018, -4.8720417, -4.1257915, -3.7543263, -3.468074, -3.223569, -4.664692, -4.73627, -3.9215488, -4.396854, -3.5604465, -4.961604, -4.2788777, -4.7874417, -4.663539, -3.993712, -5.146162, -4.7460127, -5.107827, -4.8718724, -4.540022, -4.811603}),
			),
		},
	}
}
