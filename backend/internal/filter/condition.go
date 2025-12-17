package filter

import (
	"strings"

	"gorm.io/gorm"
)

func applyCondition(db *gorm.DB, f FieldSchema, op Operator, val string) (*gorm.DB, error) {
	col := f.DBColumn

	switch op {

	case Eq:
		return db.Where(col+" = ?", val), nil
	case Ne:
		return db.Where(col+" <> ?", val), nil
	case Contains:
		return db.Where(col+" LIKE ?", "%"+val+"%"), nil
	case StartsWith:
		return db.Where(col+" LIKE ?", val+"%"), nil
	case EndsWith:
		return db.Where(col+" LIKE ?", "%"+val), nil
	case In:
		return db.Where(col+" IN ?", strings.Split(val, ",")), nil
	case Gt:
		return db.Where(col+" > ?", val), nil
	case Gte:
		return db.Where(col+" >= ?", val), nil
	case Lt:
		return db.Where(col+" < ?", val), nil
	case Lte:
		return db.Where(col+" <= ?", val), nil
	case Between:
		p := strings.Split(val, ",")
		if len(p) == 2 {
			return db.Where(col+" BETWEEN ? AND ?", p[0], p[1]), nil
		}
	case IsNull:
		if val == "true" {
			return db.Where(col + " IS NULL"), nil
		}
		return db.Where(col + " IS NOT NULL"), nil
	}

	return db, nil
}
