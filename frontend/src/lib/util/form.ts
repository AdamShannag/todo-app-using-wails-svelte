export const getFormData = (event: Event) => {
	const formElement = event.target as HTMLFormElement;
	return new FormData(formElement);
};

export const resetForm = (event: Event) => {
	const form = event.target as HTMLFormElement;
	form.reset();
};
