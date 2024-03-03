package model

type SupportedGADShapes string

const (
	SupportedGADShapesPoint                    SupportedGADShapes = "POINT"
	SupportedGADShapesPointUncertaintyCircle   SupportedGADShapes = "POINT_UNCERTAINTY_CIRCLE"
	SupportedGADShapesPointUncertaintyEllipse  SupportedGADShapes = "POINT_UNCERTAINTY_ELLIPSE"
	SupportedGADShapesPOLYGON                  SupportedGADShapes = "POLYGON"
	SupportedGADShapesPointAltitude            SupportedGADShapes = "POINT_ALTITUDE"
	SupportedGADShapesPointAltitudeUncertainty SupportedGADShapes = "POINT_ALTITUDE_UNCERTAINTY"
	SupportedGADShapesEllipsoidArc             SupportedGADShapes = "ELLIPSOID_ARC"
)

type GeographicalCoordinates struct {
	Lon float64
	Lat float64
}

type Point struct {
	Shape SupportedGADShapes
	Point GeographicalCoordinates
}

type CivicAddress struct {
	Country string
	A1      string
	A2      string
	A3      string
	A4      string
	A5      string
	A6      string
	PRD     string
	POD     string
}
