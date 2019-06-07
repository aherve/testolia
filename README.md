Testolia: a query log parser server
===================================

[![Codeship Status for aherve/testolia](https://app.codeship.com/projects/fe608280-6a8a-0137-604a-2a010c419d6a/status?branch=master)](https://app.codeship.com/projects/346454)

## Table of Contents

1. [About](#about)
2. [Installation](#installation)
3. [Testing](#testing)
4. [Initial assignment](#initialAssignement)

## About
This go program creates a server that allows to print a report of a log file on demand. It can either count the distinct request during a timeframe, or print an object listing the popular searches during a timeframe:

![count](./assets/count.png)

![popular](./assets/popular.png)

### What's inside

The program is written in go, with no external lib used.

### How it works
 
1. The file is read, every line is split in two strings `timestamp` and `query`.
2. The data is stored in a slice, that we sort against the timestamp key. The timestamps are properly formatted, so sorting alphabetically yields a proper sort by timestamp
3. At this point the server can start listening to requests
4. When receiving a request, find the first matching index of the sorted data (using a binary sort, thus `O(log(n))`), scan every line that matches the filter (simple scan until the matching condition fails) and reduce to provide the desired output (either distinct count or popular report).

## Installation & usage
This program can either be installed using a go environment, or via a provided Dockerfile

### 1. Use prebuilt image

- A prebuilt image is available on dockerhub; simply run `docker run -v <absolute-path-to-your-local-file>:/tmp/log.gz -p 8080:8080 aherve/testolia /tmp/log.gz`

Each push on github triggers a new docker build

### 2. Build docker image

- build the image with `docker build -t testolia .`
- run it with `docker run -v <your-local-file>:/tmp/log.gz -p 8080:8080 testolia /tmp/log.gz`

### 3. Install with go
run `go build` to create the `testolia` executable and run it with `./testolia <filename>`

The `<filename>` argument can either be a `.tsv` file, or a gzipped `.tsv.gz` file
 
## Testing

go test files are present and can be launched using `go test`.
The tests are also automatically run on codeship: 

[![Codeship Status for aherve/testolia](https://app.codeship.com/projects/fe608280-6a8a-0137-604a-2a010c419d6a/status?branch=master)](https://app.codeship.com/projects/346454)

## Limitations & possible improvements

- I hope I didn't step on too many naming conventions, go-specific patterns or so. This is the first time I write something in go so I'm hoping this is easily readable by more experienced go developers
- I assumed the input logs are not going to be updated during runtime so I only read it once, and do not check for any file update. To go further, we could for instance watch any changes and reload the file if need be.
- I assumed the server RAM memory is large enough so that we can keep the entire file content in memory. This allows a much faster response time, whilst raising the footprint of the program. For significantly larger datasets, a strategy where we stream the input file and filter it on the fly could be preferable.
- The error handling is really basic: just panic if anything goes wrong. I'd advocate to set at least a treshold on how many missed lines the loading script tolerates before throwing a panic
- There is close to no server logging, this would obviously be mandatory to create something production-proof.

## Initial Assignment

[ Initial assignment can be found here ](https://gist.github.com/sfriquet/55b18848d6d58b8185bbada81c620c4a)
