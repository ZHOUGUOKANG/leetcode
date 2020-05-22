package SingletonPattern

var instance *SingleObject

type SingleObject struct {
}

func (SingleObject) GetInstance() *SingleObject {
	if instance == nil {
		instance = &SingleObject{}
	}
	return instance
}

func (SingleObject) Echo() string {
	return "SingleObject Echo"
}
