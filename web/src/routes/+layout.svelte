<script lang="ts">
	import { base } from '$app/paths';
	import { QueryClient, QueryClientProvider } from '@sveltestack/svelte-query';
	import Sidebar from '$lib/components/Sidebar/Sidebar.svelte';
	import { apiBaseUrl } from '../lib/stores/api';

	import '$lib/styles/base.scss'

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
	<div class="flex">
		<Sidebar />
	{#if $apiBaseUrl}
		<slot />
	{:else if !validUrl}
		Invalid baseurl
    {:else}
        Loading...
	{/if}
	</div>
</QueryClientProvider>
