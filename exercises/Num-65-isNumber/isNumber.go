package Num_65_IsNumber

type State int
type CharType int

const (
	StateInit State = iota
	StateSignNumber
	StateNumber
	StatePoint
	StatePointWithoutNumber
	StateFraction
	StateExp
	StateExpSignNumber
	StateExpNumber
	StateEnd
)

const (
	CharNumber CharType = iota
	CharPoint
	CharExp
	CharSpace
	CharSign
	CharIllegal
)

var states = map[State]map[CharType]State{
	StateInit: {
		CharSpace:  StateInit,
		CharSign:   StateSignNumber,
		CharNumber: StateNumber,
		CharPoint:  StatePointWithoutNumber,
	},
	StateSignNumber: {
		CharNumber: StateNumber,
		CharPoint:  StatePointWithoutNumber,
	},
	StateNumber: {
		CharNumber: StateNumber,
		CharPoint:  StatePoint,
		CharExp:    StateExp,
		CharSpace:  StateEnd,
	},
	StatePoint: {
		CharNumber: StateFraction,
		CharExp:    StateExp,
		CharSpace:  StateEnd,
	},
	StatePointWithoutNumber: {
		CharNumber: StateFraction,
	},
	StateFraction: {
		CharNumber: StateFraction,
		CharExp:    StateExp,
		CharSpace:  StateEnd,
	},
	StateExp: {
		CharSign:   StateExpSignNumber,
		CharNumber: StateExpNumber,
	},
	StateExpNumber: {
		CharNumber: StateExpNumber,
		CharSpace:  StateEnd,
	},
	StateExpSignNumber: {
		CharNumber: StateExpNumber,
		CharSpace:  StateEnd,
	},
	StateEnd: {
		CharSpace: StateEnd,
	},
}

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
func isNumber(s string) bool {
	state := StateInit
	for i := 0; i < len(s); i++ {
		if v, ok := states[state][toCharType(s[i])]; ok {
			state = v
		} else {
			return false
		}
	}
	return state == StateNumber || state == StatePoint || state == StateFraction || state == StateExpNumber || state == StateEnd
}
