package entities

import "time"

type OrderItem struct {
	ID          int        `gorm:"column:id;primaryKey" json:"ID"` //type:*int      comment:訂單ID            version:2023-10-02
	OrderId     int        `gorm:"column:order_id" json:"OrderId"`
	ProductId   int        `gorm:"column:product_id" json:"ProductId"`
	Quantity    int        `gorm:"column:quantity" json:"Quantity"`
	SinglePrice int        `gorm:"column:single_price" json:"SinglePrice"`
	Sort        int        `gorm:"column:sort" json:"Sort"`                //type:int       comment:排序              version:2023-10-02
	DisplayFlag int        `gorm:"column:display_flag" json:"DisplayFlag"` //type:int       comment:顯示標誌          version:2023-10-02
	DeleteFlag  int        `gorm:"column:delete_flag" json:"DeleteFlag"`   //type:int       comment:刪除標誌          version:2023-10-02
	CreateTime  *time.Time `gorm:"column:create_time" json:"CreateTime"`   //type:*time.Time comment:創建時間         version:2023-10-02
	UpdateTime  *time.Time `gorm:"column:update_time" json:"UpdateTime"`   //type:*time.Time comment:更新時間         version:2023-10-02
}
