<script lang="ts">
	import { useQuery } from '@sveltestack/svelte-query';
	import type { DeviceData, DeviceState } from '$lib/types/device';
	import Light from './Light.svelte';
	import Switch from './Switch.svelte';
	import Tooltip from './Tooltip.svelte';
	import { faGlobe } from '@fortawesome/free-solid-svg-icons';
	import Fa from 'svelte-fa';

	export let device: DeviceData;
	export let id: string;

	const deviceQuery = useQuery(['device', 'state', id], () =>
		fetch(`http://jab-school:8080/api/v1/devices/${id}/state`).then(
			(res) => res.json() as Promise<DeviceState>
		), {staleTime: 2000}
	);

	$: console.log($deviceQuery);
</script>

<div class="device">
	<div class="info">
		<Tooltip>
			<h2>{device.name}</h2>
			<p slot="tooltip">{device.type}<br />{device.ip}</p>
		</Tooltip>
		<Fa icon={faGlobe} size='lg' color={$deviceQuery.data?.online ? '#37FF8B' : '#FF1654'} />
	</div>
	<div class="divider" />
	<!-- <h1>{device.name}</h1>
	<p>{device.ip}</p>
	<p>{device.type}</p> -->

	{#if $deviceQuery.data}
		<div class="state">
			{#each $deviceQuery.data.switches as state, i}
				<Switch {state} deviceId={id} switchId={i} />
			{/each}

			{#each $deviceQuery.data.lights as state, i}
				<Light deviceId={id} lightId={i} brightness={[ state.ison ? state.brightness : 0]} />
			{/each}
		</div>
	{/if}
</div>

<style lang="scss">
	.device {
		display: flex;
		flex-direction: column;
	}

	.info {
		padding-inline: 1rem;
		vertical-align: middle;

		display: flex;
		flex-direction: row;
		justify-content: space-between;
		align-items: center;
	}

	.divider {
		width: 100%;
		height: 2px;
		margin-top: 0.2rem;
		margin-bottom: 0.5rem;

		background-color: $light-color;
	}

	.state {
		/* padding: 1rem 2rem; */
	}
</style>
