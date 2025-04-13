<script lang="ts">
  import type { Editor } from "@tiptap/core";
  import { onMount, onDestroy } from "svelte";
  import { browser } from "$app/environment";

  let {
    content,
    editable,
    html = $bindable()
  }: { content: string; editable: boolean; html: string } = $props();

  let editor: Editor | undefined = $state(undefined);
  let element: Element;

  // WHAT IS THIS OMG
  // @ts-expect-error BUT I KNOW IT'S UNDEFINED
  let isBold = $derived(editor?.isActive("bold") ?? false);
  // @ts-expect-error BUT I KNOW IT'S UNDEFINED
  let isItalic = $derived(editor?.isActive("italic") ?? false);
  // @ts-expect-error BUT I KNOW IT'S UNDEFINED
  let isH1 = $derived(editor?.isActive("heading", { level: 1 }) ?? false);
  // @ts-expect-error BUT I KNOW IT'S UNDEFINED
  let isH2 = $derived(editor?.isActive("heading", { level: 2 }) ?? false);
  // @ts-expect-error BUT I KNOW IT'S UNDEFINED
  let isBulletList = $derived(editor?.isActive("bulletList") ?? false);
  // @ts-expect-error BUT I KNOW IT'S UNDEFINED
  let isOrderedList = $derived(editor?.isActive("orderedList") ?? false);

  onMount(async () => {
    if (browser) {
      const { Editor } = await import("@tiptap/core");
      const { default: StarterKit } = await import("@tiptap/starter-kit");

      editor = new Editor({
        element: element,
        extensions: [StarterKit],
        content: content,
        editable: editable,
        onUpdate: ({ editor }) => {
          html = editor.getHTML();
        },
        onSelectionUpdate: () => {
          editor = editor;
        }
      });

      html = editor.getHTML();
    }
  });

  $effect(() => {
    if (browser && editor && editor.getHTML() !== content) {
      editor.commands.setContent(content);
    }
  });

  onDestroy(() => {
    if (editor) {
      editor.destroy();
      editor = undefined;
    }
  });

  function toggleBold() {
    editor?.chain().focus().toggleBold().run();
  }

  function toggleItalic() {
    editor?.chain().focus().toggleItalic().run();
  }

  function toggleH1() {
    editor?.chain().focus().toggleHeading({ level: 1 }).run();
  }

  function toggleH2() {
    editor?.chain().focus().toggleHeading({ level: 2 }).run();
  }

  function toggleBulletList() {
    editor?.chain().focus().toggleBulletList().run();
  }

  function toggleOrderedList() {
    editor?.chain().focus().toggleOrderedList().run();
  }
</script>

<div class="editor-wrapper">
  {#if editable}
    <div class="flex flex-wrap gap-1 border-b border-gray-300 bg-gray-50 p-2">
      <button
        class="rounded px-3 py-1 text-sm font-medium {isBold
          ? 'border-blue-300 bg-blue-100 text-blue-700'
          : 'border-gray-300 bg-white hover:bg-gray-50'} border transition-colors"
        onclick={toggleBold}
      >
        Bold
      </button>
      <button
        class="rounded px-3 py-1 text-sm font-medium {isItalic
          ? 'border-blue-300 bg-blue-100 text-blue-700'
          : 'border-gray-300 bg-white hover:bg-gray-50'} border transition-colors"
        onclick={toggleItalic}
      >
        Italic
      </button>
      <button
        class="rounded px-3 py-1 text-sm font-medium {isH1
          ? 'border-blue-300 bg-blue-100 text-blue-700'
          : 'border-gray-300 bg-white hover:bg-gray-50'} border transition-colors"
        onclick={toggleH1}
      >
        H1
      </button>
      <button
        class="rounded px-3 py-1 text-sm font-medium {isH2
          ? 'border-blue-300 bg-blue-100 text-blue-700'
          : 'border-gray-300 bg-white hover:bg-gray-50'} border transition-colors"
        onclick={toggleH2}
      >
        H2
      </button>
      <button
        class="rounded px-3 py-1 text-sm font-medium {isBulletList
          ? 'border-blue-300 bg-blue-100 text-blue-700'
          : 'border-gray-300 bg-white hover:bg-gray-50'} border transition-colors"
        onclick={toggleBulletList}
      >
        Bullet List
      </button>
      <button
        class="rounded px-3 py-1 text-sm font-medium {isOrderedList
          ? 'border-blue-300 bg-blue-100 text-blue-700'
          : 'border-gray-300 bg-white hover:bg-gray-50'} border transition-colors"
        onclick={toggleOrderedList}
      >
        Numbered List
      </button>
    </div>
  {/if}

  <div bind:this={element} class="tiptap-editor"></div>
</div>
