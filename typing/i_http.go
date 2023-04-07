package typing

type IContext interface {
}

type RouHandle func(IRou, IContext)

type IRou interface {
	// Post(g IRouGroup, e string, f RouHandle)
	// Use(f RouHandle)
	Resp(IContext, int, interface{})
	CreateEvent(string, string, RouHandle) IEvent
}

type IMiddle interface {
	Before(IRou)
	After(IRou)
}

type IEvent interface {
	Method() string
	Uri() string
	Handle() RouHandle
}

type IRouGroup interface {
	Name() string
	Events() []IEvent
}
