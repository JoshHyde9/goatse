<script lang="ts">
  import type { CreatePost, Post } from "$lib/types";

  import { createMutation, useQueryClient } from "@tanstack/svelte-query";
  import { api } from "$lib";

  const client = useQueryClient();

  let newPostTitle = "";
  let newPostAuthor = "";

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
    <form class="mx-auto flex w-2xl flex-col space-y-3" on:submit|preventDefault={handleCreatePost}>
      <input
        type="text"
        class="rounded-md border border-indigo-400 px-2 py-1 outline-none"
        bind:value={newPostTitle}
        placeholder="Title..."
        disabled={$createPost.isPending}
      />

      <input
        type="text"
        class="rounded-md border border-indigo-400 px-2 py-1 outline-none"
        bind:value={newPostAuthor}
        placeholder="Author..."
        disabled={$createPost.isPending}
      />

      <button
        type="submit"
        class="cursor-pointer rounded-md bg-indigo-400 p-2 text-white transition disabled:bg-indigo-100"
        disabled={$createPost.isPending || !newPostTitle.trim() || !newPostAuthor.trim()}
      >
        {$createPost.isPending ? "Adding..." : "Add Todo"}
      </button>
    </form>
  </div>
</div>
