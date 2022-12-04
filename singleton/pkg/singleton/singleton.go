package singleton

type Singleton struct {
	Title string
}

var singletonInstance *Singleton

func NewSingleton(title string) *Singleton {
	if singletonInstance == nil {
		singletonInstance = &Singleton{Title: title}
	} else {
		singletonInstance.Title = title
	}
	return singletonInstance
}

func GetSingletonInstance() *Singleton {
	return singletonInstance
}
