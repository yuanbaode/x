package model

// Item struct is a row record of the item table in the  database
type Item struct {
	ID                      int64   `gorm:"primary_key;column:id;type:bigint;" json:"id"`
	Type                    int8    `gorm:"column:type;type:tinyint;default:0;" json:"type"`
	Name                    string  `gorm:"column:name;type:varchar;" json:"name"`
	ItemNo                  string  `gorm:"column:item_no;type:varchar;" json:"item_no"`
	ShopID                  int64   `gorm:"column:shop_id;type:bigint;" json:"shop_id"`
	ProductID               int64   `gorm:"column:product_id;type:bigint;" json:"product_id"`
	BrandID                 int64   `gorm:"column:brand_id;type:bigint;" json:"brand_id"`
	CategoryID              int64   `gorm:"column:category_id;type:bigint;default:0;" json:"category_id"`
	UnitID                  int64   `gorm:"column:unit_id;type:bigint;" json:"unit_id"`
	SaleKey                 string  `gorm:"column:sale_key;type:varchar;" json:"sale_key"`
	Description             string  `gorm:"column:description;type:varchar;" json:"description"`
	Image                   string  `gorm:"column:image;type:varchar;" json:"image"`
	DescriptionImages       string  `gorm:"column:description_images;type:varchar;default:[];" json:"description_images"`
	Video                   string  `gorm:"column:video;type:varchar;default:{};" json:"video"`
	PurchaseLimitation      int64   `gorm:"column:purchase_limitation;type:int;default:0;" json:"purchase_limitation"`
	PurchaseLimitationTime  int64   `gorm:"column:purchase_limitation_time;type:bigint;default:0;" json:"purchase_limitation_time"`
	PurchaseLimitationCycle int8    `gorm:"column:purchase_limitation_cycle;type:tinyint;" json:"purchase_limitation_cycle"`
	MinUnitPrice            float64 `gorm:"column:min_unit_price;type:decimal;default:0.000000;" json:"min_unit_price"`
	MaxUnitPrice            float64 `gorm:"column:max_unit_price;type:decimal;default:0.000000;" json:"max_unit_price"`
	MinLinePrice            float64 `gorm:"column:min_line_price;type:decimal;default:0.000000;" json:"min_line_price"`
	MaxLinePrice            float64 `gorm:"column:max_line_price;type:decimal;default:0.000000;" json:"max_line_price"`
	IsAlone                 int8    `gorm:"column:is_alone;type:tinyint;default:0;" json:"is_alone"`
	ShelfStatus             int8    `gorm:"column:shelf_status;type:tinyint;default:0;" json:"shelf_status"`
	ShelfTime               int64   `gorm:"column:shelf_time;type:bigint;default:0;" json:"shelf_time"`
	OffShelfRemark          string  `gorm:"column:off_shelf_remark;type:varchar;" json:"off_shelf_remark"`
	OffShelfTime            int64   `gorm:"column:off_shelf_time;type:bigint;default:0;" json:"off_shelf_time"`
	EditStatus              int8    `gorm:"column:edit_status;type:tinyint;default:0;" json:"edit_status"`
	AuditStatus             int8    `gorm:"column:audit_status;type:tinyint;default:0;" json:"audit_status"`
	SubmitAuditTime         int64   `gorm:"column:submit_audit_time;type:bigint;default:0;" json:"submit_audit_time"`
	AuditRemark             string  `gorm:"column:audit_remark;type:varchar;" json:"audit_remark"`
	AuditTime               int64   `gorm:"column:audit_time;type:bigint;default:0;" json:"audit_time"`
	CreatedAt               int64   `gorm:"column:created_at;type:bigint;default:0;" json:"created_at"`
	UpdatedAt               int64   `gorm:"column:updated_at;type:bigint;default:0;" json:"updated_at"`
	CreatedUID              int64   `gorm:"column:created_uid;type:bigint;default:0;" json:"created_uid"`
	UpdatedUID              int64   `gorm:"column:updated_uid;type:bigint;default:0;" json:"updated_uid"`
	IsDel                   int8    `gorm:"column:is_del;type:tinyint;default:0;" json:"is_del"`
	ImageShortLink          string  `gorm:"column:image_short_link;type:varchar;" json:"image_short_link"`
	ShipmentType            int8    `gorm:"column:shipment_type;type:tinyint;default:0;" json:"shipment_type"`
}

// TableName sets the insert table name for this struct type
func (i Item) TableName() string {
	return "item"
}
