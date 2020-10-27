package survivalserver

import (
	"testing"

	"github.com/ldsec/medco/connector/util"
	utilserver "github.com/ldsec/medco/connector/util/server"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func init() {
	utilserver.SetForTesting()
}

func TestBuildTimePoint(t *testing.T) {
	utilserver.TestI2B2DBConnection(t)
	timePoints, err := buildTimePoints(bigList,
		`A168`,
		"@", `A125`, "126:1", 2000)
	if err != nil {
		t.Error("Test Failed", err)
	}
	assert.Equal(t, bigTimePoints, timePoints)
}

var smallPatientList = []int64{483573, 483801}
var bigList = []int64{1000001364, 1000001364, 1000001363, 1000001363, 1000001362, 1000001362, 1000001361, 1000001361, 1000001360, 1000001360, 1000001359, 1000001359, 1000001358, 1000001358, 1000001357, 1000001357, 1000001356, 1000001356, 1000001355, 1000001355, 1000001354, 1000001354, 1000001353, 1000001353, 1000001352, 1000001352, 1000001351, 1000001351, 1000001350, 1000001350, 1000001349, 1000001349, 1000001348, 1000001348, 1000001347, 1000001347, 1000001346, 1000001346, 1000001345, 1000001345, 1000001344, 1000001344, 1000001343, 1000001343, 1000001342, 1000001342, 1000001341, 1000001341, 1000001340, 1000001340, 1000001339, 1000001339, 1000001338, 1000001338, 1000001337, 1000001337, 1000001336, 1000001336, 1000001335, 1000001335, 1000001334, 1000001334, 1000001333, 1000001333, 1000001332, 1000001332, 1000001331, 1000001331, 1000001330, 1000001330, 1000001329, 1000001329, 1000001328, 1000001328, 1000001327, 1000001327, 1000001326, 1000001326, 1000001325, 1000001325, 1000001324, 1000001324, 1000001323, 1000001323, 1000001322, 1000001322, 1000001321, 1000001321, 1000001320, 1000001320, 1000001319, 1000001319, 1000001318, 1000001318, 1000001317, 1000001317, 1000001316, 1000001316, 1000001315, 1000001315, 1000001314, 1000001314, 1000001313, 1000001313, 1000001312, 1000001312, 1000001311, 1000001311, 1000001310, 1000001310, 1000001309, 1000001309, 1000001308, 1000001308, 1000001307, 1000001307, 1000001306, 1000001306, 1000001305, 1000001305, 1000001304, 1000001304, 1000001303, 1000001303, 1000001302, 1000001302, 1000001301, 1000001301, 1000001300, 1000001300, 1000001299, 1000001299, 1000001298, 1000001298, 1000001297, 1000001297, 1000001296, 1000001296, 1000001295, 1000001295, 1000001294, 1000001294, 1000001293, 1000001293, 1000001292, 1000001292, 1000001291, 1000001291, 1000001290, 1000001290, 1000001289, 1000001289, 1000001288, 1000001288, 1000001287, 1000001287, 1000001286, 1000001286, 1000001285, 1000001285, 1000001284, 1000001284, 1000001283, 1000001283, 1000001282, 1000001282, 1000001281, 1000001281, 1000001280, 1000001280, 1000001279, 1000001279, 1000001278, 1000001278, 1000001277, 1000001277, 1000001276, 1000001276, 1000001275, 1000001275, 1000001274, 1000001274, 1000001273, 1000001273, 1000001272, 1000001272, 1000001271, 1000001271, 1000001270, 1000001270, 1000001269, 1000001269, 1000001268, 1000001268, 1000001267, 1000001267, 1000001266, 1000001266, 1000001265, 1000001265, 1000001264, 1000001264, 1000001263, 1000001263, 1000001262, 1000001262, 1000001261, 1000001261, 1000001260, 1000001260, 1000001259, 1000001259, 1000001258, 1000001258, 1000001257, 1000001257, 1000001256, 1000001256, 1000001255, 1000001255, 1000001254, 1000001254, 1000001253, 1000001253, 1000001252, 1000001252, 1000001251, 1000001251, 1000001250, 1000001250, 1000001249, 1000001249, 1000001248, 1000001248, 1000001247, 1000001247, 1000001246, 1000001246, 1000001245, 1000001245, 1000001244, 1000001244, 1000001243, 1000001243, 1000001242, 1000001242, 1000001241, 1000001241, 1000001240, 1000001240, 1000001239, 1000001239, 1000001238, 1000001238, 1000001237, 1000001237, 1000001236, 1000001236, 1000001235, 1000001235, 1000001234, 1000001234, 1000001233, 1000001233, 1000001232, 1000001232, 1000001231, 1000001231, 1000001230, 1000001230, 1000001229, 1000001229, 1000001228, 1000001228, 1000001227, 1000001227, 1000001226, 1000001226, 1000001225, 1000001225, 1000001224, 1000001224, 1000001223, 1000001223, 1000001222, 1000001222, 1000001221, 1000001221, 1000001220, 1000001220, 1000001219, 1000001219, 1000001218, 1000001218, 1000001217, 1000001217, 1000001216, 1000001216, 1000001215, 1000001215, 1000001214, 1000001214, 1000001213, 1000001213, 1000001212, 1000001212, 1000001211, 1000001211, 1000001210, 1000001210, 1000001209, 1000001209, 1000001208, 1000001208, 1000001207, 1000001207, 1000001206, 1000001206, 1000001205, 1000001205, 1000001204, 1000001204, 1000001203, 1000001203, 1000001202, 1000001202, 1000001201, 1000001201, 1000001200, 1000001200, 1000001199, 1000001199, 1000001198, 1000001198, 1000001197, 1000001197, 1000001196, 1000001196, 1000001195, 1000001195, 1000001194, 1000001194, 1000001193, 1000001193, 1000001192, 1000001192, 1000001191, 1000001191, 1000001190, 1000001190, 1000001189, 1000001189, 1000001188, 1000001188, 1000001187, 1000001187, 1000001186, 1000001186, 1000001185, 1000001185, 1000001184, 1000001184, 1000001183, 1000001183, 1000001182, 1000001182, 1000001181, 1000001181, 1000001180, 1000001180, 1000001179, 1000001179, 1000001178, 1000001178, 1000001177, 1000001177, 1000001176, 1000001176, 1000001175, 1000001175, 1000001174, 1000001174, 1000001173, 1000001173, 1000001172, 1000001172, 1000001171, 1000001171, 1000001170, 1000001170, 1000001169, 1000001169, 1000001168, 1000001168, 1000001167, 1000001167, 1000001166, 1000001166, 1000001165, 1000001165, 1000001164, 1000001164, 1000001163, 1000001163, 1000001162, 1000001162, 1000001161, 1000001161, 1000001160, 1000001160, 1000001159, 1000001159, 1000001158, 1000001158, 1000001157, 1000001157, 1000001156, 1000001156, 1000001155, 1000001155, 1000001154, 1000001154, 1000001153, 1000001153, 1000001152, 1000001152, 1000001151, 1000001151, 1000001150, 1000001150, 1000001149, 1000001149, 1000001148, 1000001148, 1000001147, 1000001147, 1000001146, 1000001146, 1000001145, 1000001145, 1000001144, 1000001144, 1000001143, 1000001143, 1000001142, 1000001142, 1000001141, 1000001141, 1000001140, 1000001140, 1000001139, 1000001139, 1000001138, 1000001138, 1000001137, 1000001137}
var bigTimePoints = util.TimePointsFromTable([][]int{{5, 1, 0}, {11, 3, 0}, {12, 1, 0}, {13, 2, 0}, {15, 1, 0}, {26, 1, 0}, {30, 1, 0}, {31, 1, 0}, {53, 2, 0}, {54, 1, 0}, {59, 1, 0}, {60, 2, 0}, {61, 1, 0}, {62, 1, 0}, {65, 2, 0}, {71, 1, 0}, {79, 1, 0}, {81, 2, 0}, {88, 2, 0}, {92, 1, 1}, {93, 1, 0}, {95, 2, 0}, {105, 1, 1}, {107, 2, 0}, {110, 1, 0}, {116, 1, 0}, {118, 1, 0}, {122, 1, 0}, {131, 1, 0}, {132, 2, 0}, {135, 1, 0}, {142, 1, 0}, {144, 1, 0}, {145, 2, 0}, {147, 1, 0}, {153, 1, 0}, {156, 2, 0}, {163, 3, 0}, {166, 2, 0}, {167, 1, 0}, {170, 1, 0}, {173, 0, 1}, {174, 0, 1}, {175, 1, 1}, {176, 1, 0}, {177, 1, 1}, {179, 2, 0}, {180, 1, 0}, {181, 2, 0}, {182, 1, 0}, {183, 1, 0}, {185, 0, 1}, {186, 1, 0}, {188, 0, 1}, {189, 1, 0}, {191, 0, 1}, {192, 0, 1}, {194, 1, 0}, {196, 0, 1}, {197, 1, 1}, {199, 1, 0}, {201, 2, 0}, {202, 1, 1}, {203, 0, 1}, {207, 1, 0}, {208, 1, 0}, {210, 1, 0}, {211, 0, 1}, {212, 1, 0}, {218, 1, 0}, {221, 0, 1}, {222, 1, 1}, {223, 1, 0}, {224, 0, 1}, {225, 0, 2}, {226, 1, 0}, {229, 1, 0}, {230, 1, 0}, {235, 0, 1}, {237, 0, 1}, {239, 2, 0}, {240, 0, 1}, {243, 0, 1}, {245, 1, 0}, {246, 1, 0}, {252, 0, 1}, {259, 0, 1}, {266, 0, 1}, {267, 1, 0}, {268, 1, 0}, {269, 1, 1}, {270, 1, 0}, {272, 0, 1}, {276, 0, 1}, {279, 0, 1}, {283, 1, 0}, {284, 1, 1}, {285, 2, 0}, {286, 1, 0}, {288, 1, 0}, {291, 1, 0}, {292, 0, 2}, {293, 1, 0}, {296, 0, 1}, {300, 0, 1}, {301, 1, 1}, {303, 1, 1}, {305, 1, 0}, {306, 1, 0}, {310, 2, 0}, {315, 0, 1}, {320, 1, 0}, {329, 1, 0}, {332, 0, 1}, {337, 1, 0}, {340, 1, 0}, {345, 1, 0}, {348, 1, 0}, {350, 1, 0}, {351, 1, 0}, {353, 2, 0}, {356, 0, 1}, {361, 1, 0}, {363, 2, 0}, {364, 1, 1}, {371, 2, 0}, {376, 0, 1}, {382, 0, 1}, {384, 0, 1}, {387, 1, 0}, {390, 1, 0}, {394, 1, 0}, {404, 0, 1}, {413, 0, 1}, {426, 1, 0}, {428, 1, 0}, {429, 1, 0}, {433, 1, 0}, {442, 1, 0}, {444, 1, 1}, {450, 1, 0}, {455, 1, 0}, {457, 1, 0}, {458, 0, 1}, {460, 1, 0}, {473, 1, 0}, {477, 1, 0}, {511, 0, 2}, {519, 1, 0}, {520, 1, 0}, {524, 2, 0}, {529, 0, 1}, {533, 1, 0}, {543, 0, 1}, {550, 1, 0}, {551, 0, 1}, {558, 1, 0}, {559, 0, 1}, {567, 1, 0}, {574, 1, 0}, {583, 1, 0}, {588, 0, 1}, {613, 1, 0}, {624, 1, 0}, {641, 1, 0}, {643, 1, 0}, {654, 1, 0}, {655, 1, 0}, {687, 1, 0}, {689, 1, 0}, {705, 1, 0}, {707, 1, 0}, {728, 1, 0}, {731, 1, 0}, {735, 1, 0}, {740, 0, 1}, {765, 1, 0}, {791, 1, 0}, {806, 0, 1}, {814, 1, 0}, {821, 0, 1}, {840, 0, 1}, {883, 1, 0}, {965, 0, 1}, {1010, 0, 1}, {1022, 0, 1}})
