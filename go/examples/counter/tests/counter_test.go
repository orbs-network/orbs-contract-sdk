package main //TODO: package name convention
import (
	"log"
	"os/exec"
	"strings"
	"testing"
)

//TODO there's run (read-only) or sendTransaction (change the blockchain state)

// Stages in the test: (action -> expected result)
// 	1. Deploy the counter contract, it also calls _init method -> "contract was successfully deployed to the blockchain" / 0
//	2. call add(5) -> 5 or nothing (success)
//	3. call get() -> 5 (success)
//	4. call add(-3) -> error (success) TODO: will it though an error?
//	5. call add(2 << 66) -> error (success) TODO: ...

// TODO: start and stop the blockchain

func prepareMsg(a, b, c, d, e string) string {
	var msg = "Test no. [1] Test method: [2] , Expected result: [3] Actual Result : [4] , Test passed: [5]"
	msg = strings.Replace(msg, "[1]", a, 1)
	msg = strings.Replace(msg, "[2]", b, 1)
	msg = strings.Replace(msg, "[3]", c, 1)
	msg = strings.Replace(msg, "[4]", d, 1)
	msg = strings.Replace(msg, "[5]", e, 1)
	return msg
}

func main() {

	// Test 1
	out, err := exec.Command("kitsat-cli", "-deploy counter_deploy.json").Output() //TODO esti

	var msg string

	if out == nil {
		msg = prepareMsg("1", "deploy", "0", "0", "Yes")
	} else {
		msg = prepareMsg("1", "deploy", "0", 0, "No")
	}
	log.Println(msg)

	// Test 2
	out2, err2 := exec.Command("kitsat-cli", "run counter_add5.json").Output()

	if out2 == nil {
		msg = prepareMsg("2", "add(5)", "5", "5", "Yes")
	} else {
		msg = prepareMsg("2", "add(5)", "5", string(out2), "No")
	}
	log.Println(msg)

	// Test 3
	out3, err3 := exec.Command("kitsat-cli", "run counter_get.json").Output()

	if out3 == nil {
		msg = prepareMsg("3", "get", "5", "5", "Yes")
	} else {
		msg = prepareMsg("3", "get", "5", string(out3), "No")
	}
	log.Println(msg)

	// Test 4
	out4, err4 := exec.Command("kitsat-cli", "run counter_add_negative.json").Output()

	if out4 == nil {
		msg = prepareMsg("4", "add(-3)", "negative number error", "negative number error", "Yes")
	} else {
		msg = prepareMsg("4", "add(-3)", "negative number error", string(out4), "No")
	}
	log.Println(msg)

	// Test 5
	out5, err5 := exec.Command("kitsat-cli", "run counter_add_overflow.json").Output()

	if out5 == nil {
		msg = prepareMsg("5", "add(2 << 66)", "overflow error", "overflow error", "Yes")
	} else {
		msg = prepareMsg("5", "add(2 << 66)", "overflow error", string(out5), "No")
	}
	log.Println(msg)
}

func TestCounterContractWorkingAsExpected(t *testing.T) {
	// TODO: fix it
}
