export type Post = {
  id: string;
  title: string;
  author: string;
  content: string;
  created_at: string;
};

export type CreatePost = Pick<Post, "title" | "author" | "content">;

export type UpdatePost = CreatePost & Pick<Post, "id">;
