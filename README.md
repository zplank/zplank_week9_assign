# zplank_week9_assign

The below was modified on the code-heim repository - 

1) Error checking for the image file input and output was done by creating a job that checks for the success of reading in an image and throws an error message if the image was not read properly. The same was done within the saveImage function to check if the image was saved properly and throws an error if it was not. 

2) Added in my own images 

3) Created a unit test to test each function within the processing pipeline. This can be found in main_test.go under the Unit Test section 

4) Added benchmark tests to the unit test file to measure throughput times for each function. This can be found in the main_test.go file under the Benchmark Test section 