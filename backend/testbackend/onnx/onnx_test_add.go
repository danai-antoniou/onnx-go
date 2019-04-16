package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Add", "TestAdd", NewTestAdd)
}

// NewTestAdd version: 3.
func NewTestAdd() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Add",
		Title:  "TestAdd",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x69, 0xa, 0x10, 0xa, 0x1, 0x78, 0xa, 0x1, 0x79, 0x12, 0x3, 0x73, 0x75, 0x6d, 0x22, 0x3, 0x41, 0x64, 0x64, 0x12, 0x8, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x5a, 0x17, 0xa, 0x1, 0x78, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x5a, 0x17, 0xa, 0x1, 0x79, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x62, 0x19, 0xa, 0x3, 0x73, 0x75, 0x6d, 0x12, 0x12, 0xa, 0x10, 0x8, 0x1, 0x12, 0xc, 0xa, 0x2, 0x8, 0x3, 0xa, 0x2, 0x8, 0x4, 0xa, 0x2, 0x8, 0x5, 0x42, 0x2, 0x10, 0x9},

		/*

		   &pb.NodeProto{
		     Input:     []string{"x", "y"},
		     Output:    []string{"sum"},
		     Name:      "",
		     OpType:    "Add",
		     Attributes: ([]*pb.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4, 5),
				tensor.WithBacking([]float32{1.7640524, 0.4001572, 0.978738, 2.2408931, 1.867558, -0.9772779, 0.95008844, -0.1513572, -0.10321885, 0.41059852, 0.14404356, 1.4542735, 0.7610377, 0.121675014, 0.44386324, 0.33367434, 1.4940791, -0.20515826, 0.3130677, -0.85409576, -2.5529897, 0.6536186, 0.8644362, -0.742165, 2.2697546, -1.4543657, 0.045758516, -0.18718386, 1.5327792, 1.4693588, 0.15494743, 0.37816253, -0.88778573, -1.9807965, -0.34791216, 0.15634897, 1.2302907, 1.2023798, -0.3873268, -0.30230275, -1.048553, -1.420018, -1.7062702, 1.9507754, -0.5096522, -0.4380743, -1.2527953, 0.7774904, -1.6138978, -0.21274029, -0.89546657, 0.3869025, -0.51080513, -1.1806322, -0.028182229, 0.42833188, 0.06651722, 0.3024719, -0.6343221, -0.36274117}),
			),

			tensor.New(
				tensor.WithShape(3, 4, 5),
				tensor.WithBacking([]float32{-0.67246044, -0.35955316, -0.8131463, -1.7262826, 0.17742614, -0.40178093, -1.6301984, 0.46278226, -0.9072984, 0.051945396, 0.7290906, 0.12898292, 1.1394007, -1.2348258, 0.40234163, -0.6848101, -0.87079716, -0.5788497, -0.31155252, 0.05616534, -1.1651498, 0.9008265, 0.46566245, -1.5362437, 1.4882522, 1.8958892, 1.1787796, -0.17992483, -1.0707526, 1.0544517, -0.40317693, 1.222445, 0.20827498, 0.97663903, 0.3563664, 0.7065732, 0.01050002, 1.7858706, 0.12691209, 0.40198937, 1.8831507, -1.347759, -1.270485, 0.9693967, -1.1731234, 1.9436212, -0.41361898, -0.7474548, 1.922942, 1.4805148, 1.867559, 0.90604466, -0.86122566, 1.9100649, -0.26800337, 0.8024564, 0.947252, -0.15501009, 0.61407936, 0.9222067}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(3, 4, 5),
				tensor.WithBacking([]float32{1.091592, 0.040604055, 0.16559172, 0.5146105, 2.044984, -1.3790588, -0.6801099, 0.31142506, -1.0105172, 0.4625439, 0.87313414, 1.5832564, 1.9004384, -1.1131508, 0.8462049, -0.35113576, 0.62328196, -0.7840079, 0.0015151799, -0.7979304, -3.7181396, 1.554445, 1.3300986, -2.2784088, 3.7580068, 0.44152343, 1.2245381, -0.3671087, 0.4620266, 2.5238104, -0.2482295, 1.6006075, -0.6795108, -1.0041574, 0.008454233, 0.8629222, 1.2407907, 2.9882503, -0.26041472, 0.09968662, 0.8345977, -2.767777, -2.9767551, 2.9201722, -1.6827755, 1.5055468, -1.6664143, 0.030035555, 0.30904424, 1.2677745, 0.9720924, 1.2929472, -1.3720307, 0.7294327, -0.2961856, 1.2307882, 1.0137691, 0.14746182, -0.02024275, 0.5594655}),
			),
		},
	}
}
