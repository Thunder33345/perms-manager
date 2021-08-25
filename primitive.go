package perms_manager

//Group represent a collection of permissions and metadata
type Group struct {
	//Name is a display friendly name of the group, it should only be used for display, it has no filtering or requirement
	Name string
	//RefName is a command friendly runtime stable name, should be unique
	RefName string
	//UID is the unique identifier for this group used for saving and referencing, must never be changed
	UID string
	//Weight dictates the overwriting precedent, where the larger overwrites the smaller
	//must be unique, otherwise behaviour is undefined
	Weight int
	//Permission is the permission that is used
	Permission Entry
	//Flags are conditional Entry that only applies in certain situations
	Flags map[string]FlagEntry
}

//Entry represent a collection of permissions and flags
type Entry struct {
	//EmptySet will discard all previously granted permissions
	EmptySet bool
	//Level is the default power level of said entry
	//Only the highest group's level is in used
	Level int
	//IgnoreLevel makes this entry's Level ignored
	IgnoreLevel bool
	//SetLevel makes overwrites the Level instead of adding or subtracting from last level
	SetLevel bool
	//Grant will add permissions to the List
	Grant []string
	//Revoke will revoke a permissions that is granted to the List by a prior group
	Revoke []string
}

//FlagEntry is an Entry but inside a Group.Flags
//its same as Entry byt with extra flag only fields
type FlagEntry struct {
	//Weight dictates the overwriting precedent, must be unique, otherwise behaviour is undefined
	//behaviour is defined by processor
	Weight int
	//Preprocess indicates that this should be processed before Group.Permission
	Preprocess bool
	Entry
}

//RawList is the raw save state for List
type RawList struct {
	//Overwrites has the highest precedent
	//Will overwrite all group based permissions
	Overwrites Entry
	//Groups are a list of group UUID to inherit permission from
	Groups []string
}

//List is the compiled result from a RawList
type List struct {
	//Level is the final applicable level
	Level int
	//Permission is th final applicable permission
	Permission []string
}
