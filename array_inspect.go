package inspect

func setArraySize[T any](array *[]T, newSize int) {
	if newSize > cap(*array) {
		*array = (*array)[:newSize]
		return
	}
	*array = make([]T, newSize)
}

func Array[PT InspectablePtr[T], T any](array *[]T, inspector *Inspector, name string,
	elementName string, description string) {

	if inspector.IsReading() {
		length := inspector.ReadArray()
		if length <= 0 {
			if length == -1 {
				setArraySize(array, 1)
				PT(&(*array)[0]).Inspect(inspector)
				for inspector.HaveNext() {
					var item T
					PT(&item).Inspect(inspector)
					*array = append(*array, item)
				}
				inspector.EndArray()
			} else {
				(*array) = (*array)[:0]
			}
			return
		}
		setArraySize(array, length)
	} else {
		inspector.WriteArray(name, elementName, len(*array), description)
	}
	for _, item := range *array {
		PT(&item).Inspect(inspector)
	}
}

func ArrayPtr[T any, PT InspectablePtr[T]](array *[]PT, inspector *Inspector, name string,
	elementName string, description string) {

	if inspector.IsReading() {
		length := inspector.ReadArray()
		if length <= 0 {
			if length == -1 {
				setArraySize(array, 1)
				(*array)[0].Inspect(inspector)
				for inspector.HaveNext() {
					item := PT(new(T))
					item.Inspect(inspector)
					*array = append(*array, item)
				}
				inspector.EndArray()
			} else {
				(*array) = (*array)[:0]
			}
			return
		}
		setArraySize(array, length)
	} else {
		inspector.WriteArray(name, elementName, len(*array), description)
	}
	for _, item := range *array {
		item.Inspect(inspector)
	}
}
