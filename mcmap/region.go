package mcmap

import (
	"errors"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

var (
	NotAvailable = errors.New("Chunk or Superchunk not available")
)

type superchunk struct {
	preChunks map[XZPos]*preChunk
	chunks    map[XZPos]*Chunk
	modified  bool
}

type Region struct {
	path             string
	superchunksAvail map[XZPos]bool

	superchunks map[XZPos]*superchunk
}

var mcaRegex = regexp.MustCompile(`^r\.([0-9-]+)\.([0-9-]+)\.mca$`)

// OpenRegion opens a region directory.
func OpenRegion(path string) (*Region, error) {
	rv := &Region{
		path:             path,
		superchunksAvail: make(map[XZPos]bool),
		superchunks:      make(map[XZPos]*superchunk),
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if !fi.IsDir() {
		return nil, fmt.Errorf("%s is not a directory", path)
	}

	names, err := f.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	for _, name := range names {
		match := mcaRegex.FindStringSubmatch(name)
		if len(match) == 3 {
			// We ignore the error here. The Regexp already ensures that the inputs are numbers.
			x, _ := strconv.ParseInt(match[1], 10, 32)
			z, _ := strconv.ParseInt(match[2], 10, 32)

			rv.superchunksAvail[XZPos{int(x), int(z)}] = true
		}
	}

	return rv, nil
}

// MaxDims calculates the approximate maximum x, z dimensions of this region in number of chunks. The actual maximum dimensions might be a bit smaller.
func (reg *Region) MaxDims() (xmin, xmax, zmin, zmax int) {
	if len(reg.superchunksAvail) == 0 {
		return 0, 0, 0, 0
	}

	xmin = math.MaxInt32
	zmin = math.MaxInt32
	xmax = math.MinInt32
	zmax = math.MinInt32

	for pos := range reg.superchunksAvail {
		if pos.X < xmin {
			xmin = pos.X
		}
		if pos.Z < zmin {
			zmin = pos.Z
		}
		if pos.X > xmax {
			xmax = pos.X
		}
		if pos.Z > zmax {
			zmax = pos.Z
		}
	}

	xmax++
	zmax++
	xmin *= 16
	xmax *= 16
	zmin *= 16
	zmax *= 16
	return
}

func chunkToSuperchunk(cx, cz int) (scx, scz, rx, rz int) {
	scx = cx >> 5
	scz = cz >> 5
	rx = ((cx % 32) + 32) % 32
	rz = ((cz % 32) + 32) % 32
	return
}

func superchunkToChunk(scx, scz, rx, rz int) (cx, cz int) {
	cx = scx*32 + rx
	cz = scz*32 + rz
	return
}

func (reg *Region) loadSuperchunk(pos XZPos) error {
	if !reg.superchunksAvail[pos] {
		return NotAvailable
	}
	fname := fmt.Sprintf("%s%cr.%d.%d.mca", reg.path, os.PathSeparator, pos.X, pos.Z)

	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	pcs, err := readRegionFile(f)
	if err != nil {
		return err
	}

	reg.superchunks[pos] = &superchunk{
		preChunks: pcs,
		chunks:    make(map[XZPos]*Chunk),
	}
	return nil
}

func (reg *Region) cleanSuperchunks() error {
	del := make(map[XZPos]bool)

	for scPos, sc := range reg.superchunks {
		if len(sc.chunks) > 0 {
			return nil
		}

		if sc.modified {
			// TODO: Save superchunk to region file
		}

		del[scPos] = true
	}

	for scPos, _ := range del {
		delete(reg.superchunks, scPos)
	}

	return nil
}

func (reg *Region) Chunk(x, z int) (*Chunk, error) {
	scx, scz, cx, cz := chunkToSuperchunk(x, z)
	scPos := XZPos{scx, scz}
	cPos := XZPos{cx, cz}

	sc, ok := reg.superchunks[scPos]
	if !ok {
		if err := reg.loadSuperchunk(scPos); err != nil {
			return nil, err
		}
		sc = reg.superchunks[scPos]
	}

	chunk, ok := sc.chunks[cPos]
	if !ok {
		pc, ok := sc.preChunks[cPos]
		if !ok {
			return nil, NotAvailable
		}

		var err error
		if chunk, err = pc.toChunk(); err != nil {
			return nil, err
		}
		sc.chunks[cPos] = chunk
	}

	if err := reg.cleanSuperchunks(); err != nil {
		return nil, err
	}

	return chunk, nil
}

// UnloadChunk marks a chunk as unused. If all chunks of a superchunk are marked as unused, the superchunk will be unloaded and saved (if needed).
func (reg *Region) UnloadChunk(x, z int) {
	scx, scz, cx, cz := chunkToSuperchunk(x, z)
	scPos := XZPos{scx, scz}
	cPos := XZPos{cx, cz}

	sc, ok := reg.superchunks[scPos]
	if !ok {
		return
	}

	chunk, ok := sc.chunks[cPos]
	if !ok {
		return
	}

	if chunk.modified {
		// TODO: Save to prechunks

		chunk.modified = false
		sc.modified = true
	}

	delete(sc.chunks, cPos)
}

// AllChunks returns a channel that will give you the positions of all possibly available chunks in an efficient order.
//
// Note the "possibly available", you still have to check, if the chunk could actually be loaded.
func (reg *Region) AllChunks() <-chan XZPos {
	ch := make(chan XZPos)
	go func(ch chan<- XZPos) {
		for spos, _ := range reg.superchunksAvail {
			scx, scz := spos.X, spos.Z
			for rx := 0; rx < 16; rx++ {
				for rz := 0; rz < 16; rz++ {
					cx, cz := superchunkToChunk(scx, scz, rx, rz)
					ch <- XZPos{cx, cz}
				}
			}
		}
		close(ch)
	}(ch)

	return ch
}
