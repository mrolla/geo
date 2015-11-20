package geo

type FeatureCollection struct {
	Type     string `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string `json:"type"`
	Geometry   Geometry `json:"geometry"`
	Properties []Property `json:"properties"`
}

type Geometry struct {
	Type        string `json:"type"`
	Coordinates interface{} `json:"coordinates"`
}

type Property struct {
}

func NewFeatureCollection() (*FeatureCollection) {
	features := make([]Feature, 0)
	return &FeatureCollection{Type: "FeatureCollection", Features: features}
}

func (f *FeatureCollection) Add(feature *Feature) {
	f.Features = append(f.Features, *feature)
}