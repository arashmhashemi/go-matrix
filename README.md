# go-matrix
## Input
The input file is a csv file containing equal rows and columns with positive integer values.

Given the following example,
```
1,2,3
4,5,6
7,8,9
```
## APIs
1. echo: returns the matrix as a string in matrix format, expected output would be
```
    1,2,3
    4,5,6
    7,8,9
``` 
2. invert: returns the matrix as a string in matrix format where the columns and rows are inverted, expected output would be
```
    1,4,7
    2,5,8
    3,6,9
``` 
3. flatten: returns the matrix as a 1 line string, with values separated by commas, expected output would be
```
    1,2,3,4,5,6,7,8,9    
``` 
4. sum: returns the sum of the integers in the matrix, expected output would be:
```
    45    
``` 
5. multiply: returns the product of the integers in the matrix. If the output is too large and won't fit in 64bit, it'll produce an "Overflow error" otherwise expected output would be:
```
    362880
``` 
## To run
### Demo: https://go-matrix.azurewebsites.net 

```bash
curl -F 'file=@./matrix.csv' "https://go-matrix.azurewebsites.net/echo"
curl -F 'file=@./matrix.csv' "https://go-matrix.azurewebsites.net/invert"
curl -F 'file=@./matrix.csv' "https://go-matrix.azurewebsites.net/flatten"
curl -F 'file=@./matrix.csv' "https://go-matrix.azurewebsites.net/sum"
curl -F 'file=@./matrix.csv' "https://go-matrix.azurewebsites.net/multiply"
```
### Docker: 
```bash
docker run --rm -p 8080:80 --name go-matrix  amhashemi/go-matrix 
```

```bash
curl -F 'file=@./matrix.csv' "localhost:8080/echo"
curl -F 'file=@./matrix.csv' "localhost:8080/invert"
curl -F 'file=@./matrix.csv' "localhost:8080/flatten"
curl -F 'file=@./matrix.csv' "localhost:8080/sum"
curl -F 'file=@./matrix.csv' "localhost:8080/multiply"
```

### Code: 

```bash
git clone https://github.com/arashmhashemi/go-matrix.git
cd go-matrix

go run .
```
```bash
curl -F 'file=@./matrix.csv' "localhost:8080/echo"
curl -F 'file=@./matrix.csv' "localhost:8080/invert"
curl -F 'file=@./matrix.csv' "localhost:8080/flatten"
curl -F 'file=@./matrix.csv' "localhost:8080/sum"
curl -F 'file=@./matrix.csv' "localhost:8080/multiply"
```
# Test: 
```bash
git clone https://github.com/arashmhashemi/go-matrix.git
cd go-matrix

go test
```
