package Num_65_IsNumber

type status int
type bType int

const (
	ByteTypeNumber bType = iota
	ByteTypeSymbol
	ByteTypePoint
	ByteTypeE
	ByteTypeSpace
	ByteTypeUnknown
)

const (
	statusBegin status = iota
	statusIntSymbol
	statusIntNumber
	statusPointWithNumber
	statusPointWithoutNumber
	statusDecimalIntNumber
	statusE
	statusESymbol
	statusENumber
	statusEnd
)

var StatusMapping = map[status]map[bType]status{
	statusBegin: {
		ByteTypeSpace:  statusBegin,
		ByteTypeNumber: statusIntNumber,
		ByteTypeSymbol: statusIntSymbol,
		ByteTypePoint:  statusPointWithoutNumber,
	},
	statusIntSymbol: {
		ByteTypeNumber: statusIntNumber,
		ByteTypePoint:  statusPointWithoutNumber,
	},
	statusPointWithNumber: {
		ByteTypeSpace:  statusEnd,
		ByteTypeNumber: statusDecimalIntNumber,
		ByteTypeE:      statusE,
	},
	statusPointWithoutNumber: {
		ByteTypeNumber: statusDecimalIntNumber,
	},
	statusIntNumber: {
		ByteTypeNumber: statusIntNumber,
		ByteTypePoint:  statusPointWithNumber,
		ByteTypeE:      statusE,
	},
	statusDecimalIntNumber: {
		ByteTypeNumber: statusDecimalIntNumber,
		ByteTypeE:      statusE,
		ByteTypeSpace:  statusEnd,
	},
	statusE: {
		ByteTypeSymbol: statusESymbol,
		ByteTypeNumber: statusENumber,
	},
	statusESymbol: {
		ByteTypeNumber: statusENumber,
	},
	statusENumber: {
		ByteTypeNumber: statusENumber,
		ByteTypeSpace:  statusEnd,
	},
	statusEnd: {
		ByteTypeSpace: statusEnd,
	},
}

func isNumber1(s string) bool {
	state := statusBegin
	for _, b := range s {
		byteType := TypeInference(byte(b))
		if v, ok := StatusMapping[state][byteType]; byteType == ByteTypeUnknown || !ok {
			return false
		} else {
			state = v
		}
	}
	return state == statusEnd || state == statusENumber || state == statusPointWithNumber || state == statusIntNumber || state == statusDecimalIntNumber
}

func TypeInference(b byte) bType {
	switch b {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return ByteTypeNumber
	case ' ':
		return ByteTypeSpace
	case '+', '-':
		return ByteTypeSymbol
	case '.':
		return ByteTypePoint
	case 'e', 'E':
		return ByteTypeE
	default:
		return ByteTypeUnknown
	}
}
