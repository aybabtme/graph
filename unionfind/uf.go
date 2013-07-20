package unionfind

// UF is a union find that knows of connections between entries. It implements
// a weighted quick union find datastucture
type UF struct {
	id    []int
	sz    []int
	count int
}

// BuildUF initialize n sites with integer names (0 to N-1)
func BuildUF(n int) UF {
	uf := UF{
		id:    make([]int, n),
		sz:    make([]int, n),
		count: n,
	}

	for i := 0; i < n; i++ {
		uf.id[i] = i
		uf.sz[i] = 1
	}
	return uf
}

// Union adds a connection between p and q
func (u *UF) Union(p, q int) {
	i := u.Find(p)
	j := u.Find(q)

	if i == j {
		return
	}

	if u.sz[i] < u.sz[j] {
		u.id[i] = j
		u.sz[j] += u.sz[i]
	} else {
		u.id[j] = i
		u.sz[i] += u.sz[j]
	}
	u.count--
}

// Find tells the component identifier for p (0 to N-1)
func (u *UF) Find(p int) int {
	for {
		if p != u.id[p] {
			p = u.id[p]
		} else {
			return p
		}
	}
}

// Connected is true if p and q are in the same component
func (u *UF) Connected(p, q int) bool { return u.Find(p) == u.Find(q) }

// Count tells the number of components
func (u *UF) Count() int { return u.count }
