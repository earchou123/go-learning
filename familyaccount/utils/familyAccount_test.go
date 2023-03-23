package utils

import (
	"testing"
)

func TestIncome(t *testing.T) {
	type Data struct {
		note  string
		input FamilyAccount
	}
	tests := []Data{
		{note: "常规用例", input: FamilyAccount{
			balance:  10000.00,
			money:    100.0,
			note:     "test1",
			loop:     true,
			details:  "收支\t账户余额\t收支金额\t 说 明",
			flag:     false,
			key:      "",
			username: "admin",
			pwd:      "123",
		},
		},
		{note: "金额为0", input: FamilyAccount{
			balance:  10000.00,
			money:    0.00,
			note:     "test2",
			loop:     true,
			details:  "收支\t账户余额\t收支金额\t 说 明",
			flag:     false,
			key:      "",
			username: "admin",
			pwd:      "123",
		},
		},
		{note: "金额为负数", input: FamilyAccount{
			balance:  10000.00,
			money:    -4.00,
			note:     "test2",
			loop:     true,
			details:  "收支\t账户余额\t收支金额\t 说 明",
			flag:     false,
			key:      "",
			username: "admin",
			pwd:      "123",
		},
		},
	}
	for _, test := range tests {
		old_balance := test.input.balance
		test.input.income()
		isPass := true
		if test.input.money < 0 {
			isPass = (test.input.balance == old_balance)

		} else {
			isPass = (test.input.balance == (old_balance + test.input.money))
		}

		if !isPass {
			t.Errorf("failed: %v\n", test.note)
		}
	}

}

func TestOncome(t *testing.T) {
	type Data struct {
		note  string
		input FamilyAccount
	}
	tests := []Data{
		{note: "常规用例", input: FamilyAccount{
			balance:  10000.00,
			money:    100.0,
			note:     "test1",
			loop:     true,
			details:  "收支\t账户余额\t收支金额\t 说 明",
			flag:     false,
			key:      "",
			username: "admin",
			pwd:      "123",
		},
		},
		{note: "金额为0", input: FamilyAccount{
			balance:  10000.00,
			money:    0.00,
			note:     "test2",
			loop:     true,
			details:  "收支\t账户余额\t收支金额\t 说 明",
			flag:     false,
			key:      "",
			username: "admin",
			pwd:      "123",
		},
		},
		{note: "金额为负数", input: FamilyAccount{
			balance:  10000.00,
			money:    -4.00,
			note:     "test3",
			loop:     true,
			details:  "收支\t账户余额\t收支金额\t 说 明",
			flag:     false,
			key:      "",
			username: "admin",
			pwd:      "123",
		},
		},
		{note: "余额不足", input: FamilyAccount{
			balance:  10000.00,
			money:    10000.10,
			note:     "test4",
			loop:     true,
			details:  "收支\t账户余额\t收支金额\t 说 明",
			flag:     false,
			key:      "",
			username: "admin",
			pwd:      "123",
		},
		},
	}
	for _, test := range tests {
		old_balance := test.input.balance
		test.input.outcome()
		isPass := true
		if test.input.money < 0 || old_balance < test.input.money {
			isPass = (test.input.balance == old_balance)
		} else {
			isPass = (test.input.balance == (old_balance - test.input.money))
		}

		if !isPass {
			t.Errorf("failed: %v\n", test.note)
		}
	}

}

func TestTransfer(t *testing.T) {
	type Data struct {
		note  string
		input FamilyAccount
	}
	tests := []Data{
		{note: "常规用例", input: FamilyAccount{
			balance:  10000.00,
			money:    100.0,
			note:     "test1",
			loop:     true,
			details:  "收支\t账户余额\t收支金额\t 说 明",
			flag:     false,
			key:      "",
			username: "admin",
			pwd:      "123",
		},
		},
		{note: "金额为0", input: FamilyAccount{
			balance:  10000.00,
			money:    0.00,
			note:     "test2",
			loop:     true,
			details:  "收支\t账户余额\t收支金额\t 说 明",
			flag:     false,
			key:      "",
			username: "admin",
			pwd:      "123",
		},
		},
		{note: "金额为负数", input: FamilyAccount{
			balance:  10000.00,
			money:    -4.00,
			note:     "test2",
			loop:     true,
			details:  "收支\t账户余额\t收支金额\t 说 明",
			flag:     false,
			key:      "",
			username: "admin",
			pwd:      "123",
		},
		},
		{note: "余额不足", input: FamilyAccount{
			balance:  10000.00,
			money:    10000.10,
			note:     "test4",
			loop:     true,
			details:  "收支\t账户余额\t收支金额\t 说 明",
			flag:     false,
			key:      "",
			username: "admin",
			pwd:      "123",
		},
		},
	}
	for _, test := range tests {
		old_balance := test.input.balance
		test.input.transfer()
		isPass := true
		if test.input.money < 0 || old_balance < test.input.money {
			isPass = (test.input.balance == old_balance)

		} else {
			isPass = (test.input.balance == (old_balance - test.input.money))
		}

		if !isPass {
			t.Errorf("failed: %v\n", test.note)
		}
	}

}
