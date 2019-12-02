package fizzbuzz

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMultiple(t *testing.T) {
	tt := []struct {
		lhr        int
		rhn        int
		isMultiple bool
	}{
		{0, 0, true},
		{12, 0, false},
		{0, 100, true},
		{120, 5, true},
		{1500, 3, true},
		{45, 7, false},
	}
	for _, tc := range tt {
		if isMultiple(tc.lhr, tc.rhn) != tc.isMultiple {
			t.Errorf("isMultiple(%v, %v) should be %v", tc.lhr, tc.rhn, tc.isMultiple)
		}
	}
}

func TestSequence(t *testing.T) {
	tt := []struct {
		parameters Parameters
		sequence   []string
	}{
		{Parameters{0, 0, 0, "niladic", "nil"}, []string{}},
		{Parameters{0, 0, 10, "Never", "SayNever"}, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}},
		{Parameters{1, 1, 10, "Salut", "Cava"}, []string{"SalutCava", "SalutCava", "SalutCava", "SalutCava", "SalutCava", "SalutCava", "SalutCava", "SalutCava", "SalutCava", "SalutCava"}},
		{Parameters{2, 3, 10, "Bonnie", "Clyde"}, []string{"1", "Bonnie", "Clyde", "Bonnie", "5", "BonnieClyde", "7", "Bonnie", "Clyde", "Bonnie"}},
	}
	for _, tc := range tt {
		seq, err := Sequence(&tc.parameters)
		assert.NoError(t, err)
		if !reflect.DeepEqual(seq, tc.sequence) {
			t.Errorf(
				"Sequence(%v) should be %v",
				tc.parameters,
				tc.sequence,
			)
		}
	}
}

func TestWithMaxLimit(t *testing.T) {
	mock := func(params *Parameters) ([]string, error) { return nil, nil }
	p := &Parameters{}
	t.Run("good limit returns no error", func(t *testing.T) {
		p.Limit = 99
		_, err := WithMaxLimit(mock, 100)(p)
		assert.NoError(t, err)
	})
	t.Run("exceeding limit must return an error", func(t *testing.T) {
		p.Limit = 101
		_, err := WithMaxLimit(mock, 100)(p)
		assert.Error(t, err)
	})
}
