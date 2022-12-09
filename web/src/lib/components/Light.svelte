<script lang="ts">
	import { faLightbulb } from '@fortawesome/free-solid-svg-icons';
	import RangeSlider from 'svelte-range-slider-pips';
	import { useMutation, useQueryClient } from '@sveltestack/svelte-query';
	import Fa from 'svelte-fa';

	export let deviceId: string;
	export let lightId: number;
	export let brightness: number[];

	// let brightness: number[] = [state.brightness];

	// $: {
	// 	brightness = [state.brightness];
	// 	// console.log(brightness);
	// }

	// $: console.log({brightness})

	const queryClient = useQueryClient();

	const mutation = useMutation(
		() => {
			return fetch(`http://jab-school:8080/api/v1/devices/${deviceId}/light`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					on: brightness[0] > 0,
					brightness: brightness[0],
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

	let timer: ReturnType<typeof setTimeout>;

	const debounce = () => {
		clearTimeout(timer);
		timer = setTimeout(() => {
			$mutation.mutate();
			console.log({
				on: brightness[0] > 0,
				brightness: brightness[0],
				light: lightId.toString()
			});
		}, 500);
	};
</script>

<!-- <div class="base-card card">
	<h2>Light {lightId}</h2>
	<p>IsOn: {state.ison}</p>
	<p>Brightness: {state.brightness}</p>
	<button
		on:click={() => {
			forceOn = false;
			$mutation.mutate();
		}}>Toggle</button
	>
</div> -->

<div class="base-card card">
	<div class="top">
		<div class="icon">
			<Fa
				icon={faLightbulb}
				color={'#FA9F42'}
				style={`filter: brightness(${brightness[0] > 0 ? brightness[0] * 0.8 + 20 : 0}%)`}
				size="2x"
			/>
		</div>
		<h2>Light {lightId}</h2>
	</div>
	<div class="test">
		<RangeSlider min={0} max={100} bind:values={brightness} on:change={debounce} />
	</div>
</div>

<style lang="scss">
	.card {
		display: flex;
		flex-direction: column;
		flex-wrap: wrap;
		align-items: center;
		gap: 1rem;
	}

	.top {
		width: 100%;
		display: flex;
		align-items: center;
		gap: 1rem;
		flex: 1;
	}

	.test {
		width: 100%;
		flex: 1;
	}
</style>
