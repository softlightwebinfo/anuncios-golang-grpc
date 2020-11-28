package structs

import "cientosdeanuncios.com/backend/enums"

type Error struct {
	Message string
	Code    enums.EnumCode
	Error   error
}
