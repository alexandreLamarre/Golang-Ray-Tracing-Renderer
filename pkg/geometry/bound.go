package geometry

import (
	"github.com/alexandreLamarre/Golang-Ray-Tracing-Renderer/pkg/algebra"
	"math"
)

//Bounds handles AABB(axis-aligned bounding boxes) ray intersects to speed up complex Shape Groups rendering
type Bounds struct{
	maximum *algebra.Vector
	minimum *algebra.Vector
}

//NewBounds Initializer for Bounds
func NewBounds (min, max *algebra.Vector) *Bounds{
	return &Bounds{minimum: min, maximum: max}
}

//GetBoundsTransform takes a min,max bounding points of a cube and applies the transform to its vertices
// Then returns a bounding box of the new min and max
func GetBoundsTransform(min, max *algebra.Vector, transform *algebra.Matrix) *Bounds{
	minX := min.Get()[0]; minY := min.Get()[1]; minZ := min.Get()[2]
	maxX := max.Get()[0]; maxY := max.Get()[1]; maxZ := max.Get()[2]
	distX := maxX - minX; distY := maxY - minY; distZ := maxZ - minZ
	vertices := []*algebra.Vector{
		min,
		algebra.NewPoint(minX + distX, minY, minZ),
		algebra.NewPoint(minX, minY + distY, minZ),
		algebra.NewPoint(minX, minY, minZ + distZ),
		max,
		algebra.NewPoint(maxX - distX, maxY, maxZ),
		algebra.NewPoint(maxX, maxY - distY, maxZ),
		algebra.NewPoint(maxX, maxY, maxZ - distZ),
	}
	var newMinX float64 = math.Inf(1); var newMinY float64 = math.Inf(1); var newMinZ float64 = math.Inf(1)
	var newMaxX float64 = math.Inf(-1); var newMaxY float64 = math.Inf(-1); var newMaxZ float64 = math.Inf(-1)
	for i := 0; i < len(vertices); i++{

		newVertex := transform.MultiplyByVec(vertices[i])
		newX := newVertex.Get()[0]; newY := newVertex.Get()[1]; newZ := newVertex.Get()[2]
		newMinX = math.Min(newMinX, newX); newMinY = math.Min(newMinY, newY); newMinZ = math.Min(newMinZ, newZ)
		newMaxX = math.Max(newMaxX, newX); newMaxY = math.Max(newMaxY, newY); newMaxZ = math.Max(newMaxZ, newZ)
	}
	return NewBounds(algebra.NewPoint(newMinX, newMinY, newMinZ),
		           algebra.NewPoint(newMaxX, newMaxY, newMaxZ))
}

//Intersect Checks ray intersection with a bounding box
func (b *Bounds) Intersect(ray *algebra.Ray) bool{
	origin := ray.Get()["origin"]; direction := ray.Get()["direction"]
	minX := b.minimum.Get()[0]; minY := b.minimum.Get()[1] ; minZ := b.minimum.Get()[2];
	maxX := b.maximum.Get()[0]; maxY := b.maximum.Get()[1]; maxZ := b.maximum.Get()[2];

	xtmin, xtmax := checkCustomAxis(origin.Get()[0], direction.Get()[0], minX, maxX)
	ytmin, ytmax := checkCustomAxis(origin.Get()[1], direction.Get()[1], minY, maxY)
	ztmin, ztmax := checkCustomAxis(origin.Get()[2], direction.Get()[2], minZ, maxZ)

	tmin := max(xtmin, ytmin, ztmin)
	tmax := min(xtmax, ytmax, ztmax)

	if tmin > tmax{
		return false
	}
	return true
}

// helpers

func checkCustomAxis(origin, direction, min, max float64) (float64, float64){
	tminNumerator := min -origin
	tmaxNumerator := max - origin

	var tmin float64
	var tmax float64
	EPSILON := 0.0001
	if math.Abs(direction) >= EPSILON {
		tmin = tminNumerator/direction
		tmax = tmaxNumerator/direction
	} else {
		tmin = math.Abs(tminNumerator) * math.Inf(-1)
		tmax = tmaxNumerator * math.Inf(1)
	}

	if tmin > tmax {
		temp := tmin
		tmin = tmax
		tmax = temp
	}

	return tmin, tmax
}

func max(values ...float64) float64{
	if len(values) == 1{
		return values[0]
	}
	maxVal := math.Inf(-1)
	for i:= 0; i < len(values); i++{
		maxVal = math.Max(values[i], maxVal)
	}
	return maxVal
}

func min(values ...float64) float64{
	if len(values) == 1{
		return values[0]
	}
	minVal := math.Inf(1)
	for i:= 0; i < len(values); i++{
		minVal = math.Min(values[i], minVal)
	}
	return minVal
}