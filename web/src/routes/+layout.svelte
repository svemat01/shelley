<script lang="ts">
	import { QueryClient, QueryClientProvider } from '@sveltestack/svelte-query';
	import Header from '../components/header.svelte';
	import { apiBaseUrl } from '../stores/api';

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
	<Header />
	{#if $apiBaseUrl}
		<slot />
	{:else if !validUrl}
		Invalid baseurl
    {:else}
        Loading...
	{/if}
</QueryClientProvider>
