package rawImage2png

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"

	// "os"
	"os/exec"
	"time"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
)

func init() {
	activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	// Reading in a raw image file

	fileContents := ctx.GetInput("file")

	if fileContents == "" {
		return true, activity.NewError("File content must not be empty", "", nil)
	}
	msg, _ := coerce.ToString(fileContents)
	buffer := []byte(msg)

	// Saving raw file to disk
	rand.Seed(time.Now().UnixNano())
	rnum := rand.Intn(1000000)
	// fmt.Println("blah:",rnum)
	writefilename := fmt.Sprintf("tmprawfile%06d", rnum)
	ioutil.WriteFile(writefilename, buffer, 0777)

	// converting raw image to png
	cmd := "ffmpeg"
	args := []string{"-vcodec", "rawvideo", "-f", "rawvideo", "-pix_fmt", "rgb565", "-s", "640x480", "-i", writefilename, writefilename + ".png"}
	err = exec.Command(cmd, args...).Run()
	if err != nil {
		// fmt.Println(err)
		return true, err
	}

	// Opening PNG file to pass on
	pngf, err := ioutil.ReadFile(writefilename + ".png")
	if err != nil {
		log.Fatal(err)
	}

	ctx.Logger().Infof("file %q is open as []byte.", writefilename+".png")

	// Cleaning up files
	cmd = "rm"
	args = []string{writefilename, writefilename + ".png"}
	err = exec.Command(cmd, args...).Run()
	if err != nil {
		return true, err
	}

	// Outputting files
	output := &Output{OutFilePNG: pngf}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}
	return true, nil
}