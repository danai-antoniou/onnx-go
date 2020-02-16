package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("BatchNormalization", "TestBatchnormEpsilon", NewTestBatchnormEpsilon)
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
                Input:     {"x", "s", "bias", "mean", "var"},
                Output:    {"y"},
                Name:      "",
                OpType:    "BatchNormalization",
                Domain:    "",
                Attribute: {
                    &ir.AttributeProto{
                        Name:        "epsilon",
                        RefAttrName: "",
                        DocString:   "",
                        Type:        1,
                        F:           0.009999999776482582,
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
        Name:        "test_batchnorm_epsilon",
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
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
                Name: "s",
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
                                },
                            },
                        },
                    },
                    Denotation: "",
                },
                DocString: "",
            },
            &ir.ValueInfoProto{
                Name: "bias",
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
                                },
                            },
                        },
                    },
                    Denotation: "",
                },
                DocString: "",
            },
            &ir.ValueInfoProto{
                Name: "mean",
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
                                },
                            },
                        },
                    },
                    Denotation: "",
                },
                DocString: "",
            },
            &ir.ValueInfoProto{
                Name: "var",
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
                                        Value:      &ir.TensorShapeProto_Dimension_DimValue{DimValue:2},
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
        },
        ValueInfo:              nil,
        QuantizationAnnotation: nil,
    },
    MetadataProps: nil,
}
*/

