<script lang="ts">
	import { useMutation, useQueryClient } from '@sveltestack/svelte-query';
	import type { SwitchState } from '../types/device';

	export let state: SwitchState;
	export let deviceId: string;
    export let switchId: number;

    const queryClient = useQueryClient();

	const mutation = useMutation(() => {
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
	}, {
        onSuccess: () => {
            queryClient.invalidateQueries(['device', 'state', deviceId]);
        },
    });
</script>

<div class="switch">
    <h2>Switch {switchId}</h2>
    <p>IsOn: {state.ison}</p>
    <button on:click={()=>{
        $mutation.mutate();
    }}>Toggle</button>
</div>
