package geo

import (
	"errors"
	"bytes"
	"encoding/binary"
)

var (
	ErrInvalidWkb = errors.New("geo: invalid WKB")
	ErrUnsupportedDataType = errors.New("geo: scan value must be []byte")
	ErrWrongGeometryType = errors.New("geo: wrong geometry type")
)

type Point [2]float64

type Polygon struct {
	LinearRings []LinearRing
}

type LinearRing []Point

func (p *Polygon) SetLinearRings(rings []LinearRing) *Polygon {
	p.LinearRings = rings
	return p
}

func (p *Polygon) AddLinearRing(ring LinearRing) *Polygon {
	p.LinearRings = append(p.LinearRings, ring)
	return p
}

func (p *Polygon) ToGeoJson() (*Feature) {
	geometry := Geometry{Type: "Polygon", Coordinates: p.LinearRings}
	properties := make([]Property, 0)
	return &Feature{Type: "Feature", Geometry: geometry, Properties: properties}
}

func (p *Polygon) Scan(value interface{}) (err error) {
	data, ok := value.([]byte)
	if !ok {
		return ErrUnsupportedDataType
	}

	if len(data) == 0 {
		return
	}

	if len(data) < 9 {
		return ErrInvalidWkb
	}

	buffer := bytes.NewBuffer(data)
	var temp uint32

	// First get the byte order. 0 for BigEndian, 1 for LittleEndian
	first, err := buffer.ReadByte()

	if err != nil {
		return err
	}

	var order binary.ByteOrder = binary.BigEndian
	if (first == 1) {
		order = binary.LittleEndian
	}

	binary.Read(buffer, order, &temp)

	typeCode := int(temp)
	if typeCode != 3 {
		return ErrWrongGeometryType
	}

	// Read the number of rings.
	binary.Read(buffer, order, &temp)
	ringCount := int(temp)

	for j := 0; j < ringCount; j++ {

		binary.Read(buffer, order, &temp)
		length := int(temp)

		points := make([]Point, length)
		for i := 0; i < length; i++ {
			binary.Read(buffer, order, &points[i][0])
			binary.Read(buffer, order, &points[i][1])
		}

		p.AddLinearRing(points)
	}

	return
}