func getTime() int64 {
	return time.Now().UTC().Unix()
}

func AssertRangeMin(w http.ResponseWriter, min float64, value interface{}) bool {
	if err := assertRangeMin(min, value); err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return false
	}
	return true
}

func AssertRangeMax(w http.ResponseWriter, max float64, value interface{}) bool {
	if err := assertRangeMax(max, value); err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return false
	}
	return true
}

func assertRangeMin(minimum float64, value interface{}) error {

	var val float64
	switch v := value.(type) {
	case int:
		val = float64(v)
	case float64:
		val = v
	case string:
		val = float64(len(v))
	default:
		log.Println("assertRange: ignoring range assertion for unknown type")
	}

	err := fmt.Errorf("assertRange: value %v exceeded value of range min: %v", value, minimum)
	if val < minimum {
		return err
	}
	return nil
}

func assertRangeMax(maximum float64, value interface{}) error {

	var val float64
	switch v := value.(type) {
	case int:
		val = float64(v)
	case float64:
		val = v
	case string:
		val = float64(len(v))
	default:
		log.Println("assertRange: ignoring range assertion for unknown type")
	}

	err := fmt.Errorf("assertRange: value %v exceeded value of range max: %v", value, maximum)
	if val > maximum && val != -1 {
		return err
	}
	return nil
}

func assertMAPSTRINGINT(m map[string]interface{}, key string) (map[string]int, error) {
	result := map[string]int{}
	object := m[key].(map[string]interface{})
	for k, v := range object {
		if f, ok := v.(float64); ok {
			result[k] = int(f)
		}
	}
	if len(object) != len(result) {
		return nil, fmt.Errorf("assertMAPSTRINGINT: '%s' is required for this request", key)
	}
	return result, nil
}

func AssertMAPSTRINGINT(w http.ResponseWriter, m map[string]interface{}, key string) (map[string]int, bool) {
	data, err := assertMAPSTRINGINT(m, key)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return nil, false
	}
	return data, true
}

func AssertSTRING(w http.ResponseWriter, m map[string]interface{}, key string) (string, bool) {
	s, err := assertSTRING(m, key)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return s, false
	}
	return s, true
}

func assertSTRING(m map[string]interface{}, key string) (string, error) {
	s, ok := m[key].(string)
	if !ok {
		return s, fmt.Errorf("assertSTRING: '%s' is required for this request", key)
	}
	return s, nil
}

func AssertARRAYSTRING(w http.ResponseWriter, m map[string]interface{}, key string) ([]string, bool) {
	s, err := assertARRAYSTRING(m, key)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return nil, false
	}
	return s, true
}

func assertARRAYSTRING(m map[string]interface{}, key string) ([]string, error) {
	a, ok := m[key].([]interface{})
	if !ok {
		return nil, fmt.Errorf("'%s' is required for this request", key)
	}
	b := []string{}
	for _, v := range a {
		s, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("strings are required for this request: %s", key)
		}
		b = append(b, s)
	}
	return b, nil
}

func AssertFLOAT64(w http.ResponseWriter, m map[string]interface{}, key string) (float64, bool) {
	f, err := assertFLOAT64(m, key)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return 0, false
	}
	return f, true
}

func assertFLOAT64(m map[string]interface{}, key string) (float64, error) {
	f, ok := m[key].(float64)
	if !ok {
		return 0, fmt.Errorf("assertFLOAT64: '%s' is required for this request", key)
	}
	return f, nil
}

func AssertBOOL(w http.ResponseWriter, m map[string]interface{}, key string) (bool, bool) {
	b, err := assertBOOL(m, key)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return false, false
	}
	return b, true
}

func assertBOOL(m map[string]interface{}, key string) (bool, error) {
	v, ok := m[key].(bool)
	if !ok {
		return false, fmt.Errorf("assertBOOL: '%s' is required for this request", key)
	}
	return v, nil
}

func AssertINT(w http.ResponseWriter, m map[string]interface{}, key string) (int, bool) {
	x, err := assertINT(m, key)
	if err != nil {
		cloudfunc.HttpError(w, err, http.StatusBadRequest)
		return 0, false
	}
	return x, true
}

func assertINT(m map[string]interface{}, key string) (int, error) {
	v, ok := m[key]
	if !ok {
		PrettyPrint(m)
		return 0, fmt.Errorf("assertINT: '%s' is required for this request", key)
	}
	switch v := m[key].(type) {
	case string:
		num, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		return num, nil
	case float64:
		return int(v), nil
	case int:
	default:
		panic("fail assert int: " + key)
	}
	return v.(int), nil
}
