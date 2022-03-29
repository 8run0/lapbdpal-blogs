package blogs

import (
	"github.com/labstack/echo/v4"
)

type BlogImpl struct {
}

// (GET /blogs)
func (b *BlogImpl) GetBlogs(ctx echo.Context) error {
	return nil
}

// (POST /blogs)
func (b *BlogImpl) PostBlogs(ctx echo.Context) error {
	return nil
}

// (DELETE /blogs/{id})
func (b *BlogImpl) DeleteBlogsId(ctx echo.Context, id string) error {
	return nil
}

// (GET /blogs/{id})
func (b *BlogImpl) GetBlogsId(ctx echo.Context, id string) error {
	return nil
}

// (GET /blogs/{username})
func (b *BlogImpl) GetBlogsUsername(ctx echo.Context, username string) error {
	return nil
}
