package myprivate

// ExportGetMessageFunc is an unboud func to getMessage()
var ExportGetMessageFunc = (*Secret).getMessage

// ExportGetMessage is a proxy func to getMessage()
func (s *Secret) ExportGetMessage() string {
	return s.getMessage()
}
