package inspect

func clearMap[K comparable, V any](value map[K]V) {
	for key := range value {
		delete(value, key)
	}
}

func makeOrClearMap[K comparable, V any](value *map[K]V, length int) {
	if *value != nil {
		clearMap(*value)
	} else {
		*value = make(map[K]V, length)
	}
}

func StringMap[T Inspectable](value *map[string]T, inspector *Inspector,
	name string, elementName string, description string) {

	if !inspector.IsReading() {
		inspector.WriteMap(name, elementName, len(*value), description)
		for key, item := range *value {
			inspector.WriteNextKey(key)
			item.Inspect(inspector)
		}
		return
	}

	length := inspector.ReadMap()
	if length > 0 {
		makeOrClearMap(value, length)
		for i := 0; i < length; i++ {
			var item T
			key := inspector.ReadNextKey()
			item.Inspect(inspector)
			(*value)[key] = item
		}
		return
	}
	if length == 0 {
		if *value != nil {
			clearMap(*value)
		}
		return
	}
	makeOrClearMap(value, 1)
	key := inspector.ReadNextKey()
	for len(key) > 0 {
		var item T
		item.Inspect(inspector)
		(*value)[key] = item
		key = inspector.ReadNextKey()
	}
	inspector.EndMap()
}

func Map[K comparable, PK InspectablePtr[K], T any, PT InspectablePtr[T]](value *map[K]T, inspector *Inspector,
	name string, keyName string, elementName string, description string) {

	itemName := name + ".item"
	if !inspector.IsReading() {
		inspector.WriteArray(name, itemName, len(*value), description)
		for key, item := range *value {
			o := inspector.StartObject(itemName, "key value pair")
			PK(&key).Inspect(o.Property("k", true, "key"))
			PT(&item).Inspect(o.Property("v", true, "value"))
			o.End()
		}
		inspector.EndArray()
		return
	}
	length := inspector.ReadArray()
	if length > 0 {
		makeOrClearMap(value, length)
		for i := 0; i < length; i++ {
			var key K
			var item T
			o := inspector.StartObject(itemName, "key value pair")
			PK(&key).Inspect(o.Property("k", true, "key"))
			PT(&item).Inspect(o.Property("v", true, "value"))
			o.End()
			(*value)[key] = item
		}
		return
	}
	if length == 0 {
		if *value != nil {
			clearMap(*value)
		}
		return
	}
	makeOrClearMap(value, 1)
	var key K
	var item T
	o := inspector.StartObject(itemName, "key value pair")
	PK(&key).Inspect(o.Property("k", true, "key"))
	PT(&item).Inspect(o.Property("v", true, "value"))
	o.End()
	(*value)[key] = item
	for inspector.HaveNext() {
		var key K
		var item T
		o := inspector.StartObject(itemName, "key value pair")
		PK(&key).Inspect(o.Property("k", true, "key"))
		PT(&item).Inspect(o.Property("v", true, "value"))
		o.End()
		(*value)[key] = item
	}
}

func MapPtr[K comparable, PK InspectablePtr[K], T any, PT InspectablePtr[T]](value *map[K]PT, inspector *Inspector,
	name string, keyName string, elementName string, description string) {

	itemName := name + ".item"
	if !inspector.IsReading() {
		inspector.WriteArray(name, itemName, len(*value), description)
		for key, item := range *value {
			o := inspector.StartObject(itemName, "key value pair")
			PK(&key).Inspect(o.Property("k", true, "key"))
			item.Inspect(o.Property("v", true, "value"))
			o.End()
		}
		inspector.EndArray()
		return
	}
	length := inspector.ReadArray()
	if length > 0 {
		makeOrClearMap(value, length)
		for i := 0; i < length; i++ {
			var key K
			item := PT(new(T))
			o := inspector.StartObject(itemName, "key value pair")
			PK(&key).Inspect(o.Property("k", true, "key"))
			item.Inspect(o.Property("v", true, "value"))
			o.End()
			(*value)[key] = item
		}
		return
	}
	if length == 0 {
		if *value != nil {
			clearMap(*value)
		}
		return
	}
	makeOrClearMap(value, 1)
	var key K
	item := PT(new(T))
	o := inspector.StartObject(itemName, "key value pair")
	PK(&key).Inspect(o.Property("k", true, "key"))
	item.Inspect(o.Property("v", true, "value"))
	o.End()
	(*value)[key] = item
	for inspector.HaveNext() {
		var key K
		item := PT(new(T))
		o := inspector.StartObject(itemName, "key value pair")
		PK(&key).Inspect(o.Property("k", true, "key"))
		item.Inspect(o.Property("v", true, "value"))
		o.End()
		(*value)[key] = item
	}
}
