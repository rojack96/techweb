package entities

const (
	Geofence  = "geofence"
	Sightings = "sightings"
	Topology  = "topology"
	Users     = "users"
)

type Region struct {
	GID       uint32   `db:"gid"`
	CodRip    *float64 `db:"cod_rip"`
	CodReg    *float64 `db:"cod_reg"`
	DenReg    *string  `db:"den_reg"`
	ShapeLeng *float64 `db:"shape_leng"`
	ShapeArea *float64 `db:"shape_area"`
	Geom      []byte   `db:"geom"`
	GeomWGS84 []byte   `db:"geom_wgs84"`
}
