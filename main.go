package main

import (
	"fmt"
	"image"
	"image/color"
	"net/http"
	"sync"

	"gocv.io/x/gocv"
)

var (
	webcam   *gocv.VideoCapture
	img      *gocv.Mat
	imgMutex sync.Mutex
)

func captureLoop() {
	newMat := gocv.NewMat()
	img = &newMat
	for {
		imgMutex.Lock()
		webcam.Read(img)
		imgMutex.Unlock()
	}
}

func applyMask(img gocv.Mat) gocv.Mat {
	width := img.Cols()
	height := img.Rows()

	// Create mask (single channel)
	mask := gocv.NewMatWithSize(height, width, gocv.MatTypeCV8U)
	defer mask.Close()

	// Fill mask with black
	gocv.Rectangle(&mask, image.Rect(0, 0, width, height), color.RGBA{0, 0, 0, 0}, -1)

	// Fill center area with white (active region)
	gocv.Rectangle(
		&mask,
		image.Rect(
			int(0.1*float64(width)),
			int(0.1*float64(height)),
			int(0.9*float64(width)),
			int(0.9*float64(height)),
		),
		color.RGBA{255, 255, 255, 0}, // blue=255 for single channel
		-1,
	)

	// Create output and copy only masked region
	result := gocv.NewMat()
	img.CopyToWithMask(&result, mask)
	return result
}

func frameHandler(w http.ResponseWriter, r *http.Request) {
	imgMutex.Lock()
	defer imgMutex.Unlock()

	if img.Empty() {
		http.Error(w, "No image", http.StatusInternalServerError)
		return
	}

	// Apply Gaussian blur to the image
	gocv.GaussianBlur(*img, img, image.Pt(15, 15), 0, 0, gocv.BorderDefault)

	// Convert the image to grayscale
	gocv.CvtColor(*img, img, gocv.ColorBGRToGray)

	// Mask outside 0.1 to 0.9 range (width and height both)
	mask := applyMask(*img)
	defer mask.Close() // Close the mask to free resources
	buf, _ := gocv.IMEncode(".jpg", mask)

	if buf == nil {
		http.Error(w, "Failed to encode image", http.StatusInternalServerError)
	}

	frameBytes := buf.GetBytes()
	defer buf.Close()
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(frameBytes)
}

func main() {
	list_cameras()
	var err error
	webcam, err = gocv.OpenVideoCapture(1)
	if err != nil {
		panic(err)
	}
	defer webcam.Close()

	go captureLoop()

	http.HandleFunc("/frame", frameHandler)
	fmt.Println("Serving at http://localhost:8080/frame")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
