<script lang="ts">
	import { faBoltLightning } from '@fortawesome/free-solid-svg-icons';
	import { useMutation, useQueryClient } from '@sveltestack/svelte-query';
	import Fa from 'svelte-fa';
	import type { SwitchState } from '../types/device';

	export let state: SwitchState;
	export let deviceId: string;
	export let switchId: number;

	const queryClient = useQueryClient();

	const mutation = useMutation(
		() => {
			return fetch(`http://jab-school:8080/api/v1/devices/${deviceId}/switch`, {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					on: !state.ison,
					switch: switchId.toString()
				})
			});
		},
		{
			onSuccess: () => {
				queryClient.invalidateQueries(['device', 'state', deviceId]);
			}
		}
	);

	const toggle = () => {
		$mutation.mutate();
	};
</script>

<div class="base-card card" on:click={toggle} on:keyup={toggle}>
	<div class="icon">
		<Fa icon={faBoltLightning} color={state.ison ? '#FA9F42' : '#1b2347'} size="2x" />
	</div>
	<h2>Switch {switchId}</h2>
</div>

<style lang="scss">
	.card {
		display: flex;
		align-items: center;
		gap: 1rem;
	}
</style>
