//go:build linux || darwin

package dark

// png, 32bit
var Neutral []byte = []byte{
	0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d,
	0x49, 0x48, 0x44, 0x52, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x20,
	0x08, 0x06, 0x00, 0x00, 0x00, 0x73, 0x7a, 0x7a, 0xf4, 0x00, 0x00, 0x04,
	0x01, 0x7a, 0x54, 0x58, 0x74, 0x52, 0x61, 0x77, 0x20, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x20, 0x74, 0x79, 0x70, 0x65, 0x20, 0x65, 0x78,
	0x69, 0x66, 0x00, 0x00, 0x78, 0xda, 0xb5, 0x57, 0x5b, 0x76, 0xac, 0x38,
	0x0c, 0xfc, 0xf7, 0x2a, 0x66, 0x09, 0x96, 0xe4, 0x87, 0xbc, 0x1c, 0x83,
	0xf1, 0x39, 0xb3, 0x83, 0x59, 0xfe, 0x94, 0x8d, 0x21, 0x74, 0x37, 0xa4,
	0xe9, 0xe4, 0x5e, 0x48, 0xf0, 0x4b, 0x48, 0xe5, 0x2a, 0xc9, 0x24, 0x66,
	0xf9, 0xef, 0xdf, 0x6a, 0xfe, 0xc1, 0xc5, 0xec, 0xd4, 0x38, 0x1f, 0x35,
	0xa4, 0x10, 0x2c, 0x2e, 0x97, 0x5c, 0xe2, 0x8c, 0x8e, 0xda, 0xf5, 0x9a,
	0xfa, 0x93, 0xac, 0xeb, 0xcf, 0x7e, 0xf1, 0x58, 0xc2, 0xf8, 0x61, 0xde,
	0xec, 0x0b, 0x8c, 0x29, 0x41, 0x2b, 0xeb, 0x30, 0x6d, 0x0b, 0x0b, 0xe6,
	0xd1, 0xa7, 0x31, 0x4e, 0x23, 0x08, 0x6d, 0xf6, 0x9b, 0xa3, 0xad, 0x43,
	0x19, 0x3d, 0xff, 0xb5, 0x90, 0xf3, 0x98, 0x9f, 0x1e, 0xe7, 0xa7, 0xe1,
	0x90, 0xf5, 0xd9, 0xd1, 0x40, 0x20, 0xb4, 0x46, 0xb6, 0x65, 0xbc, 0x30,
	0x1c, 0x09, 0x0f, 0x44, 0x6e, 0x1d, 0xcf, 0x03, 0x51, 0x48, 0x1a, 0x1f,
	0xb6, 0x56, 0xe6, 0x11, 0xd9, 0x8d, 0x29, 0xfd, 0xfa, 0x75, 0x12, 0x39,
	0xf8, 0x40, 0xd1, 0xe1, 0xe9, 0xd8, 0xc6, 0x18, 0x12, 0xfa, 0xca, 0xd6,
	0x45, 0xf0, 0x59, 0x1a, 0xd0, 0x3a, 0x73, 0xea, 0x8e, 0xfc, 0x4a, 0xe8,
	0x3e, 0xb1, 0x8d, 0x37, 0x53, 0x06, 0x26, 0x5e, 0x84, 0xc4, 0xe2, 0x29,
	0xc2, 0x2b, 0x4a, 0x69, 0xbf, 0x2c, 0x19, 0x6d, 0xea, 0x4f, 0x35, 0xcd,
	0x10, 0x06, 0x6d, 0x10, 0xf1, 0x64, 0xd1, 0x55, 0x1a, 0x48, 0x09, 0x08,
	0x40, 0x9a, 0x46, 0xa0, 0x6c, 0x77, 0x32, 0x8f, 0xdc, 0x7c, 0x71, 0x74,
	0x71, 0xdd, 0xd9, 0x96, 0x45, 0x90, 0xba, 0x34, 0xe3, 0x83, 0x6a, 0x7b,
	0xfb, 0x94, 0x37, 0x7b, 0x8f, 0x2e, 0xe6, 0x47, 0x1a, 0xec, 0xaa, 0x69,
	0x18, 0x0b, 0xf2, 0x28, 0xab, 0x0d, 0x7b, 0x7b, 0x3a, 0x4f, 0x7e, 0x73,
	0xb4, 0x2d, 0xc8, 0x1e, 0x87, 0x8f, 0x91, 0x75, 0xde, 0x23, 0x3f, 0xcc,
	0x8b, 0x6f, 0x0c, 0x7e, 0x5d, 0xe6, 0x28, 0x77, 0xad, 0x45, 0x6b, 0xdf,
	0x34, 0x76, 0x91, 0x5d, 0x00, 0x17, 0x61, 0x6c, 0x6a, 0xdb, 0x4a, 0xef,
	0xc1, 0x6e, 0x6a, 0x2c, 0xf6, 0xb7, 0x02, 0xee, 0x68, 0x83, 0x41, 0xd6,
	0x2a, 0x3a, 0xed, 0x4e, 0xb8, 0xd5, 0x66, 0x3b, 0x23, 0xa7, 0x8a, 0x9d,
	0x51, 0x69, 0x13, 0xfa, 0x89, 0x18, 0xda, 0x57, 0x72, 0x54, 0x28, 0x53,
	0xa5, 0xa5, 0xb7, 0x33, 0xcd, 0x80, 0xe8, 0x78, 0xe1, 0x88, 0x96, 0x79,
	0x36, 0x2c, 0x7d, 0x52, 0x21, 0x52, 0xe2, 0x19, 0x49, 0x40, 0xe2, 0xda,
	0x4d, 0x95, 0x23, 0x72, 0xa4, 0x88, 0x22, 0x25, 0xe6, 0x9e, 0x43, 0x4e,
	0x78, 0xc7, 0x42, 0x3d, 0x6c, 0xea, 0xe1, 0x66, 0x52, 0x5b, 0x8c, 0x2d,
	0x04, 0x53, 0x26, 0x38, 0xa3, 0x9e, 0x64, 0x3f, 0xbc, 0xcd, 0x5d, 0xc3,
	0x5a, 0x5b, 0x2d, 0x11, 0x59, 0xdd, 0xb9, 0x02, 0x2e, 0x6e, 0x99, 0x0e,
	0x14, 0x96, 0x20, 0x7f, 0x6b, 0x60, 0x06, 0x45, 0xa8, 0x0e, 0x52, 0x7d,
	0x27, 0x78, 0xbb, 0x9f, 0xaf, 0xa6, 0xab, 0x40, 0x41, 0xdf, 0x69, 0x56,
	0x6c, 0x30, 0xdb, 0xc9, 0xac, 0x2e, 0x26, 0x4f, 0x5f, 0xc9, 0x25, 0x5d,
	0x68, 0x81, 0xa1, 0x47, 0xbb, 0x56, 0x3d, 0xc5, 0x32, 0x1c, 0x80, 0x22,
	0x84, 0xf6, 0x00, 0x43, 0x02, 0x05, 0x6c, 0x20, 0xf1, 0x14, 0x80, 0x28,
	0x32, 0x47, 0x22, 0x10, 0xa9, 0x10, 0x28, 0x03, 0x3a, 0x8b, 0xe3, 0x09,
	0x0a, 0x90, 0xf7, 0x5c, 0x00, 0x92, 0x9d, 0x48, 0x80, 0x38, 0xa8, 0x0e,
	0xc4, 0xc6, 0x3b, 0x91, 0xba, 0x29, 0x7b, 0x5e, 0xa7, 0x71, 0xaa, 0x8a,
	0x33, 0xe2, 0x25, 0xa0, 0x5c, 0xb5, 0x55, 0x31, 0xc4, 0x72, 0xce, 0x23,
	0x7f, 0xa2, 0x53, 0xe4, 0x50, 0xf6, 0xe2, 0x9d, 0xf7, 0x3e, 0xf8, 0xe8,
	0xd5, 0x27, 0x9f, 0x83, 0x84, 0x56, 0x79, 0x21, 0xc4, 0xd0, 0x8e, 0xe7,
	0x1c, 0x25, 0xba, 0xe8, 0x63, 0x88, 0x31, 0xaa, 0x89, 0x29, 0x66, 0x15,
	0x75, 0xea, 0x35, 0x68, 0x54, 0xd5, 0xa4, 0x39, 0x71, 0x12, 0x1c, 0xdf,
	0x3e, 0xa1, 0x4e, 0x93, 0xa6, 0x94, 0x72, 0x46, 0xd0, 0x0c, 0xcf, 0x19,
	0x6f, 0x67, 0x18, 0xe4, 0x3c, 0xf1, 0x24, 0x93, 0x9b, 0xfc, 0x14, 0xa6,
	0x38, 0xe9, 0x94, 0xcc, 0x94, 0x67, 0xa4, 0xcf, 0xec, 0x66, 0x3f, 0x87,
	0x39, 0xce, 0x3a, 0xa7, 0x39, 0x17, 0x2e, 0x52, 0x50, 0xe0, 0x25, 0x94,
	0x58, 0xb4, 0xa4, 0x92, 0x17, 0x5a, 0x90, 0x4a, 0x8b, 0x5b, 0xfc, 0x12,
	0x96, 0xb8, 0xe8, 0x92, 0x96, 0x5c, 0x91, 0x6a, 0x55, 0xaa, 0xab, 0xbe,
	0x86, 0x1a, 0x4d, 0xd5, 0x9a, 0x6a, 0xde, 0x55, 0x1b, 0xb2, 0xbe, 0xdc,
	0x1f, 0xa8, 0x46, 0x43, 0x35, 0xee, 0x4a, 0x35, 0xc3, 0xb8, 0xab, 0x86,
	0xd9, 0x18, 0x37, 0x17, 0xd4, 0x8e, 0x13, 0xdf, 0x34, 0x83, 0x62, 0xec,
	0x08, 0x82, 0x47, 0xa8, 0x06, 0xc5, 0x90, 0xd8, 0x4d, 0x33, 0xab, 0xe4,
	0x1c, 0x37, 0xe5, 0x9a, 0x66, 0xf8, 0x1e, 0xa1, 0x2a, 0x3c, 0x03, 0xa4,
	0x6f, 0xe2, 0x14, 0x6a, 0x8a, 0x41, 0x41, 0xb7, 0x10, 0xfb, 0x4a, 0xbb,
	0x76, 0x43, 0x39, 0x03, 0x16, 0xff, 0x88, 0x6e, 0x26, 0x6a, 0xd7, 0x8d,
	0x7f, 0xab, 0x9c, 0x69, 0xd2, 0xdd, 0x54, 0xee, 0x55, 0xb7, 0x33, 0xd5,
	0x4a, 0xfb, 0x4a, 0xcc, 0x5d, 0xb1, 0xb5, 0x0c, 0x1b, 0xa9, 0x56, 0x50,
	0x7d, 0x35, 0xd6, 0xcc, 0x8a, 0x1f, 0x7c, 0x7e, 0xaf, 0x5b, 0xf3, 0xce,
	0xe0, 0xa5, 0xad, 0xda, 0x3e, 0xd7, 0xb2, 0xd4, 0xba, 0x76, 0xc7, 0x82,
	0x79, 0xb0, 0xec, 0xcb, 0x3c, 0x2d, 0x7d, 0x62, 0x8c, 0x6e, 0x86, 0x30,
	0x1f, 0x60, 0xe9, 0xed, 0x11, 0xc5, 0xb1, 0x35, 0x67, 0xb6, 0x53, 0x5d,
	0x3a, 0xaa, 0xd6, 0xe9, 0x0b, 0x8f, 0x18, 0x4f, 0x47, 0xe6, 0xbb, 0xc5,
	0x1b, 0xa3, 0x73, 0xb2, 0x65, 0x79, 0xc1, 0x77, 0xc2, 0xd3, 0x6a, 0x35,
	0x61, 0x61, 0xf3, 0xd8, 0xc9, 0x7e, 0x9e, 0xd0, 0x8f, 0x65, 0x3c, 0xaa,
	0x76, 0x00, 0xbd, 0xf3, 0x72, 0x20, 0xf7, 0x4e, 0x30, 0x73, 0x0d, 0xfe,
	0x33, 0x4c, 0xe6, 0x62, 0xe1, 0xe0, 0xfe, 0xca, 0xed, 0x8a, 0x7d, 0x63,
	0xcc, 0x6c, 0xa0, 0xed, 0x7d, 0x54, 0xa7, 0x89, 0x6a, 0xee, 0x72, 0xf0,
	0xca, 0xe3, 0xa3, 0xad, 0xb9, 0x8e, 0x71, 0x91, 0x3f, 0x9f, 0x90, 0x7d,
	0xaa, 0xd9, 0x5a, 0xac, 0x97, 0x75, 0x60, 0xce, 0x81, 0xde, 0x38, 0x04,
	0x9e, 0xde, 0x32, 0xf7, 0xa8, 0xdc, 0x8f, 0x8f, 0x6f, 0xe4, 0x7f, 0x29,
	0xd1, 0x33, 0xfa, 0xdf, 0x6c, 0xef, 0x4c, 0xfe, 0x9f, 0x16, 0x8a, 0xb9,
	0x9b, 0x27, 0x37, 0x0f, 0xb6, 0xdf, 0x15, 0xec, 0x46, 0x76, 0x27, 0xe7,
	0x17, 0x3e, 0x3a, 0x0c, 0x73, 0x6d, 0xf3, 0x19, 0xe1, 0x07, 0xd5, 0x7e,
	0xca, 0xce, 0xda, 0x9a, 0x4f, 0x8b, 0xf4, 0xaa, 0x36, 0xcd, 0xfb, 0xe2,
	0xbc, 0x87, 0xc9, 0xfc, 0xf6, 0x00, 0x79, 0x3c, 0x8f, 0xde, 0x14, 0xe4,
	0x1d, 0xfe, 0xcc, 0x3b, 0x59, 0xef, 0x62, 0x33, 0x1f, 0xbc, 0xf6, 0x2d,
	0x6f, 0xe6, 0x7d, 0xaa, 0xdd, 0xc3, 0x64, 0x7e, 0x7a, 0x90, 0x3d, 0xf3,
	0x66, 0x7e, 0x59, 0x62, 0xaf, 0x09, 0xf9, 0xd7, 0x1c, 0x49, 0xc5, 0x5f,
	0x71, 0xf8, 0x77, 0xdc, 0xfc, 0x0f, 0xd0, 0x3e, 0x49, 0xeb, 0xa1, 0xb6,
	0x82, 0xf0, 0x00, 0x00, 0x01, 0x85, 0x69, 0x43, 0x43, 0x50, 0x49, 0x43,
	0x43, 0x20, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x00, 0x00, 0x78,
	0x9c, 0x7d, 0x91, 0x3d, 0x48, 0xc3, 0x40, 0x1c, 0xc5, 0x5f, 0x53, 0xc5,
	0x2a, 0x15, 0x07, 0x33, 0x88, 0x38, 0x64, 0x68, 0x9d, 0x2c, 0x88, 0x8a,
	0x08, 0x2e, 0x52, 0xc5, 0x22, 0x58, 0x28, 0x6d, 0x85, 0x56, 0x1d, 0x4c,
	0x2e, 0xfd, 0x82, 0x26, 0x0d, 0x49, 0x8b, 0x8b, 0xa3, 0xe0, 0x5a, 0x70,
	0xf0, 0x63, 0xb1, 0xea, 0xe0, 0xe2, 0xac, 0xab, 0x83, 0xab, 0x20, 0x08,
	0x7e, 0x80, 0x38, 0x3a, 0x39, 0x29, 0xba, 0x48, 0x89, 0xff, 0x4b, 0x0a,
	0x2d, 0x62, 0x3c, 0x38, 0xee, 0xc7, 0xbb, 0x7b, 0x8f, 0xbb, 0x77, 0x80,
	0xd0, 0x28, 0x33, 0xcd, 0xea, 0x1a, 0x07, 0x34, 0xbd, 0x6a, 0x26, 0x63,
	0x51, 0x29, 0x93, 0x5d, 0x95, 0x7a, 0x5e, 0xd1, 0x8b, 0x00, 0x44, 0xcc,
	0x22, 0x2c, 0x33, 0xcb, 0x88, 0xa7, 0x16, 0xd3, 0xf0, 0x1c, 0x5f, 0xf7,
	0xf0, 0xf1, 0xf5, 0x2e, 0xc2, 0xb3, 0xbc, 0xcf, 0xfd, 0x39, 0xfa, 0xd5,
	0x9c, 0xc5, 0x00, 0x9f, 0x44, 0x3c, 0xc7, 0x0c, 0xb3, 0x4a, 0xbc, 0x41,
	0x3c, 0xbd, 0x59, 0x35, 0x38, 0xef, 0x13, 0x8b, 0xac, 0x28, 0xab, 0xc4,
	0xe7, 0xc4, 0x63, 0x26, 0x5d, 0x90, 0xf8, 0x91, 0xeb, 0x8a, 0xcb, 0x6f,
	0x9c, 0x0b, 0x0e, 0x0b, 0x3c, 0x53, 0x34, 0xd3, 0xc9, 0x79, 0x62, 0x91,
	0x58, 0x2a, 0x74, 0xb0, 0xd2, 0xc1, 0xac, 0x68, 0x6a, 0xc4, 0x53, 0xc4,
	0x21, 0x55, 0xd3, 0x29, 0x5f, 0xc8, 0xb8, 0xac, 0x72, 0xde, 0xe2, 0xac,
	0x95, 0x6b, 0xac, 0x75, 0x4f, 0xfe, 0xc2, 0x60, 0x4e, 0x5f, 0x49, 0x71,
	0x9d, 0xe6, 0x08, 0x62, 0x58, 0x42, 0x1c, 0x09, 0x48, 0x50, 0x50, 0x43,
	0x09, 0x65, 0x54, 0x11, 0xa1, 0x55, 0x27, 0xc5, 0x42, 0x92, 0xf6, 0xa3,
	0x1e, 0xfe, 0x61, 0xc7, 0x9f, 0x20, 0x97, 0x42, 0xae, 0x12, 0x18, 0x39,
	0x16, 0x50, 0x81, 0x06, 0xd9, 0xf1, 0x83, 0xff, 0xc1, 0xef, 0x6e, 0xad,
	0xfc, 0xe4, 0x84, 0x9b, 0x14, 0x8c, 0x02, 0xdd, 0x2f, 0xb6, 0xfd, 0x11,
	0x06, 0x7a, 0x76, 0x81, 0x66, 0xdd, 0xb6, 0xbf, 0x8f, 0x6d, 0xbb, 0x79,
	0x02, 0xf8, 0x9f, 0x81, 0x2b, 0xbd, 0xed, 0xaf, 0x34, 0x80, 0x99, 0x4f,
	0xd2, 0xeb, 0x6d, 0x2d, 0x74, 0x04, 0x0c, 0x6c, 0x03, 0x17, 0xd7, 0x6d,
	0x4d, 0xd9, 0x03, 0x2e, 0x77, 0x80, 0xa1, 0x27, 0x43, 0x36, 0x65, 0x47,
	0xf2, 0xd3, 0x14, 0xf2, 0x79, 0xe0, 0xfd, 0x8c, 0xbe, 0x29, 0x0b, 0x0c,
	0xde, 0x02, 0x7d, 0x6b, 0x6e, 0x6f, 0xad, 0x7d, 0x9c, 0x3e, 0x00, 0x69,
	0xea, 0x6a, 0xf9, 0x06, 0x38, 0x38, 0x04, 0x46, 0x0b, 0x94, 0xbd, 0xee,
	0xf1, 0xee, 0x40, 0x67, 0x6f, 0xff, 0x9e, 0x69, 0xf5, 0xf7, 0x03, 0xdb,
	0x6d, 0x72, 0xd1, 0x56, 0xb0, 0x77, 0xf2, 0x00, 0x00, 0x0d, 0x76, 0x69,
	0x54, 0x58, 0x74, 0x58, 0x4d, 0x4c, 0x3a, 0x63, 0x6f, 0x6d, 0x2e, 0x61,
	0x64, 0x6f, 0x62, 0x65, 0x2e, 0x78, 0x6d, 0x70, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x3c, 0x3f, 0x78, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x20, 0x62,
	0x65, 0x67, 0x69, 0x6e, 0x3d, 0x22, 0xef, 0xbb, 0xbf, 0x22, 0x20, 0x69,
	0x64, 0x3d, 0x22, 0x57, 0x35, 0x4d, 0x30, 0x4d, 0x70, 0x43, 0x65, 0x68,
	0x69, 0x48, 0x7a, 0x72, 0x65, 0x53, 0x7a, 0x4e, 0x54, 0x63, 0x7a, 0x6b,
	0x63, 0x39, 0x64, 0x22, 0x3f, 0x3e, 0x0a, 0x3c, 0x78, 0x3a, 0x78, 0x6d,
	0x70, 0x6d, 0x65, 0x74, 0x61, 0x20, 0x78, 0x6d, 0x6c, 0x6e, 0x73, 0x3a,
	0x78, 0x3d, 0x22, 0x61, 0x64, 0x6f, 0x62, 0x65, 0x3a, 0x6e, 0x73, 0x3a,
	0x6d, 0x65, 0x74, 0x61, 0x2f, 0x22, 0x20, 0x78, 0x3a, 0x78, 0x6d, 0x70,
	0x74, 0x6b, 0x3d, 0x22, 0x58, 0x4d, 0x50, 0x20, 0x43, 0x6f, 0x72, 0x65,
	0x20, 0x34, 0x2e, 0x34, 0x2e, 0x30, 0x2d, 0x45, 0x78, 0x69, 0x76, 0x32,
	0x22, 0x3e, 0x0a, 0x20, 0x3c, 0x72, 0x64, 0x66, 0x3a, 0x52, 0x44, 0x46,
	0x20, 0x78, 0x6d, 0x6c, 0x6e, 0x73, 0x3a, 0x72, 0x64, 0x66, 0x3d, 0x22,
	0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x77,
	0x33, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x31, 0x39, 0x39, 0x39, 0x2f, 0x30,
	0x32, 0x2f, 0x32, 0x32, 0x2d, 0x72, 0x64, 0x66, 0x2d, 0x73, 0x79, 0x6e,
	0x74, 0x61, 0x78, 0x2d, 0x6e, 0x73, 0x23, 0x22, 0x3e, 0x0a, 0x20, 0x20,
	0x3c, 0x72, 0x64, 0x66, 0x3a, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x72, 0x64, 0x66, 0x3a, 0x61, 0x62, 0x6f,
	0x75, 0x74, 0x3d, 0x22, 0x22, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x78, 0x6d,
	0x6c, 0x6e, 0x73, 0x3a, 0x78, 0x6d, 0x70, 0x4d, 0x4d, 0x3d, 0x22, 0x68,
	0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6e, 0x73, 0x2e, 0x61, 0x64, 0x6f,
	0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x61, 0x70, 0x2f, 0x31,
	0x2e, 0x30, 0x2f, 0x6d, 0x6d, 0x2f, 0x22, 0x0a, 0x20, 0x20, 0x20, 0x20,
	0x78, 0x6d, 0x6c, 0x6e, 0x73, 0x3a, 0x73, 0x74, 0x45, 0x76, 0x74, 0x3d,
	0x22, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6e, 0x73, 0x2e, 0x61,
	0x64, 0x6f, 0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x61, 0x70,
	0x2f, 0x31, 0x2e, 0x30, 0x2f, 0x73, 0x54, 0x79, 0x70, 0x65, 0x2f, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x23, 0x22, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x78, 0x6d, 0x6c, 0x6e, 0x73,
	0x3a, 0x64, 0x63, 0x3d, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f,
	0x70, 0x75, 0x72, 0x6c, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x64, 0x63, 0x2f,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x31, 0x2e, 0x31,
	0x2f, 0x22, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x78, 0x6d, 0x6c, 0x6e, 0x73,
	0x3a, 0x47, 0x49, 0x4d, 0x50, 0x3d, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3a,
	0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x69, 0x6d, 0x70, 0x2e, 0x6f,
	0x72, 0x67, 0x2f, 0x78, 0x6d, 0x70, 0x2f, 0x22, 0x0a, 0x20, 0x20, 0x20,
	0x20, 0x78, 0x6d, 0x6c, 0x6e, 0x73, 0x3a, 0x74, 0x69, 0x66, 0x66, 0x3d,
	0x22, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6e, 0x73, 0x2e, 0x61,
	0x64, 0x6f, 0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x66,
	0x66, 0x2f, 0x31, 0x2e, 0x30, 0x2f, 0x22, 0x0a, 0x20, 0x20, 0x20, 0x20,
	0x78, 0x6d, 0x6c, 0x6e, 0x73, 0x3a, 0x78, 0x6d, 0x70, 0x3d, 0x22, 0x68,
	0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6e, 0x73, 0x2e, 0x61, 0x64, 0x6f,
	0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x61, 0x70, 0x2f, 0x31,
	0x2e, 0x30, 0x2f, 0x22, 0x0a, 0x20, 0x20, 0x20, 0x78, 0x6d, 0x70, 0x4d,
	0x4d, 0x3a, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44,
	0x3d, 0x22, 0x67, 0x69, 0x6d, 0x70, 0x3a, 0x64, 0x6f, 0x63, 0x69, 0x64,
	0x3a, 0x67, 0x69, 0x6d, 0x70, 0x3a, 0x34, 0x35, 0x37, 0x31, 0x65, 0x34,
	0x34, 0x32, 0x2d, 0x65, 0x31, 0x31, 0x66, 0x2d, 0x34, 0x63, 0x66, 0x30,
	0x2d, 0x39, 0x65, 0x61, 0x35, 0x2d, 0x31, 0x38, 0x30, 0x36, 0x37, 0x36,
	0x66, 0x64, 0x30, 0x31, 0x37, 0x63, 0x22, 0x0a, 0x20, 0x20, 0x20, 0x78,
	0x6d, 0x70, 0x4d, 0x4d, 0x3a, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x44, 0x3d, 0x22, 0x78, 0x6d, 0x70, 0x2e, 0x69, 0x69, 0x64,
	0x3a, 0x63, 0x62, 0x61, 0x38, 0x32, 0x39, 0x33, 0x66, 0x2d, 0x34, 0x64,
	0x66, 0x31, 0x2d, 0x34, 0x38, 0x63, 0x31, 0x2d, 0x39, 0x64, 0x36, 0x63,
	0x2d, 0x63, 0x34, 0x30, 0x64, 0x61, 0x61, 0x34, 0x62, 0x61, 0x65, 0x62,
	0x39, 0x22, 0x0a, 0x20, 0x20, 0x20, 0x78, 0x6d, 0x70, 0x4d, 0x4d, 0x3a,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x3d, 0x22, 0x78, 0x6d, 0x70, 0x2e,
	0x64, 0x69, 0x64, 0x3a, 0x37, 0x39, 0x64, 0x63, 0x39, 0x65, 0x37, 0x39,
	0x2d, 0x65, 0x32, 0x33, 0x38, 0x2d, 0x34, 0x35, 0x38, 0x62, 0x2d, 0x38,
	0x32, 0x39, 0x36, 0x2d, 0x38, 0x38, 0x30, 0x61, 0x64, 0x39, 0x34, 0x65,
	0x31, 0x62, 0x39, 0x31, 0x22, 0x0a, 0x20, 0x20, 0x20, 0x64, 0x63, 0x3a,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x3d, 0x22, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x2f, 0x70, 0x6e, 0x67, 0x22, 0x0a, 0x20, 0x20, 0x20, 0x47, 0x49,
	0x4d, 0x50, 0x3a, 0x41, 0x50, 0x49, 0x3d, 0x22, 0x32, 0x2e, 0x30, 0x22,
	0x0a, 0x20, 0x20, 0x20, 0x47, 0x49, 0x4d, 0x50, 0x3a, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x3d, 0x22, 0x57, 0x69, 0x6e, 0x64, 0x6f,
	0x77, 0x73, 0x22, 0x0a, 0x20, 0x20, 0x20, 0x47, 0x49, 0x4d, 0x50, 0x3a,
	0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x6d, 0x70, 0x3d, 0x22, 0x31,
	0x36, 0x36, 0x32, 0x36, 0x37, 0x31, 0x32, 0x34, 0x38, 0x31, 0x34, 0x33,
	0x30, 0x38, 0x31, 0x22, 0x0a, 0x20, 0x20, 0x20, 0x47, 0x49, 0x4d, 0x50,
	0x3a, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3d, 0x22, 0x32, 0x2e,
	0x31, 0x30, 0x2e, 0x33, 0x32, 0x22, 0x0a, 0x20, 0x20, 0x20, 0x74, 0x69,
	0x66, 0x66, 0x3a, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x3d, 0x22, 0x31, 0x22, 0x0a, 0x20, 0x20, 0x20, 0x78, 0x6d,
	0x70, 0x3a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x54, 0x6f, 0x6f,
	0x6c, 0x3d, 0x22, 0x47, 0x49, 0x4d, 0x50, 0x20, 0x32, 0x2e, 0x31, 0x30,
	0x22, 0x0a, 0x20, 0x20, 0x20, 0x78, 0x6d, 0x70, 0x3a, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x44, 0x61, 0x74, 0x65, 0x3d, 0x22, 0x32,
	0x30, 0x32, 0x32, 0x3a, 0x30, 0x39, 0x3a, 0x30, 0x38, 0x54, 0x32, 0x33,
	0x3a, 0x30, 0x37, 0x3a, 0x32, 0x38, 0x2b, 0x30, 0x32, 0x3a, 0x30, 0x30,
	0x22, 0x0a, 0x20, 0x20, 0x20, 0x78, 0x6d, 0x70, 0x3a, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x44, 0x61, 0x74, 0x65, 0x3d, 0x22, 0x32, 0x30, 0x32,
	0x32, 0x3a, 0x30, 0x39, 0x3a, 0x30, 0x38, 0x54, 0x32, 0x33, 0x3a, 0x30,
	0x37, 0x3a, 0x32, 0x38, 0x2b, 0x30, 0x32, 0x3a, 0x30, 0x30, 0x22, 0x3e,
	0x0a, 0x20, 0x20, 0x20, 0x3c, 0x78, 0x6d, 0x70, 0x4d, 0x4d, 0x3a, 0x48,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20,
	0x3c, 0x72, 0x64, 0x66, 0x3a, 0x53, 0x65, 0x71, 0x3e, 0x0a, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x3c, 0x72, 0x64, 0x66, 0x3a, 0x6c, 0x69, 0x0a, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x73, 0x74, 0x45, 0x76, 0x74, 0x3a, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3d, 0x22, 0x73, 0x61, 0x76, 0x65, 0x64,
	0x22, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x73, 0x74, 0x45, 0x76,
	0x74, 0x3a, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x3d, 0x22, 0x2f,
	0x22, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x73, 0x74, 0x45, 0x76,
	0x74, 0x3a, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x44,
	0x3d, 0x22, 0x78, 0x6d, 0x70, 0x2e, 0x69, 0x69, 0x64, 0x3a, 0x66, 0x30,
	0x65, 0x34, 0x38, 0x64, 0x64, 0x35, 0x2d, 0x33, 0x37, 0x66, 0x64, 0x2d,
	0x34, 0x30, 0x30, 0x34, 0x2d, 0x38, 0x36, 0x36, 0x64, 0x2d, 0x35, 0x32,
	0x33, 0x36, 0x63, 0x35, 0x61, 0x32, 0x32, 0x33, 0x32, 0x61, 0x22, 0x0a,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x73, 0x74, 0x45, 0x76, 0x74, 0x3a,
	0x73, 0x6f, 0x66, 0x74, 0x77, 0x61, 0x72, 0x65, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x3d, 0x22, 0x47, 0x69, 0x6d, 0x70, 0x20, 0x32, 0x2e, 0x31, 0x30,
	0x20, 0x28, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x73, 0x29, 0x22, 0x0a,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x73, 0x74, 0x45, 0x76, 0x74, 0x3a,
	0x77, 0x68, 0x65, 0x6e, 0x3d, 0x22, 0x32, 0x30, 0x32, 0x32, 0x2d, 0x30,
	0x39, 0x2d, 0x30, 0x38, 0x54, 0x32, 0x33, 0x3a, 0x30, 0x37, 0x3a, 0x32,
	0x38, 0x22, 0x2f, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x72,
	0x64, 0x66, 0x3a, 0x53, 0x65, 0x71, 0x3e, 0x0a, 0x20, 0x20, 0x20, 0x3c,
	0x2f, 0x78, 0x6d, 0x70, 0x4d, 0x4d, 0x3a, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x3e, 0x0a, 0x20, 0x20, 0x3c, 0x2f, 0x72, 0x64, 0x66, 0x3a,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3e,
	0x0a, 0x20, 0x3c, 0x2f, 0x72, 0x64, 0x66, 0x3a, 0x52, 0x44, 0x46, 0x3e,
	0x0a, 0x3c, 0x2f, 0x78, 0x3a, 0x78, 0x6d, 0x70, 0x6d, 0x65, 0x74, 0x61,
	0x3e, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x0a,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x0a, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x0a,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x0a, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x0a, 0x3c, 0x3f,
	0x78, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x20, 0x65, 0x6e, 0x64, 0x3d,
	0x22, 0x77, 0x22, 0x3f, 0x3e, 0x3f, 0xf9, 0x2f, 0x0a, 0x00, 0x00, 0x00,
	0x06, 0x62, 0x4b, 0x47, 0x44, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf9,
	0x43, 0xbb, 0x7f, 0x00, 0x00, 0x00, 0x09, 0x70, 0x48, 0x59, 0x73, 0x00,
	0x00, 0x2e, 0x23, 0x00, 0x00, 0x2e, 0x23, 0x01, 0x78, 0xa5, 0x3f, 0x76,
	0x00, 0x00, 0x00, 0x07, 0x74, 0x49, 0x4d, 0x45, 0x07, 0xe6, 0x09, 0x08,
	0x15, 0x07, 0x1c, 0x85, 0xc0, 0xb3, 0xb2, 0x00, 0x00, 0x00, 0x19, 0x74,
	0x45, 0x58, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x00, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20,
	0x47, 0x49, 0x4d, 0x50, 0x57, 0x81, 0x0e, 0x17, 0x00, 0x00, 0x02, 0x24,
	0x49, 0x44, 0x41, 0x54, 0x58, 0xc3, 0xed, 0xd7, 0x4f, 0x88, 0x8d, 0x51,
	0x18, 0x06, 0xf0, 0xdf, 0xd5, 0x5d, 0x21, 0xa2, 0xb9, 0x58, 0x18, 0x84,
	0x14, 0x0b, 0x59, 0x48, 0xd4, 0x50, 0x16, 0x16, 0x16, 0x92, 0x1d, 0x36,
	0xb3, 0x61, 0xd2, 0x6c, 0x84, 0x46, 0x62, 0x63, 0xb2, 0x53, 0xb3, 0x98,
	0x92, 0xa2, 0xac, 0x88, 0x95, 0x9a, 0x14, 0x65, 0x4a, 0x93, 0x26, 0x59,
	0x90, 0x3f, 0x89, 0x21, 0x21, 0xd9, 0x68, 0x46, 0xa2, 0xc6, 0x48, 0xdc,
	0xb9, 0x36, 0xe7, 0xd6, 0xe9, 0xeb, 0x7c, 0x33, 0x77, 0xee, 0x9d, 0x7b,
	0x6d, 0x3c, 0x75, 0x3a, 0xe7, 0xbe, 0xe7, 0xfd, 0xde, 0xf3, 0x9c, 0x3f,
	0xcf, 0x7b, 0xce, 0x2d, 0x68, 0x2e, 0xe6, 0x61, 0x1b, 0x56, 0xe3, 0x17,
	0x1e, 0xe1, 0x95, 0x16, 0x60, 0x0e, 0xce, 0x61, 0x1c, 0x95, 0x4c, 0x79,
	0x89, 0x5d, 0xcd, 0x26, 0x70, 0x31, 0x31, 0x70, 0x5c, 0x26, 0xb1, 0xa7,
	0x59, 0x83, 0xaf, 0x8a, 0x06, 0x2a, 0xa3, 0x13, 0x45, 0x2c, 0xc5, 0x95,
	0xa8, 0xef, 0x7d, 0xb3, 0x08, 0x1c, 0x88, 0x06, 0x19, 0x4c, 0xf4, 0x0f,
	0xe1, 0x71, 0x28, 0x4b, 0x8a, 0x38, 0x83, 0x7d, 0x35, 0x04, 0x2e, 0x63,
	0x04, 0xaf, 0x31, 0x10, 0xea, 0x14, 0x46, 0xa2, 0xf6, 0x8a, 0x44, 0xff,
	0xce, 0xac, 0xe1, 0xd2, 0x34, 0xfb, 0x95, 0x57, 0x06, 0xb1, 0x36, 0x87,
	0xc4, 0xf5, 0x8c, 0x5f, 0x27, 0x76, 0xe7, 0x10, 0xaa, 0x9b, 0x40, 0x25,
	0x9c, 0xf2, 0xed, 0x39, 0x24, 0xba, 0x31, 0x9a, 0xf1, 0xef, 0x4e, 0xc9,
	0xa5, 0x51, 0x9d, 0xdf, 0x44, 0x5b, 0x26, 0xe6, 0x79, 0xf4, 0xa1, 0x14,
	0x6c, 0x93, 0xf8, 0x89, 0x3f, 0xd9, 0x00, 0xc5, 0x9c, 0xc0, 0xd7, 0x30,
	0x9c, 0xb1, 0x6d, 0xc0, 0xc1, 0x28, 0x68, 0x15, 0xa5, 0x70, 0x8e, 0x8e,
	0x85, 0xdf, 0xa7, 0xd0, 0x13, 0xda, 0x9f, 0xc2, 0xf2, 0x0f, 0x4d, 0x35,
	0x8b, 0xd4, 0x16, 0x1c, 0x9e, 0x62, 0xc6, 0xf7, 0x13, 0xfe, 0xa3, 0x91,
	0xcf, 0xd3, 0xc8, 0x7e, 0xbc, 0x96, 0x8c, 0x35, 0x13, 0xfc, 0x40, 0x57,
	0xc2, 0x5e, 0xc2, 0xfa, 0x44, 0xcc, 0x85, 0x09, 0xdf, 0xb7, 0x11, 0xc1,
	0x43, 0xc5, 0x3a, 0xf6, 0xfd, 0x0d, 0xbe, 0x27, 0x82, 0x2f, 0x08, 0xf5,
	0x03, 0x6c, 0x0c, 0xed, 0x9e, 0x20, 0xdf, 0x3b, 0x58, 0x84, 0x23, 0x91,
	0x72, 0xca, 0xb8, 0x3b, 0xd3, 0x2d, 0xa8, 0x62, 0x2c, 0xf1, 0xcd, 0xd6,
	0xd0, 0xb7, 0x38, 0x64, 0xb9, 0xe9, 0x14, 0x74, 0xb4, 0x5e, 0x15, 0x2c,
	0xcf, 0x9c, 0xfa, 0x2a, 0x26, 0x42, 0xfd, 0x15, 0x9b, 0xd0, 0x1f, 0x56,
	0x2a, 0x8b, 0xe7, 0x21, 0xf1, 0xf5, 0xd7, 0x73, 0x08, 0xe7, 0xe2, 0x76,
	0xc2, 0xff, 0x1b, 0x0a, 0x09, 0xff, 0x42, 0x20, 0xd3, 0x81, 0x2d, 0x68,
	0xaf, 0x55, 0x86, 0x3b, 0x12, 0x01, 0xd7, 0x84, 0x3c, 0xdf, 0x9e, 0xf0,
	0xbf, 0x11, 0x88, 0x64, 0x51, 0xc1, 0xb3, 0xe9, 0x96, 0xb4, 0x91, 0x4c,
	0x58, 0x09, 0x4b, 0xbe, 0xac, 0x91, 0x87, 0x43, 0x23, 0x18, 0xc7, 0x5e,
	0x7c, 0xfe, 0x17, 0x04, 0x06, 0x82, 0xdc, 0x86, 0x1b, 0x99, 0x41, 0xde,
	0x19, 0x78, 0x11, 0xd2, 0x68, 0x15, 0x1d, 0x91, 0xce, 0xab, 0xb8, 0x87,
	0x0f, 0xb3, 0xf1, 0x80, 0xa8, 0x45, 0x05, 0x27, 0x13, 0x3e, 0x63, 0x98,
	0x3f, 0x1b, 0x8f, 0xc7, 0x5a, 0x70, 0x21, 0x0c, 0x18, 0xa3, 0x2d, 0xba,
	0x74, 0x9a, 0x4e, 0x60, 0x02, 0xbd, 0x09, 0xfb, 0x89, 0xc4, 0xed, 0xd8,
	0xb4, 0x43, 0x78, 0x19, 0xef, 0x12, 0xb7, 0xe3, 0xd9, 0x56, 0x11, 0xf8,
	0x9d, 0xb3, 0x0a, 0x5d, 0x58, 0xd7, 0x2a, 0x19, 0x5e, 0x0d, 0x7f, 0x2c,
	0xb2, 0x4a, 0xea, 0x6d, 0x65, 0x1e, 0x38, 0x9d, 0xb0, 0xed, 0xc7, 0xe6,
	0x56, 0x11, 0xb8, 0x85, 0x87, 0x09, 0x7b, 0x5f, 0x3d, 0x04, 0x0a, 0x58,
	0x99, 0xb8, 0x5e, 0x3f, 0xe2, 0xcb, 0x14, 0xdf, 0x95, 0x72, 0x9e, 0xd8,
	0x4f, 0x72, 0x2e, 0xa5, 0xff, 0xc8, 0xc5, 0x5f, 0xd0, 0xd6, 0xcd, 0x71,
	0x56, 0x00, 0x89, 0x16, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44,
	0xae, 0x42, 0x60, 0x82,
}