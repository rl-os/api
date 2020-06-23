package entity

type UserRelation struct {
	ID       uint `json:"id" db:"id"`
	UserId   uint `json:"user_id" db:"user_id"`
	TargetId uint `json:"target_id" db:"target_id"`
	Friend   bool `json:"friend" db:"friend"`
}
