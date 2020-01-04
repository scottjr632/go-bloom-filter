# GoLang Bloom Filter 
[![codecov](https://codecov.io/gh/scottjr632/go-bloom-filter/branch/master/graph/badge.svg)](https://codecov.io/gh/scottjr632/go-bloom-filter) [![CircleCI](https://circleci.com/gh/scottjr632/go-bloom-filter.svg?style=svg)](https://circleci.com/gh/scottjr632/go-bloom-filter) 

Simple bloom filter implementation for the GoLang programming language  

## Installing
```bash
go get github.com/scottjr632/go-bloom-filter
```

### Importing
```go
import "github.com/scottjr632/go-bloom-filter"
```

## Creating a new optimal filter 
```go
estimatedNumberOfItems := 3000
maxFailureRate := 0.01
bf := bloomfilter.NewFromEstimate(estimatedNumberOfItems, maxFailureRate)
```

### Creating a new filter
```go
sizeOfFilter, numHashFns := 128, 3
bf := bloomfilter.New(sizeOfFilter, numHashFns)
```

## Handling values

#### Adding a value to the filter
```go
valueToAdd := 'https://maliciousurl.xyz'
bf.Add(valueToAdd)
```

#### Checking if a value has been set
```go
valueToCheck := 'https://maliciousurl.xyz'
bf.Check(valueToCheck) // -> true

bf.Check('value that has not been set') // -> false
```
