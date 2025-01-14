// Define a simple calculator object
const calculator = {
  add: (a, b) => a + b,
  subtract: (a, b) => a - b,
  multiply: (a, b) => a * b
};

// Do some calculations and store results
const results = {
  sum: calculator.add(5, 3),
  difference: calculator.subtract(10, 4),
  product: calculator.multiply(6, 2)
};

// Store these in a global variable so we can access it from Go
globalResults = results;
