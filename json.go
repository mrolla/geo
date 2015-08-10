package geo

type FeatureCollection struct {
	Type     string `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type     string `json:"type"`
	Geometry Geometry `json:"geometry"`
}

type Geometry struct {
	Type        string `json:"type"`
	Coordinates interface{} `json:"coordinates"`
}

func NewFeatureCollection() (*FeatureCollection) {
	return &FeatureCollection{Type: "FeatureCollection"}
}

func (f *FeatureCollection) Add(feature *Feature) {
	f.Features = append(f.Features, *feature)
}