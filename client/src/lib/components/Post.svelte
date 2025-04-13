<script lang="ts">
  import type { Post } from "$lib/types";

  import { createMutation, createQuery, useQueryClient } from "@tanstack/svelte-query";
  import { api } from "$lib";
  import { goto } from "$app/navigation";
  import UpdatePost from "./UpdatePost.svelte";
  import TipTap from "./TipTap.svelte";

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

<article class="mx-auto my-2 max-w-prose">
  {#if !postId || $post.isLoading}
    <span>Loading...</span>
  {/if}
  {#if $post.isError}
    <span>Error: {$post.error.message}</span>
  {/if}

  {#if $post.isSuccess}
    <header class="my-12">
      <a href="/" class="mb-12 flex gap-2 text-zinc-700 transition-all duration-300 hover:gap-4">
        <span>←</span>
        <span>Back to home</span>
      </a>

      <h1
        class="text-4xl leading-[53px] tracking-tighter break-words md:my-6 md:text-6xl md:leading-[73px]"
      >
        {$post.data.title}
      </h1>

      <time
        datetime={$post.data.created_at.toString()}
        class="block text-base leading-6 text-balance text-zinc-700"
        >{new Date($post.data.created_at)
          .toLocaleDateString("en-AU", {
            hour: "2-digit",
            minute: "2-digit",
            day: "2-digit",
            month: "2-digit",
            year: "numeric",
            hour12: false
          })
          .replace(",", "-")}</time
      >

      <p class="text-xl font-medium break-words italic">
        - {$post.data.author}
      </p>
    </header>
    <div
      class="prose-img:breakout prose-base prose-zinc lg:prose-xl prose-img:mb-24 prose-img:rounded-md prose-img:shadow-lg max-w-none"
    >
      <TipTap content={$post.data.content} editable={false} html={$post.data.content} />
      <button
        class="w-full cursor-pointer rounded-md bg-gradient-to-r from-purple-600 to-blue-500 py-2 text-base text-white transition duration-300 hover:bg-gradient-to-r hover:from-blue-500 hover:to-purple-600 disabled:cursor-not-allowed disabled:bg-gradient-to-r disabled:from-purple-300 disabled:to-blue-300 disabled:opacity-70 disabled:hover:from-purple-300 disabled:hover:to-blue-300"
        on:click={handleDeletePost}>Delete</button
      >
    </div>

    <UpdatePost post={$post.data} />
  {/if}
</article>
