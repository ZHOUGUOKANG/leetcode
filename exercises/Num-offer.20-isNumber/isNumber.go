package Num_offer_20_isNumber

type State int
type CharType int

const (
	StateInitial State = iota
	StateIntSign
	StateInteger
	StatePoint
	StatePointWithoutInt
	StateFraction
	StateExp
	StateExpSign
	StateExpNumber
	StateEnd
)

const (
	CharNumber CharType = iota
	CharExp
	CharPoint
	CharSign
	CharSpace
	CharIllegal
)

func toCharType(ch byte) CharType {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CharNumber
	case 'e', 'E':
		return CharExp
	case '.':
		return CharPoint
	case '+', '-':
		return CharSign
	case ' ':
		return CharSpace
	default:
		return CharIllegal
	}
}

var transfer = map[State]map[CharType]State{
	StateInitial: {
		CharSpace:  StateInitial,
		CharNumber: StateInteger,
		CharPoint:  StatePointWithoutInt,
		CharSign:   StateIntSign,
	},
	StateIntSign: {
		CharNumber: StateInteger,
		CharPoint:  StatePointWithoutInt,
	},
	StateInteger: {
		CharNumber: StateInteger,
		CharExp:    StateExp,
		CharPoint:  StatePoint,
		CharSpace:  StateEnd,
	},
	StatePoint: {
		CharNumber: StateFraction,
		CharExp:    StateExp,
		CharSpace:  StateEnd,
	},
	StatePointWithoutInt: {
		CharNumber: StateFraction,
	},
	StateFraction: {
		CharNumber: StateFraction,
		CharExp:    StateExp,
		CharSpace:  StateEnd,
	},
	StateExp: {
		CharNumber: StateExpNumber,
		CharSign:   StateExpSign,
	},
	StateExpSign: {
		CharNumber: StateExpNumber,
	},
	StateExpNumber: {
		CharNumber: StateExpNumber,
		CharSpace:  StateEnd,
	},
	StateEnd: {
		CharSpace: StateEnd,
	},
}

func isNumber(s string) bool {
	state := StateInitial
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i])
		if v, ok := transfer[state][typ]; !ok {
			return false
		} else {
			state = v
		}
	}
	return state == StateInteger || state == StatePoint || state == StateFraction || state == StateExpNumber || state == StateEnd
}
