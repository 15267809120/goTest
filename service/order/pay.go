package order

const ORDER_TYPE_ALIPAY = 1
const ORDER_TYPE_WECHAT = 2

type PayOrder struct {
	Id int
	price float64
}

func getOrderInfoById(id int){

}