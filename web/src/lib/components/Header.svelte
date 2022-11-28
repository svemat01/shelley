<script lang="ts">
    import { apiBaseUrl } from '../stores/api';

    const updateApiBaseUrl = (event: Event) => {
        const inputUrl = prompt('Enter the base url of the API', $apiBaseUrl);

        if (!inputUrl) {
            alert('Not updated');
            return;
        }

        try {
            new URL(inputUrl);
        } catch {
            alert('Invalid url');
            return;
        }

        apiBaseUrl.set(inputUrl);
    };

</script>

<header>
    <h1>My App</h1>
    <button on:click={updateApiBaseUrl}>Set API URL</button>
</header>