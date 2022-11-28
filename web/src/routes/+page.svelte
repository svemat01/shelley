<script lang="ts">
	// Example.svelte
	import { Queries, useQueries, useQuery, type UseQueryStoreResult } from '@sveltestack/svelte-query';
	import type { DeviceData } from '../types/device';
	import { apiBaseUrl } from '$lib/stores/api';

	import Device from '$lib/components/Device.svelte';

	let devicesQuery: UseQueryStoreResult<Record<string, DeviceData>, unknown, Record<string, DeviceData>, "devices">

	$: devicesQuery = useQuery('devices', () =>
		fetch(`${$apiBaseUrl}/api/v1/devices/all`).then(
			(res) => res.json() as Promise<Record<string, DeviceData>>
		)
	);
</script>

{#if $devicesQuery.isLoading}
	<p>Loading...</p>
{:else if $devicesQuery.isError}
	<p>Error: {$devicesQuery.error.message}</p>
{:else if $devicesQuery.isSuccess}
	{#each Object.entries($devicesQuery.data) as [id, device]}
		<Device {device} {id} />
	{/each}
{/if}

