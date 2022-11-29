<script lang="ts">
	import shelley from '$lib/images/shelley.svg';
	import NavButton from './NavButton.svelte';
	import { faHouse } from '@fortawesome/free-solid-svg-icons';
</script>

<aside>
	<section class="top">
		<img src={shelley} alt="shelley icon" width="50px" />
		<h1>Shelley</h1>
	</section>
	<div class="divider" />

	<section class="main">
		<nav>
			<NavButton url="/" title="Home" icon={faHouse} />
			<NavButton url="/test" title="TEST" />
		</nav>
	</section>

	<!-- <section class="footer">
		<p>Made by Jakob Helgesson</p>
	</section> -->
</aside>

<style lang="scss">
	aside {
		background-color: $grey-color;

		border-right: 1px solid $grey-color-dark;

		padding: 1rem 0.8rem;

		grid-area: sidebar;

		display: flex;
		flex-direction: column;
		gap: 1rem;
	}

	.top {
		display: flex;
		align-items: center;

		gap: 1rem;

		padding-inline: 0.5rem;

		h1 {
			color: $primary-color;
			font-size: 1.8rem;
		}
	}

	.divider {
		/* margin-top: 1rem; */
		margin-bottom: .5rem;

		border-bottom: 2px solid;
	}

	.main {
		padding-inline: 0.5rem;

		nav {
			display: flex;
			flex-direction: column;
			gap: 1rem;
			align-items: center;
		}
	}

	/* .footer {
		margin-top: auto;

		p {
			color: #fff;
			font-size: 1.2rem;
		}
	} */
</style>
