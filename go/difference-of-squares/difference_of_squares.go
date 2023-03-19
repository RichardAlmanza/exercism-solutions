package diffsquares

type memoizationMap map[int]int
type sumMemoization struct{
	squareOfSum memoizationMap
	sumOfSquares memoizationMap
}

var (
	singleton *sumMemoization
)

func NewSumMemoization() *sumMemoization {
	if singleton == nil {
		singleton = new(sumMemoization)
		singleton.squareOfSum = make(memoizationMap)
		singleton.sumOfSquares = make(memoizationMap)
	}

	return singleton
}

func (mm memoizationMap) get(n int, formula func(int)int) int{
	value, exists := mm[n]

	if !exists {
		value = formula(n)

		mm[n] = value
	}

	return value
}

func (sm *sumMemoization) GetSquareOfSum(n int) int {
	return sm.squareOfSum.get(n, func(i int) int {
		value := i * (i + 1) / 2
		value = value * value

		return value
	})
}

func (sm *sumMemoization) GetSumOfSquares(n int) int {
	return sm.sumOfSquares.get(n, func(i int) int {
		return i * (i + 1) * (2 * i + 1) / 6
	})
}

func SquareOfSum(n int) int {
	return NewSumMemoization().GetSquareOfSum(n)
}

func SumOfSquares(n int) int {
	return NewSumMemoization().GetSumOfSquares(n)
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
