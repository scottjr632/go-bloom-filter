# GoLang Bloom Filter 
[![codecov](https://codecov.io/gh/scottjr632/go-bloom-filter/branch/master/graph/badge.svg)](https://codecov.io/gh/scottjr632/go-bloom-filter) [![CircleCI](https://circleci.com/gh/scottjr632/go-bloom-filter.svg?style=svg)](https://circleci.com/gh/scottjr632/go-bloom-filter) 

Simple bloom filter implementation for the GoLang programming language  

## Creating a new filter
```go
sizeOfFilter, numHashFns := 128, 3
bf := bloomfilter.New(sizeOfFilter, numHashFns)
```

## Creating a new optimal filter 
```go
estimatedNumberOfItems := 3000
maxFailureRate := 0.01
bf := NewFromEstimate(estimatedNumberOfItems, maxFailureRate)
```

### Adding a value to the filter
```go
valueToAdd := 'https://maliciousurl.xyz'
bf.Add(valueToAdd)
```

### Checking if value has been set
```go
valueToCheck := 'https://maliciousurl.xyz'
bf.Check(valueToCheck) // -> true

bf.Check('value that has not been set') // -> false
```
