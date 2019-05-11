# How to execute ?

1. Run following command. 

```
go get -u github.com/sanksons/riverntorch
```

This will install the binary in $GOPATH/bin.

2. Now execute the binary, using following command:

```
riverntorch --file=/tmp/data.yml
```

3. Please check sample `data.yml` for data format to be used.

## Approach Explaination

## 1. Example One

```
A:1, B:2, C:5, D:10
```

### Fastest Bearer Approach

fastestBearerApproach uses the idea of fastest person to always act as the torchbearer. This minimizes the torch returning time.

```
side A -> sideB : A + D
side A <- sideB : A
side A -> sideB : A + C
side A <- sideB : A
side A -> sideB : A + B

Total Time: 19
```

### Club Slowest Approach

clubSlowestApproach uses the idea of clubbing the Slowest members together, so that time of forward journey can be reduced.

However, this incurs an extra cost on backward journey. So, instead of always using the fastest rider to return the torch, it uses the combination of fastest and second fastest riders.
```
side A -> sideB : A + B
side A <- sideB : A
side A -> sideB : C + D
side A <- sideB : B
side A -> sideB : A + B

Total Time: 17
```

## 2. Example Two

```
A:1, B:10, C:11, D:12, E:13, F14, G15
```

### Fastest Bearer Approach

```
side A -> sideB : A + G
side A <- sideB : A
side A -> sideB : A + F
side A <- sideB : A
side A -> sideB : A + E
side A <- sideB : A
side A -> sideB : A + D
side A <- sideB : A
side A -> sideB : A + C
side A <- sideB : A
side A -> sideB : A + B

Total Time: 80
```

### Club Slowest Approach

```
side A -> sideB : A + B
side A <- sideB : A
side A -> sideB : F + G
side A <- sideB : B
side A -> sideB : A + B
side A -> sideB : A
side A -> sideB : D + E
side A <- sideB : B
side A -> sideB : A + B
side A -> sideB : A
side A -> sideB : A + C

Total Time: 92
```