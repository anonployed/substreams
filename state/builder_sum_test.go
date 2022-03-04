package state

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/big"
	"strconv"
	"testing"
)

func TestBuilderSumBigInt(t *testing.T) {
	tests := []struct {
		name          string
		builder       *Builder
		key           string
		existingValue *big.Int
		value         *big.Int
		expectedValue *big.Int
	}{
		{
			name:          "found",
			builder:       NewBuilder("b", "", "", "", nil),
			key:           "key",
			existingValue: big.NewInt(3),
			value:         big.NewInt(4),
			expectedValue: big.NewInt(7),
		},
		{
			name:          "not found",
			builder:       NewBuilder("b", "", "", "", nil),
			key:           "key",
			existingValue: nil,
			value:         big.NewInt(4),
			expectedValue: big.NewInt(4),
		},
	}

	initTestBuilder := func(b *Builder, key string, value *big.Int) {
		b.KV = map[string][]byte{}
		if value != nil {
			b.KV[key] = []byte(value.String())
		}
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			initTestBuilder(test.builder, test.key, test.existingValue)

			test.builder.SumBigInt(0, test.key, test.value)
			actual, found := test.builder.GetAt(0, test.key)
			if !found {
				t.Errorf("value not found")
			}

			actualInt := new(big.Int).SetBytes(actual)

			assert.Equal(t, 0, actualInt.Cmp(test.expectedValue))
		})
	}
}

func TestBuilderSumInt64(t *testing.T) {
	int64ptr := func(i int64) *int64 {
		var p *int64
		p = new(int64)
		*p = i
		return p
	}

	tests := []struct {
		name          string
		builder       *Builder
		key           string
		existingValue *int64
		value         int64
		expectedValue int64
	}{
		{
			name:          "found",
			builder:       NewBuilder("b", "", "", "", nil),
			key:           "key",
			existingValue: int64ptr(3),
			value:         4,
			expectedValue: 7,
		},
		{
			name:          "not found",
			builder:       NewBuilder("b", "", "", "", nil),
			key:           "key",
			existingValue: nil,
			value:         4,
			expectedValue: 4,
		},
	}

	initTestBuilder := func(b *Builder, key string, value *int64) {
		b.KV = map[string][]byte{}
		if value != nil {
			b.KV[key] = []byte(fmt.Sprintf("%d", *value))
		}
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			initTestBuilder(test.builder, test.key, test.existingValue)

			test.builder.SumInt64(0, test.key, test.value)
			actual, found := test.builder.GetAt(0, test.key)
			if !found {
				t.Errorf("value not found")
			}

			actualInt, err := strconv.ParseInt(string(actual), 10, 64)
			assert.NoError(t, err)

			assert.Equal(t, test.expectedValue, actualInt)
		})
	}
}

func TestBuilderSumFloat64(t *testing.T) {
	float64ptr := func(i float64) *float64 {
		var p *float64
		p = new(float64)
		*p = i
		return p
	}

	tests := []struct {
		name          string
		builder       *Builder
		key           string
		existingValue *float64
		value         float64
		expectedValue float64
	}{
		{
			name:          "found",
			builder:       NewBuilder("b", "", "", "", nil),
			key:           "key",
			existingValue: float64ptr(3.0),
			value:         4.0,
			expectedValue: 7.0,
		},
		{
			name:          "not found",
			builder:       NewBuilder("b", "", "", "", nil),
			key:           "key",
			existingValue: nil,
			value:         4.0,
			expectedValue: 4.0,
		},
	}

	initTestBuilder := func(b *Builder, key string, value *float64) {
		b.KV = map[string][]byte{}
		if value != nil {
			b.KV[key] = []byte(strconv.FormatFloat(*value, 'g', 100, 64))
		}
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			initTestBuilder(test.builder, test.key, test.existingValue)

			test.builder.SumFloat64(0, test.key, test.value)
			actual, found := test.builder.GetAt(0, test.key)
			if !found {
				t.Errorf("value not found")
			}

			actualInt, err := strconv.ParseFloat(string(actual), 64)
			assert.NoError(t, err)

			assert.Equal(t, test.expectedValue, actualInt)
		})
	}
}

func TestBuilderSumBigFloat(t *testing.T) {
	tests := []struct {
		name          string
		builder       *Builder
		key           string
		existingValue *big.Float
		value         *big.Float
		expectedValue *big.Float
	}{
		{
			name:          "found",
			builder:       NewBuilder("b", "", "", "", nil),
			key:           "key",
			existingValue: big.NewFloat(3),
			value:         big.NewFloat(4),
			expectedValue: big.NewFloat(7),
		},
		{
			name:          "not found",
			builder:       NewBuilder("b", "", "", "", nil),
			key:           "key",
			existingValue: nil,
			value:         big.NewFloat(4),
			expectedValue: big.NewFloat(4),
		},
	}

	initTestBuilder := func(b *Builder, key string, value *big.Float) {
		b.KV = map[string][]byte{}
		if value != nil {
			b.KV[key] = []byte(value.Text('g', -1))
		}
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			initTestBuilder(test.builder, test.key, test.existingValue)

			test.builder.SumBigFloat(0, test.key, test.value)
			actual, found := test.builder.GetAt(0, test.key)
			if !found {
				t.Errorf("value not found")
			}

			actualInt, _, err := big.ParseFloat(string(actual), 10, 100, big.ToNearestEven)
			assert.NoError(t, err)

			assert.Equal(t, 0, actualInt.Cmp(test.expectedValue))
		})
	}
}
