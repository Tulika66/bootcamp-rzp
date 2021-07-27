package Cache

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)
const (
	frequencyChecker=1
)

var CacheLocal *cache.Cache
func Init(){
	CacheLocal:= cache.New(time.Duration(int64(frequencyChecker))*time.Minute,cache.NoExpiration)
	fmt.Println("Cache setup Done. Set values to it for better use. cache:=",CacheLocal)
}


