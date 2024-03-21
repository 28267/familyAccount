package util

import (
	"fmt"
	"os"
)

type FamilyAccount struct {
	balance float64
	money   float64
	note    string
	key     int
	// index   bool
	details string
	judge   bool
}

func (fa *FamilyAccount) Screen() {
	for {
		fmt.Println("\n---------------家庭收支记账软件---------------")
		fmt.Println()
		fmt.Println("               1.收支明细                    ")
		fmt.Println("               2.登记收入                    ")
		fmt.Println("               3.登记支出                    ")
		fmt.Println("               4.退出                        ")
		fmt.Println()
		fmt.Println("               请选择1-4                     ")
		fmt.Scanln(&fa.key)

		switch fa.key {
		case 1:
			fa.Detail()
		case 2:
			fa.Income()
		case 3:
			fa.Spending()
		case 4:
			fa.Quit()
		default:
			fmt.Println("您输入有误，请重新输入")
		}
		// if !fa.index {
		// 	break
		// }

	}

}

func (fa *FamilyAccount) Detail() {
	fmt.Println("---------------当前收支明细记录---------------")
	fmt.Println(fa.details)
	if !fa.judge {
		fmt.Println("您当前没有收支记录，来一笔吧！")
	}
}

func (fa *FamilyAccount) Income() {
	fmt.Println("本次收入金额")
	fmt.Scanln(&fa.money)
	fa.balance += fa.money //修改账户余额
	fmt.Println("本次收入说明")
	fmt.Scanln(&fa.note)
	fa.details += fmt.Sprintf("\n收入\t%v     \t+ %v     \t%v", fa.balance, fa.money, fa.note)
	fmt.Println(fa.details)
	fa.judge = true
}

func (fa *FamilyAccount) Spending() {
	fmt.Println("本次支出金额")
	fmt.Scanln(&fa.money)
	if fa.money > fa.balance {
		fmt.Println("余额不足")
	} else {
		fa.balance -= fa.money //修改账户余额
		fmt.Println("本次收入说明")
		fmt.Scanln(&fa.note)
		fa.details += fmt.Sprintf("\n支出\t%v     \t- %v     \t%v", fa.balance, fa.money, fa.note)
		fmt.Println(fa.details)
		fa.judge = true
	}
}

func (fa *FamilyAccount) Quit() {
	fmt.Println("您确定要退出吗（输入yes/no)")
	for {
		flag := ""
		fmt.Scanln(&flag)
		if flag == "yes" || flag == "YES" {
			fmt.Println("您已退出家庭记账软件~")
			os.Exit(0)
			// break
		} else if flag == "no" || flag == "NO" {
			fmt.Println("您已取消退出")
			break
		} else {
			fmt.Println("请输入YES/NO")
		}
	}
}

func Account() *FamilyAccount {
	return &FamilyAccount{
		balance: 10000.0,
		money:   0.0,
		note:    "",
		key:     0,
		// index:   true,
		details: "收支\t账户金额\t收支金额\t说明",
		judge:   false,
	}
}
