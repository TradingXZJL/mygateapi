package mygateapi

//返回格式
//状态码 200
//
//名称	类型	描述
//» currency	string	币种信息
//» available	string	可用金额
//» locked	string	冻结金额
//» update_id	integer	版本号

type PrivateRestSpotAccountsResRow struct {
	Currency  string `json:"currency"`  //string 币种信息
	Available string `json:"available"` //string 可用金额
	Locked    string `json:"locked"`    //string 冻结金额
	UpdateID  int    `json:"update_id"` //integer 版本号
}

type PrivateRestSpotAccountsRes []PrivateRestSpotAccountsResRow

//返回
//状态码	含义	描述	格式
//201	Created(opens new window)	交易成功执行。具体执行结果根据下单时的 Time in force 策略来决定	Order
//Order
//
//现货单详情
//
//#属性
//名称	类型	必选	限制	描述
//id	string	false	只读	订单 ID
//text	string	false	none	订单自定义信息，用户可以用该字段设置自定义 ID，用户自定义字段必须满足以下条件：
//
//1. 必须以 t- 开头
//2. 不计算 t- ，长度不能超过 28 字节
//3. 输入内容只能包含数字、字母、下划线(_)、中划线(-) 或者点(.)
//
//除用户自定义信息以外，以下为内部保留字段，标识订单来源:
//101 代表 android 下单
//102 代表 IOS 下单
//103 代表 IPAD 下单
//104 代表 webapp 下单
//3 代表 web 下单
//2 代表 apiv2 下单
//apiv4 代表 apiv4 下单
//amend_text	string	false	只读	用户修改订单时备注的信息
//create_time	string	false	只读	订单创建时间
//update_time	string	false	只读	订单最新修改时间
//create_time_ms	integer(int64)	false	只读	订单创建时间，毫秒精度
//update_time_ms	integer(int64)	false	只读	订单最近修改时间，毫秒精度
//status	string	false	只读	订单状态。
//
//- open: 等待处理
//- closed: 全部成交
//- cancelled: 订单撤销
//currency_pair	string	true	none	交易货币对
//type	string	false	none	订单类型
//
//- limit : 限价单
//- market : 市价单
//account	string	false	none	账户类型，spot - 现货账户，margin - 杠杆账户，cross_margin - 全仓杠杆账户，unified - 统一账户
//统一账户（旧）只能设置 cross_margin
//side	string	true	none	买单或者卖单
//amount	string	true	none	交易数量
//type为limit时，指交易货币，即需要交易的货币，如BTC_USDT中指BTC。
//type为market时，根据买卖不同指代不同
//- side : buy 指代计价货币，BTC_USDT中指USDT
//- side : sell 指代交易货币，BTC_USDT中指BTC
//price	string	false	none	交易价,type=limit时必填
//time_in_force	string	false	none	Time in force 策略。
//
//- gtc: GoodTillCancelled
//- ioc: ImmediateOrCancelled，立即成交或者取消，只吃单不挂单
//- poc: PendingOrCancelled，被动委托，只挂单不吃单
//- fok: FillOrKill，全部成交或者全部取消
//
//type=market时仅支持ioc和fok
//iceberg	string	false	none	冰山下单显示的数量，不指定或传 0 都默认为普通下单。目前不支持全部冰山。
//auto_borrow	boolean	false	只写	杠杆(包括逐仓全仓)交易时，如果账户余额不足，是否由系统自动借入不足部分
//auto_repay	boolean	false	none	全仓杠杆下单是否开启自动还款，默认关闭。需要注意的是:
//
//1. 此字段仅针对全仓杠杆有效。逐仓杠杆不支持订单级别的自动还款设置，只能通过 POST /margin/auto_repay 修改用户级别的设置
//2. auto_borrow 与 auto_repay 支持同时开启
//left	string	false	只读	交易货币未成交数量
//filled_amount	string	false	只读	交易货币已成交数量
//fill_price	string	false	只读	已成交的计价币种总额，该字段废弃，建议使用相同意义的 filled_total
//filled_total	string	false	只读	已成交总金额
//avg_deal_price	string	false	只读	平均成交价
//fee	string	false	只读	成交扣除的手续费
//fee_currency	string	false	只读	手续费计价单位
//point_fee	string	false	只读	手续费抵扣使用的点卡数量
//gt_fee	string	false	只读	手续费抵扣使用的 GT 数量
//gt_maker_fee	string	false	只读	手续费maker抵扣使用的 GT 数量
//gt_taker_fee	string	false	只读	手续费taker抵扣使用的 GT 数量
//gt_discount	boolean	false	只读	是否开启GT抵扣
//rebated_fee	string	false	只读	返还的手续费
//rebated_fee_currency	string	false	只读	返还手续费计价单位
//stp_id	integer	false	只读	订单所属的STP用户组id，同一个STP用户组内用户之间的订单不允许发生自成交。
//
//1. 如果撮合时两个订单的 stp_id 非 0 且相等，则不成交，而是根据 taker 的 stp_act 执行相应策略。
//2. 没有设置STP用户组成交的订单，stp_id 默认返回 0。
//stp_act	string	false	none	Self-Trading Prevention Action,用户可以用该字段设置自定义限制自成交策略。
//
//1. 用户在设置加入STP用户组后，可以通过传递 stp_act 来限制用户发生自成交的策略，没有传递 stp_act 默认按照 cn 的策略。
//2. 用户在没有设置加入STP用户组时，传递 stp_act 参数会报错。
//3. 用户没有使用 stp_act 发生成交的订单，stp_act 返回 -。
//
//- cn: Cancel newest,取消新订单，保留老订单
//- co: Cancel oldest,取消⽼订单，保留新订单
//- cb: Cancel both,新旧订单都取消
//finish_as	string	false	只读	订单结束方式，包括：
//
//- open: 等待处理
//- filled: 完全成交
//- cancelled: 用户撤销
//- liquidate_cancelled: 爆仓撤销
//- small: 订单数量太小
//- depth_not_enough: 深度不足导致撤单
//- trader_not_enough: 对手方不足导致撤单
//- ioc: 未立即成交，因为 tif 设置为 ioc
//- poc: 未满足挂单策略，因为 tif 设置为 poc
//- fok: 未立即完全成交，因为 tif 设置为 fok
//- stp: 订单发生自成交限制而被撤销
//- unknown: 未知
//action_mode	string	false	只写	处理模式:
//下单时根据action_mode返回不同的字段, 该字段只在请求时有效，响应结果中不包含该字段
//ACK: 异步模式，只返回订单关键字段
//RESULT: 无清算信息
//FULL: 完整模式（默认）

