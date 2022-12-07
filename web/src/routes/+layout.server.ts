import type { LayoutServerLoad } from './$types';

// import package.json
import pkg from '../../package.json';

export const load: LayoutServerLoad = () => {
	return {
		version: pkg.version
	};
};
