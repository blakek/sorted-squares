import test from "ava";
import { sortedSquares, sortedSquares__TCO } from "./index.js";

test("squares and sorts the input", (t) => {
  t.deepEqual(sortedSquares([]), []);
  t.deepEqual(sortedSquares([3]), [9]);
  t.deepEqual(sortedSquares([1, 2, 3]), [1, 4, 9]);
  t.deepEqual(sortedSquares([-2, 0, 1]), [0, 1, 4]);
  t.deepEqual(sortedSquares([-5, -3, -2]), [4, 9, 25]);
});

test("TCO version also works", (t) => {
  t.deepEqual(sortedSquares__TCO([]), []);
  t.deepEqual(sortedSquares__TCO([3]), [9]);
  t.deepEqual(sortedSquares__TCO([1, 2, 3]), [1, 4, 9]);
  t.deepEqual(sortedSquares__TCO([-2, 0, 1]), [0, 1, 4]);
  t.deepEqual(sortedSquares__TCO([-5, -3, -2]), [4, 9, 25]);

  // If JavaScript engines would follow the ES2015 standard, this would pass:
  // const overflowArray = Array.from({ length: 10000 }, (_, i) => i);
  // t.notThrows(() => sortedSquares(overflowArray));
});
