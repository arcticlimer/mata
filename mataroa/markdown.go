package mataroa

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
)

type PostFrontmatter struct {
	Title       string    `yaml:"title"`
	Slug        string    `yaml:"slug"`
	PublishedAt time.Time `yaml:"published_at"`
}

func NewPost(filePath string) (Post, error) {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Post{}, fmt.Errorf("error reading markdown file: %s", err)
	}

	var metadata PostFrontmatter
	rest, err := frontmatter.Parse(strings.NewReader(string(f)), &metadata)
	if err != nil {
		return Post{}, fmt.Errorf("error parsing markdown file frontmatter: %s", err)
	}

	if metadata.Title == "" {
		return Post{}, fmt.Errorf("post '%s' missing 'title' attribute", filePath)
	}

	if metadata.Slug == "" {
		return Post{}, fmt.Errorf("post '%s' missing 'slug' attribute", filePath)
	}

	return Post{
		Title:       metadata.Title,
		PublishedAt: metadata.PublishedAt.Format("2006-01-02"),
		Slug:        metadata.Slug,
		Body:        string(rest),
	}, nil
}
