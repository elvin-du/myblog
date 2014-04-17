package database

func Map2Struct(data map[string]interface{}) ([]string, []interface{}) {
	ks, vs := []string{}, []interface{}{}
	for k, v := range data {
		ks = append(ks, k)
		vs = append(vs, v)
	}

	return ks, vs
}
