package rawImage2png

import "os"

type Input struct {
	File interface{} `md:"file,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	r.File = values["file"].(*os.File)
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"file": r.File,
	}
}

type Output struct {
	OutFilePNG []byte `md:"outFilePNG,required"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	o.OutFilePNG = values["outFilePNG"].([]byte)
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"outFilePNG": o.OutFilePNG,
	}
}
