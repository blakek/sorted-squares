/**
 * Creates a sorted list of squares from the input
 * @param {number[]} input - a sorted list of numbers
 * @return {number[]} a sorted list of numbers squared
 */
export function sortedSquares(input) {
  const first = input[0];
  const last = input[input.length - 1];

  if (first === undefined) return [];

  if (Math.abs(first) > last) {
    return [...sortedSquares(input.slice(1)), first ** 2];
  }

  return [...sortedSquares(input.slice(0, -1)), last ** 2];
}

/**
 * Creates a sorted list of squares from the input
 * @param {number[]} input - a sorted list of numbers
 * @param {number[]} accum
 * @return {number[]} a sorted list of numbers squared
 */
export function sortedSquares__TCO(input, accum = []) {
  const first = input[0];
  const last = input[input.length - 1];

  if (first === undefined) return accum;

  if (Math.abs(first) > last) {
    return sortedSquares__TCO(input.slice(1), [first ** 2, ...accum]);
  }

  return sortedSquares__TCO(input.slice(0, -1), [last ** 2, ...accum]);
}
