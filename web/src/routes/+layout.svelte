<script lang="ts">
	import type { PageData } from './$types';
	import { QueryClient, QueryClientProvider } from '@sveltestack/svelte-query';
	import Sidebar from '$lib/components/Sidebar/Sidebar.svelte';

	import '$lib/styles/base.scss';

	export let data: PageData;

	const queryClient = new QueryClient();
</script>

<QueryClientProvider client={queryClient}>
	<main class="container">
		<Sidebar version={data.version} />
		<section>
			<slot />
		</section>
		<!-- <Footer /> -->
	</main>
</QueryClientProvider>

<style lang="scss">
	.container {
		display: grid;

		grid-template-columns: 300px 1fr;
		/* grid-template-rows: 1fr 50px; */
		grid-template-rows: 1fr;
		grid-template-areas: 'sidebar main';
		/* 'sidebar footer'; */

		height: 100vh;

		@media (max-width: $tablet) {
			grid-template-columns: 60px 1fr;
		}
	}

	section {
		grid-area: main;
	}
</style>
