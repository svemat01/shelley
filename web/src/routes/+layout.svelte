<script lang="ts">
	import { base } from '$app/paths';
	import { QueryClient, QueryClientProvider } from '@sveltestack/svelte-query';
	import Sidebar from '$lib/components/Sidebar/Sidebar.svelte';
	import Footer from '$lib/components/Footer/Footer.svelte';
	import { apiBaseUrl } from '../lib/stores/api';

	import '$lib/styles/base.scss';

	const queryClient = new QueryClient();

	let validUrl = false;
	$: {
		try {
			new URL($apiBaseUrl);
			validUrl = true;
		} catch {
			validUrl = false;
		}
	}
</script>

<QueryClientProvider client={queryClient}>
	<main class="container">
		<Sidebar />
		<section>
			{#if $apiBaseUrl}
				<slot />
			{:else if !validUrl}
				Invalid baseurl
			{:else}
				Loading...
			{/if}
		</section>
		<Footer />
	</main>
</QueryClientProvider>

<style lang="scss">
	.container {
		display: grid;

		grid-template-columns: 300px 1fr;
		grid-template-rows: 1fr 50px;
		grid-template-areas:
			'sidebar main'
			'sidebar footer';

		height: 100vh;
	}

	section {
		grid-area: main;
	}
</style>
