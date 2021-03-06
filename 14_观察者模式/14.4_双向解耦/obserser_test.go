package obserser

func ExampleObserser() {
	//前台小姐童子喆
	tongzizhe := new(Secretary)
	//看股票的同事
	tongshi1 := NewStockObserver("魏关姹", tongzizhe)
	//看NBA的同事
	tongshi2 := NewNBAObserver("易管查", tongzizhe)

	//前台记下了两位同事
	tongzizhe.Attach(tongshi1)
	tongzizhe.Attach(tongshi2)
	//发现老板回来
	tongzizhe.Action = "老板回来了！"
	//通知两个同事
	tongzizhe.Notify()

	//老板胡汉三
	boos := new(Boss)
	//看股票的同事
	tongshi1 = NewStockObserver("魏关姹", boos)
	//看NBA的同事
	tongshi2 = NewNBAObserver("易管查", boos)

	boos.Attach(tongshi1)
	boos.Attach(tongshi2)

	boos.Detach(tongshi1)

	//老板回来了
	boos.Action = "我胡汉三回来了！"
	//发出通知
	boos.Notify()

	// OutPut:
	// 老板回来了！ 魏关姹 关闭股票行情，继续工作！
	// 老板回来了！ 易管查 关闭NBA直播，继续工作！
	// 我胡汉三回来了！ 易管查 关闭NBA直播，继续工作！
}
