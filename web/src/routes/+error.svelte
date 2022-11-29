<script>
	import { page } from '$app/stores';
    
    export const prerender = true;
</script>

<h1>{$page.status}: {$page.error?.message}</h1>
