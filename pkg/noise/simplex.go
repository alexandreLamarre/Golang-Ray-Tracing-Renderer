package noise

import (
	"math"
)

// Implements the Open Simplex version of the simplex noise algorithm
// Code adapted from the Open Simplex 2 Algorithm courtesy of github user KdotJPG
// Source code at : https://github.com/KdotJPG/OpenSimplex2/blob/master/java/OpenSimplex2F.java#L190

var LOOKUP3d *lattice = newLattice()
var PMASK int = 2047
var PSIZE int = 2048
var N3 float64 = 0.030485933181293584
var GRADIENTS3d [][3]float64 = newGradient3D()

func Simplex(x, y, z float64, seed int64) float64 {
	r := 0.6 * (x + y + z) // use r = 0.5 for avoiding discontiniuties otherwise use r = 0.6 for a more smooth feel
	xr := r - x
	yr := r - y
	zr := r - z

	return Simplex3Noise(xr, yr, zr, seed)
}

func Simplex3Noise(xr, yr, zr float64, seed int64) float64 {
	perm, permGrad := newHashes(seed)
	xrb := math.Floor(xr)
	yrb := math.Floor(yr)
	zrb := math.Floor(zr)

	xri := xr - xrb
	yri := yr - yrb
	zri := zr - zrb

	xht := int(xri + 0.5)
	yht := int(yri + 0.5)
	zht := int(zri + 0.5)
	index := (xht << 0) | (yht << 1) | (zht << 2)

	value := 0.0
	c := LOOKUP3d[index]
	for c != nil {
		dxr := xri + c.dxr
		dyr := yri + c.dyr
		dzr := zri + c.dzr
		attn := 0.5 - dxr*dxr - dyr*dyr - dzr*dzr
		if attn < 0 {
			c = c.nextOnFailure
		} else {
			pxm := int(xrb) + c.xrv&PMASK
			pym := (int(yrb) + c.yrv) & PMASK
			pzm := (int(zrb) + c.zrv) & PMASK
			// TODO : implemet gradient/ permutation function
			grad := permGrad[perm[perm[pxm]^pym]^pzm]
			extrapolation := grad[0]*dxr + grad[1]*dyr + grad[2]*dzr
			attn *= attn
			value += attn * attn * extrapolation
			c = c.nextOnSuccess
		}
	}
	return value
}

type lattice [8]*latticePoint

type latticePoint struct {
	nextOnFailure *latticePoint
	nextOnSuccess *latticePoint
	dxr, dyr, dzr float64
	xrv, yrv, zrv int
}

func newLatticePoint(xrv, yrv, zrv, lattice int) *latticePoint {
	dxr := float64(-xrv) + float64(lattice)*0.5
	dyr := float64(-yrv) + float64(lattice)*0.5
	dzr := float64(-zrv) + float64(lattice)*0.5
	xrv = xrv + lattice*1024
	yrv = yrv + lattice*1024
	zrv = zrv + lattice*1024
	return &latticePoint{dxr: dxr, dyr: dyr, dzr: dzr, xrv: int(xrv), yrv: int(yrv), zrv: int(zrv)}
}

func newLattice() *lattice {
	lat := &lattice{}
	for i := 0; i < 8; i++ {
		i1 := (i >> 0) & 1
		j1 := (i >> 1) & 1
		k1 := (i >> 2) & 1
		i2 := i1 ^ 1
		j2 := j1 ^ 1
		k2 := k1 ^ 1

		// The two points within this octant, one from each of the two cubic half-lattices.
		c0 := newLatticePoint(i1, j1, k1, 0)
		c1 := newLatticePoint(i1+i2, j1+j2, k1+k2, 1)

		// Each single step away on the first half-lattice.
		c2 := newLatticePoint(i1^1, j1, k1, 0)
		c3 := newLatticePoint(i1, j1^1, k1, 0)
		c4 := newLatticePoint(i1, j1, k1^1, 0)

		// Each single step away on the second half-lattice.
		c5 := newLatticePoint(i1+(i2^1), j1+j2, k1+k2, 1)
		c6 := newLatticePoint(i1+i2, j1+(j2^1), k1+k2, 1)
		c7 := newLatticePoint(i1+i2, j1+j2, k1+(k2^1), 1)

		// First two are guaranteed.
		c0.nextOnSuccess = c1
		c0.nextOnFailure = c0.nextOnSuccess
		c1.nextOnSuccess = c2
		c1.nextOnFailure = c1.nextOnSuccess

		// Once we find one on the first half-lattice, the rest are out.
		// In addition, knowing c2 rules out c5.
		c2.nextOnFailure = c3
		c2.nextOnSuccess = c6
		c3.nextOnFailure = c4
		c3.nextOnSuccess = c5
		c4.nextOnSuccess = c5
		c4.nextOnFailure = c4.nextOnSuccess

		// Once we find one on the second half-lattice, the rest are out.
		c5.nextOnFailure = c6
		c5.nextOnSuccess = nil
		c6.nextOnFailure = c7
		c6.nextOnSuccess = nil
		c7.nextOnSuccess = nil
		c7.nextOnFailure = c7.nextOnSuccess

		lat[i] = c0
	}
	return lat
}

