package helper

type Entry struct {
	Key   string
	Value int64
}

type Entries []Entry

func (e Entries) Len() int           { return len(e) }
func (e Entries) Less(i, j int) bool { return e[i].Value > e[j].Value }
func (e Entries) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
