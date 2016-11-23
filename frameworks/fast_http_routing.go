package frameworks

// this router is not has handler like Controller,
// sometime we need use Controller to Simplify ours code.
func RouterController(r *routing.RouteGroup, path string, controller interface{}, handlers ...routing.Handler) *routing.RouteGroup {
	stru := reflect.ValueOf(controller)
	typ := stru.Type()
	c := r.Group(path)
	// get all func from controller, then add it to router group if sign is 'router.handler'
	for i := stru.NumMethod() - 1; i >= 0; i-- {
		fun := stru.Method(i)
		ifun, ok := fun.Interface().(func(*routing.Context) error)

		if ok {
			name := typ.Method(i).Name
			// to lower the initial
			name = strings.ToLower(string(name[0])) + name[1:]
			c.Any("/" + name, func(p *routing.Context) error {
				p.Set("method", name)
				return ifun(p)
			}, handlers...)
		}
	}
	return c
}


// i think the routing is not perfect :
// 1. 'Use' like Middleware but is not, it can handle incoming request, but can not handle outcoming response 
// 2. when i have too many router group, the code format is not beautiful. is better use index format?