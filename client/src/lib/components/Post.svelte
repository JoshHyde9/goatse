<script lang="ts">
  import type { Post } from "$lib/types";

  import { createMutation, createQuery, useQueryClient } from "@tanstack/svelte-query";
  import { api } from "$lib";
  import { goto } from "$app/navigation";

  export let postId: string;

  const client = useQueryClient();

  const post = createQuery<Post>({
    queryKey: ["post", postId],
    queryFn: () => api().getPostById(postId)
  });

  const deletePost = createMutation<Post[]>({
    mutationFn: () => api().deletePostById(postId),
    onSuccess: async () => {
      await client.invalidateQueries({ queryKey: ["posts"] });
    }
  });

  const handleDeletePost = () => {
    $deletePost.mutate();
    goto("/");
  };
</script>

<div>
  {#if !postId || $post.isLoading}
    <span>Loading...</span>
  {/if}
  {#if $post.isError}
    <span>Error: {$post.error.message}</span>
  {/if}

  {#if $post.isSuccess}
    <a href="/" class="mb-10 block">Back to deez</a>
    <div class="flex flex-col space-y-1 px-4">
      <h1>{$post.data.title}</h1>
      <p>{$post.data.author}</p>
      <time datetime="">{$post.data.created_at.toString()}</time>
      <button
        class="cursor-pointer bg-indigo-500 py-2 text-white transition hover:bg-indigo-400"
        on:click={handleDeletePost}>Delete</button
      >
    </div>
  {/if}
</div>
