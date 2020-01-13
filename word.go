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
	Base  []int // 添字をズラす量,葉ノードの場合はずらす量の代わりに単語に対応する値(単語ID)
	Check []int // 親のID(遷移元)
}

func (d *DoubleArray) NewDoubleArray() {
	d.Base = []int{1}
	d.Check = []int{0}
}

func (d *DoubleArray) Build(word string) {
	for _, c := range word {
		d.Base[s] + d.Code(c)
	}
}

func (d *DoubleArray) Code(r rune) {

}
