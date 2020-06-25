@testable import sortedSquares
import XCTest

final class swiftTests: XCTestCase {
	func testSortedSquares() {
		XCTAssertEqual(sortedSquares([]), [])
		XCTAssertEqual(sortedSquares([1, 2, 3]), [1, 4, 9])
		XCTAssertEqual(sortedSquares([-3, -2, -1]), [1, 4, 9])
		XCTAssertEqual(sortedSquares([-9, -5, -1, 0, 1, 88]), [0, 1, 1, 25, 81, 7744])
	}

	static var allTests = [
		("returns sorted squares", testSortedSquares),
	]
}
