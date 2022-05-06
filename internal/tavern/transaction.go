package tavern

import (
	"time"

	"github.com/google/uuid"
)

// Value - object - immutable; no have ID
// Состояние таких объектов не изменяется после создания
// В реальном приложении ID для транзакции необходим, но для учебных целей этого примера достаточно.

// 1 time set, then no change;
type Transaction struct {
	// all values lowercase since they are immutable
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
