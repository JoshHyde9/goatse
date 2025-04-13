<script lang="ts">
  import type { CreatePost, Post } from "$lib/types";

  import { createMutation, useQueryClient } from "@tanstack/svelte-query";
  import { api } from "$lib";
  import Input from "./Input.svelte";

  const client = useQueryClient();

  let newPostTitle = $state("");
  let newPostAuthor = $state("");

  const createPost = createMutation<Post, Error, CreatePost>({
    mutationFn: (data) => api().createPost(data),
    onSuccess: async () => {
      await client.invalidateQueries({ queryKey: ["posts"] });

      newPostTitle = "";
      newPostAuthor = "";
    }
  });

  const handleCreatePost = () => {
    if (!newPostTitle.trim()) return;
    if (!newPostAuthor.trim()) return;

    const newPost: CreatePost = {
      title: newPostTitle,
      author: newPostAuthor
    };

    $createPost.mutate(newPost);
  };
</script>

<div>
  <div class="flex flex-col space-y-1 px-4">
    <form class="mx-auto flex w-2xl flex-col space-y-3" onsubmit={handleCreatePost}>
      <Input isPending={$createPost.isPending} placeholder="Title..." bind:value={newPostTitle} />

      <Input isPending={$createPost.isPending} placeholder="Author..." bind:value={newPostAuthor} />

      <button
        type="submit"
        class="cursor-pointer rounded-md bg-gradient-to-r from-purple-600 to-blue-500 p-2 text-white transition duration-300 hover:bg-gradient-to-r hover:from-blue-500 hover:to-purple-600 disabled:cursor-not-allowed disabled:bg-gradient-to-r disabled:from-purple-300 disabled:to-blue-300 disabled:opacity-70 disabled:hover:from-purple-300 disabled:hover:to-blue-300"
        disabled={$createPost.isPending || !newPostTitle.trim() || !newPostAuthor.trim()}
      >
        {$createPost.isPending ? "Adding..." : "Create Post"}
      </button>
    </form>
  </div>
</div>
