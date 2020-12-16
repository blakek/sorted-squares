#include "./int-vector.h"
#include <stdio.h>

int abs(int num) { return num < 0 ? -num : num; }

vector *sortedSquaresRecursive(vector *nums, vector *accum) {
	// Given an empty input.
	if (nums->size == 0) {
		return accum;
	}

	// Only one element remains. Add it and return the result.
	if (nums->size == 1) {
		vector_unshift(accum, nums->data[0] * nums->data[0]);
		return accum;
	}

	int first = nums->data[0];
	int last = nums->data[nums->size - 1];

	if (abs(first) > last) {
		vector_unshift(accum, first * first);
		vector_shift(nums);
	} else {
		vector_unshift(accum, last * last);
		vector_pop(nums);
	}

	return sortedSquaresRecursive(nums, accum);
}

vector *sortedSquares(vector *nums) {
	vector *initialAccumulator = vector_create_empty();
	vector *result = sortedSquaresRecursive(nums, initialAccumulator);

	vector_free(initialAccumulator);

	return result;
}

int main(int argc, char **argv) {
	int data[] = {-3, 1, 2, 5, 9};
	vector *nums = vector_create(5, data);
	vector *result = sortedSquares(nums);

	vector_print(result);
}
