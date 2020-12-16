#include <stdio.h>
#include <stdlib.h>

typedef struct vector_ {
	int *data;
	size_t capacity;
	size_t size;
} vector;

vector *vector_create_empty() {
	vector *v = malloc(sizeof(vector));
	v->capacity = 0;
	v->size = 0;
	return v;
}

vector *vector_create(int size, int *data) {
	vector *v = malloc(sizeof(vector));

	v->capacity = size;
	v->size = size;
	v->data = malloc(size * sizeof(int));

	for (int i = 0; i < size; i++) {
		v->data[i] = data[i];
	}

	return v;
}

void vector_grow(vector *v) {
	size_t new_capacity = v->capacity == 0 ? 2 : v->capacity * 2;
	v->data = realloc(v->data, sizeof(int *) * new_capacity);
	v->capacity = new_capacity;
}

int vector_pop(vector *v) {
	if (v->size > 0) {
		v->size--;
		return v->data[v->size - 1];
	}
	return 0;
}

void vector_push(vector *v, int element) {
	while (v->size + 1 >= v->capacity) {
		vector_grow(v);
	}

	v->data[v->size] = element;
	v->size++;
}

int vector_shift(vector *v) {
	if (v->size > 0) {

		int removedValue = v->data[0];

		for (int i = 1; i < v->size; i++) {
			v->data[i - 1] = v->data[i];
		}

		v->size--;

		return removedValue;
	}

	return 0;
}

void vector_unshift(vector *v, int element) {
	size_t nextSize = v->size + 1;

	while (nextSize >= v->capacity) {
		vector_grow(v);
	}

	for (int i = v->size; i > 0; i--) {
		v->data[i] = v->data[i - 1];
	}

	v->data[0] = element;
	v->size = nextSize;
}

void vector_free(vector *v) {
	free(v->data);
	free(v);
}

void vector_print(vector *v) {
	printf("{");

	for (int i = 0; i < v->size; i++) {
		printf(" %d", v->data[i]);
	}

	printf(" }\n");
}
