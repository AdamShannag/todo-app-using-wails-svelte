import { loadTodos } from '$lib/store/todos';
import type { PageLoad } from './$types';

export const load: PageLoad = async () => {
	await loadTodos();
};
