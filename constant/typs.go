package constant

type Ips struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}
type Item struct {
	UniqueID    string `json:"unique_id"`
	IP          string `json:"ip"`
	Port        string `json:"port"`
	Country     string `json:"country"`
	IPAddress   string `json:"ip_address"`
	Anonymity   int    `json:"anonymity"`
	Protocol    string `json:"protocol"`
	Isp         string `json:"isp"`
	Speed       int    `json:"speed"`
	ValidatedAt string `json:"validated_at"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
type Data struct {
	CurrentPage  int         `json:"current_page"`
	Data         []Item      `json:"data"`
	FirstPageURL string      `json:"first_page_url"`
	From         int         `json:"from"`
	LastPage     int         `json:"last_page"`
	LastPageURL  string      `json:"last_page_url"`
	NextPageURL  string      `json:"next_page_url"`
	Path         string      `json:"path"`
	PerPage      int         `json:"per_page"`
	PrevPageURL  interface{} `json:"prev_page_url"`
	To           int         `json:"to"`
	Total        int         `json:"total"`
}
