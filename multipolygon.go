package geogoth

// MultiPolygon ...
type MultiPolygon struct {
	Coords [][][][]float64
}

// NewMultiPolygon creates MultiPolygon
func NewMultiPolygon(coords [][][][]float64) MultiPolygon {
	return MultiPolygon{
		Coords: coords,
	}
}

// Coordinates returns array of longitude, latitude of the MultiPolygon
func (m MultiPolygon) Coordinates() interface{} {
	return m.Coords
}

// GetCoordinates returns array of longitude, latitude of Three-dimensional arrays (MultiLineString, Polygon)
// coordnum - index of coordinate arr
func (m MultiPolygon) GetCoordinates(setsnum, setnum, coordnum int) (float64, float64) {
	coords := (m.Coordinates()).([][][][]float64)

	lon := coords[setsnum][setnum][coordnum][0]
	lat := coords[setsnum][setnum][coordnum][1]

	return lon, lat // longitude (Y), latitude (X)
}

// Type returns type of the MultiPolygon (MultiPolygon)
func (m MultiPolygon) Type() string {
	return "MultiPolygon"
}

// Length returns length of the MultiPolygon
func (m MultiPolygon) Length() float64 {
	return MultipolygonLength(m)
}

// DistanceTo returns distance between two geo objects
func (m MultiPolygon) DistanceTo(f Feature) float64 {

	var distance float64

	switch f.Type() {
	case "Point":
		point := f.(*Point)
		distance = DistancePointMultiPolygon(point, &m)

	case "MultiPoint":
		mpoint := f.(*MultiPoint)
		distance = DistanceMultiPointMultiPolygon(mpoint, &m)

	case "LineString":
		ln := f.(*LineString)
		distance = DistanceLineStringMultiPolygon(ln, &m)

	case "MultiLineString":
		mlstr := f.(*MultiLineString)
		distance = DistanceMultiLineStringMultiPolygon(mlstr, &m)

	case "Polygon":
		plgn := f.(*Polygon)
		distance = DistancePolygonMultiPolygon(plgn, &m)

	case "MultiPolygon":
		mpolyg := f.(*MultiPolygon)
		distance = DistanceMultiPolygonMultiPolygon(&m, mpolyg)
	}

	return distance
}

// IntersectsWith returns true if geoObject intersects with Feature
func (m MultiPolygon) IntersectsWith(f Feature) bool {
	var intersection bool

	switch f.Type() {
	case "Point":
		// point := f.(*Point)

	case "MultiPoint":
		// mpoint := f.(*MultiPoint)

	case "LineString":
		// lstr := f.(*LineString)

	case "MultiLineString":
		// mlinestr := f.(*MultiLineString)

	case "Polygon":
		// polygon := f.(*Polygon)

	case "MultiPolygon":
		// mpolyg := f.(*MultiPolygon)
	}

	return intersection

}