type PrivateRestSpotOrdersPostRes struct {
	ID                 string `json:"id"`                   //string false 只读 订单 ID
	Text               string `json:"text"`                 //string false none 订单自定义信息，用户可以用该字段设置自定义 ID 用户自定义字段必须满足以下条件： 1. 必须以 t- 开头 2. 不计算 t- ，长度不能超过 28 字节 3. 输入内容只能包含数字、字母、下划线(_)、中划线(-) 或者点(.) 除用户自定义信息以外，以下为内部保留字段，标识订单来源: 101 代表 android 下单 102 代表 IOS 下单 103 代表 IPAD 下单 104 代表 webapp 下单 3 代表 web 下单 2 代表 apiv2 下单 apiv4 代表 apiv4 下单
	AmendText          string `json:"amend_text"`           //string false 只读 用户修改订单时备注的信息
	CreateTime         string `json:"create_time"`          //string false 只读 订单创建时间
	UpdateTime         string `json:"update_time"`          //string false 只读 订单最新修改时间
	CreateTimeMs       int    `json:"create_time_ms"`       //integer(int64) false 只读 订单创建时间，毫秒精度
	UpdateTimeMs       int    `json:"update_time_ms"`       //integer(int64) false 只读 订单最近修改时间，毫秒精度
	Status             string `json:"status"`               //string false 只读 订单状态 open: 等待处理 closed: 全部成交 cancelled: 订单撤销
	CurrencyPair       string `json:"currency_pair"`        //string true none 交易货币对
	Type               string `json:"type"`                 //string false none 订单类型
	Account            string `json:"account"`              //string false none 账户类型 spot - 现货账户，margin - 杠杆账户，cross_margin - 全仓杠杆账户，unified - 统一账户 统一账户（旧）只能设置 cross_margin
	Side               string `json:"side"`                 //string true none 买单或者卖单 buy/sell
	Amount             string `json:"amount"`               //string true none 交易数量
	Price              string `json:"price"`                //string false none 交易价,type=limit时必填
	TimeInForce        string `json:"time_in_force"`        //string false none Time in force 策略 gtc-GoodTillCancelled；ioc-ImmediateOrCancelled，立即成交或者取消，只吃单不挂单；poc-PendingOrCancelled，被动委托，只挂单不吃单；fok-FillOrKill，全部成交或者全部取消 type=market时仅支持ioc和fok
	Iceberg            string `json:"iceberg"`              //string false none 冰山下单显示的数量
	AutoBorrow         bool   `json:"auto_borrow"`          //boolean false 只写 杠杆(包括逐仓全仓)交易时，如果账户余额不足，是否由系统自动借入不足部分
	AutoRepay          bool   `json:"auto_repay"`           //boolean false none 全仓杠杆下单是否开启自动还款，默认关闭 1. 此字段仅针对全仓杠杆有效。逐仓杠杆不支持订单级别的自动还款设置，只能通过 POST /margin/auto_repay 修改用户级别的设置 2. auto_borrow 与 auto_repay 支持同时开启
	Left               string `json:"left"`                 //string false 只读 交易货币未成交数量
	FilledAmount       string `json:"filled_amount"`        //string false 只读 交易货币已成交数量
	FillPrice          string `json:"fill_price"`           //string false 只读 已成交的计价币种总额，该字段废弃，建议使用相同意义的 filled_total
	FilledTotal        string `json:"filled_total"`         //string false 只读 已成交总金额
	AvgDealPrice       string `json:"avg_deal_price"`       //string false 只读 平均成交价
	Fee                string `json:"fee"`                  //string false 只读 成交扣除的手续费
	FeeCurrency        string `json:"fee_currency"`         //string false 只读 手续费计价单位
	PointFee           string `json:"point_fee"`            //string false 只读 手续费抵扣使用的点卡数量
	GtFee              string `json:"gt_fee"`               //string false 只读 手续费抵扣使用的 GT 数量
	GtMakerFee         string `json:"gt_maker_fee"`         //string false 只读 手续费maker抵扣使用的 GT 数量
	GtTakerFee         string `json:"gt_taker_fee"`         //string false 只读 手续费taker抵扣使用的 GT 数量
	GtDiscount         bool   `json:"gt_discount"`          //boolean false 只读 是否开启GT抵扣
	RebatedFee         string `json:"rebated_fee"`          //string false 只读 返还的手续费
	RebatedFeeCurrency string `json:"rebated_fee_currency"` //string false 只读 返还手续费计价单位
	StpID              int    `json:"stp_id"`               //integer false 只读 订单所属的STP用户组id 1、如果撮合时两个订单的 stp_id 非 0 且相等，则不成交，而是根据 taker 的 stp_act 执行相应策略。 2、没有设置STP用户组成交的订单，stp_id 默认返回 0
	StpAct             string `json:"stp_act"`              //string false none Self-Trading Prevention Action cn-Cancel newest 取消新订单，保留老订单；co-Cancel oldest 取消⽼订单，保留新订单；cb-Cancel both 新旧订单都取消
	FinishAs           string `json:"finish_as"`            //string false 只读 订单结束方式 open-等待处理 filled-全部成交 cancelled-订单撤销 liquidate_cancelled-爆仓撤销 small-订单数量太小 depth_not_enough-深度不足导致撤单 trader_not_enough-对手方不足导致撤单 ioc-未立即成交，因为 tif 设置为 ioc poc-未满足挂单策略，因为 tif 设置为 poc fok-未立即完全成交，因为 tif 设置为 fok stp-订单发生自成交限制而被撤销 unknown-未知
	ActionMode         string `json:"action_mode"`          //string false 只写 处理模式 ACK: 异步模式，只返回订单关键字段 RESULT: 无清算信息 FULL: 完整模式
}