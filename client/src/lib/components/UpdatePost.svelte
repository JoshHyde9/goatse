<script lang="ts">
  import type { UpdatePost, Post } from "$lib/types";

  import { createMutation, useQueryClient } from "@tanstack/svelte-query";
  import { api } from "$lib";
  import Input from "./Input.svelte";
  import TipTap from "./TipTap.svelte";

  let { post }: { post: UpdatePost } = $props();

  const client = useQueryClient();

  let postTitle = $state(post.title);
  let postAuthor = $state(post.author);
  let postContent = $state(post.content);
  let editPost = $state(false);

  const updatePost = createMutation<Post, Error, UpdatePost>({
    mutationFn: (data) => api().updatePost(data),
    onSuccess: async () => {
      editPost = false;
      await client.invalidateQueries({ queryKey: ["posts"] });
      await client.invalidateQueries({ queryKey: ["post", post.id] });

      postTitle = "";
      postAuthor = "";
    }
  });

  const handleUpdatePost = (event: SubmitEvent) => {
    event.preventDefault();

    if (!postTitle.trim()) return;
    if (!postAuthor.trim()) return;

    const newPost: UpdatePost = {
      id: post.id,
      title: postTitle,
      content: postContent,
      author: postAuthor
    };

    $updatePost.mutate(newPost);
  };
</script>

<div class="my-4">
  {#if editPost}
    <div class="space-y-1">
      <form class="mx-auto flex flex-col space-y-3" onsubmit={handleUpdatePost}>
        <label for="Title">Title: </label>
        <Input isPending={$updatePost.isPending} placeholder="Title..." bind:value={postTitle} />

        <label for="Author">Author: </label>
        <Input isPending={$updatePost.isPending} placeholder="Author..." bind:value={postAuthor} />

        <div>
          <label for="content">Content: </label>
          <TipTap content={postContent} editable={true} bind:html={postContent} />
        </div>

        <button
          type="submit"
          class="cursor-pointer rounded-md bg-gradient-to-r from-purple-600 to-blue-500 p-2 text-white transition duration-300 hover:bg-gradient-to-r hover:from-blue-500 hover:to-purple-600 disabled:cursor-not-allowed disabled:bg-gradient-to-r disabled:from-purple-300 disabled:to-blue-300 disabled:opacity-70 disabled:hover:from-purple-300 disabled:hover:to-blue-300"
          disabled={$updatePost.isPending ||
            !postTitle.trim() ||
            !postAuthor.trim() ||
            (postTitle === post.title &&
              postAuthor === post.author &&
              postContent === post.content)}
        >
          {$updatePost.isPending ? "Updating..." : "Update Post"}
        </button>
      </form>
      <button
        onclick={() => (editPost = false)}
        class="group relative mb-2 flex w-full cursor-pointer items-center justify-center overflow-hidden rounded-lg bg-gradient-to-r from-purple-600 to-blue-500 p-0.5 text-sm text-gray-900 group-hover:from-purple-600 group-hover:to-blue-500 hover:from-blue-500 hover:to-purple-600 hover:text-white focus:ring-blue-300 focus:outline-none dark:text-white dark:focus:ring-blue-800"
      >
        <span class="relative w-full rounded-md bg-white px-5 py-2.5 text-center text-black"
          >Cancel</span
        >
      </button>
    </div>
  {:else}
    <div class="mt-5">
      <button
        onclick={() => (editPost = true)}
        aria-label="Edit"
        class="cursor-pointer rounded-md bg-gradient-to-r from-purple-600 to-blue-500 p-2 text-white transition duration-300 hover:bg-gradient-to-r hover:from-blue-500 hover:to-purple-600 disabled:cursor-not-allowed disabled:bg-gradient-to-r disabled:from-purple-300 disabled:to-blue-300 disabled:opacity-70 disabled:hover:from-purple-300 disabled:hover:to-blue-300"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-6 w-6"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
          stroke-width="2"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"
          />
        </svg>
      </button>
    </div>
  {/if}
</div>
