package typing

type IContext interface {
}

type RouHandle func(IRou, IContext)

type IRou interface {
	// Post(g IRouGroup, e string, f RouHandle)
	// Use(f RouHandle)
	CreateEvent(string, string, RouHandle) IEvent
	Resp(IContext, int, interface{})
	Resp200(IContext, interface{})
	Set(c IContext, key string, arg interface{})
	Get(c IContext, key string) (interface{}, bool)
	Body(c IContext, arg interface{}) error
	Header(c IContext, arg interface{}) error
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
