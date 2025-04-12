<script lang="ts">
  import type { Post } from "$lib/types";

  import { createQuery } from "@tanstack/svelte-query";
  import { api } from "$lib";
  import CreatePost from "./CreatePost.svelte";

  const posts = createQuery<Post[]>({ queryKey: ["posts"], queryFn: () => api().getAllPosts() });
</script>

<div class="my-2 flex flex-col items-center space-y-2">
  <h1 class="text-2xl font-semibold">Create a Post</h1>
  <CreatePost />
</div>
<div class="container mx-auto mt-8 flex justify-center">
  {#if $posts.status === "pending"}
    <span>Loading...</span>
  {:else if $posts.status === "error"}
    <span>Error: {$posts.error.message}</span>
  {:else}
    <ul class="flex flex-wrap justify-center gap-4">
      {#each $posts.data as post (post.id)}
        <a href={`/${post.id}`}>
          <div
            class="relative flex w-96 flex-col rounded-lg border border-slate-200 bg-white shadow-sm transition hover:shadow-md"
          >
            <div class="p-4">
              <h5 class="mb-2 text-xl font-semibold break-words text-slate-800">
                {post.title}
              </h5>
              <p class="leading-normal font-light break-words text-slate-600">
                {post.author}
              </p>
            </div>
            <div class="mx-3 mt-auto border-t border-slate-200 px-1 pt-2 pb-3">
              <time
                datetime={post.created_at.toString()}
                class="ml-auto text-xs text-gray-700 dark:text-gray-400"
                >{new Date(post.created_at)
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
            </div>
          </div>
        </a>
      {/each}
    </ul>
  {/if}
</div>
