# Interesting Numbers Analyzer

One of my friends is interested in numerology, and often shares "sigificant" numbers that she sees out in the world. This is obviously due to random chance, but for fun, I decided to see what percentage of all numbers would be considered "interesting".

## Features

- Analyzes numbers up to a specified number of digits
- Determines if a number is interesting based on three criteria:
  1. Palindrome: The number reads the same forwards and backwards
  2. Straight: The number contains a sequence of consecutive ascending or descending digits
  3. Repeat: The number contains a sequence of repeated digits
- Configurable strictness for straight and repeat tests
- Calculates and displays the percentage of interesting numbers in the given range

## Usage

To run the program:
go run main.go
Copy
The output will show the percentage of interesting numbers for the specified number of digits.

## Customization

You can adjust the following parameters in the code:

- `digits`: Change this to analyze different ranges of numbers (e.g., 3-digit numbers, 7-digit numbers, etc.)
- `straightStrictness`: Adjust how many consecutive ascending/descending digits are required for a number to be considered a "straight"
- `repeatStrictness`: Adjust how many repeated digits are required for a number to be considered a "repeat"

## Future Improvements

- Accept command-line arguments for number of digits and strictness levels
- Add more criteria for "interesting" numbers
- Optimize performance for larger ranges of numbers
- Implement graphing functionality to visualize the distribution of interesting numbers

## License

[MIT License](LICENSE)