package addDimension

// type Settings struct {
// 	ASetting string `md:"aSetting,required"`
// }

type Input struct {
	Data interface{} `md:"data"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	// strVal, _ := coerce.ToString(values["anInput"])
	r.Data = values["data"]
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"data": r.Data,
	}
}

type Output struct {
	Output []interface{} `md:"output"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	// strVal, _ := coerce.ToString()
	o.Output = values["output"].([]interface{})
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"output": o.Output,
	}
}
