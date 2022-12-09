<script lang="ts">
	// Example.svelte
	import {
		Queries,
		useQueries,
		useQuery,
		type UseQueryStoreResult
	} from '@sveltestack/svelte-query';
	import type { DeviceData } from '../lib/types/device';
	import { apiBaseUrl } from '$lib/stores/api';

	import Device from '$lib/components/Device.svelte';

	let devicesQuery: UseQueryStoreResult<
		Record<string, DeviceData>,
		unknown,
		Record<string, DeviceData>,
		'devices'
	>;

	$: devicesQuery = useQuery('devices', () =>
		fetch(`${$apiBaseUrl}/api/v1/devices/all`).then(
			(res) => res.json() as Promise<Record<string, DeviceData>>
		)
	);
</script>

{#if $devicesQuery.isLoading}
	<div class="loading">
		<p>Loading...</p>
	</div>
{:else if $devicesQuery.isError}
	<p>Error: {$devicesQuery.error.message}</p>
{:else if $devicesQuery.isSuccess}
	<div class="devices">
		{#each Object.entries($devicesQuery.data) as [id, device]}
			<Device {device} {id} />
		{/each}
	</div>
{/if}


<style lang="scss">
	.devices {
		display: grid;
		grid-template-columns: repeat(4, 1fr);
		gap: 2rem;

		padding: 2rem;

		@media (max-width: $desktop-large) {
			grid-template-columns: repeat(3, 1fr);
		}
		@media (max-width: $desktop) {
			grid-template-columns: repeat(2, 1fr);
		}
		@media (max-width: $tablet) {
			grid-template-columns: 1fr;
		}
	}
</style>
