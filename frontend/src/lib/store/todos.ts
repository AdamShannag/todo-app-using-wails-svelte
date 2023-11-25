import { AddTodo, ClearAll, GetTodos, RemoveTodo } from '$lib/wailsjs/go/service/TodoService';
import { writable } from 'svelte/store';

export type Todo = {
	id: string;
	text: string;
	completed: boolean;
};

export const todos = writable<Todo[]>([]);

export const loadTodos = async () => {
	try {
		const c = await GetTodos();
		todos.update(() => {
			return [...c];
		});
	} catch (error) {
		console.log(error);
	}
};

export const addTodo = async (todo: string) => {
	try {
		await AddTodo(todo);
		loadTodos();
	} catch (error) {
		console.log(error);
	}
};

export const removeTodo = async (id: string) => {
	try {
		await RemoveTodo(id);
		loadTodos();
	} catch (error) {
		console.log(error);
	}
};

export const clearTodos = async () => {
	try {
		await ClearAll();
		loadTodos();
	} catch (error) {
		console.log(error);
	}
};
