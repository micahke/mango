package galaxymap

import (
	"container/list"
	"math"
	"math/rand"
	"time"

	// "github.com/micahke/infinite-universe/mango/logging"
	"github.com/micahke/infinite-universe/mango/util"
)

// import glm "github.com/go-gl/mathgl/mgl32"

// RegionManager is a struct that manages the regions of the galaxy map.


type RegionManager struct {

  regionSize [2]int

}


type Region struct {

  coords [2]int
  
  color util.Color

}


var region_colors []util.Color = []util.Color{
  util.NewColorRGBi(48, 42, 47),
  util.ELECTRON_BLUE,
  util.PINK_GLAMOUR,
}


const COLOR_BASELINE int = 40


var regionManager *RegionManager = &RegionManager{
  regionSize: [2]int{100, 100},
}


// Runs every frame  
func (manager *RegionManager) Update() {

  // Figure out what the x and y coordinates are 
  // currentTile := tilemap.tilePositions[0]

  // logging.Log(currentTile)

}

var regionCache *RegionCache = NewRegionCache(15*time.Second, 100000)


type RegionCache struct {
    maxAge time.Duration
    size   int
    list   *list.List
    cache  map[[2]int]*list.Element
}

func NewRegionCache(maxAge time.Duration, size int) *RegionCache {
    return &RegionCache{
        maxAge: maxAge,
        size:   size,
        list:   list.New(),
        cache:  make(map[[2]int]*list.Element),
    }
}

func (c *RegionCache) Get(x, y int) *Region {
    if elem, ok := c.cache[[2]int{x, y}]; ok {
        c.list.MoveToBack(elem)
        return elem.Value.(*Region)
    }
    return nil
}

func (c *RegionCache) Put(x, y int, region *Region) {
    // Remove the least recently used pair if the cache is full
    if c.list.Len() >= c.size {
        c.evict()
    }
    // Add the new pair to the tail of the list and the map
    elem := c.list.PushBack(region)
    c.cache[[2]int{x, y}] = elem
}

func (c *RegionCache) evict() {
    // Remove the head of the list (which is the least recently used pair) and its corresponding key-value pair from the map
    elem := c.list.Front()
    delete(c.cache, elem.Value.(*Region).coords)
    c.list.Remove(elem)
}

func (manager *RegionManager) CalculateRegionForTile(x, y int) *Region {
    if region := regionCache.Get(x, y); region != nil {
        return region
    }
    region := new(Region)
    region.coords[0] = int(math.Floor(float64(x) / float64(manager.regionSize[0])))
    region.coords[1] = int(math.Floor(float64(y) / float64(manager.regionSize[1])))
    seed := (region.coords[0]&0xFFFF)<<16 | (region.coords[1] & 0xFFFF)
    rand.Seed(int64(seed))
    r := COLOR_BASELINE + rand.Intn(40)
    g := COLOR_BASELINE + rand.Intn(40)
    b := COLOR_BASELINE + rand.Intn(40)
    region.color = util.NewColorRGBi(r, g, b)
    regionCache.Put(x, y, region)
    return region
}

