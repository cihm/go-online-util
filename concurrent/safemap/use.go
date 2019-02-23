package concurrent

var ParameterGmap *SafeMap

func init() {
	ParameterGmap = NewSafeMap()
	//sample
	use()
}

func use() {
	ParameterGmap.ReadMap("key")
	ParameterGmap.WriteMap("key", 12)
}
