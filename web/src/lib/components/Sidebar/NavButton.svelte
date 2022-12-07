<script lang="ts">
	import Fa from 'svelte-fa';
	import type { IconDefinition } from '@fortawesome/fontawesome-common-types';
	import { page } from '$app/stores';
	import { fade } from 'svelte/transition';

	export let url: string;
	export let title: string;
	export let icon: IconDefinition;
</script>

<a href={url} class:selected={url == $page.route.id} transition:fade>
	<div class="icon">
		<Fa {icon} size="sm" />
	</div>
	<div class="text">
		{title}
	</div>
</a>

<style lang="scss">
	a {
		color: $light-color;
		font-size: 1.4rem;
		font-weight: 600;
		text-align: center;
		vertical-align: middle;

		width: 100%;
		border-radius: 0.5rem;
		padding: 0.8rem 0.8rem;

		transition: all 0.2s ease-in-out;

		display: flex;
		align-items: center;
		justify-content: center;
		gap: 0.8rem;

		&.selected {
			color: $primary-color;

			background-color: $dark-color;
			/* border: $light-color solid 3px; */
		}

		@media (max-width: $tablet) {
			padding: 0.2rem 0.2rem;
			aspect-ratio: 1 / 1;

			display: flex;
			align-items: center;
			justify-content: center;

			.text {
				display: none;

				font-size: 1.3rem;
			}

			.icon {
				font-size: 1.8rem;
			}
		}
	}
</style>
