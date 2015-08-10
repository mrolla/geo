package geo
import "testing"

var testPolygon = []byte{0x01, 0x03, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0xE0, 0x21, 0x7C, 0x4B, 0xE0, 0x36, 0x22, 0x40, 0xD8, 0x5F, 0x3B, 0x5A, 0x0D, 0xC1, 0x46, 0x40}

func TestScanPolygon(t *testing.T) {

	p := &Polygon{}

	if err := p.Scan(testPolygon); err != nil {
		t.Errorf("incorrect scan, got %v", err)
	}

	if len(p.LinearRings) != 1 {
		t.Errorf("incorrect scan result, expecting 1 ring, got %d", len(p.LinearRings))
	}

	feature := p.ToGeoJson()

	if feature.Geometry.Type != "Polygon" {
		t.Errorf("incorrect feature type, expecting 'Polygon', received '%s'", feature.Type)
	}

	rings, ok := feature.Geometry.Coordinates.([]LinearRing)

	if !ok {
		t.Fail()
	}

	if len(rings) != 1 {
		t.Fail()
	}
}
