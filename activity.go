package rawImage2png

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/project-flogo/core/activity"
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
	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	// Making the file a byte array
	rawfile := input.File.(*os.File)
	fileinfo, err := rawfile.Stat()
	buffer := make([]byte, fileinfo.Size())
	rawfile.Read(buffer)

	// Saving raw file to disk
	rand.Seed(time.Now().UnixNano())
	rnum := rand.Intn(1000000)
	fmt.Println(rnum)
	writefilename := fmt.Sprintf("tmprawfile%06d", rnum)
	ioutil.WriteFile(writefilename, buffer, 0777)

	// converting raw image to png
	cmd := "ffmpeg"
	args := []string{"-vcodec", "rawvideo", "-f", "rawvideo", "-pix_fmt", "rgb565", "-s", "640x480", "-i", writefilename, writefilename + ".png"}
	err = exec.Command(cmd, args...).Run()
	if err != nil {
		fmt.Println(err)
		return true, err
	}

	// Opening PNG file to pass on
	pngf, err := os.Open(writefilename + ".png") // For read access.
	if err != nil {
		log.Fatal(err)
	}

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
