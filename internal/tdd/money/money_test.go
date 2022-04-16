package main_test

import "testing"

// 问题: 我想跟踪我的资产
// 我的资产由3支股票组成,每股多少钱,乘以总股数就是单支股票的资产
// 股票是由不同公司发行的,币种分别是欧元/美元/韩元
// 汇率为:1欧元=1.2美元,1美元=1100韩元

// tdd 思考, 问题拆为两个:单支股的资产计算; 汇率转换
// 开始时,一定要从小处开始,这是非常有必要的.

func TestMultiplicationUSD(t *testing.T) {
	t.Parallel()
	testMultiplication(t, "USD")
}

func TestMultiplicationEUR(t *testing.T) {
	t.Parallel()
	testMultiplication(t, "EUR")
}

func TestMultiplicationKRW(t *testing.T) {
	t.Parallel()
	testMultiplication(t, "KRW")
}

func testMultiplication(t *testing.T, currency string) {
	type testcase struct {
		name  string
		m     Money
		times int
		want  Money
	}

	table := []testcase{
		{
			name:  "5元的2倍是10元",
			m:     Money{amount: 5, currency: currency},
			times: 2,
			want:  Money{amount: 10, currency: currency},
		},
		{
			name:  "5元的6倍是30元",
			m:     Money{amount: 5, currency: currency},
			times: 6,
			want:  Money{amount: 30, currency: currency},
		},
	}

	for _, tc := range table {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.m.Times(tc.times)
			assertEqual(t, tc.want, got)
		})
	}
}

func TestDivisionUSD(t *testing.T) {
	t.Parallel()
	testDivision(t, "USD")
}

func TestDivisionEUR(t *testing.T) {
	t.Parallel()
	testDivision(t, "EUR")
}

func TestDivisionKRW(t *testing.T) {
	t.Parallel()
	testDivision(t, "KRW")
}

func testDivision(t *testing.T, currency string) {
	type testcase struct {
		m       Money
		divisor int
		want    Money
	}

	table := []testcase{
		{
			m:       Money{amount: 5, currency: currency},
			divisor: 2,
			want:    Money{amount: 2.5, currency: currency},
		},
		{
			m:       Money{amount: 4002, currency: currency},
			divisor: 4,
			want:    Money{amount: 1000.5, currency: currency},
		},
	}

	for _, tc := range table {
		got := tc.m.Divide(tc.divisor)
		assertEqual(t, tc.want, got)
	}
}

func TestAddition(t *testing.T) {
	t.Parallel()

	fiveDollars := Money{amount: 5, currency: "USD"}
	tenDollars := Money{amount: 10, currency: "USD"}
	fifteenDollars := Money{amount: 15, currency: "USD"}

	var portfolio Portfolio
	portfolio = portfolio.Add(fiveDollars).Add(tenDollars)
	portfolioInDollars := portfolio.Evaluate("USD")

	assertEqual(t, fifteenDollars, portfolioInDollars)
}

type Money struct {
	amount   float64
	currency string
}

func (m Money) Times(multiplier int) Money {
	return Money{amount: m.amount * float64(multiplier), currency: m.currency}
}

func (m Money) Divide(divisor int) Money {
	return Money{
		amount:   m.amount / float64(divisor),
		currency: m.currency,
	}
}

type Portfolio []Money

func (p Portfolio) Add(m Money) Portfolio {
	p = append(p, m)
	return p
}

func (p Portfolio) Evaluate(currency string) Money {
	total := 0.0
	for _, m := range p {
		total += m.amount
	}
	return Money{
		amount:   total,
		currency: currency,
	}
}

func assertEqual(t *testing.T, want, got Money) {
	t.Helper()

	if want != got {
		t.Errorf("want: %+v, got: %+v", want, got)
	}
}
