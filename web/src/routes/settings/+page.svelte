<div class="container">
	<div class="base-card card">
		Test
	</div>
</div>

<style lang="scss">
	.container {
		display: grid;
		grid-template-columns: repeat(4, 1fr);
		gap: 2rem;

		padding: 2rem;

		@media (max-width: $desktop-large) {
			grid-template-columns: repeat(3, 1fr);
		}
		@media (max-width: $desktop) {
			grid-template-columns: repeat(2, 1fr);
		}
		@media (max-width: $tablet) {
			grid-template-columns: 1fr;
		}
	}
</style>
