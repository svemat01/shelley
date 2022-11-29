<script lang="ts">
    import Fa from 'svelte-fa';
    import type { IconDefinition } from '@fortawesome/fontawesome-common-types'
    import { page } from '$app/stores';

	export let url: string;
	export let title: string;
    export let icon: IconDefinition | undefined = undefined;
</script>

<a href={url} class:selected={url == $page.route.id}>
    {#if icon}
        <Fa icon={icon} size='md'/>
    {/if}
    {title}
</a>

<style lang="scss">
    a {
        color: $light-color;
        font-size: 1.4rem;
        font-weight: 600;
        text-align: center;
        vertical-align: middle;

        width: 100%;
        border-radius: .5rem;
        padding: 0.5rem 0.8rem;

        &.selected {
            color: $primary-color;

            background-color: $dark-color;
            outline: $light-color solid 3px;
        }
    }
</style>
