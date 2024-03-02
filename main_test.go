package main

import (
	"image"
	"strings"
	"testing"
)


var testImagePaths = []string{
	"go_21_goroutines_pipeline/images/image1.jpeg",
	"go_21_goroutines_pipeline/images/image2.jpeg",
	"go_21_goroutines_pipeline/images/image3.jpeg",
	"go_21_goroutines_pipeline/images/image4.jpeg",
}

/////UNIT TESTS////


// TestLoadImage function test
func TestLoadImage(t *testing.T) {
	jobs := loadImage(testImagePaths)

	for _, path := range testImagePaths {
		job := <-jobs
		if job.Err != nil {
			t.Errorf("Failed to load image %s: %v", path, job.Err)
		}
		if job.Image == nil {
			t.Errorf("No image returned for %s", path)
		}
		if !strings.Contains(job.OutPath, "output") {
			t.Errorf("Output path %s does not contain 'output' directory", job.OutPath)
		}
	}
}

// TestResize function test
func TestResize(t *testing.T) {
	input := make(chan Job, 1)
	outputSize := image.Point{500, 500}

	// Load a single image for testing
	for _, path := range testImagePaths[:1] {
		job, _ := loadImage([]string{path}) 
		inputJob := <-job
		input <- inputJob
	}
	close(input)

	resizedJobs := resize(input)

	for job := range resizedJobs {
		if job.Err != nil {
			t.Errorf("Failed to resize image: %v", job.Err)
		}
		if job.Image.Bounds().Size() != outputSize {
			t.Errorf("Image was not resized correctly, got size %v, want %v", job.Image.Bounds().Size(), outputSize)
		}
	}
}

// TestConvertToGrayscale function test
func TestConvertToGrayscale(t *testing.T) {
	input := make(chan Job, 1)

	// Load a single image for testing
	for _, path := range testImagePaths[:1] {
		job, _ := loadImage([]string{path}) 
		inputJob := <-job
		input <- inputJob
	}
	close(input)

	grayscaleJobs := convertToGrayscale(input)

	for job := range grayscaleJobs {
		if job.Err != nil {
			t.Errorf("Failed to convert image to grayscale: %v", job.Err)
		}
		_, ok := job.Image.(*image.Gray)
		if !ok {
			t.Errorf("Image is not grayscale")
		}
	}
}

// TestSaveImage function test
func TestSaveImage(t *testing.T) {
	input := make(chan Job, 1)

	// Load a single image for testing
	for _, path := range testImagePaths[:1] {
		job, _ := loadImage([]string{path}) 
		inputJob := <-job
		input <- inputJob
	}
	close(input)

	success := saveImage(input)

	if <-success != true {
		t.Errorf("Failed to save image")
	}
}

////BENCHMARK TESTS////

// BenchmarkLoadImage benchmark test
func BenchmarkLoadImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := loadImage(testImagePaths)
		for range ch {
		}
	}
}

// BenchmarkResize benchmark test
func BenchmarkResize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch1 := loadImage(testImagePaths)
		ch2 := resize(ch1)
		for range ch2 {

		}
	}
}

// BenchmarkConvertToGrayscale benchmark test
func BenchmarkConvertToGrayscale(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch1 := loadImage(testImagePaths)
		ch2 := resize(ch1)
		ch3 := convertToGrayscale(ch2)
		for range ch3 {
		}
	}
}

// BenchmarkSaveImage benchmark test
func BenchmarkSaveImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch1 := loadImage(testImage
