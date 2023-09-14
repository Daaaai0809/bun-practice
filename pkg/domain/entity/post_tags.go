package entity

type PostTags struct {
	PostId	int64 `bun:"post_id,pk" json:"post_id"`
	Post	*Post `bun:"rel:belongs-to,join:post_id=id" json:"post"`
	TagId	int64 `bun:"tag_id,pk" json:"tag_id"`
	Tag	*Tag `bun:"rel:belongs-to,join:tag_id=id" json:"tag"`
}