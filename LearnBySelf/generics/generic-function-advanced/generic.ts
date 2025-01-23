// Searchable interface defines methods for searching
interface Searchable {
  isMatch(other: any): boolean;
}

// FindFirst returns the first occurrence of x in array s
function findFirst<T>(s: T[], x: T): number {
  return s.findIndex((v) => v === x);
}

// FindAll returns all indices where x appears in array s
function findAll<T>(s: T[], x: T): number[] {
  return s.reduce((indices: number[], v, i) => {
    if (v === x) {
      indices.push(i);
    }
    return indices;
  }, []);
}

// FindBy returns the first element that matches the predicate
function findBy<T>(
  s: T[],
  predicate: (item: T) => boolean
): [T | undefined, boolean] {
  const found = s.find(predicate);
  return [found, found !== undefined];
}

// Filter returns a new array containing only elements that match the predicate
function filter<T>(s: T[], predicate: (item: T) => boolean): T[] {
  return s.filter(predicate);
}

// Max returns the maximum element in an array
function max<T extends number | string | bigint>(
  s: T[]
): [T | undefined, Error | null] {
  if (s.length === 0) {
    return [undefined, new Error("empty array")];
  }
  if (typeof s[0] === "number") {
    return [Math.max(...(s as number[])) as T, null];
  }
  return [s.reduce((a, b) => (a > b ? a : b)), null];
}

// Person class with enhanced functionality
class Person implements Searchable {
  constructor(public name: string, public age: number, public id: string) {}

  isMatch(other: any): boolean {
    return other instanceof Person && this.id === other.id;
  }
}

// Main function to demonstrate usage
function main() {
  // Basic type examples
  const numbers = [1, 2, 3, 2, 4, 2, 5];
  console.log(`First occurrence of 2: ${findFirst(numbers, 2)}`);
  console.log(`All occurrences of 2: ${findAll(numbers, 2)}`);

  // Using findBy with a predicate
  const [evenNumber, found] = findBy(numbers, (n) => n % 2 === 0);
  console.log(`First even number: ${evenNumber}, found: ${found}`);

  // Using filter
  const evenNumbers = filter(numbers, (n) => n % 2 === 0);
  console.log(`Even numbers: ${evenNumbers}`);

  // Using max
  const [maxNum, err] = max(numbers);
  if (!err) {
    console.log(`Maximum number: ${maxNum}`);
  }

  // Complex type example with Person
  const people = [
    new Person("Alice", 25, "001"),
    new Person("Bob", 30, "002"),
    new Person("Charlie", 35, "003"),
    new Person("David", 30, "004"),
  ];

  // Find people by age
  const thirtyYearOlds = filter(people, (p) => p.age === 30);
  console.log("People who are 30:", thirtyYearOlds);

  // Find person by custom predicate
  const [person, personFound] = findBy(people, (p) => p.name === "Alice");
  if (personFound) {
    console.log("Found person:", person);
  }
}

// Execute main function
main();
