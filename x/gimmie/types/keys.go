package types

const (
	// ModuleName defines the module name
	ModuleName = "gimmie"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_gimmie"

    
)

var (
	ParamsKey = []byte("p_gimmie")
)



func KeyPrefix(p string) []byte {
    return []byte(p)
}
