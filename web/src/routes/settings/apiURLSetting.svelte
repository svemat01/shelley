<script lang="ts">
	import { apiBaseUrl } from '$lib/stores/api.js';

	let editBaseUrl = $apiBaseUrl;
	const saveBaseUrl = () => {
		try {
			new URL(editBaseUrl);
		} catch {
			alert('Invalid URL');
			return;
		}

		apiBaseUrl.set(editBaseUrl);
	};
</script>

<div class="base-card card">
	<h2>API URL</h2>
	<div class="horizontal">
		<input type="text" bind:value={editBaseUrl} />
		<button on:click={saveBaseUrl}>Save</button>
	</div>
</div>

<style lang="scss">
	.card {
		display: flex;
		flex-direction: column;
		gap: 0.3rem;
	}

	.horizontal {
		display: flex;
		flex-direction: row;
		align-items: center;
		gap: 0.5rem;
        flex-wrap: wrap;
	}

	input {
		flex: 1;

		padding: 0.5rem;
		/* border: 2px solid $light-color; */
		border: 0;
		border-radius: 0.5rem;
		background-color: $grey-color-dark;
		color: $light-color;

		font-size: 1rem;
	}

	button {
		padding: 0.5rem 1rem;
		/* border: 2px solid $light-color; */
		border: 0;
		border-radius: 0.5rem;
		background-color: $grey-color-dark;
		color: $light-color;

		font-size: 1rem;
	}
</style>
