package ex2_test

import (
	"testing"

	"github.com/StasMerzlyakov/yagoadvanst/cmd/sprint7/task4/ex2"
)

func BenchmarkBalancer(b *testing.B) {

	const connsN = 100
	var cons []*ex2.Connection
	for i := 0; i < connsN; i++ {
		cons = append(cons, &ex2.Connection{})
	}

	bln1 := ex2.NewLoadBalancerChan(cons)
	bln1.Init()
	defer bln1.Close()

	bln2 := ex2.NewLoadBalancerAtomic(cons)

	bln3 := ex2.NewLoadBalancerMutex(cons)

	b.ResetTimer()
	b.Run("chan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bln1.NextConn()
		}

	})

	b.Run("atom", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bln2.NextConn()
		}
	})

	b.Run("mtx", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bln3.NextConn()
		}
	})
}

func BenchmarkLoadBalancer(b *testing.B) {
	const connsN = 100
	const triesN = 1000

	conns := make([]*ex2.Connection, 0, connsN)
	for i := 0; i < connsN; i++ {
		conns = append(conns, &ex2.Connection{})
	}

	lbChan := ex2.NewLoadBalancerChan(conns)
	lbChan.Init()
	defer lbChan.Close()
	lbAtomic := ex2.NewLoadBalancerAtomic(conns)
	lbMutex := ex2.NewLoadBalancerMutex(conns)

	b.ResetTimer()

	b.Run("chan", func(b *testing.B) {
		for i := 0; i < triesN; i++ {
			lbChan.NextConn()
		}
	})

	b.Run("atomic", func(b *testing.B) {
		for i := 0; i < triesN; i++ {
			lbAtomic.NextConn()
		}
	})

	b.Run("mutex", func(b *testing.B) {
		for i := 0; i < triesN; i++ {
			lbMutex.NextConn()
		}
	})
}
