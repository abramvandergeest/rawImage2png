package addDimension

import (
	"fmt"
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
	// fmt.Println("GETS HERE")

	// iCtx := test.NewActivityInitContext(nil, nil)
	// act, err := New(iCtx)
	act := &Activity{}
	// assert.Nil(t, err)
	input := &Input{Data: []float64{1.1, 2.1}}
	tc := test.NewActivityContext(act.Metadata())

	tc.SetInputObject(input)

	done, err := act.Eval(tc)
	assert.True(t, done)
	assert.Nil(t, err)

	output := &Output{}
	tc.GetOutputObject(output)
	fmt.Println(output.Output)
}
