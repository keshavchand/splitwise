package simplifier

type MinOwes []Owes

func (m MinOwes) Len() int {
	return len(m)
}

func (m MinOwes) Less(i, j int) bool {
	return m[i].Amount < m[j].Amount
}

func (m MinOwes) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinOwes) Pop() any {
	x := (*m)[0]
	*m = (*m)[1:]
	return x
}

func (m *MinOwes) Push(_x any) {
	x := _x.(Owes)
	*m = append(*m, x)
}

type MaxOwes []Owes

func (m MaxOwes) Len() int {
	return len(m)
}

func (m MaxOwes) Less(i, j int) bool {
	return m[i].Amount > m[j].Amount
}

func (m MaxOwes) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MaxOwes) Pop() any {
	x := (*m)[0]
	*m = (*m)[1:]
	return x
}

func (m *MaxOwes) Push(_x any) {
	x := _x.(Owes)
	*m = append(*m, x)
}
