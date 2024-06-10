# Math-skills

Math-skills is a program written in Go that reads data from a file and prints the results of its average, median, variance and standard deviation.

## Features

- Read data from a file.
- Execute calculation for average, median, variance and standard deviation.
- Print the results as rounded integers.

## Installation

1. Clone the repository:

    ```bash
    git clone https://learn.zone01kisumu.ke/git/togondol/math-skills
    ```

2. Navigate to the project directory:

    ```bash
    cd math-skills
    ```

## Usage

To get the results of the data in the file, run the following command:

```go
$ go run . "file name"
```

Example:

```go
$ go run . data.txt
```

Output:

```bash
Average: 132
Median: 117
Variance: 785
Standard Deviation: 28              
```
The above numbers are only examples.

## Testing 
To run the tests present do the following:

1. Open the folder using this command:

```
cd math-skills/
```
2. Run the test using this command:

```
go test -v
```


## Contributing

If you have suggestions for improvements, bug fixes, or new features, feel free to open an issue or submit a pull request.

## Author

This project was build and maintained by:

[Thadeus Ogondola](https://learn.zone01kisumu.ke/git/togondol/)
