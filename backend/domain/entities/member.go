package entities

import "time"

type Member struct {
	ID          int        `gorm:"column:id;primaryKey" json:"ID"`             //type:*int      comment:會員ID          version:2023-10-02
	Name        string     `gorm:"column:name" json:"Name" binding:"required"` //type:string    comment:會員名稱        version:2023-10-02
	Account     string     `gorm:"column:account" json:"Account"`
	Password    string     `gorm:"column:password" json:"Password"`
	PhoneNumber int        `gorm:"column:phone_number" json:"PhoneNumber" binding:"required"` //type:int       comment:電話號碼        version:2023-10-02
	Address     string     `gorm:"column:address" json:"Address" binding:"required"`          //type:string    comment:地址            version:2023-10-02
	Sort        int        `gorm:"column:sort" json:"Sort"`                                   //type:int       comment:排序            version:2023-10-02
	DisplayFlag int        `gorm:"column:display_flag" json:"DisplayFlag"`                    //type:int       comment:顯示標誌        version:2023-10-02
	DeleteFlag  int        `gorm:"column:delete_flag" json:"DeleteFlag"`                      //type:int       comment:刪除標誌        version:2023-10-02
	CreateTime  *time.Time `gorm:"column:create_time" json:"CreateTime"`                      //type:*time.Time comment:創建時間       version:2023-10-02
	UpdateTime  *time.Time `gorm:"column:update_time" json:"UpdateTime"`                      //type:*time.Time comment:更新時間       version:2023-10-02
}
