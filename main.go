package main

import (
	"fmt"
	"image"
	"strings"
	imageprocessing "zplank_week9_assign/go_21_goroutines_pipeline/image_processing"
)

type Job struct {
	InputPath string
	Image     image.Image
	OutPath   string
	Err       error
}

func loadImage(paths []string) <-chan Job {
	out := make(chan Job)
	go func() {
		for _, p := range paths {
			job := Job{
				InputPath: p,
				OutPath:   strings.Replace(p, "go_21_goroutines_pipeline/images/", "go_21_goroutines_pipeline/images/output/", 1),
			}
			job.Image, job.Err = imageprocessing.ReadImage(p)
			out <- job
		}
		close(out)
	}()
	return out
}

func resize(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input {
			if job.Err == nil {
				job.Image, job.Err = imageprocessing.Resize(job.Image)
			}
			out <- job
		}
		close(out)
	}()
	return out
}

func convertToGrayscale(input <-chan Job) <-chan Job {
	out := make(chan Job)
	go func() {
		for job := range input {
			if job.Err == nil {
				job.Image, job.Err = imageprocessing.Grayscale(job.Image)
			}
			out <- job
		}
		close(out)
	}()
	return out
}

func saveImage(input <-chan Job) <-chan bool {
	out := make(chan bool)
	go func() {
		for job := range input {
			if job.Err == nil {
				job.Err = imageprocessing.WriteImage(job.OutPath, job.Image)
			}
			out <- (job.Err == nil)
		}
		close(out)
	}()
	return out
}

func main() {
	imagePaths := []string{
		"go_21_goroutines_pipeline/images/image1.jpeg",
		"go_21_goroutines_pipeline/images/image2.jpeg",
		"go_21_goroutines_pipeline/images/image3.jpeg",
		"go_21_goroutines_pipeline/images/image4.jpeg",
	}

	channel1 := loadImage(imagePaths)
	channel2 := resize(channel1)
	channel3 := convertToGrayscale(channel2)
	writeResults := saveImage(channel3)

	// Collect all jobs to check for errors at the end.
	var jobs []Job

	for job := range channel3 {
		jobs = append(jobs, job)
	}

	for _, job := range jobs {
		if job.Err != nil {
			fmt.Printf("Failed processing %s: %v\n", job.InputPath, job.Err)
		} else {
			fmt.Printf("Successfully processed %s\n", job.InputPath)
		}
	}

	for success := range writeResults {
		if success {
			fmt.Println("Image saved successfully!")
		} else {
			fmt.Println("Failed to save image!")
		}
	}
}
