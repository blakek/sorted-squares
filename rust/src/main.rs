fn sorted_squares_recursive(mut input: Vec<i64>, mut accum: Vec<i64>) -> Vec<i64> {
    // The entire input has been used. Return the result.
    if input.is_empty() {
        return accum;
    }

    let (&first, rest) = input.split_first().unwrap();

    if rest.is_empty() {
        accum.insert(0, first.pow(2));
        return accum;
    }
    let (&last, _) = rest.split_last().unwrap();

    if first.abs() > last {
        input.remove(0);
        accum.insert(0, first.pow(2));
    } else {
        input.pop();
        accum.insert(0, last.pow(2));
    }

    return sorted_squares_recursive(input, accum);
}

fn sorted_squares(input: Vec<i64>) -> Vec<i64> {
    sorted_squares_recursive(input, vec![])
}

fn main() {
    println!("{:?}", sorted_squares(vec![-3, 1, 2, 5]));
}
