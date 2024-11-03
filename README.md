# Amicable Numbers

This project is an **amicable numbers calculator** that explores pairs of numbers within a specified range, identifying pairs for which each number is the sum of the other’s proper divisors. The calculator has been implemented in two versions, **Python** and **Go**, to run both sequentially and in parallel.

## Project Description

Two numbers are defined as **amicable numbers** if, given a number \(A\), the sum of its proper divisors (excluding the number itself) equals a number \(B\), and vice versa. Therefore, the pair (A, B) is said to be amicable. A classic example is the pair 220 and 284:

- The proper divisors of 220 are: 1, 2, 4, 5, 10, 11, 20, 22, 44, 55, and 110, summing up to 284.
- The proper divisors of 284 are: 1, 2, 4, 71, and 142, summing up to 220.

### How It Works

The calculator checks the sum of each number’s proper divisors within a user-specified range and determines if they form an amicable pair. The algorithm has been optimized to avoid duplicate checks and reduce unnecessary calculations.

## Requirements

- **Python**: Version 3.8 or later, to run the Python version.
- **Go**: Version 1.16 or later, to run the Go version.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/Loki-it/Amicable-Numbers.git
   cd Amicable-Numbers
   ```

## Running the Code

### Running the Python Version

1. Navigate to the project directory.
2. Run the Python script:
   ```bash
   python3 main.py
   ```
3. The program will display the found amicable pairs and the execution time.

### Running the Go Version

1. Navigate to the project directory.
2. Compile and run the Go script:
   ```bash
   go run main.go
   ```
3. Upon completion, the program will display the found amicable pairs and the execution time.

## Project Notes

This implementation of the amicable numbers calculator was developed in two different languages, **Python** and **Go**, to explore performance differences in intensive computing scenarios. The use of `multiprocessing` in Python and **goroutines** in Go allowed optimization of the calculations in both languages, leveraging parallelization to improve execution times.

## Performance Comparison

In this project, execution times between Python and Go are compared to demonstrate performance differences between the two languages:

| Search Limit      | Go Time (seconds) | Python Time (seconds) | Go Speed Advantage |
|-------------------|-------------------|------------------------|--------------------|
| \(10^6\)          | 0.130             | 3.238                 | ~25x              |
| \(10^7\)          | 4.031             | 94.182                | ~23x              |

Developed by [Loki-it](https://github.com/Loki-it) - Tested on a Ryzen 9 5900X
