package tokyohot

import (
	"testing"

	"github.com/HDuoDuo/metatube/provider/internal/testkit"
)

func TestTokyoHot_GetMovieInfoByID(t *testing.T) {
	testkit.Test(t, New, []string{
		"s2mbd-002",
		"n1633",
		"n1624",
		"kb1624",
	})
}

func TestTokyoHot_SearchMovie(t *testing.T) {
	testkit.Test(t, New, []string{
		"1624",
		"n0238",
	})
}