func newHashes(seed int64) ([]int, [][3]float64) {
	perm := make([]int, PSIZE, PSIZE)
	permGrad := make([][3]float64, PSIZE, PSIZE)
	source := make([]int, PSIZE, PSIZE)

	for i := PSIZE - 1; i >= 0; i-- {
		source[i] = i
	}

	for i := PSIZE - 1; i >= 0; i-- {
		seed = seed*int64(6364136223846793005) + int64(1442695040888963407)
		r := int(seed+int64(31)) % (i + 1)
		if r < 0 {
			r += i + 1
		}
		perm[i] = source[r]
		permGrad[i] = GRADIENTS3d[perm[i]]
		source[r] = source[i]
	}

	return perm, permGrad
}

func newGradient3D() [][3]float64 {
	res := make([][3]float64, PSIZE, PSIZE)
	grad3 := [][3]float64{
		{-2.22474487139, -2.22474487139, -1.0},
		{-2.22474487139, -2.22474487139, 1.0},
		{-3.0862664687972017, -1.1721513422464978, 0.0},
		{-1.1721513422464978, -3.0862664687972017, 0.0},
		{-2.22474487139, -1.0, -2.22474487139},
		{-2.22474487139, 1.0, -2.22474487139},
		{-1.1721513422464978, 0.0, -3.0862664687972017},
		{-3.0862664687972017, 0.0, -1.1721513422464978},
		{-2.22474487139, -1.0, 2.22474487139},
		{-2.22474487139, 1.0, 2.22474487139},
		{-3.0862664687972017, 0.0, 1.1721513422464978},
		{-1.1721513422464978, 0.0, 3.0862664687972017},
		{-2.22474487139, 2.22474487139, -1.0},
		{-2.22474487139, 2.22474487139, 1.0},
		{-1.1721513422464978, 3.0862664687972017, 0.0},
		{-3.0862664687972017, 1.1721513422464978, 0.0},
		{-1.0, -2.22474487139, -2.22474487139},
		{1.0, -2.22474487139, -2.22474487139},
		{0.0, -3.0862664687972017, -1.1721513422464978},
		{0.0, -1.1721513422464978, -3.0862664687972017},
		{-1.0, -2.22474487139, 2.22474487139},
		{1.0, -2.22474487139, 2.22474487139},
		{0.0, -1.1721513422464978, 3.0862664687972017},
		{0.0, -3.0862664687972017, 1.1721513422464978},
		{-1.0, 2.22474487139, -2.22474487139},
		{1.0, 2.22474487139, -2.22474487139},
		{0.0, 1.1721513422464978, -3.0862664687972017},
		{0.0, 3.0862664687972017, -1.1721513422464978},
		{-1.0, 2.22474487139, 2.22474487139},
		{1.0, 2.22474487139, 2.22474487139},
		{0.0, 3.0862664687972017, 1.1721513422464978},
		{0.0, 1.1721513422464978, 3.0862664687972017},
		{2.22474487139, -2.22474487139, -1.0},
		{2.22474487139, -2.22474487139, 1.0},
		{1.1721513422464978, -3.0862664687972017, 0.0},
		{3.0862664687972017, -1.1721513422464978, 0.0},
		{2.22474487139, -1.0, -2.22474487139},
		{2.22474487139, 1.0, -2.22474487139},
		{3.0862664687972017, 0.0, -1.1721513422464978},
		{1.1721513422464978, 0.0, -3.0862664687972017},
		{2.22474487139, -1.0, 2.22474487139},
		{2.22474487139, 1.0, 2.22474487139},
		{1.1721513422464978, 0.0, 3.0862664687972017},
		{3.0862664687972017, 0.0, 1.1721513422464978},
		{2.22474487139, 2.22474487139, -1.0},
		{2.22474487139, 2.22474487139, 1.0},
		{3.0862664687972017, 1.1721513422464978, 0.0},
		{1.1721513422464978, 3.0862664687972017, 0.0},
	}
	for i := 0; i < len(grad3); i++ {
		grad3[i][0] /= N3
		grad3[i][1] /= N3
		grad3[i][2] /= N3
	}
	for i := 0; i < PSIZE; i++ {
		res[i] = grad3[i%len(grad3)]
	}
	return res
}
