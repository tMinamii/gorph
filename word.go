package goan

type Word struct {
	SurfaceForm       string // 表層形
	LeftContextID     int    // 左文脈ID
	RightContextID    int    // 右文脈ID
	Cost              int    // コスト
	WordClass         string // 品詞
	ClassDetail1      string // 品詞細分類1
	ClassDetail2      string // 品詞細分類2
	ClassDetail3      string // 品詞細分類3
	ConjugationalType string // 活用型
	ConjugationalForm string // 活用形
	OriginalForm      string // 原型
	Katakana          string // 読み
	Pronunciation     string // 発音
}

type Dict struct {
	MaxLength int
	TrieTree  *Trie
}

type Trie struct {
	Rune rune
	Word *Word
	Next map[rune]*Trie
}

type DoubleArray struct {
	Base  []int // 添字をズラす量
	Check []int
}
