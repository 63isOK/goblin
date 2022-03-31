package main

import "testing"

// 问题: 我想跟踪我的资产
// 我的资产由3支股票组成,每股多少钱,乘以总股数就是单支股票的资产
// 股票是由不同公司发行的,币种分别是欧元/美元/韩元
// 汇率为:1欧元=1.2美元,1美元=1100韩元

// tdd 思考, 问题拆为两个:单支股的资产计算; 汇率转换
// 开始时,一定要从小处开始,这是非常有必要的.

func TestMultiplication(t *testing.T) {
	// 美元股的资产计算,两股5美元的股票是10美元的资产
	fiver := Dollar{
		amount: 5,
	}
	tenner := fiver.Times(2)
	if tenner.amount != 10 {
		t.Fatalf("want 10, get %d", tenner.amount)
	}
}

type Dollar struct {
	amount int
}

func (d Dollar) Times(multiplier int) Dollar {
	return Dollar{amount: d.amount * multiplier}
}
