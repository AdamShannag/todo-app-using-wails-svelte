<script lang="ts">
	import { addTodo } from '$lib/store/todos';
	import { getFormData, resetForm } from '$lib/util/form';

	const addItem = async (event: Event) => {
		const formData = String(getFormData(event).get('todo'));
		if (formData) {
			await addTodo(formData);
		}
		resetForm(event);
	};
</script>

<form on:submit|preventDefault={addItem}>
	<div class="flex justify-between space-x-3">
		<input
			type="text"
			name="todo"
			class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
			placeholder="Task"
		/>

		<button
			class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800"
			type="submit">Add</button
		>
	</div>
</form>