// NewTestBatchnormEpsilon version: 3.
func NewTestBatchnormEpsilon() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "BatchNormalization",
		Title:  "TestBatchnormEpsilon",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0xe1, 0x1, 0xa, 0x41, 0xa, 0x1, 0x78, 0xa, 0x1, 0x73, 0xa, 0x4, 0x62, 0x69, 0x61, 0x73, 0xa, 0x4, 0x6d, 0x65, 0x61, 0x6e, 0xa, 0x3, 0x76, 0x61, 0x72, 0x12, 0x1, 0x79, 0x22, 0x12, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x11, 0xa, 0x7, 0x65, 0x70, 0x73, 0x69, 0x6c, 0x6f, 0x6e, 0x15, 0xa, 0xd7, 0x23, 0x3c, 0xa0, 0x1, 0x1, 0x12, 0x16, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x6e, 0x6f, 0x72, 0x6d, 0x5f, 0x65, 0x70, 0x73, 0x69, 0x6c, 0x6f, 0x6e, 0x5a, 0x1b, 0xa, 0x1, 0x78, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x5a, 0xf, 0xa, 0x1, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x5a, 0x12, 0xa, 0x4, 0x62, 0x69, 0x61, 0x73, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x5a, 0x12, 0xa, 0x4, 0x6d, 0x65, 0x61, 0x6e, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x5a, 0x11, 0xa, 0x3, 0x76, 0x61, 0x72, 0x12, 0xa, 0xa, 0x8, 0x8, 0x1, 0x12, 0x4, 0xa, 0x2, 0x8, 0x3, 0x62, 0x1b, 0xa, 0x1, 0x79, 0x12, 0x16, 0xa, 0x14, 0x8, 0x1, 0x12, 0x10, 0xa, 0x2, 0x8, 0x2, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x", "s", "bias", "mean", "var"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "BatchNormalization",
		     Attributes: ([]*ir.AttributeProto) (len=1 cap=1) {
		    (*ir.AttributeProto)(0xc000176620)(name:"epsilon" type:FLOAT f:0.01 )
		   }
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 3, 4, 5),
				tensor.WithBacking([]float32{1.7640524, 0.4001572, 0.978738, 2.2408931, 1.867558, -0.9772779, 0.95008844, -0.1513572, -0.10321885, 0.41059852, 0.14404356, 1.4542735, 0.7610377, 0.121675014, 0.44386324, 0.33367434, 1.4940791, -0.20515826, 0.3130677, -0.85409576, -2.5529897, 0.6536186, 0.8644362, -0.742165, 2.2697546, -1.4543657, 0.045758516, -0.18718386, 1.5327792, 1.4693588, 0.15494743, 0.37816253, -0.88778573, -1.9807965, -0.34791216, 0.15634897, 1.2302907, 1.2023798, -0.3873268, -0.30230275, -1.048553, -1.420018, -1.7062702, 1.9507754, -0.5096522, -0.4380743, -1.2527953, 0.7774904, -1.6138978, -0.21274029, -0.89546657, 0.3869025, -0.51080513, -1.1806322, -0.028182229, 0.42833188, 0.06651722, 0.3024719, -0.6343221, -0.36274117, -0.67246044, -0.35955316, -0.8131463, -1.7262826, 0.17742614, -0.40178093, -1.6301984, 0.46278226, -0.9072984, 0.051945396, 0.7290906, 0.12898292, 1.1394007, -1.2348258, 0.40234163, -0.6848101, -0.87079716, -0.5788497, -0.31155252, 0.05616534, -1.1651498, 0.9008265, 0.46566245, -1.5362437, 1.4882522, 1.8958892, 1.1787796, -0.17992483, -1.0707526, 1.0544517, -0.40317693, 1.222445, 0.20827498, 0.97663903, 0.3563664, 0.7065732, 0.01050002, 1.7858706, 0.12691209, 0.40198937, 1.8831507, -1.347759, -1.270485, 0.9693967, -1.1731234, 1.9436212, -0.41361898, -0.7474548, 1.922942, 1.4805148, 1.867559, 0.90604466, -0.86122566, 1.9100649, -0.26800337, 0.8024564, 0.947252, -0.15501009, 0.61407936, 0.9222067}),
			),

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]float32{0.37642553, -1.0994008, 0.2982382}),
			),

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]float32{1.3263859, -0.69456786, -0.14963454}),
			),

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]float32{-0.43515354, 1.8492638, 0.67229474}),
			),

			tensor.New(
				tensor.WithShape(3),
				tensor.WithBacking([]float32{0.9755215, 0.8558034, 0.011714084}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2, 3, 4, 5),
				tensor.WithBacking([]float32{2.160282, 1.6431195, 1.8625059, 2.3410907, 2.1995292, 1.1208229, 1.8516426, 1.433996, 1.452249, 1.6470785, 1.5460061, 2.0428197, 1.7799582, 1.5375245, 1.6596919, 1.6179104, 2.0579133, 1.4135956, 1.6100968, 1.1675311, 4.506847, 0.71812826, 0.46903986, 2.3672945, -1.1913923, 3.208784, 1.4363358, 1.7115653, -0.32063043, -0.24569702, 1.3073254, 1.0435891, 2.5393505, 3.8307803, 1.9014714, 1.3056693, 0.036770284, 0.069747865, 1.9480412, 1.8475821, -3.6324859, -4.3842998, -4.9636493, 2.4379027, -2.541796, -2.396928, -4.045855, 0.06327261, -4.7766957, -1.940871, -3.3226516, -0.72724444, -2.5441291, -3.899803, -1.5673411, -0.6433948, -1.3756773, -0.8981249, -2.794117, -2.24446, 1.2364037, 1.355052, 1.1830584, 0.8368149, 1.5586641, 1.33904, 0.87324816, 1.6668656, 1.1473577, 1.5110843, 1.7678446, 1.5402954, 1.9234262, 1.0231657, 1.6439477, 1.2317209, 1.1611983, 1.2718991, 1.373253, 1.5126843, 2.8670657, 0.42604357, 0.9402048, 3.305526, -0.26802018, -0.74965733, 0.09763235, 1.7029884, 2.7555318, 0.24453002, 1.9667685, 0.046040177, 1.244317, 0.3364684, 1.0693419, 0.65556055, 1.4779949, -0.6196666, 1.3404498, 1.0154369, 2.3010361, -4.238054, -4.081658, 0.45167488, -3.8846056, 2.4234233, -2.3474326, -3.0230885, 2.3815703, 1.4861348, 2.2694798, 0.32345566, -3.2533512, 2.3555083, -2.0527189, 0.11380169, 0.40685582, -1.8240302, -0.26745757, 0.35616624}),
			),
		},
	}
}
