package dizhi

import (
	"github.com/chonglou/destiny/yao"
)

//YinYang 阴阳
func (p Type) YinYang() yao.Type {
	switch p {
	case Zi:
		return yao.Yang
	case Chou:
		return yao.Yin
	case Yin:
		return yao.Yang
	case Mao:
		return yao.Yin
	case Chen:
		return yao.Yang
	case Si:
		return yao.Yin
	case Wu:
		return yao.Yang
	case Wei:
		return yao.Yin
	case Shen:
		return yao.Yang
	case You:
		return yao.Yin
	case Xu:
		return yao.Yang
	case Hai:
		return yao.Yin
	default:
		return ""
	}
}
