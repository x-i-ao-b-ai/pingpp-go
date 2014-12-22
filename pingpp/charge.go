package pingpp

type ChargeParams struct {
	order_no  string
	appid     string
	channel   string
	amount    uint64
	currency  string
	client_ip string
	subject   string
	body      string
	metadata  map[string]string
	Extra     Extra `json:"extra"`
}

type ChargeListParams struct {
	Limit       uint64
	Start_after string
	End_before  string
	Createdgt   string "created[gt]"
	Createdgte  string "created[gte]"
	Createdlt   string "created[lt]"
	Createdlte  string "created[lte]"
	Appid       string "app[id]"
	Channel     string
	Paid        uint64
	Refunded    uint64
}

type Charge struct {
	Id              string `json:"id"`
	Object          string `json:"object"`
	Created         uint64 `json:"created"`
	Livemode        bool   `json:"livemode"`
	Paid            bool   `json:"paid"`
	Refunded        bool   `json:"refunded"`
	Order_no        string `json:"order_no"`
	App             string `json:"app"`
	Channel         string `json:"channel"`
	Amount          int    `json:"amount"`
	Amount_settle   uint64 `json:"amount_settle"`
	Amount_refunded uint64 `json:"amount_refunded"`
	Time_expire     uint64 `json:"time_expire"`
	Time_settle     uint64 `json:"time_settle"`
	Transaction_no  string `json:"transaction_no"`
	Currency        string `json:"currency"`
	Client_ip       string `json:"client_ip"`
	Subject         string `json:"subject"`
	Body            string `json:"body"`
	Failure_code    int    `json:"failure_code"`
	Failure_msg     string `json:"failure_msg"`
	Extra_data      Extra  `json:"extra"`
	//Metadata        map[string]string `json:"metadata"`
	Refunds    RefundList  `json:"refunds"`
	Credential interface{} `json:"credential"`
}

type ChargeList struct {
	Object   string   `json:"object"`
	Url      string   `json:"url"`
	Has_more bool     `json:"has_more"`
	charges  []Charge `json:"charges"`
}

type Extra struct {
	Result_url  string `json:"result_url"`
	Success_url string `json:"success_url"`
	Cancel_url  string `json:"cancel_url"`
	Trade_type  string `json:"trade_type"`
	Open_id     string `json:"openid"`
	Bfb_login   bool   `json:"bfb_login"`
}

// func (charge *Charge) UnmarshalJSON(data []byte) error {
// 	var temp_charge Charge
// 	err := json.Unmarshal(data, &temp_charge)
// 	if err == nil {
// 		*charge = Charge(temp_charge)
// 	}
// 	return err
// }
