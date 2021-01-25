package main

import (
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

type Curve struct {
	CachedSize   int       `json:"-"`
	CachedPoints []*Vec3   `json:"-"`
	Inx          uint32 `json:"-"`
	TotalDur     uint32
	RawPoints    string
	CurveLen     float64        `json:"-"`
	BezierPoints []*BezierPoint `json:"-"`
}

type BezierPoint struct {
	Point   *Vec3
	Handle1 *Vec3
	Handle2 *Vec3
}

/*ApproximateLength(p1, p2) {
// console.log( p1, p2 );
}

GetPoint(p1, p2, t) {
t = this.clamp(0, 1, t);

var t1 = Math.pow(1 - t, 3);
var t2 = 3 * Math.pow(1 - t, 2) * t;
var t3 = 3 * (1 - t) * Math.pow(t, 2);
var t4 = Math.pow(t, 3);

var l1 = operateArrays([t1, t1, t1], [ p1.position.x, p1.position.y, p1.position.z ], multiply);
var l2 = operateArrays([t2, t2, t2], [ p1.handle2.x, p1.handle2.y, p1.handle2.z ], multiply);
var l3 = operateArrays([t3, t3, t3], [ p2.handle1.x, p2.handle1.y, p2.handle1.z ], multiply);
var l4 = operateArrays([t4, t4, t4], [ p2.position.x, p2.position.y, p2.position.z ], multiply);

var p1 = operateArrays(l1, l2, sum);
var p2 = operateArrays(p1, l3, sum);
var p3 = operateArrays(p2, l4, sum);

return p3;
}
*/
const resolution = 8

func (c *Curve) Length() (l float64) {
	for i := 0; i < len(c.BezierPoints)-1; i++ {
		l += c.PointsDist(c.BezierPoints[i], c.BezierPoints[i+1])
	}

	return
}


func (c *Curve) GetPosAtProp(prop float64) *Vec3 {
	var firstPoint *BezierPoint

	secondPoint := &BezierPoint{}
	distance := prop * c.CurveLen

	if distance < 0 {
		x := *c.BezierPoints[0].Point
		return &x
	}
	totalLength := 0.
	var curveLength float64
	for i := 0; i < len(c.BezierPoints) - 1; i++ {

		curveLength = c.PointsDist(c.BezierPoints[i], c.BezierPoints[i + 1])// this.ApproximateLength(points[i], points[i + 1]);
		if totalLength + curveLength >= distance {
		firstPoint = c.BezierPoints[i];
		secondPoint = c.BezierPoints[i + 1];
		break;
		} else {
			totalLength += curveLength;
		}
	}

	if firstPoint == nil  {
		firstPoint = c.BezierPoints[len(c.BezierPoints)- 1];
		secondPoint = c.BezierPoints[0];

		curveLength = c.PointsDist(firstPoint, secondPoint)// this.ApproximateLength();
	}

	distance -= totalLength;

	position := c.GetPoint(firstPoint, secondPoint, distance / curveLength);
	return position
}

func (c *Curve) PointsDist(p1 *BezierPoint, p2 *BezierPoint) (total float64) {
	lastPosition := p1.Point

	for i := 0; i < resolution+1; i++ {
		currentPosition := c.GetPoint(p1, p2, float64(i)/resolution)
		total += currentPosition.Sub(lastPosition).Length()
		lastPosition = currentPosition
	}

	return
}

func clamp(min, max, val float64) float64 {
	return math.Min(math.Max(val, min), max)
}

func (v *Vec3) Sub(vec3s *Vec3) *Vec3 {
	return &Vec3{v[0] - vec3s[0], v[1] - vec3s[1], v[2] - vec3s[2]}
}

func (v *Vec3) Muls(s float64) *Vec3 {
	return &Vec3{v[0] * s, v[1] * s, v[2] * s}
}

func (v *Vec3) Mul(vv *Vec3) *Vec3 {
	return &Vec3{v[0] * vv[0], v[1] * vv[1], v[2] * vv[2]}
}

func (v *Vec3) Sum(vv *Vec3) *Vec3 {
	return &Vec3{v[0] + vv[0], v[1] + vv[1], v[2] + vv[2]}
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

func (v *Vec3) DistSq(vec3s *Vec3) float64 {
	return (v[0] - vec3s[0])*(v[0] - vec3s[0]) +
		(v[1] - vec3s[1])*(v[1] - vec3s[1]) +
		(v[2] - vec3s[2])*(v[2] - vec3s[2])
}

func (c *Curve) GetPoint(p1, p2 *BezierPoint, t float64) *Vec3 {
	t = clamp(0, 1, t)

	return p1.Point.Muls(math.Pow(1-t, 3)).Sum(
		p1.Handle2.Muls(3 * math.Pow(1-t, 2) * t)).Sum(
		p2.Handle1.Muls(3 * (1 - t) * math.Pow(t, 2))).Sum(
		p2.Point.Muls(math.Pow(t, 3)))
}

func (c *Curve) GetCachedPoint(t float64)  *Vec3 {
	if t >= 1 {
		return c.CachedPoints[c.CachedSize - 1]
	}
	return c.CachedPoints[int(math.Floor(t * float64(c.CachedSize)))]
}


func (c *Curve) Parse() (err error) {
	data := strings.Split(c.RawPoints, " ")
	var p1 float64
	for i := 0; i < len(data)/9; i++ {
		b := BezierPoint{
			Point:   &Vec3{},
			Handle1: &Vec3{},
			Handle2: &Vec3{},
		}

		for j := 0; j < 3; j++ {
			p1, err = strconv.ParseFloat(data[i*9+j], 32)
			if err != nil {
				return
			}
			b.Point[j%3] = p1
		}

		for j := 3; j < 6; j++ {
			p1, err = strconv.ParseFloat(data[i*9+j], 32)
			if err != nil {
				return
			}
			b.Handle1[j%3] = p1
		}

		for j := 6; j < 9; j++ {
			p1, err = strconv.ParseFloat(data[i*9+j], 32)
			if err != nil {
				return
			}
			b.Handle2[j%3] = p1
		}

		c.BezierPoints = append(c.BezierPoints, &b)
	}
	c.CurveLen = c.Length()

	return
}
const CACHED_SIZE = 1000

func (c *Curve) Cache() []*Vec3 {
	start := time.Now()

	for i := 0; i < CACHED_SIZE; i++ {
		c.CachedPoints = append(c.CachedPoints, c.GetPosAtProp(float64(i) / float64(c.CachedSize)))
	}
	c.CachedSize = len(c.CachedPoints)

	log.Printf("Calculating caching points %s", time.Since(start))
	return c.CachedPoints
}

