package pack

import (
	"github.com/lizaiganshenmo/mixStew/cmd/interaction/dal/db"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/interaction"
)

func Comments(comments []*db.CommentInfo) (commentList []*interaction.Comment) {
	commentList = make([]*interaction.Comment, len(comments))
	for i, v := range comments {
		commentList[i] = &interaction.Comment{
			CommentId:        v.CommentId,
			Uid:              v.Uid,
			TargertArticleId: v.TargetArticleId,
			TargertUid:       v.TargetUid,
			CommentBody:      v.Body,
			CreatedAt:        v.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:        v.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
	}

	return

}
