package geogoth

import (
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestNewMultiPoint(t *testing.T) {
	assert := assert.New(t)

	mpoint := &Objects.Object3
	CreateObj3()

	assert.True(mpoint.Coords[0][0] == 37.570792444518)
	assert.True(mpoint.Coords[2][1] == 55.77134814581)
}

func TestGetCoordinatesMPoint(t *testing.T) {
	assert := assert.New(t)

	mpoint := &Objects.Object3
	CreateObj3()

	coords := (mpoint.GetCoordinates()).([][]float64)

	assert.True(coords[1][0] == 37.624941407105)
	assert.True(coords[1][1] == 55.738048529295)
}

func TestGetTypeMPoint(t *testing.T) {
	assert := assert.New(t)

	mpoint := &Objects.Object3
	CreateObj3()

	assert.True(mpoint.GetType() == "MultiPoint")
}

func TestGetLengthMPoint(t *testing.T) {
	assert := assert.New(t)

	mpoint := &Objects.Object3
	CreateObj3()

	assert.True(mpoint.GetLength() == 0)
}

func TestDistanceToMPoint(t *testing.T) {
	assert := assert.New(t)

	mpoint := &Objects.Object3
	CreateObj3()

	// MultiPoint - Point
	point := &Objects.Object1
	CreateObj1()

	assert.True(int(mpoint.DistanceTo(point)) == 1619)
}