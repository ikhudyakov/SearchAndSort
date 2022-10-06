package binarysearch

import (
	"testing"
)

type binarySearchTest struct {
	arg      int
	expected bool
}

var binarySearchTests = []binarySearchTest{
	{47500, true},
	{72557, true},
	{30045, false},
	{-123, false},
}

var arr []int = []int{83600, 5336, 65699, 59656, 95739, 20097, 24041, 45562, 21780, 28665, 57533, 2184, 27350, 20225, 43524,
	1974, 84738, 47840, 32920, 73326, 57289, 4325, 98189, 50156, 5062, 57340, 38282, 24879, 37907, 18607, 31676, 63955,
	3901, 1719, 99096, 33312, 2015, 92988, 95534, 16897, 19376, 85526, 90023, 84855, 25490, 74136, 81028, 93391, 42578,
	17447, 21557, 97708, 43869, 33987, 28458, 8700, 51168, 41770, 92336, 68740, 8039, 75175, 36490, 8047, 93889, 92179,
	87472, 68420, 23439, 33813, 47190, 38005, 77078, 90036, 50506, 3606, 65790, 8780, 81035, 15649, 24719, 54262, 71194,
	13411, 15470, 65216, 18132, 10845, 69045, 45372, 90146, 49247, 43639, 20946, 62849, 47199, 35859, 88982, 11233, 23976,
	8340, 58843, 53338, 70706, 62864, 40233, 31786, 94969, 57745, 45680, 19352, 22543, 88696, 26456, 59121, 9366, 26744,
	26000, 90518, 58081, 51099, 44513, 15185, 3019, 59125, 64610, 83641, 62684, 88082, 73376, 23998, 6815, 83708, 25242,
	90508, 99161, 38831, 61469, 29192, 76472, 47500, 47839, 61203, 87030, 58590, 38150, 65681, 4973, 32467, 15104, 3705,
	83068, 63314, 41424, 49013, 40994, 9289, 72891, 36715, 86068, 99132, 215, 20560, 64309, 45658, 68722, 17588, 68704,
	55231, 95762, 75023, 75224, 73302, 6208, 65517, 25719, 51786, 9599, 84272, 34147, 54232, 85660, 38178, 48874, 57413,
	65467, 4241, 17582, 19277, 76284, 82712, 10774, 92003, 50131, 80711, 98607, 30044, 24527, 50066, 81436, 43084, 22579,
	47218, 73791, 76821, 18209, 5721, 226, 31735, 11029, 67974, 75263, 55975, 43816, 58845, 86171, 54706, 2855, 83690,
	30146, 17222, 56219, 31243, 33556, 54300, 26149, 53194, 86657, 46736, 49876, 72557, 46318, 35151, 56281, 15240, 69103,
	88460, 91362, 97495, 21287, 93094, 32879, 75889, 44920, 40002, 17223, 23209, 46949, 74860, 56374, 97145, 47358, 95336,
	70597, 77739, 20198, 11052, 28728, 87072, 6633, 41365, 7837, 12780, 73616, 19664, 79567, 50971, 57297, 44463, 59419,
	84895, 47309, 96858, 59767, 92565, 90944, 90883, 13612, 52725, 8111, 86004, 47901, 71036, 17866, 53655, 7976, 21002,
	43726, 28750, 8033, 87495, 64311, 9143, 11026, 65099, 43240, 65673, 5658, 7930, 78786, 1}

func TestBinarySearch(t *testing.T) {

	for _, test := range binarySearchTests {
		if output := BinarySearch(arr, test.arg); output != test.expected {
			t.Errorf("Output %v not equal to expected %v", output, test.expected)
		}
	}
}
