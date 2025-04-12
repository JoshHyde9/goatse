export type Post = {
  id: string;
  title: string;
  author: string;
  created_at: Date;
};

export type CreatePost = Pick<Post, "title" | "author">;
