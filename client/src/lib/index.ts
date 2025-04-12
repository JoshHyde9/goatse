import type { CreatePost, Post } from "./types";

export const api = (customFetch = fetch) => ({
  getAllPosts: async () => {
    const response = await customFetch("http://localhost:5000/posts");

    const data = (await response.json()) as Array<Post>;
    return data;
  },
  getPostById: async (id: string): Promise<Post> => {
    const response = await customFetch(`http://localhost:5000/posts/${id}`);
    const data = (await response.json()) as Post;
    return data;
  },
  deletePostById: async (id: string) => {
    const response = await customFetch(`http://localhost:5000/posts/${id}`, {
      method: "DELETE"
    });
    const data = (await response.json()) as Promise<Post[]>;
    return data;
  },
  createPost: async (data: CreatePost) => {
    const response = await customFetch("http://localhost:5000/posts/create", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(data)
    });
    const newPost = (await response.json()) as Promise<Post>;
    return newPost;
  }
});
