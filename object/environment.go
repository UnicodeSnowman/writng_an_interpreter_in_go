package object

func NewEnvironment() *Environment {
	return &Environment{store: map[string]Object{}, outer: nil}
}

func NewInnerEnvironment(outer *Environment, local map[string]Object) *Environment {
	s := map[string]Object{}
	for k, v := range local {
		s[k] = v
	}
	return &Environment{store: s, outer: outer}
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
