<script lang="ts">
	import TodoItem from '$lib/components/TodoItem.svelte';

	import { todos } from '$lib/store/todos';
	import type { Todo } from '$lib/store/todos';

	let todoItems: Todo[];

	todos.subscribe((todos) => {
		todoItems = todos;
	});
</script>

<hr />
<ul>
	{#if todoItems.length > 0}
		{#each todoItems as { id, text }}
			<TodoItem {id} {text} />
		{/each}
	{:else}
		<p>Nothing Todo!</p>
	{/if}
</ul>
<hr />
