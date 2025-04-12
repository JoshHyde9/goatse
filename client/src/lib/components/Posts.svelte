<script lang="ts">
  import type { Post } from "$lib/types";

  import { createQuery } from "@tanstack/svelte-query";
  import { api } from "$lib";
  import CreatePost from "./CreatePost.svelte";

  const posts = createQuery<Post[]>({ queryKey: ["posts"], queryFn: () => api().getAllPosts() });
</script>

<div>
  {#if $posts.status === "pending"}
    <span>Loading...</span>
  {:else if $posts.status === "error"}
    <span>Error: {$posts.error.message}</span>
  {:else}
    <ul>
      {#each $posts.data as post (post.id)}
        <a href={post.id}>
          <p>{post.title}</p>
        </a>
      {/each}
    </ul>
  {/if}
</div>

<div>
  <CreatePost />
</div>
