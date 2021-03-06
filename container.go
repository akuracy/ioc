package ioc

var AllowAbsent = true

type Container interface {
	// RegisterValue registers value with name
	RegisterValue(name string, value interface{}) bool

	// RegisterSingleton register a singleton value of prototype
	// Return its corresponded name
	// Only one value will be created
	RegisterSingleton(prototype interface{}) string

	// RegisterTransient register a transient value of prototype
	// Return its corresponded name
	// New value will be created in every resolve
	RegisterTransient(prototype interface{}) string

	// RegisterTransientCreator register a new transient of name. It will be created through creator.
	RegisterTransientCreator(name string, creator Creator) bool

	// RegisterSingletonCreator register a new singleton of name. It will be created through creator.
	RegisterSingletonCreator(name string, creator Creator) bool

	// Contains returns true if name is already registered
	Contains(name string) bool

	// RegisterAliases adds aliases to origin
	RegisterAliases(origin interface{}, aliases ...interface{}) bool

	// GetAliases returns all aliases of name which is also included in the result
	GetAliases(origin interface{}) []string

	// Resolve finds or creates value by prototype, and inject all dependencies
	Resolve(prototype interface{}) interface{}

	// Inject ptrToObj's fields with inject tag
	Inject(ptrToObj interface{})
}

var rootContainer = NewContainer()

func RootContainer() Container {
	return rootContainer
}

func NewContainer() Container {
	c := &containerImpl{}
	c.nameToRegistryIndex = make(map[string]int, 10)
	c.factory = NewFactory()
	return c
}

func RegisterValue(name string, value interface{}) bool {
	return rootContainer.RegisterValue(name, value)
}

func RegisterSingleton(prototype interface{}) string {
	return rootContainer.RegisterSingleton(prototype)
}

func RegisterTransient(prototype interface{}) string {
	return rootContainer.RegisterTransient(prototype)
}

func RegisterTransientCreator(name string, creator Creator) bool {
	return rootContainer.RegisterTransientCreator(name, creator)
}

func RegisterSingletonCreator(name string, creator Creator) bool {
	return rootContainer.RegisterSingletonCreator(name, creator)
}

func Contains(name string) bool {
	return rootContainer.Contains(name)
}

func RegisterAliases(origin interface{}, alias ...interface{}) bool {
	return rootContainer.RegisterAliases(origin, alias...)
}

func GetAliases(origin interface{}) []string {
	return rootContainer.GetAliases(origin)
}

func Resolve(prototype interface{}) interface{} {
	return rootContainer.Resolve(prototype)
}

func Inject(ptrToObj interface{}) {
	rootContainer.Inject(ptrToObj)
}
