<script lang="ts">
	import shelley from '$lib/images/shelley.svg';
	import NavButton from './NavButton.svelte';
	import { faHouse } from '@fortawesome/free-solid-svg-icons';
	import CloseButton from './ToggleButton.svelte';
	import { sidebar } from '$lib/stores/sidebar';
	import { fade } from 'svelte/transition';

	const open = () => {
		sidebar.set(true);
	};
</script>

<aside class:collapsed={!$sidebar} transition:fade>
	<!-- {#if $sidebar}
		<CloseButton on:click={() => ($sidebar = false)} />
	{/if} -->

	<section class="top" on:click={() => ($sidebar = !$sidebar)}>
		<img src={shelley} alt="shelley icon" width="50px" />
		{#if $sidebar}
			<h1>Shelley</h1>
		{/if}
	</section>
	<div class="divider" />

	<section class="main">
		<NavButton url="/" title="Home" icon={faHouse} />
		<NavButton url="/test" title="TEST" />
	</section>

	<section class="footer">
		<p>Made by Jakob Helgesson</p>
	</section>
</aside>

<style lang="scss">
	aside {
		background-color: $grey-color;

		/* padding: 1rem 0.8rem; */
		padding-top: 1rem;

		grid-area: sidebar;

		display: flex;
		flex-direction: column;
		gap: 1rem;

		position: relative;

		transition: all 0.2s ease-in-out;

		& > section {
			padding-inline: 1.2rem;
		}
	}

	.collapsed {
		/* padding: 0.8rem 0.8rem; */
		padding-top: 0.8rem;

		& > section {
			padding-inline: 0.2rem;
		}

		.top {
			justify-content: center;
		}

		.footer {
			display: none;
		}
	}

	.top {
		display: flex;
		align-items: center;

		gap: 1rem;

		cursor: pointer;

		h1 {
			color: $primary-color;
			font-size: 1.8rem;
		}
	}

	.divider {
		/* margin-top: 1rem; */
		margin-bottom: 0.5rem;
		margin-inline: auto;

		width: 90%;

		border-bottom: 2px solid $light-color;
	}

	.main {
		display: flex;
		flex-direction: column;
		gap: 1rem;
		align-items: center;
	}

	.footer {
		display: flex;
		justify-content: center;
		align-items: center;

		width: 100%;
		height: 60px;

		margin-top: auto;

		background-color: $grey-color-dark;
		/* background-color: $dark-color; */
		color: $light-color;
		border-top: 2px solid $light-color;

		p {
			font-size: 1.2rem;
		}
	}
</style>
