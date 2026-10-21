package calc

import (
	"math/rand"
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2,3) = %d, want 5", got)
	}
}

func TestMul(t *testing.T) {
	if got := Mul(4, 5); got != 20 {
		t.Fatalf("Mul(4,5) = %d, want 20", got)
	}
}

func TestIsEven(t *testing.T) {
	if !IsEven(6) || IsEven(7) {
		t.Fatal("IsEven returned wrong result")
	}
}

func TestFlakyAddUnderLoad(t *testing.T) {
	if rand.Float64() < 0.4 {
		t.Fatal("flaky: simulated timeout while adding")
	}
	if got := Add(1, 1); got != 2 {
		t.Fatalf("Add(1,1) = %d, want 2", got)
	}
}

func TestFlakyMulUnderContention(t *testing.T) {
	if rand.Float64() < 0.4 {
		t.Fatal("flaky: simulated contention in multiplier")
	}
	if got := Mul(3, 3); got != 9 {
		t.Fatalf("Mul(3,3) = %d, want 9", got)
	}
}

func TestFlakyFailsOnFirstAttemptOnly(t *testing.T) {
	if os.Getenv("GITHUB_RUN_ATTEMPT") == "1" {
		t.Fatal("flaky: simulated first-attempt failure (GITHUB_RUN_ATTEMPT=1)")
	}
	if got := Add(1, 1); got != 2 {
		t.Fatalf("Add(1,1) = %d, want 2", got)
	}
}
