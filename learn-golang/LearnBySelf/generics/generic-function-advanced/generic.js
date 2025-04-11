// FindFirst returns the first occurrence of x in array s
function findFirst_1(s, x) {
  return s.findIndex((v) => v === x);
}

// FindAll returns all indices where x appears in array s
function findAll_1(s, x) {
  return s.reduce((indices, v, i) => {
    if (v === x) {
      indices.push(i);
    }
    return indices;
  }, []);
}

// FindBy returns the first element that matches the predicate
function findBy_1(s, predicate) {
  const found = s.find(predicate);
  return [found, found !== undefined];
}

// Filter returns a new array containing only elements that match the predicate
function filter_1(s, predicate) {
  return s.filter(predicate);
}

// Max returns the maximum element in an array
function max_1(s) {
  if (s.length === 0) {
    return [undefined, new Error("empty array")];
  }
  if (typeof s[0] === "number") {
    return [Math.max(...s), null];
  }
  return [s.reduce((a, b) => (a > b ? a : b)), null];
}

// Person_1 class with enhanced functionality
class Person_1 {
  constructor(name, age, id) {
    this.name = name;
    this.age = age;
    this.id = id;
  }

  isMatch(other) {
    return other instanceof Person_1 && this.id === other.id;
  }
}

// Main function to demonstrate usage
function main_1() {
  // Basic type examples
  const numbers = [1, 2, 3, 2, 4, 2, 5];
  console.log(`First occurrence of 2: ${findFirst_1(numbers, 2)}`);
  console.log(`All occurrences of 2: ${findAll_1(numbers, 2)}`);

  // Using findBy_1 with a predicate
  const [evenNumber, found] = findBy_1(numbers, (n) => n % 2 === 0);
  console.log(`First even number: ${evenNumber}, found: ${found}`);

  // Using filter_1
  const evenNumbers = filter_1(numbers, (n) => n % 2 === 0);
  console.log(`Even numbers: ${evenNumbers}`);

  // Using max_1
  const [maxNum, err] = max_1(numbers);
  if (!err) {
    console.log(`Maximum number: ${maxNum}`);
  }

  // Complex type example with Person_1
  const people = [
    new Person_1("Alice", 25, "001"),
    new Person_1("Bob", 30, "002"),
    new Person_1("Charlie", 35, "003"),
    new Person_1("David", 30, "004"),
  ];

  // Find people by age
  const thirtyYearOlds = filter_1(people, (p) => p.age === 30);
  console.log("People who are 30:", thirtyYearOlds);

  // Find person by custom predicate
  const [person, personFound] = findBy_1(people, (p) => p.name === "Alice");
  if (personFound) {
    console.log("Found person:", person);
  }
}

// Execute main_1 function
main_1();
