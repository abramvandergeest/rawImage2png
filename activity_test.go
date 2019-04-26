package rawImage2png

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	raw, err := os.Open("/Users/avanderg@tibco.com/working/coffee_carafe_demo/Jabil_Image_Classification/Camera_Capture/Cup/Image0") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	input := &Input{File: raw}

	tc := test.NewActivityContext(act.Metadata())

	tc.SetInputObject(input)

	done, err := act.Eval(tc)
	assert.True(t, done)
	assert.Nil(t, err)

	output := &Output{}
	tc.GetOutputObject(output)
	fmt.Println(output.OutFilePNG)

}
