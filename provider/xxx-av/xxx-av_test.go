package xxx_av

import (
	"testing"

	"github.com/HDuoDuo/metatube/provider/internal/testkit"
)

func TestXXXAV_GetMovieInfoByID(t *testing.T) {
	testkit.Test(t, New, []string{
		"24719",
		"23395",
		"19337",
	})
}
