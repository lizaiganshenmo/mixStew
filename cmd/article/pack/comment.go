package pack

import (
	"github.com/lizaiganshenmo/mixStew/kitex_gen/article"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/interaction"
	"github.com/lizaiganshenmo/mixStew/kitex_gen/user"
)

func Comment(comment *interaction.Comment, userInfo *user.User, following bool) (cmt *article.Comment) {
	if comment == nil {
		return
	}
	cmt = &article.Comment{
		Id:        comment.CommentId,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		Body:      comment.CommentBody,
	}

	cmt.Author = &article.Author{
		Following: following,
	}

	if userInfo != nil {
		cmt.Author.Username = userInfo.UserRegister.Username
		cmt.Author.Bio = userInfo.Bio
		cmt.Author.Image = userInfo.Image
	}

	return

}
