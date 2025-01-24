package badge

import (
	"bytes"
	_ "embed"

	"github.com/jellydator/ttlcache/v3"

	"github.com/HDuoDuo/metatube/imageutil"
)

//go:embed zimu.png
var zimu []byte

func init() {
	badge, _, _ := imageutil.Decode(bytes.NewReader(zimu))
	badgeCache.Set("zimu.png", badge, ttlcache.NoTTL)
}
