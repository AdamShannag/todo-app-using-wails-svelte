<script lang="ts">
	import { removeTodo } from '$lib/store/todos';
	import { getFormData } from '$lib/util/form';
	import DeleteButton from './DeleteButton.svelte';

	export let id: string;
	export let text: string;

	const deleteItem = async (event: Event) => await removeTodo(String(getFormData(event).get('id')));
</script>

<li
	class="flex px-3 py-2 mb-2 justify-between bg-gray-800 border border-gray-700 rounded-lg shadow"
>
	<p>{text}</p>
	<form on:submit|preventDefault={deleteItem}>
		<input type="hidden" name="id" value={id} />
		<DeleteButton/>
	</form>
</li>
