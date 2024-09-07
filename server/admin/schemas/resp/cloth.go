package resp

import "Awesome/core"

type ClothStyleResp struct {
	ID          string        `json:"id" structs:"id"`
	Name        string        `json:"name" structs:"name"`
	Code        string        `json:"code" structs:"code"`
	Picture     string        `json:"picture" structs:"picture"`
	Colors      []string      `json:"colors" structs:"colors"`
	Sizes       []string      `json:"sizes" structs:"sizes"`
	Year        string        `json:"year" structs:"year"`
	Season      string        `json:"season" structs:"season"`
	UnitPrice   float64       `json:"unit_price" structs:"unit_price"`
	Procedures  []Procedure   `json:"procedures" structs:"procedures"`
	Consumption []Consumption `json:"consumption" structs:"consumption"`
	Crafts      []Craft       `json:"crafts" structs:"crafts"`
	Status      uint8         `json:"status" structs:"status"`
	Remark      string        `json:"remark" structs:"remark"`
	CreateTime  core.TsTime   `json:"create_time" structs:"create_time"`
	UpdateTime  core.TsTime   `json:"update_time" structs:"update_time"`
}

type Procedure struct {
	Name             string  `json:"name" structs:"name"`
	Sort             uint    `json:"sort" structs:"sort"`
	Price            float64 `json:"price" structs:"price"`
	Indiscriminately bool    `json:"indiscriminately" structs:"indiscriminately"`
	IsCompleted      bool    `json:"is_completed" structs:"is_completed"`
}

type Consumption struct {
	Name       string  `json:"name" structs:"name"`
	Part       string  `json:"part" structs:"part"`
	Nums       uint    `json:"nums" structs:"nums"`
	Unit       string  `json:"unit" structs:"unit"`
	UnitPrice  float64 `json:"unit_price" structs:"unit_price"`
	Size       string  `json:"size" structs:"size"`
	SizeAllow  bool    `json:"size_allow" structs:"size_allow"`
	Color      string  `json:"color" structs:"color"`
	ColorAllow bool    `json:"color_allow" structs:"color_allow"`
}

type Craft struct {
	Name        string  `json:"name,omitempty" structs:"name"`
	Sort        uint    `json:"sort,omitempty" structs:"sort"`
	Description string  `json:"description,omitempty" structs:"description"`
	Children    []Craft `json:"children" structs:"children"`
}

type ClothOrderResp struct {
	ID           string         `json:"id" structs:"id"`
	Name         string         `json:"name" structs:"name"`
	Code         string         `json:"code" structs:"code"`
	CustomerId   uint8          `json:"customer_id" structs:"customer_id"`
	DeliveryDate int64          `json:"delivery_date"	structs:"delivery_date"`
	OrderType    string         `json:"order_type"	 structs:"order_type"`
	Salesman     string         `json:"salesman"      structs:"salesman"`
	Cloth        *OrderStyle    `json:"cloth" structs:"cloth"`
	Total        uint           `json:"total" structs:"total"`
	TotalPrice   float64        `json:"total_price" structs:"total_price"`
	Contains     []OrderContain `json:"contains" structs:"contains"`
	Procedure    []Procedure    `json:"procedure" structs:"procedure"`
	Consumption  []Consumption  `json:"consumption" structs:"consumption"`
	Crafts       []Craft        `json:"crafts" structs:"crafts"`
	Status       uint8          `json:"status" structs:"status"`
	// 进度
	Progress   ClothOrderProgress `json:"progress" structs:"progress"`
	Remark     string             `json:"remark"	 structs:"remark"`
	CreateTime core.TsTime        `json:"create_time" structs:"create_time"`
	UpdateTime core.TsTime        `json:"update_time" structs:"update_time"`
}

type ClothOrderProgress struct {
	Total   int64   `json:"total" structs:"total"`
	Done    int64   `json:"done" structs:"done"`
	Percent float64 `json:"percent" structs:"percent"`
	Cuted   int64   `json:"cuted" structs:"cuted"`
}

type OrderStyle struct {
	Name        string        `json:"name" structs:"name"`
	Code        string        `json:"code" structs:"code"`
	Picture     string        `json:"picture" structs:"picture"`
	Colors      []string      `json:"colors" structs:"colors"`
	Sizes       []string      `json:"sizes" structs:"sizes"`
	Year        string        `json:"year" structs:"year"`
	Season      string        `json:"season" structs:"season"`
	UnitPrice   float64       `json:"unitPrice" structs:"unit_price"`
	Procedures  []Procedure   `json:"procedures" structs:"procedures"`
	Crafts      []Craft       `json:"crafts" structs:"crafts"`
	Consumption []Consumption `json:"consumption" structs:"consumption"`
	Status      uint8         `json:"status" structs:"status"`
	Remark      string        `json:"remark" structs:"remark"`
}

type OrderContain struct {
	Size  string `json:"size" structs:"size"`
	Color string `json:"color" structs:"color"`
	Nums  uint   `json:"nums" structs:"nums"`
}

type ClothTailor struct {
	OrderNo string `json:"order_no"`
	OrderId uint   `json:"order_id"`
	Order   struct {
		Name    string `json:"name"`
		Code    string `json:"code"`
		Picture string `json:"picture"`
	} `json:"order,omitempty"`
	ClothNo  string `json:"cloth_no"`
	Customer string `json:"customer"`
	Info     struct {
	} `json:"info"`
	Piece      []TailorPiece `json:"piece,omitempty"`
	CreateTime core.TsTime   `json:"create_time" structs:"create_time"`
	UpdateTime core.TsTime   `json:"update_time" structs:"update_time"`
}

type TailorPiece struct {
	Size  string  `json:"size"`
	Color string  `json:"color"`
	Nums  uint    `json:"nums"`
	Price float64 `json:"price"`
}
