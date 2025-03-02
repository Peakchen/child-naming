package naming

import (
	"fmt"

	"github.com/Peakchen/child-naming/pkg/types"
)

func GetWordPinyin(word string, ps []*types.Pinyin) (*types.Pinyin, error) {
	for _, p := range ps {
		if word == p.Word {
			return p, nil
		}
	}
	return nil, fmt.Errorf("%s pinyin not find.", word)
}
