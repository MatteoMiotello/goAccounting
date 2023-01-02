package bootstrap

var initializers = [...]Initializer{
	Config{},
	Db{},
	Api{},
}

func Run() {
	for _, initializer := range initializers {
		err := initializer.Init()
		if err != nil {
			panic(err.Error())
			return
		}
	}
}
