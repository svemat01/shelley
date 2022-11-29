import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';

export const load: PageLoad = (event) => {
	throw error(404, 'Not Found');
};
