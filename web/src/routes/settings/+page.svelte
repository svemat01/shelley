<script lang="ts">
	import ApiUrlSetting from './apiURLSetting.svelte';

</script>

<div class="container">
	<div class="title">
		<h1>Settings:</h1>
	</div>
	<ApiUrlSetting />
</div>

<style lang="scss">
	.container {
		display: grid;
		grid-template-columns: repeat(3, 1fr);
		gap: 2rem;

		padding: 2rem;
		@media (max-width: $desktop) {
			grid-template-columns: repeat(2, 1fr);
		}
		@media (max-width: $tablet) {
			grid-template-columns: 1fr;
			padding: 1rem;
			gap: 1rem;
		}
	}

	.title {
		grid-column: 1 / -1;

		padding-inline: 1rem;

		font-size: 1.5rem;
	}
</style>
