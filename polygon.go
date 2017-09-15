package geom

// Polygon is a geometry consisting of multiple closed LineStrings.
// There must be only one exterior LineString with a clockwise winding order.
// There may be one or more interior LineStrings with a counterclockwise winding orders.
type Polygon [][][2]float64

// SubLineStrings returns the coordinates of the lineStrings
func (p *Polygon) SubLineStrings() [][][2]float64 {
	return *p
}
