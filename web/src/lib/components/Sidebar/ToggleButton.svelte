<script lang="ts">
	import { faXmark } from '@fortawesome/free-solid-svg-icons';
	import Fa from 'svelte-fa';
	
</script>

<div class="button" on:click>
	<Fa icon={faXmark} size="2x" />
</div>

<style lang="scss">
	.button {
		color: $light-color;
		cursor: pointer;

		position: absolute;
        top: 1rem;
        right: 1rem;
	}
</style>
