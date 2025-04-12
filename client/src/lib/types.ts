export type Post = {
  id: string;
  title: string;
  author: string;
  created_at: string;
};

export type CreatePost = Pick<Post, "title" | "author">;
