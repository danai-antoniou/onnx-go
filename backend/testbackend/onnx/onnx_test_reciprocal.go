package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Reciprocal", "TestReciprocal", NewTestReciprocal)
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
                OpType:    "Reciprocal",
                Domain:    "",
                Attribute: nil,
                DocString: "",
            },
        },
        Name:        "test_reciprocal",
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

// NewTestReciprocal version: 3.
func NewTestReciprocal() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Reciprocal",
		Title:  "TestReciprocal",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x57, 0xa, 0x12, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0xa, 0x52, 0x65, 0x63, 0x69, 0x70, 0x72, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0xf, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x69, 0x70, 0x72, 0x6f, 0x63, 0x61, 0x6c, 0x5a, 0x17, 0xa, 0x1, 0x78, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x62, 0x17, 0xa, 0x1, 0x79, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0x9},

		/*

		   &ir.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Reciprocal",
		     Attributes: ([]*ir.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4, 5),
				tensor.WithBacking([]float32{1.0488136, 1.2151893, 1.1027634, 1.0448833, 0.9236548, 1.145894, 0.9375872, 1.391773, 1.4636627, 0.8834415, 1.291725, 1.0288949, 1.0680445, 1.4255967, 0.57103604, 0.5871293, 0.5202184, 1.3326199, 1.2781568, 1.3700122, 1.4786184, 1.2991586, 0.96147937, 1.2805293, 0.61827445, 1.139921, 0.6433533, 1.4446689, 1.0218483, 0.91466194, 0.7645556, 1.2742337, 0.9561503, 1.068434, 0.5187898, 1.1176355, 1.1120957, 1.1169341, 1.443748, 1.1818203, 0.8595079, 0.937032, 1.1976311, 0.5602255, 1.1667666, 1.1706378, 0.7103826, 0.6289263, 0.8154284, 0.86371076, 1.0701967, 0.9386015, 1.4883738, 0.6020448, 0.7088767, 0.6613095, 1.1531084, 0.7532916, 0.96631074, 0.7444256}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4, 5),
				tensor.WithBacking([]float32{0.9534583, 0.82291704, 0.90681285, 0.9570447, 1.0826555, 0.872681, 1.0665674, 0.718508, 0.6832175, 1.1319369, 0.77415854, 0.97191656, 0.93629056, 0.70146066, 1.751203, 1.7032024, 1.9222697, 0.7504015, 0.78237665, 0.7299205, 0.676307, 0.76972896, 1.040064, 0.7809271, 1.6174047, 0.87725383, 1.554356, 0.6922001, 0.9786188, 1.0933001, 1.3079494, 0.7847854, 1.0458606, 0.93594927, 1.9275628, 0.8947461, 0.8992032, 0.895308, 0.6926417, 0.84615237, 1.1634564, 1.0671995, 0.8349816, 1.7849956, 0.85706943, 0.8542352, 1.4076922, 1.5900115, 1.2263492, 1.157795, 0.93440765, 1.0654149, 0.6718742, 1.6610059, 1.4106826, 1.5121514, 0.8672212, 1.3275071, 1.0348638, 1.3433176}),
			),
		},
	}
}
