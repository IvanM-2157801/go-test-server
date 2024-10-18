package model

type Planet struct {
	Climate        *string      `json:"climate,omitempty"`
	Diameter       *string      `json:"diameter,omitempty"`
	Gravity        *string      `json:"gravity,omitempty"`
	Name           *string      `json:"name,omitempty"`
	OrbitalPeriod  *string      `json:"orbital_period,omitempty"`
	Population     *string      `json:"population,omitempty"`
	ResidentURLs   []*string    `json:"residents,omitempty"`
	Residents      []*Character `json:"-"`
	RotationPeriod *string      `json:"rotation_period,omitempty"`
	SurfaceWater   *string      `json:"surface_water,omitempty"`
	Terrain        *string      `json:"terrain,omitempty"`
	URL            *string      `json:"url,omitempty"`
	ID             *string      `json:"id,omitempty"`
}
