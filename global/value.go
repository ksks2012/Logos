package global

var (
	CurrentUnitID int64
)

func InitCurrentUnitID() {
	CurrentUnitID = 0
}

func IncrementCurrentUnitID() {
	CurrentUnitID++
}
