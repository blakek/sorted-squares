func sq(_ n: Int) -> Int { return n * n }

func sortedSquares(_ numbers: [Int]) -> [Int] {
	func _sortedSquares(view: ArraySlice<Int>, accum: [Int]) -> [Int] {
		if view.count == 0 {
			return accum
		}

		if abs(view.first!) > view.last! {
			return _sortedSquares(
				view: view[view.startIndex + 1 ..< view.endIndex],
				accum: [sq(view.first!)] + accum
			)
		}

		return _sortedSquares(
			view: view[view.startIndex ..< view.endIndex - 1],
			accum: [sq(view.last!)] + accum
		)
	}

	return _sortedSquares(view: numbers[0...], accum: [])
}
