package models

type CabalVersion struct {
	IsLatest    bool
	GreaterThan string
	LessThan    string
}

type CabalDependency struct {
	Name    string
	Version CabalVersion
}

type CabalPackage struct {
	Dependencies []CabalDependency
}
