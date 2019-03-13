package oop

//定義空類 以產生 工廠模式 的方法
type Factory struct{}

func (f *Factory) FacMethod(a, b int, operator string) (value int) {
	var i Opter
	switch operator {
	case "+":
		var AddNum Add = Add{BaseNum{a, b}}
		i = &AddNum
	case "-":
		var SubNum Sub = Sub{BaseNum{a, b}}
		i = &SubNum
	}
	//介面實現 :value = i.Opt()
	value = MultiState(i) //多型實現
	return
}
