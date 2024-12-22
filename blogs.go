package main

// BlogPost represents a single blog post
type BlogPost struct {
	Title           string
	Content         string
	Slug            string
	MetaDescription string
}

// BlogPosts is a slice containing all blog posts
var BlogPosts = []BlogPost{
	{Title: "First Post", Content: "This is the first blog post.", Slug: "first-post", MetaDescription: "This is a brief summary of the first blog post."},
	{Title: "Second Post", Content: "This is the second blog post.", Slug: "second-post", MetaDescription: "This is a brief summary of the second blog post."},
}