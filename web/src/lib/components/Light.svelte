<script lang="ts">
	import { useMutation, useQueryClient } from '@sveltestack/svelte-query';
	import type { LightState } from '../types/device';

	export let state: LightState;
	export let deviceId: string;
	export let lightId: number;

	let brightness: number = state.brightness;
	let forceOn = false;

	const queryClient = useQueryClient();

	const mutation = useMutation(
		() => {
			return fetch(`http://jab-school:8080/api/v1/devices/${deviceId}/light`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					on: forceOn || !state.ison,
					brightness,
					light: lightId.toString()
				})
			});
		},
		{
			onSuccess: () => {
				queryClient.invalidateQueries(['device', 'state', deviceId]);
			}
		}
	);
</script>

<div class="light">
	<h2>Light {lightId}</h2>
	<p>IsOn: {state.ison}</p>
	<p>Brightness: {state.brightness}</p>
	<button
		on:click={() => {
			forceOn = false;
			$mutation.mutate();
		}}>Toggle</button
	>
	<button on:click={() => {
		const inputBrightness = prompt('Enter brightness (0-100):');
		if (inputBrightness) {
			const tempBrightness = parseInt(inputBrightness);
			if (tempBrightness >= 0 && tempBrightness <= 100) {
				brightness = tempBrightness;
				forceOn = true;
				$mutation.mutate();
			} else {
				alert('Brightness must be between 0 and 100');
			}
		} else {
			alert('Invalid brightness');
			return;
		}
	}}>Brightness</button>
</div>
