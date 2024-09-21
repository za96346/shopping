package entities

import "time"

type Product struct {
	ID          int        `gorm:"column:id;primaryKey" json:"ID"`                     //type:*int      comment:產品ID            version:2023-10-02
	Name        string     `gorm:"column:name" json:"Name" binding:"required"`         //type:string    comment:產品名稱          version:2023-10-02
	Title       string     `gorm:"column:title" json:"Title" binding:"required"`       //type:string    comment:產品標題          version:2023-10-02
	Description string     `gorm:"column:description" json:"Description"`              //type:string    comment:產品描述          version:2023-10-02
	Category    string     `gorm:"column:category" json:"Category"`                    //type:string    comment:產品類別          version:2023-10-02
	Price       int        `gorm:"column:price" json:"Price" binding:"required"`       //type:int       comment:價格              version:2023-10-02
	Quantity    int        `gorm:"column:quantity" json:"Quantity" binding:"required"` //type:int       comment:庫存數量          version:2023-10-02
	Sort        int        `gorm:"column:sort" json:"Sort"`                            //type:int       comment:排序              version:2023-10-02
	DisplayFlag int        `gorm:"column:display_flag" json:"DisplayFlag"`             //type:int       comment:顯示標誌          version:2023-10-02
	DeleteFlag  int        `gorm:"column:delete_flag" json:"DeleteFlag"`               //type:int       comment:刪除標誌          version:2023-10-02
	CreateTime  *time.Time `gorm:"column:create_time" json:"CreateTime"`               //type:*time.Time comment:創建時間         version:2023-10-02
	UpdateTime  *time.Time `gorm:"column:update_time" json:"UpdateTime"`               //type:*time.Time comment:更新時間         version:2023-10-02
}
