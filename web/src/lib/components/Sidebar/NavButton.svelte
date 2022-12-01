<script lang="ts">
	import Fa from 'svelte-fa';
	import type { IconDefinition } from '@fortawesome/fontawesome-common-types';
	import { page } from '$app/stores';
	import { fade } from 'svelte/transition';
	import { sidebar } from '$lib/stores/sidebar';

	export let url: string;
	export let title: string;
	export let icon: IconDefinition | undefined = undefined;
</script>

<a href={url} class:selected={url == $page.route.id} transition:fade class:collapsed={!$sidebar}>
	{#if icon}
		<Fa {icon} size="md" />
	{/if}
	{#if $sidebar || !icon}
		{title}
	{/if}
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

		&.selected {
			color: $primary-color;

			background-color: $dark-color;
			/* border: $light-color solid 3px; */
		}

        &.collapsed {
            padding: 0.2rem 0.2rem;
            aspect-ratio: 1 / 1;

            display: flex;
            align-items: center;
            justify-content: center;

            font-size: 100%;
        }
	}
</style>
