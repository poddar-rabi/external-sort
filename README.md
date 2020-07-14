# Search Engine Company Sort
Sort search terms from log files in ascending alphabetical order. The log files are plain text, all lower case, one search term per line.

The application needs to be able to process upto 1TB of input log files stored on a single hard drive. No limitations on hard drive space including for temporary and resultant files. The input could be a single log file or multiple log files.

Not sufficient RAM to load all the search items simultaneously and can load at most K search terms in memory, such that  
`2 >= K < infinity`


----

The application disk sort is build as a go module

# Dependency

`Go 1.12+`

# Input Parameters

**Input Directory:** Directory where the input log files are present.
It assumes that the directory contains only the files to be sorted and nothing else and processes everything inside this folder. 
Specify full path to the directory

**Output Directory:** Directory in which the intermediate files would be created. 
After the program run, this directory would contain only one file, the sorted output file.
Specify full path to the directory

**Limit:** The number of lines that can be kept in the memory

#Output
A single text file with sorted content in the output directory specified in input.


# Run unit test

` GO111MODULE=on go test ./...`

# Program Execution

`cd disksort`
` GO111MODULE=on go run cmd/main.go`

