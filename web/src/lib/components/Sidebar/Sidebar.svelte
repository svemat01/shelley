<script lang="ts">
	import shelley from '$lib/images/shelley.svg';
	import NavButton from './NavButton.svelte';
	import { faGear, faHouse } from '@fortawesome/free-solid-svg-icons';
	import { fade } from 'svelte/transition';

	export let version: string;
</script>

<aside transition:fade>
	<!-- {#if $sidebar}
		<CloseButton on:click={() => ($sidebar = false)} />
	{/if} -->

	<section class="top">
		<img src={shelley} alt="shelley icon" width="50px" />
		<h1>Shelley</h1>
	</section>
	<div class="divider" />

	<section class="main">
		<NavButton url="/" title="Home" icon={faHouse} />
		<NavButton url="/settings" title="Settings" icon={faGear} />
	</section>

	<section class="footer">
		<p>Made by Jakob Helgesson</p>
		<p>Version: {version}</p>
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

		@media (max-width: $tablet) {
			/* padding: 0.8rem 0.8rem; */
			padding-top: 0.8rem;

			& > section {
				padding-inline: 0.2rem;
			}

			.top {
				justify-content: center;

				& > h1 {
					display: none;
				}
			}

			.footer {
				display: none;
			}
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
		flex-direction: column;
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
