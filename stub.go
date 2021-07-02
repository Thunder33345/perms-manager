package ranker

type ProcessorFunc func(r RawPermissible, pr GroupProvider) (Permissible, error)

func (p ProcessorFunc) Process(r RawPermissible, pr GroupProvider) (Permissible, error) {
	return p(r, pr)
}

type Processor interface {
	Process(r RawPermissible, pr GroupProvider) (Permissible, error)
}

type GroupProviderFunc func(uid string) (Group, bool)

func (g GroupProviderFunc) GetGroup(uid string) (Group, bool) {
	return g(uid)
}

type GroupProvider interface {
	GetGroup(uid string) (Group, bool)
}

type Judge interface {
	HasPermission(p Permissible, node string) bool
	HasPermissionWithLevel(p Permissible, node string,level int) bool
	IsHigherLevel(source Permissible, subject Permissible) bool
}
