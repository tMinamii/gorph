package gorph

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

var words = []string{
	"プリン",
	"プリン体",
	"プディング",
	"プリニー",
	"プリンセス",
	"プリティ",
	"プラージ",
	"プルミエール",
}

var w1 = []string{
	"a",
	"ac",
	"b",
	"cab",
	"cd",
}

type Dict struct {
	MaxLength int
	TrieTree  *Trie
}

// メモリ効率は、あまり考えずトライ木を抽象的に実装する
type Trie struct {
	Rune rune
	Word *Word
	Next map[rune]*Trie
}
