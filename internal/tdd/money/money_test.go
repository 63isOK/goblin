package main

import "testing"

// 问题: 我想跟踪我的资产
// 我的资产由3支股票组成,每股多少钱,乘以总股数就是单支股票的资产
// 股票是由不同公司发行的,币种分别是欧元/美元/韩元
// 汇率为:1欧元=1.2美元,1美元=1100韩元

// tdd 思考, 问题拆为两个:单支股的资产计算; 汇率转换
// 开始时,一定要从小处开始,这是非常有必要的.

func TestMultiplication(t *testing.T) {
	type testcase struct {
		oneStock Money
		times    int
		allStock Money
	}
	table := []testcase{
		{
			oneStock: Money{amount: 5, currency: "USD"},
			times:    2,
			allStock: Money{amount: 10, currency: "USD"},
		},
		{
			oneStock: Money{amount: 5, currency: "EUR"},
			times:    2,
			allStock: Money{amount: 10, currency: "EUR"},
		},
	}

	for _, tc := range table {
		all := tc.oneStock.Times(tc.times)
		if all.amount != tc.allStock.amount {
			t.Fatalf("want %d, got %d", tc.allStock.amount, all.amount)
		}
		if all.currency != tc.allStock.currency {
			t.Fatalf("want %s, got %s", tc.allStock.currency, all.currency)
		}
	}
}

type Money struct {
	amount   int
	currency string
}

func (m Money) Times(multiplier int) Money {
	return Money{amount: m.amount * multiplier, currency: m.currency}
}
