package renderfile

import (
	"context"
	"fmt"
	"strings"

	"github.com/a-h/templ"
	"github.com/muhwyndhamhp/marknotes/config"
	"github.com/muhwyndhamhp/marknotes/pkg/admin"
	"github.com/muhwyndhamhp/marknotes/pkg/models"
	"github.com/muhwyndhamhp/marknotes/pkg/post/values"
	"github.com/muhwyndhamhp/marknotes/pub"
	pub_post_detail "github.com/muhwyndhamhp/marknotes/pub/pages/post_detail/post_detail"
	pub_variables "github.com/muhwyndhamhp/marknotes/pub/variables"
	"github.com/muhwyndhamhp/marknotes/template"
	"github.com/muhwyndhamhp/marknotes/utils/fileman"
	"github.com/muhwyndhamhp/marknotes/utils/scopes"
)

func RenderPost(ctx context.Context, post *models.Post) {
	userID := uint(0)

	post.FormMeta = map[string]interface{}{
		"UserID": userID,
	}

	baseURL := strings.Split(config.Get(config.OAUTH_URL), "/callback")[0]
	canonURL := fmt.Sprintf("%s/articles/%s.html", baseURL, post.Slug)

	bodyOpts := pub_variables.BodyOpts{
		HeaderButtons: admin.AppendHeaderButtons(userID),
		Component:     nil,
		ExtraHeaders: []templ.Component{
			pub.CannonicalRel(canonURL),
		},
	}

	postDetail := pub_post_detail.PostDetail(bodyOpts, *post)

	if err := fileman.CheckDir("public/articles"); err != nil {
		fmt.Println(err)
	}

	err := template.RenderPost(postDetail, "public/articles", post.Slug, post.ID)
	if err != nil {
		fmt.Println(err)
	}
}

func RenderPosts(ctx context.Context, repo models.PostRepository) {
	err := fileman.DeletAllFiles("public/articles")
	if err != nil {
		fmt.Println(err)
	}

	posts, err := repo.Get(ctx, scopes.Where("status = ?", values.Published))
	if err != nil {
		fmt.Println(err)
	}

	for _, post := range posts {
		RenderPost(ctx, &post)
	}
}
