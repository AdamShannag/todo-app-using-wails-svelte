<script lang="ts">
	import AddTodoItem from '$lib/components/AddTodoItem.svelte';
	import ClearTodos from '$lib/components/ClearTodos.svelte';
	import TodosList from '$lib/components/TodosList.svelte';
</script>

<div class="container p-5 h-full mx-auto flex justify-center items-center">
	<div class="space-y-5 mt-5 w-full">
		<h1 class="h1">Todos</h1>
		<TodosList />
		<AddTodoItem />
		<ClearTodos />
	</div>
</div>
