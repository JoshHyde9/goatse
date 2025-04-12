import { api } from "$lib";
import type { PageLoad } from "./$types";

export const load: PageLoad = async ({ parent, fetch }) => {
  const { queryClient } = await parent();

  await queryClient.prefetchQuery({
    queryKey: ["posts"],
    queryFn: () => api(fetch).getAllPosts()
  });
};
