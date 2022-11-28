<script lang="ts">
	import { useQuery } from '@sveltestack/svelte-query';
	import type { DeviceData, DeviceState } from '../types/device';
	import Light from './Light.svelte';
	import Switch from './Switch.svelte';

	export let device: DeviceData;
	export let id: string;

	const deviceQuery = useQuery(['device', 'state', id], () =>
		fetch(`http://jab-school:8080/api/v1/devices/${id}/state`).then(
			(res) => res.json() as Promise<DeviceState>
		)
	);
</script>

<div>
	<h1>{device.name} ({id})</h1>
	<p>{device.ip}</p>
	<p>{device.type}</p>

	{#if $deviceQuery.data}
		<div class="state">
			<p>Online: {$deviceQuery.data.online}</p>

			{#each $deviceQuery.data.switches as state, i}
				<Switch {state} deviceId={id} switchId={i} />
			{/each}

            {#each $deviceQuery.data.lights as state, i}
                <Light {state} deviceId={id} lightId={i} />
            {/each}
		</div>
	{/if}
</div>
