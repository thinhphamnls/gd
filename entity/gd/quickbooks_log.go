package gdentity

import (
	"time"
)

type QuickBooksLog struct {
	QuickBooksLogId    uint      `gorm:"column:quickbooks_log_id;primaryKey;autoIncrement"`
	QuickBooksTicketId uint      `gorm:"column:quickbooks_ticket_id;index"`
	Batch              uint      `gorm:"column:batch;not null;index"`
	EntityId           uint      `gorm:"column:entityId;index"`
	Msg                string    `gorm:"column:msg;type:text;not null"`
	LogDatetime        time.Time `gorm:"column:log_datetime;not null"`
}
