package gdentity

import (
	"time"

	"gorm.io/datatypes"
)

type QuickBooksTicketStatus int32

const (
	QuickBooksTicketStatusFail    QuickBooksTicketStatus = 0
	QuickBooksTicketStatusSuccess QuickBooksTicketStatus = 1
)

type QuickBooksTicketSync int32

const (
	QuickBooksTicketNotSync QuickBooksTicketSync = 0
	QuickBooksTicketSynced  QuickBooksTicketSync = 1
)

type QuickBooksTicket struct {
	QuickBooksTicketId uint                   `gorm:"column:quickbooks_ticket_id;primaryKey;autoIncrement"`
	QBUsername         string                 `gorm:"column:qb_username;size:255;index"`
	Ticket             string                 `gorm:"column:ticket;size:200;index"`
	Processed          uint                   `gorm:"column:processed;default:0"`
	Ids                datatypes.JSON         `gorm:"column:ids;type:text"`
	IDsOccurs          datatypes.JSON         `gorm:"column:ids_occurs;type:text"`
	Type               string                 `gorm:"column:type;size:200"`
	LastErrorNum       string                 `gorm:"column:lasterror_num;size:32"`
	Total              int                    `gorm:"column:total"`
	LastErrorMsg       datatypes.JSON         `gorm:"column:lasterror_msg;type:text"`
	IPAddr             string                 `gorm:"column:ipaddr;size:15"`
	WriteDatetime      time.Time              `gorm:"column:write_datetime;not null;autoCreateTime"`
	TouchDatetime      time.Time              `gorm:"column:touch_datetime;not null;autoCreateTime"`
	Status             QuickBooksTicketStatus `gorm:"column:status;default:0"`
	Synced             QuickBooksTicketSync   `gorm:"column:synced;default:0"`
	SyncedBy           uint                   `gorm:"column:synced_by"`
}
