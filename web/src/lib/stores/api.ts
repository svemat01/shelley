import { createLocalStorage, persist } from '@macfja/svelte-persistent-store';
import { writable } from 'svelte/store';

export const apiBaseUrl = persist(writable(''), createLocalStorage(), 'apiBaseUrl');
